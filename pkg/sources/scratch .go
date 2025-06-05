// Package sqlitehelper provides a lightweight, type‐safe way to create SQLite tables from Go structs
// and bulk‐import TSV files into those tables quickly. It uses reflection to infer column names/types,
// slog for structured logging, and buffered I/O + transactions for high throughput.
//
// This version does NOT use Go generics; instead, you pass a “model” (e.g. Artist{}) to both CreateTable
// and ImportTSV, and it discovers the struct fields at runtime via reflection.
//
// Usage example:
//
//	// 1) Define your struct with `db` tags to override column names or skip fields.
//	//    And `tsv` tags to indicate which TSV column (zero-based) maps to which field.
//	type Artist struct {
//	    ID   int     `db:"id,PRIMARY KEY AUTOINCREMENT"`
//	    MBID string  `db:"mbid,NOT NULL,UNIQUE" tsv:"0"`
//	    Name string  `db:"name"                tsv:"1"`
//	    Age  int     `db:"age"                 tsv:"2"`
//	    Note *string `db:"note"` // no tsv tag → not populated from TSV
//	}
//
//	// 2) In main or init, open a slog.Logger and a DBHelper.
//	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
//	logger := slog.New(handler)
//	helper, err := sqlitehelper.NewDBHelper("artists.db", logger)
//	if err != nil {
//	    log.Fatalf("cannot open DB: %v", err)
//	}
//	defer helper.Close()
//
//	// 3) Create (or recreate) the “artists” table based on the Artist{} struct.
//	if err := helper.CreateTable("artists", Artist{}, true); err != nil {
//	    log.Fatalf("create table failed: %v", err)
//	}
//
//	// 4) Import a TSV file into “artists” in batches of 5000 rows,
//	//    ignoring the “id” column on insert (because it’s AUTOINCREMENT).
//	ctx := context.Background()
//	if err := helper.ImportTSV(
//	    "artists",
//	    Artist{},
//	    "artists.tsv",
//	    5000,
//	    []string{"id"},
//	    ctx,
//	); err != nil {
//	    log.Fatalf("import TSV failed: %v", err)
//	}
package sources

import (
	"bufio"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"log/slog"

	// Make sure to include this in your go.mod:
	//   github.com/mattn/go-sqlite3 v1.14.20 (or latest)
	_ "github.com/mattn/go-sqlite3"
)

const (
	// Size of the buffered reader (4 MB).
	defaultReadBufferSize = 4 * 1024 * 1024

	// Max capacity per line (10 MB).
	maxLineCapacity = 10 * 1024 * 1024

	// How often (seconds) to log import progress.
	defaultLoggerInterval = 5
)

// DBHelper wraps *sql.DB + *slog.Logger for creating tables and importing TSVs.
type DBHelper struct {
	DB     *sql.DB
	Logger *slog.Logger
}

// NewDBHelper opens (or creates) a SQLite file at dbPath and returns a DBHelper.
// It also sets WAL mode and foreign_keys=ON by default.
func NewDBHelper(dbPath string, logger *slog.Logger) (*DBHelper, error) {
	if logger == nil {
		return nil, errors.New("sqlitehelper: logger cannot be nil")
	}
	// Enable WAL and foreign keys.
	connStr := dbPath + "?_journal_mode=WAL&_foreign_keys=ON"
	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, fmt.Errorf("sqlitehelper: open db: %w", err)
	}
	// SQLite is happiest with a single writer at a time.
	db.SetMaxOpenConns(1)
	return &DBHelper{
		DB:     db,
		Logger: logger,
	}, nil
}

// Close closes the underlying *sql.DB.
func (h *DBHelper) Close() error {
	return h.DB.Close()
}

// -----------------------------------------------------------------------------
// CreateTable
//
//	Recreate (or create if not exists) a table based on the struct type of `model`.
//	tableName:      the name of the table to create.
//	model:          a zero‐value of your struct (e.g. Artist{}). Must be a struct (or pointer to struct).
//	dropIfExists:   if true, first does “DROP TABLE IF EXISTS tableName;”
//
// It uses reflection on “model” to generate a CREATE TABLE statement.  Fields must be exported.
//
// Supported Go → SQLite type mappings (nullable if pointer‐type):
//
//	string     → TEXT
//	bool       → INTEGER
//	int*, uint* → INTEGER
//	float32/64 → REAL
//	[]byte     → BLOB
//	time.Time  → DATETIME
//
// Struct tags on each field of model (exported fields only):
//
//	`db:"<column_name>[,<SQL extras>]"`
//	   • If `<column_name>` is empty, uses the field name (lowercased).
//	   • If tag=="-", skips that field entirely.
//	   • After the comma, you can put “PRIMARY KEY”, “AUTOINCREMENT”, “NOT NULL”, etc.
//
// Example:
//
//	type Artist struct {
//	    ID   int    `db:"id,PRIMARY KEY AUTOINCREMENT"`
//	    MBID string `db:"mbid,NOT NULL,UNIQUE"`
//	    Name string `db:"name"`
//	    Age  int    // tagless → column "age" of type INTEGER
//	    Note *string `db:"note"`
//	}
//
// -----------------------------------------------------------------------------
func (h *DBHelper) CreateTable(tableName string, model interface{}, dropIfExists bool) error {
	// 1) Resolve reflect.Type of the struct:
	rt := reflect.TypeOf(model)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		return fmt.Errorf("sqlitehelper: CreateTable: model must be a struct or pointer to struct, got %s", rt.Kind())
	}

	// 2) Optionally DROP TABLE IF EXISTS
	if dropIfExists {
		dropSQL := fmt.Sprintf("DROP TABLE IF EXISTS %s;", tableName)
		if _, err := h.DB.Exec(dropSQL); err != nil {
			return fmt.Errorf("sqlitehelper: drop table: %w", err)
		}
		h.Logger.Info("Dropped table", "table", tableName)
	}

	// 3) Build the column definitions by walking through rt’s fields:
	var columns []string
	numFields := rt.NumField()
	for i := 0; i < numFields; i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue // skip unexported fields
		}
		tag := field.Tag.Get("db")
		if tag == "-" {
			continue // skip
		}

		// Determine column name and extras:
		var colName, extras string
		if tag != "" {
			parts := strings.SplitN(tag, ",", 2)
			colName = strings.TrimSpace(parts[0])
			if len(parts) == 2 {
				extras = strings.TrimSpace(parts[1])
			}
		}
		if colName == "" {
			// default to lowercase of field name
			colName = strings.ToLower(field.Name)
		}

		sqliteType, err := mapGoTypeToSQLite(field.Type)
		if err != nil {
			return fmt.Errorf("sqlitehelper: unsupported field '%s' (%s): %w",
				field.Name, field.Type.String(), err)
		}

		colDef := fmt.Sprintf("%s %s", colName, sqliteType)
		if extras != "" {
			colDef += " " + extras
		}
		columns = append(columns, colDef)
	}

	if len(columns) == 0 {
		return fmt.Errorf("sqlitehelper: CreateTable: no exported struct fields found in %s", rt.Name())
	}

	createSQL := fmt.Sprintf(
		"CREATE TABLE IF NOT EXISTS %s (\n    %s\n);",
		tableName,
		strings.Join(columns, ",\n    "),
	)
	if _, err := h.DB.Exec(createSQL); err != nil {
		return fmt.Errorf("sqlitehelper: create table: %w", err)
	}
	h.Logger.Info("Created table (or already exists)", "table", tableName)
	return nil
}

// mapGoTypeToSQLite infers a SQLite column type from a Go reflect.Type.
// Supports pointer types (nullable) and basic kinds.
func mapGoTypeToSQLite(rt reflect.Type) (string, error) {
	nullable := false
	if rt.Kind() == reflect.Ptr {
		nullable = true
		rt = rt.Elem()
	}

	var sqliteType string
	switch rt.Kind() {
	case reflect.String:
		sqliteType = "TEXT"
	case reflect.Bool:
		sqliteType = "INTEGER"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		sqliteType = "INTEGER"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		sqliteType = "INTEGER"
	case reflect.Float32, reflect.Float64:
		sqliteType = "REAL"
	case reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			sqliteType = "BLOB"
		} else {
			return "", fmt.Errorf("slice of %s not supported", rt.Elem().Kind())
		}
	default:
		// special: time.Time
		if rt == reflect.TypeOf(time.Time{}) {
			sqliteType = "DATETIME"
		} else {
			return "", fmt.Errorf("unsupported Go type: %s", rt.String())
		}
	}

	// (Nullable is automatic in SQLite unless you append NOT NULL in the tag.)
	_ = nullable
	return sqliteType, nil
}

// -----------------------------------------------------------------------------
// ImportTSV
//
//	Reads a TSV file (no header) and bulk‐inserts rows into `tableName`.  You must
//	annotate each field in `model` (struct) that you want to import with a `tsv:"<index>"` tag,
//	where <index> is a zero-based column index in the TSV file.
//
//	tableName:    the existing table you already created with CreateTable.
//	model:        a zero-value struct (or pointer) with field tags `tsv:"<index>"`.
//	filePath:     path to the TSV file (tab-separated, no header).
//	batchSize:    number of rows per transaction (e.g. 5000).
//	ignoreCols:   slice of column names to skip in the INSERT (e.g. ["id"] if id is AUTOINCREMENT).
//	ctx:          context for cancellation.
//
// Requirements on `model` and its tags:
//   - model must be a struct (or pointer to struct).
//   - Each field you want to import from the TSV must have a tag `tsv:"<zero-based index>"`.
//   - If a field has no `tsv` tag, it is omitted from INSERT (e.g. ID).
//   - The struct may also have fields without TSV tags (e.g. an AUTOINCREMENT PK).
//   - Conversion rules for each field’s Go type: string, bool, ints, uints, floats, time.Time
//     (parsed via time.RFC3339).  Pointer types become nullable: an empty TSV cell → nil pointer.
//
// -----------------------------------------------------------------------------
func (h *DBHelper) ImportTSV(
	tableName string,
	model interface{},
	filePath string,
	batchSize int,
	ignoreCols []string,
	ctx context.Context,
) error {
	// 0) Reflect on model to gather field info
	rt := reflect.TypeOf(model)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	if rt.Kind() != reflect.Struct {
		return fmt.Errorf("sqlitehelper: ImportTSV: model must be a struct or pointer to struct, got %s", rt.Kind())
	}

	// Build a map: tsvIndex → fieldInfo
	type fieldInfo struct {
		fieldIndex int          // index within the struct
		colName    string       // column name in SQLite
		colType    reflect.Type // Go type for conversion
	}
	tsvFields := make(map[int]fieldInfo)
	maxTSVIndex := -1

	// Turn ignoreCols into a set for quick lookup
	ignoreSet := make(map[string]struct{}, len(ignoreCols))
	for _, c := range ignoreCols {
		ignoreSet[c] = struct{}{}
	}

	numFields := rt.NumField()
	for i := 0; i < numFields; i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		tsvTag := field.Tag.Get("tsv")
		if tsvTag == "" {
			continue // this field isn’t loaded from TSV
		}
		idx, err := strconv.Atoi(tsvTag)
		if err != nil {
			return fmt.Errorf("sqlitehelper: ImportTSV: invalid tsv tag on field '%s': %w", field.Name, err)
		}

		// Determine column name using the same db‐tag logic as CreateTable:
		dbTag := field.Tag.Get("db")
		var colName string
		if dbTag != "" {
			parts := strings.SplitN(dbTag, ",", 2)
			colName = strings.TrimSpace(parts[0])
			if colName == "" || colName == "-" {
				return fmt.Errorf("sqlitehelper: ImportTSV: field '%s' has invalid db tag", field.Name)
			}
		}
		if colName == "" {
			colName = strings.ToLower(field.Name)
		}
		if _, skip := ignoreSet[colName]; skip {
			continue
		}
		if idx > maxTSVIndex {
			maxTSVIndex = idx
		}
		tsvFields[idx] = fieldInfo{
			fieldIndex: i,
			colName:    colName,
			colType:    field.Type,
		}
	}
	if len(tsvFields) == 0 {
		return errors.New("sqlitehelper: ImportTSV: no fields with `tsv` tag found")
	}

	// Build a sorted list of indexes
	orderedIndexes := make([]int, 0, len(tsvFields))
	for idx := range tsvFields {
		orderedIndexes = append(orderedIndexes, idx)
	}
	sortInts(orderedIndexes)

	// Build column names and “?” placeholders for the INSERT statement:
	colNames := make([]string, 0, len(orderedIndexes))
	placeholders := make([]string, 0, len(orderedIndexes))
	for _, idx := range orderedIndexes {
		colNames = append(colNames, tsvFields[idx].colName)
		placeholders = append(placeholders, "?")
	}
	insertSQL := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s);",
		tableName,
		strings.Join(colNames, ", "),
		strings.Join(placeholders, ", "),
	)

	stmt, err := h.DB.Prepare(insertSQL)
	if err != nil {
		return fmt.Errorf("sqlitehelper: prepare insert: %w", err)
	}
	defer stmt.Close()

	// 1) Open the TSV file and get its size (for progress logging)
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("sqlitehelper: open file: %w", err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("sqlitehelper: fstat: %w", err)
	}
	totalSize := stat.Size()

	h.Logger.Info("Starting TSV import",
		"table", tableName,
		"file", stat.Name(),
		"size_mb", totalSize/(1024*1024),
		"batch_size", batchSize,
	)

	// 2) Background goroutine for progress logging
	var totalInserted atomic.Int64
	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(defaultLoggerInterval * time.Second)
		defer ticker.Stop()
		var lastCount int64
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				pos, _ := file.Seek(0, io.SeekCurrent)
				percent := float64(pos) / float64(totalSize) * 100
				currentCount := totalInserted.Load()
				rowsPerSec := float64(currentCount-lastCount) / float64(defaultLoggerInterval)
				lastCount = currentCount

				h.Logger.Info("Import progress",
					"percent", fmt.Sprintf("%.2f%%", percent),
					"rows_per_sec", fmt.Sprintf("%.0f", rowsPerSec),
					"total_inserted", currentCount,
				)
			}
		}
	}()

	// 3) Read & parse the TSV in batches
	reader := bufio.NewReaderSize(file, defaultReadBufferSize)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, defaultReadBufferSize), maxLineCapacity)

	// Begin the first transaction
	tx, err := h.DB.BeginTx(ctx, nil)
	if err != nil {
		close(done)
		return fmt.Errorf("sqlitehelper: begin tx: %w", err)
	}
	insertedThisBatch := 0

	// helper to roll back on error
	parseErr := func(e error) error {
		_ = tx.Rollback()
		close(done)
		return e
	}

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		cols := bytesSplitNoAlloc(line, '\t')
		if len(cols)-1 < maxTSVIndex {
			return parseErr(fmt.Errorf(
				"sqlitehelper: fewer TSV columns (%d) than expected index %d",
				len(cols), maxTSVIndex))
		}

		// Build args in the correct order:
		args := make([]interface{}, 0, len(orderedIndexes))
		for _, idx := range orderedIndexes {
			info := tsvFields[idx]
			raw := cols[idx]
			val, err := convertBytesToGo(raw, info.colType)
			if err != nil {
				return parseErr(fmt.Errorf("sqlitehelper: convert col %d to %s: %w", idx, info.colType, err))
			}
			args = append(args, val)
		}

		// Exec inside current TX
		if _, err := tx.StmtContext(ctx, stmt).ExecContext(ctx, args...); err != nil {
			return parseErr(fmt.Errorf("sqlitehelper: exec insert: %w", err))
		}
		insertedThisBatch++
		totalInserted.Add(1)

		// If batch full, commit and start a new TX
		if insertedThisBatch >= batchSize {
			if err := tx.Commit(); err != nil {
				return parseErr(fmt.Errorf("sqlitehelper: commit: %w", err))
			}
			tx, err = h.DB.BeginTx(ctx, nil)
			if err != nil {
				close(done)
				return fmt.Errorf("sqlitehelper: begin tx: %w", err)
			}
			insertedThisBatch = 0
		}

		// Check cancellation
		select {
		case <-ctx.Done():
			return parseErr(ctx.Err())
		default:
		}
	}

	if err := scanner.Err(); err != nil {
		return parseErr(fmt.Errorf("sqlitehelper: scanner error: %w", err))
	}

	// Commit any remaining rows
	if insertedThisBatch > 0 {
		if err := tx.Commit(); err != nil {
			return parseErr(fmt.Errorf("sqlitehelper: final commit: %w", err))
		}
	} else {
		// No rows in this partial TX → just rollback to close it
		_ = tx.Rollback()
	}

	close(done)
	h.Logger.Info("Import complete",
		"table", tableName,
		"total_inserted", totalInserted.Load(),
	)
	return nil
}

// convertBytesToGo converts a raw TSV cell (as []byte) to an appropriate Go value
// for a given destType.  Supports pointer types (nullable) and basic kinds.
func convertBytesToGo(raw []byte, destType reflect.Type) (interface{}, error) {
	// Check if destType is a pointer
	isPtr := false
	if destType.Kind() == reflect.Ptr {
		isPtr = true
		destType = destType.Elem()
	}
	str := string(raw)
	if str == "" {
		// Empty cell → nil if pointer, else zero value
		if isPtr {
			return nil, nil
		}
		switch destType.Kind() {
		case reflect.String:
			return "", nil
		case reflect.Bool:
			return false, nil
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int64(0), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return uint64(0), nil
		case reflect.Float32, reflect.Float64:
			return float64(0), nil
		default:
			if destType == reflect.TypeOf(time.Time{}) {
				return time.Time{}, nil
			}
			return nil, fmt.Errorf("sqlitehelper: unsupported zero value for type %s", destType)
		}
	}

	// Non-empty cell: parse according to destType.Kind()
	switch destType.Kind() {
	case reflect.String:
		if isPtr {
			v := str
			return &v, nil
		}
		return str, nil

	case reflect.Bool:
		v, err := strconv.ParseBool(str)
		if err != nil {
			return nil, err
		}
		if isPtr {
			b := v
			return &b, nil
		}
		return v, nil

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		if isPtr {
			switch destType.Kind() {
			case reflect.Int:
				tmp := int(v)
				return &tmp, nil
			case reflect.Int8:
				tmp := int8(v)
				return &tmp, nil
			case reflect.Int16:
				tmp := int16(v)
				return &tmp, nil
			case reflect.Int32:
				tmp := int32(v)
				return &tmp, nil
			case reflect.Int64:
				tmp := int64(v)
				return &tmp, nil
			}
		}
		return v, nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uv, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return nil, err
		}
		if isPtr {
			switch destType.Kind() {
			case reflect.Uint:
				tmp := uint(uv)
				return &tmp, nil
			case reflect.Uint8:
				tmp := uint8(uv)
				return &tmp, nil
			case reflect.Uint16:
				tmp := uint16(uv)
				return &tmp, nil
			case reflect.Uint32:
				tmp := uint32(uv)
				return &tmp, nil
			case reflect.Uint64:
				tmp := uint64(uv)
				return &tmp, nil
			}
		}
		return uv, nil

	case reflect.Float32, reflect.Float64:
		fv, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		if isPtr {
			if destType.Kind() == reflect.Float32 {
				tmp := float32(fv)
				return &tmp, nil
			}
			tmp := float64(fv)
			return &tmp, nil
		}
		return fv, nil

	default:
		// Maybe time.Time?
		if destType == reflect.TypeOf(time.Time{}) {
			tm, err := time.Parse(time.RFC3339, str)
			if err != nil {
				return nil, err
			}
			if isPtr {
				return &tm, nil
			}
			return tm, nil
		}
		return nil, fmt.Errorf("sqlitehelper: unsupported conversion to Go type %s", destType)
	}
}

// bytesSplitNoAlloc splits a []byte by delimiter into subslices without extra allocations.
// It’s slightly “hacky but fast” for TSV parsing.
func bytesSplitNoAlloc(b []byte, delim byte) [][]byte {
	var result [][]byte
	start := 0
	for i := 0; i < len(b); i++ {
		if b[i] == delim {
			result = append(result, b[start:i])
			start = i + 1
		}
	}
	result = append(result, b[start:])
	return result
}

// sortInts is a simple insertion sort to sort small []int ascending.
func sortInts(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

package sources

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync/atomic"
	"time"

	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

const READ_BUFFER_SIZE = 1024 * 1024 * 10
const MAX_LINE_CAPACITY = 64 * 1024 * 10
const LOGGER_INTERVAL = 2

func ImportArtists(filePath string, db *ent.Client, batchSize int, ctx context.Context) error {
	start := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}
	slog.Info("Importing Artists",
		"name", fileStat.Name(),
		"size_mb", fileStat.Size()/(1024*1024),
	)

	var totalInserts atomic.Int64
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(LOGGER_INTERVAL * time.Second)
		defer ticker.Stop()

		var lastCount int64
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				pos, _ := file.Seek(0, io.SeekCurrent)
				percent := float64(pos) / float64(fileStat.Size()) * 100

				currentCount := totalInserts.Load()
				insertsPerSecond := float64(currentCount-lastCount) / float64(LOGGER_INTERVAL)
				lastCount = currentCount

				slog.Info("Import progress",
					"percent", fmt.Sprintf("%.2f%%", percent),
					"inserts_per_sec", fmt.Sprintf("%.0f", insertsPerSecond),
					"total_inserted", currentCount,
				)
			}
		}
	}()

	reader := bufio.NewReaderSize(file, READ_BUFFER_SIZE)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 0, READ_BUFFER_SIZE), MAX_LINE_CAPACITY)

	batch := make([]*ent.ArtistCreate, 0, batchSize)
	for scanner.Scan() {
		line := scanner.Bytes()
		fields := bytes.Split(line, []byte("\t"))

		mbid, err := uuid.ParseBytes(fields[1])
		if err != nil {
			slog.Error("uuid parse failed", "err", err, "line", line, "fields", fields)
			close(done)
			return err
		}
		ac := db.Artist.Create().SetMbid(mbid).SetName(string(fields[2]))
		batch = append(batch, ac)

		if len(batch) >= batchSize {
			err = db.Artist.CreateBulk(batch...).Exec(ctx)
			if err != nil {
				close(done)
				return err
			}
			totalInserts.Add(int64(len(batch)))
			batch = batch[:0] // reuse the same slice
		}
	}

	if err := scanner.Err(); err != nil && err != io.EOF {
		close(done)
		return err
	}

	if len(batch) > 0 {
		err = db.Artist.CreateBulk(batch...).Exec(ctx)
		if err != nil {
			close(done)
			return err
		}
		totalInserts.Add(int64(len(batch)))
	}

	close(done)
	slog.Info("Import completed",
		"total_inserted", totalInserts.Load(),
		"total_duration", time.Since(start),
	)
	return nil
}

func ImportReleases(filePath string, db *ent.Client, batchSize int, ctx context.Context) error {
	return nil
}

func GetCmd() *cobra.Command {
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "data import",
	}

	ImportArtistsCmd := &cobra.Command{
		Use:   "artists",
		Short: "imparts all artists",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.NewDatabase("file:./data/database.db?_fk=1&_journal_mode=WAL")
			if err != nil {
				return err
			}
			return ImportArtists(args[0], db.Client, 500, context.Background())
		},
	}
	ImportReleasesCmd := &cobra.Command{
		Use:   "releases",
		Short: "imparts all releases",
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.NewDatabase("file:./data/database.db?_fk=1&_journal_mode=WAL")
			if err != nil {
				return err
			}
			return ImportArtists(args[0], db.Client, 500, context.Background())
		},
	}

	importCmd.AddCommand(ImportArtistsCmd)
	importCmd.AddCommand(ImportReleasesCmd)

	return importCmd
}

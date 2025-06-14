// Code generated by ent, DO NOT EDIT.

package image

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

const (
	// Label holds the string label denoting the image type in the database.
	Label = "image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFile holds the string denoting the file field in the database.
	FieldFile = "file"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldDimentions holds the string denoting the dimentions field in the database.
	FieldDimentions = "dimentions"
	// FieldSizeBits holds the string denoting the size_bits field in the database.
	FieldSizeBits = "size_bits"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeRelease holds the string denoting the release edge name in mutations.
	EdgeRelease = "release"
	// EdgeUploader holds the string denoting the uploader edge name in mutations.
	EdgeUploader = "uploader"
	// Table holds the table name of the image in the database.
	Table = "images"
	// ReleaseTable is the table that holds the release relation/edge.
	ReleaseTable = "images"
	// ReleaseInverseTable is the table name for the Release entity.
	// It exists in this package in order to avoid circular dependency with the "release" package.
	ReleaseInverseTable = "releases"
	// ReleaseColumn is the table column denoting the release relation/edge.
	ReleaseColumn = "release_image"
	// UploaderTable is the table that holds the uploader relation/edge.
	UploaderTable = "images"
	// UploaderInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UploaderInverseTable = "users"
	// UploaderColumn is the table column denoting the uploader relation/edge.
	UploaderColumn = "user_images"
)

// Columns holds all SQL columns for image fields.
var Columns = []string{
	FieldID,
	FieldFile,
	FieldType,
	FieldNote,
	FieldDimentions,
	FieldSizeBits,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"release_image",
	"user_images",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// FileValidator is a validator for the "file" field. It is called by the builders before save.
	FileValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pid.ID
)

// Type defines the type for the "type" enum field.
type Type string

// Type values.
const (
	TypeWebp Type = "webp"
	TypePng  Type = "png"
	TypeJpg  Type = "jpg"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeWebp, TypePng, TypeJpg:
		return nil
	default:
		return fmt.Errorf("image: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Image queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFile orders the results by the file field.
func ByFile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFile, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByNote orders the results by the note field.
func ByNote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNote, opts...).ToFunc()
}

// BySizeBits orders the results by the size_bits field.
func BySizeBits(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSizeBits, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByReleaseField orders the results by release field.
func ByReleaseField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReleaseStep(), sql.OrderByField(field, opts...))
	}
}

// ByUploaderField orders the results by uploader field.
func ByUploaderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUploaderStep(), sql.OrderByField(field, opts...))
	}
}
func newReleaseStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReleaseInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ReleaseTable, ReleaseColumn),
	)
}
func newUploaderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UploaderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UploaderTable, UploaderColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package processedimage

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

const (
	// Label holds the string label denoting the processedimage type in the database.
	Label = "processed_image"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldDimentions holds the string denoting the dimentions field in the database.
	FieldDimentions = "dimentions"
	// FieldSizeBits holds the string denoting the size_bits field in the database.
	FieldSizeBits = "size_bits"
	// FieldThumb holds the string denoting the thumb field in the database.
	FieldThumb = "thumb"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// EdgeSource holds the string denoting the source edge name in mutations.
	EdgeSource = "source"
	// Table holds the table name of the processedimage in the database.
	Table = "processed_images"
	// SourceTable is the table that holds the source relation/edge.
	SourceTable = "processed_images"
	// SourceInverseTable is the table name for the Image entity.
	// It exists in this package in order to avoid circular dependency with the "image" package.
	SourceInverseTable = "images"
	// SourceColumn is the table column denoting the source relation/edge.
	SourceColumn = "image_proccesed_image"
)

// Columns holds all SQL columns for processedimage fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldDimentions,
	FieldSizeBits,
	FieldThumb,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "processed_images"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"image_proccesed_image",
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
	// DimentionsValidator is a validator for the "dimentions" field. It is called by the builders before save.
	DimentionsValidator func(int) error
	// ThumbValidator is a validator for the "thumb" field. It is called by the builders before save.
	ThumbValidator func([]byte) error
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
	TypeWEBP Type = "WEBP"
	TypePNG  Type = "PNG"
	TypeJPG  Type = "JPG"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeWEBP, TypePNG, TypeJPG:
		return nil
	default:
		return fmt.Errorf("processedimage: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the ProcessedImage queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByDimentions orders the results by the dimentions field.
func ByDimentions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDimentions, opts...).ToFunc()
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

// BySourceField orders the results by source field.
func BySourceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourceStep(), sql.OrderByField(field, opts...))
	}
}
func newSourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, SourceTable, SourceColumn),
	)
}

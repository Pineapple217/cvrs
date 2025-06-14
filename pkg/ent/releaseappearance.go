// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/Pineapple217/cvrs/pkg/ent/artist"
	"github.com/Pineapple217/cvrs/pkg/ent/release"
	"github.com/Pineapple217/cvrs/pkg/ent/releaseappearance"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

// ReleaseAppearance is the model entity for the ReleaseAppearance schema.
type ReleaseAppearance struct {
	config `json:"-"`
	// ReleaseID holds the value of the "release_id" field.
	ReleaseID pid.ID `json:"release_id,omitempty"`
	// ArtistID holds the value of the "artist_id" field.
	ArtistID pid.ID `json:"artist_id,omitempty"`
	// Order holds the value of the "order" field.
	Order int `json:"order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ReleaseAppearanceQuery when eager-loading is set.
	Edges        ReleaseAppearanceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ReleaseAppearanceEdges holds the relations/edges for other nodes in the graph.
type ReleaseAppearanceEdges struct {
	// Artist holds the value of the artist edge.
	Artist *Artist `json:"artist,omitempty"`
	// Release holds the value of the release edge.
	Release *Release `json:"release,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ArtistOrErr returns the Artist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReleaseAppearanceEdges) ArtistOrErr() (*Artist, error) {
	if e.Artist != nil {
		return e.Artist, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: artist.Label}
	}
	return nil, &NotLoadedError{edge: "artist"}
}

// ReleaseOrErr returns the Release value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ReleaseAppearanceEdges) ReleaseOrErr() (*Release, error) {
	if e.Release != nil {
		return e.Release, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: release.Label}
	}
	return nil, &NotLoadedError{edge: "release"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ReleaseAppearance) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case releaseappearance.FieldReleaseID, releaseappearance.FieldArtistID, releaseappearance.FieldOrder:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ReleaseAppearance fields.
func (ra *ReleaseAppearance) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case releaseappearance.FieldReleaseID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field release_id", values[i])
			} else if value.Valid {
				ra.ReleaseID = pid.ID(value.Int64)
			}
		case releaseappearance.FieldArtistID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field artist_id", values[i])
			} else if value.Valid {
				ra.ArtistID = pid.ID(value.Int64)
			}
		case releaseappearance.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				ra.Order = int(value.Int64)
			}
		default:
			ra.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ReleaseAppearance.
// This includes values selected through modifiers, order, etc.
func (ra *ReleaseAppearance) Value(name string) (ent.Value, error) {
	return ra.selectValues.Get(name)
}

// QueryArtist queries the "artist" edge of the ReleaseAppearance entity.
func (ra *ReleaseAppearance) QueryArtist() *ArtistQuery {
	return NewReleaseAppearanceClient(ra.config).QueryArtist(ra)
}

// QueryRelease queries the "release" edge of the ReleaseAppearance entity.
func (ra *ReleaseAppearance) QueryRelease() *ReleaseQuery {
	return NewReleaseAppearanceClient(ra.config).QueryRelease(ra)
}

// Update returns a builder for updating this ReleaseAppearance.
// Note that you need to call ReleaseAppearance.Unwrap() before calling this method if this ReleaseAppearance
// was returned from a transaction, and the transaction was committed or rolled back.
func (ra *ReleaseAppearance) Update() *ReleaseAppearanceUpdateOne {
	return NewReleaseAppearanceClient(ra.config).UpdateOne(ra)
}

// Unwrap unwraps the ReleaseAppearance entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ra *ReleaseAppearance) Unwrap() *ReleaseAppearance {
	_tx, ok := ra.config.driver.(*txDriver)
	if !ok {
		panic("ent: ReleaseAppearance is not a transactional entity")
	}
	ra.config.driver = _tx.drv
	return ra
}

// String implements the fmt.Stringer.
func (ra *ReleaseAppearance) String() string {
	var builder strings.Builder
	builder.WriteString("ReleaseAppearance(")
	builder.WriteString("release_id=")
	builder.WriteString(fmt.Sprintf("%v", ra.ReleaseID))
	builder.WriteString(", ")
	builder.WriteString("artist_id=")
	builder.WriteString(fmt.Sprintf("%v", ra.ArtistID))
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", ra.Order))
	builder.WriteByte(')')
	return builder.String()
}

// ReleaseAppearances is a parsable slice of ReleaseAppearance.
type ReleaseAppearances []*ReleaseAppearance

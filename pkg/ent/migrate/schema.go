// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArtistsColumns holds the columns for the "artists" table.
	ArtistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// ArtistsTable holds the schema information for the "artists" table.
	ArtistsTable = &schema.Table{
		Name:       "artists",
		Columns:    ArtistsColumns,
		PrimaryKey: []*schema.Column{ArtistsColumns[0]},
	}
	// ImagesColumns holds the columns for the "images" table.
	ImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "file", Type: field.TypeString},
		{Name: "original_name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WEBP", "PNG", "JPG"}},
		{Name: "note", Type: field.TypeString, Nullable: true},
		{Name: "dimention_width", Type: field.TypeInt},
		{Name: "dimention_height", Type: field.TypeInt},
		{Name: "size_bits", Type: field.TypeUint32},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "artist_image", Type: field.TypeInt64, Unique: true, Nullable: true},
		{Name: "image_data", Type: field.TypeInt, Nullable: true},
		{Name: "release_image", Type: field.TypeInt64, Unique: true, Nullable: true},
		{Name: "user_images", Type: field.TypeInt64},
	}
	// ImagesTable holds the schema information for the "images" table.
	ImagesTable = &schema.Table{
		Name:       "images",
		Columns:    ImagesColumns,
		PrimaryKey: []*schema.Column{ImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "images_artists_image",
				Columns:    []*schema.Column{ImagesColumns[11]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_image_data_data",
				Columns:    []*schema.Column{ImagesColumns[12]},
				RefColumns: []*schema.Column{ImageDataColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_releases_image",
				Columns:    []*schema.Column{ImagesColumns[13]},
				RefColumns: []*schema.Column{ReleasesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "images_users_images",
				Columns:    []*schema.Column{ImagesColumns[14]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ImageDataColumns holds the columns for the "image_data" table.
	ImageDataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "avr_r", Type: field.TypeInt},
		{Name: "avr_g", Type: field.TypeInt},
		{Name: "avr_b", Type: field.TypeInt},
		{Name: "avg_brightness", Type: field.TypeInt},
		{Name: "avg_saturation", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ImageDataTable holds the schema information for the "image_data" table.
	ImageDataTable = &schema.Table{
		Name:       "image_data",
		Columns:    ImageDataColumns,
		PrimaryKey: []*schema.Column{ImageDataColumns[0]},
	}
	// ProcessedImagesColumns holds the columns for the "processed_images" table.
	ProcessedImagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"WEBP", "PNG", "JPG"}},
		{Name: "dimentions", Type: field.TypeInt},
		{Name: "size_bits", Type: field.TypeUint32},
		{Name: "thumb", Type: field.TypeBytes},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "image_proccesed_image", Type: field.TypeInt64},
	}
	// ProcessedImagesTable holds the schema information for the "processed_images" table.
	ProcessedImagesTable = &schema.Table{
		Name:       "processed_images",
		Columns:    ProcessedImagesColumns,
		PrimaryKey: []*schema.Column{ProcessedImagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "processed_images_images_proccesed_image",
				Columns:    []*schema.Column{ProcessedImagesColumns[8]},
				RefColumns: []*schema.Column{ImagesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "processedimage_image_proccesed_image",
				Unique:  false,
				Columns: []*schema.Column{ProcessedImagesColumns[8]},
			},
		},
	}
	// ReleasesColumns holds the columns for the "releases" table.
	ReleasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"album", "single", "EP", "compilation", "unknown"}},
		{Name: "release_date", Type: field.TypeTime},
	}
	// ReleasesTable holds the schema information for the "releases" table.
	ReleasesTable = &schema.Table{
		Name:       "releases",
		Columns:    ReleasesColumns,
		PrimaryKey: []*schema.Column{ReleasesColumns[0]},
	}
	// ReleaseAppearancesColumns holds the columns for the "release_appearances" table.
	ReleaseAppearancesColumns = []*schema.Column{
		{Name: "order", Type: field.TypeInt},
		{Name: "artist_id", Type: field.TypeInt64},
		{Name: "release_id", Type: field.TypeInt64},
	}
	// ReleaseAppearancesTable holds the schema information for the "release_appearances" table.
	ReleaseAppearancesTable = &schema.Table{
		Name:       "release_appearances",
		Columns:    ReleaseAppearancesColumns,
		PrimaryKey: []*schema.Column{ReleaseAppearancesColumns[1], ReleaseAppearancesColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "release_appearances_artists_artist",
				Columns:    []*schema.Column{ReleaseAppearancesColumns[1]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "release_appearances_releases_release",
				Columns:    []*schema.Column{ReleaseAppearancesColumns[2]},
				RefColumns: []*schema.Column{ReleasesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TasksColumns holds the columns for the "tasks" table.
	TasksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"scale_img"}},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "working", "error", "done"}, Default: "pending"},
		{Name: "error", Type: field.TypeString, Nullable: true},
		{Name: "payload", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TasksTable holds the schema information for the "tasks" table.
	TasksTable = &schema.Table{
		Name:       "tasks",
		Columns:    TasksColumns,
		PrimaryKey: []*schema.Column{TasksColumns[0]},
	}
	// TracksColumns holds the columns for the "tracks" table.
	TracksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "position", Type: field.TypeInt},
		{Name: "release_tracks", Type: field.TypeInt64},
	}
	// TracksTable holds the schema information for the "tracks" table.
	TracksTable = &schema.Table{
		Name:       "tracks",
		Columns:    TracksColumns,
		PrimaryKey: []*schema.Column{TracksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tracks_releases_tracks",
				Columns:    []*schema.Column{TracksColumns[3]},
				RefColumns: []*schema.Column{ReleasesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TrackAppearancesColumns holds the columns for the "track_appearances" table.
	TrackAppearancesColumns = []*schema.Column{
		{Name: "order", Type: field.TypeInt},
		{Name: "artist_id", Type: field.TypeInt64},
		{Name: "track_id", Type: field.TypeInt64},
	}
	// TrackAppearancesTable holds the schema information for the "track_appearances" table.
	TrackAppearancesTable = &schema.Table{
		Name:       "track_appearances",
		Columns:    TrackAppearancesColumns,
		PrimaryKey: []*schema.Column{TrackAppearancesColumns[1], TrackAppearancesColumns[2]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "track_appearances_artists_artist",
				Columns:    []*schema.Column{TrackAppearancesColumns[1]},
				RefColumns: []*schema.Column{ArtistsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "track_appearances_tracks_track",
				Columns:    []*schema.Column{TrackAppearancesColumns[2]},
				RefColumns: []*schema.Column{TracksColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt64, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 32},
		{Name: "password", Type: field.TypeBytes},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArtistsTable,
		ImagesTable,
		ImageDataTable,
		ProcessedImagesTable,
		ReleasesTable,
		ReleaseAppearancesTable,
		TasksTable,
		TracksTable,
		TrackAppearancesTable,
		UsersTable,
	}
)

func init() {
	ImagesTable.ForeignKeys[0].RefTable = ArtistsTable
	ImagesTable.ForeignKeys[1].RefTable = ImageDataTable
	ImagesTable.ForeignKeys[2].RefTable = ReleasesTable
	ImagesTable.ForeignKeys[3].RefTable = UsersTable
	ProcessedImagesTable.ForeignKeys[0].RefTable = ImagesTable
	ReleaseAppearancesTable.ForeignKeys[0].RefTable = ArtistsTable
	ReleaseAppearancesTable.ForeignKeys[1].RefTable = ReleasesTable
	TracksTable.ForeignKeys[0].RefTable = ReleasesTable
	TrackAppearancesTable.ForeignKeys[0].RefTable = ArtistsTable
	TrackAppearancesTable.ForeignKeys[1].RefTable = TracksTable
}

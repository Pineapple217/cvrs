package sources

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/Pineapple217/cvrs/pkg/ent/artist"
	"github.com/Pineapple217/cvrs/pkg/ent/release"
	"github.com/Pineapple217/cvrs/pkg/pid"
	"github.com/spf13/cobra"
)

const LOGGER_INTERVAL = 2 * time.Second

type decodeFunc[T any] func(decoder *xml.Decoder, se xml.StartElement) (T, error)
type bulkInsertFunc[T any] func(db *ent.Client, ctx context.Context, items []T) error

func ImportEntities[T any](
	filePath string,
	elementName string,
	db *ent.Client,
	batchSize int,
	decode decodeFunc[T],
	bulkInsert bulkInsertFunc[T],
	ctx context.Context,
) error {
	start := time.Now()
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)
	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("stat file: %w", err)
	}

	slog.Info("Importing",
		"name", fileInfo.Name(),
		"size_mb", fileInfo.Size()/(1024*1024),
		"element", elementName,
	)

	var total atomic.Int64
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(LOGGER_INTERVAL)
		defer ticker.Stop()

		var last int64
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				pos := decoder.InputOffset()
				percent := float64(pos) * 100 / float64(fileInfo.Size())

				current := total.Load()
				rate := float64(current-last) / LOGGER_INTERVAL.Seconds()
				last = current

				slog.Info("Progress",
					"percent", fmt.Sprintf("%.2f%%", percent),
					"rate_per_sec", int(rate),
					"total", current,
				)
			}
		}
	}()

	batch := make([]T, 0, batchSize)

	for {
		tok, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			close(done)
			return fmt.Errorf("read token: %w", err)
		}

		se, ok := tok.(xml.StartElement)
		if !ok || se.Name.Local != elementName {
			continue // OPTIMIZED: merge check for fewer branches
		}
		item, err := decode(decoder, se)
		if err != nil {
			slog.Error("decode error", "err", err)
			continue
		}
		batch = append(batch, item)
		if len(batch) >= batchSize {
			if err := bulkInsert(db, ctx, batch); err != nil {
				close(done)
				return fmt.Errorf("bulk insert: %w", err)
			}
			total.Add(int64(len(batch)))
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := bulkInsert(db, ctx, batch); err != nil {
			close(done)
			return err
		}
		total.Add(int64(len(batch)))
	}
	close(done)

	slog.Info("Import completed",
		"total", total.Load(),
		"duration", time.Since(start).Truncate(time.Millisecond),
	)
	return nil
}

type DiscogsArtist struct {
	ID   int64  `xml:"id"`
	Name string `xml:"name"`
}

func decodeArtist(decoder *xml.Decoder, se xml.StartElement) (DiscogsArtist, error) {
	var a DiscogsArtist
	return a, decoder.DecodeElement(&a, &se)
}

func insertArtists(db *ent.Client, ctx context.Context, items []DiscogsArtist) error {
	builders := make([]*ent.ArtistCreate, 0, len(items))
	for _, a := range items {
		if a.Name == "" {
			slog.Warn("no name found", "id", a.ID)
			continue
		}
		builders = append(builders, db.Artist.Create().SetDid(a.ID).SetName(a.Name))
	}
	return db.Artist.CreateBulk(builders...).Exec(ctx)
}

type DiscogsRelease struct {
	ID      string          `xml:"id,attr"`
	Title   string          `xml:"title"`
	Artists []DiscogsArtist `xml:"artists>artist"`
	// ExtraArtists []DiscogsExtraArtist `xml:"extraartists>artist"`
	Formats   []DiscogsFormat `xml:"formats>format"`
	Tracklist []DiscogsTrack  `xml:"tracklist>track"`
	Released  string          `xml:"released"`
	// Country      string               `xml:"country"`
	// DataQuality  string               `xml:"data_quality"`
}

func (r *DiscogsRelease) FlattenDescriptions() []string {
	var descriptions []string
	for _, format := range r.Formats {
		descriptions = append(descriptions, format.Descriptions...)
	}
	return descriptions
}

type DiscogsExtraArtist struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
	Role string `xml:"role"`
}

type DiscogsFormat struct {
	Name         string   `xml:"name,attr"`
	Qty          string   `xml:"qty,attr"`
	Text         string   `xml:"text,attr"`
	Descriptions []string `xml:"descriptions>description"`
}

type DiscogsTrack struct {
	Position     string               `xml:"position"`
	Title        string               `xml:"title"`
	Duration     string               `xml:"duration"`
	ExtraArtists []DiscogsExtraArtist `xml:"extraartists>artist"`
}

func decodeRelease(decoder *xml.Decoder, se xml.StartElement) (DiscogsRelease, error) {
	var a DiscogsRelease
	return a, decoder.DecodeElement(&a, &se)
}

func parseDate(d string) time.Time {
	if len(d) == 4 {
		if year, err := strconv.Atoi(d); err == nil {
			return time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		}
	}
	parts := strings.Split(d, "-")
	if len(parts) == 3 {
		year, err1 := strconv.Atoi(parts[0])
		month, err2 := strconv.Atoi(parts[1])
		day, err3 := strconv.Atoi(parts[2])

		if err1 == nil && err2 == nil && err3 == nil {
			if day == 0 {
				return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
			}
			return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
		}
	}

	return time.Date(0000, time.January, 1, 0, 0, 0, 0, time.UTC)
}

func insertRelease(db *ent.Client, ctx context.Context, items []DiscogsRelease) error {
	builders := make([]*ent.ReleaseCreate, 0, len(items))
	buildersRA := make([]*ent.ReleaseAppearanceCreate, 0)
	buildersTrack := make([]*ent.TrackCreate, 0)
	buildersRA := make([]*ent.TrackAppearanceCreate, 0)
	for _, r := range items {
		t := release.TypeUnknown
		for _, d := range r.FlattenDescriptions() {
			dl := strings.ToLower(d)
			if strings.Contains(dl, "single") {
				t = release.TypeSingle
				break
			}
			if strings.Contains(dl, "album") {
				t = release.TypeAlbum
				break
			}
			if strings.Contains(dl, "ep") {
				t = release.TypeEP
				break
			}
			if strings.Contains(dl, "compilation") {
				t = release.TypeCompilation
				break
			}
		}
		if t == release.TypeUnknown {
			slog.Debug("No release type found", "type", r.FlattenDescriptions(), "id", r.ID)
		}

		aIds := []int64{}
		for _, ar := range r.Artists {
			aIds = append(aIds, ar.ID)
		}
		as, err := db.Artist.Query().
			Select(artist.FieldID).
			Where(artist.DidIn(aIds...)).
			All(ctx)
		if err != nil {
			return err
		}

		releaseID := pid.New()
		for i, a := range as {
			buildersRA = append(buildersRA, db.ReleaseAppearance.Create().
				SetArtistID(a.ID).
				SetReleaseID(releaseID).
				SetOrder(i),
			)
		}

		for i, t := range r.Tracklist {
			tId := pid.New()
			buildersTrack = append(buildersTrack, db.Track.Create().
				SetTitle(t.Title).
				SetID(tId),
			)
		}

		builders = append(builders, db.Release.Create().
			SetID(releaseID).
			SetName(r.Title).
			SetType(t).
			SetReleaseDate(parseDate(r.Released)),
		)

	}
	err := db.Release.CreateBulk(builders...).Exec(ctx)
	if err != nil {
		return err
	}

	err = db.ReleaseAppearance.CreateBulk(buildersRA...).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func GetCmd() *cobra.Command {
	importCmd := &cobra.Command{
		Use:   "import",
		Short: "data import",
	}

	ImportArtistsCmd := &cobra.Command{
		Use:   "artists",
		Short: "imports all artists",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.NewDatabase("file:./data/database.db?_fk=1&_journal_mode=WAL")
			if err != nil {
				return err
			}
			c, err := db.Client.Artist.Delete().Where(artist.DidNotNil()).Exec(cmd.Context())
			if err != nil {
				return err
			}
			slog.Info("removed old record", "count", c)
			return ImportEntities(
				args[0], "artist", db.Client, 500,
				decodeArtist, insertArtists,
				cmd.Context(),
			)
		},
	}
	importCmd.AddCommand(ImportArtistsCmd)

	ImportReleaseCmd := &cobra.Command{
		Use:   "releases",
		Short: "imports all releases",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			db, err := database.NewDatabase("file:./data/database.db?_fk=1&_journal_mode=WAL")
			if err != nil {
				return err
			}
			return ImportEntities(
				args[0], "release", db.Client, 500,
				decodeRelease, insertRelease,
				cmd.Context(),
			)
		},
	}
	importCmd.AddCommand(ImportReleaseCmd)

	return importCmd
}

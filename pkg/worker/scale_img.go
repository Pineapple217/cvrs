package worker

import (
	"context"
	"encoding/json"
	"image"
	"log/slog"
	"path/filepath"

	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

var sizes = []int{64, 265, 1024}

func ScaleImg(t *ent.Task, db *database.Database, ctx context.Context) error {
	var ti database.TaskScaleImg
	err := json.Unmarshal(t.Payload, &ti)
	if err != nil {
		return err
	}

	i, err := db.Client.Image.Get(ctx, ti.ImageId)
	if err != nil {
		return err
	}
	img, err := imgio.Open(filepath.Join(database.IMG_DIR, i.File))
	if err != nil {
		return err
	}

	if i.DimentionHeight != i.DimentionWidth {
		img = CropToSquare(img)
	}

	imgs := []image.Image{}
	for _, size := range sizes {
		smallImg := transform.Resize(img, size, size, transform.Lanczos)
		imgs = append(imgs, smallImg)
	}
	dbImgs, err := db.SaveProcedImgs(ctx, i.ID, imgs)
	if err != nil {
		return err
	}
	_ = dbImgs

	slog.Info("done scaling img", "img", ti.ImageId)
	return nil
}

func CropToSquare(img image.Image) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var size int
	if width < height {
		size = width
	} else {
		size = height
	}

	startX := (width - size) / 2
	startY := (height - size) / 2

	return transform.Crop(img, image.Rect(startX, startY, startX+size, startY+size))
}

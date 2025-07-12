package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"os"
	"path"
	"slices"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	"github.com/Pineapple217/cvrs/pkg/ent"
	entImage "github.com/Pineapple217/cvrs/pkg/ent/image"
	"github.com/Pineapple217/cvrs/pkg/ent/processedimage"
	"github.com/Pineapple217/cvrs/pkg/ent/task"
	"github.com/Pineapple217/cvrs/pkg/pid"
	"github.com/chai2010/webp"
)

var AllowedMIME = []string{"image/png", "image/jpeg", "image/webp"}

const MAX_IMG_SIZE = 1024 * 1024 * 5 // MB
const TEMP_DIR = "tmp"
const IMG_DIR = "img"

func (d Database) SaveImg(ctx context.Context, f *multipart.FileHeader, uploader pid.ID) (*ent.Image, error) {
	var err error
	// Check file type
	mimeType := f.Header.Get("Content-Type")
	if !slices.Contains(AllowedMIME, mimeType) {
		return nil, fmt.Errorf("%s is not an allowed img format", mimeType)
	}

	// Check file size
	if f.Size > MAX_IMG_SIZE {
		return nil, fmt.Errorf("img is to large: %d", f.Size)
	}

	// Write img to temp file
	tempFile, err := os.CreateTemp(path.Join(d.Conf.DataLocation, TEMP_DIR), "img*")
	if err != nil {
		return nil, fmt.Errorf("failed to create tmp file, %s", err)
	}
	defer tempFile.Close()
	sourceFile, err := f.Open()
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(tempFile, sourceFile)
	if err != nil {
		return nil, err
	}
	defer sourceFile.Close()

	// Get img info
	_, err = tempFile.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}
	img, format, err := image.Decode(tempFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode img: %s", err)
	}
	imgType, err := ParseImageType(format)
	if err != nil {
		return nil, err
	}
	bounds := img.Bounds()

	// Write to DB
	stats, err := tempFile.Stat()
	if err != nil {
		return nil, err
	}
	err = tempFile.Close()
	if err != nil {
		return nil, err
	}

	id := pid.New()
	taskData, err := json.Marshal(TaskScaleImg{
		ImageId: id,
	})
	if err != nil {
		return nil, err
	}

	// Start transaction ================================
	tx, err := d.Client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	DBimg, err := tx.Image.Create().
		SetType(imgType).
		SetDimentionWidth(bounds.Dx()).
		SetDimentionHeight(bounds.Dy()).
		SetOriginalName(f.Filename).
		SetSizeBits(uint32(stats.Size())).
		SetFile(id.String()).
		SetUploaderID(uploader).
		SetID(id).
		Save(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}
	err = os.Rename(tempFile.Name(), path.Join(d.Conf.DataLocation, IMG_DIR, id.String()))
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	_, err = tx.Task.Create().
		SetType(task.TypeScaleImg).
		SetPayload(taskData).
		Save(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	// End transaction ===================================

	// TODO: Clean up temp file that are stale
	return DBimg.Unwrap(), nil
}

func (d Database) SaveProcedImgs(ctx context.Context, source pid.ID, imgs []image.Image) ([]*ent.ProcessedImage, error) {
	temps := []*os.File{}
	for _, i := range imgs {
		tempFile, err := os.CreateTemp(path.Join(d.Conf.DataLocation, TEMP_DIR), "proc_img*")
		if err != nil {
			return nil, fmt.Errorf("failed to create tmp file, %s", err)
		}
		defer tempFile.Close()
		temps = append(temps, tempFile)

		err = webp.Encode(tempFile, i, &webp.Options{
			Quality: 90,
		})
		if err != nil {
			return nil, err
		}
	}

	// Start transaction ================================
	tx, err := d.Client.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	imgCreates := []*ent.ProcessedImageCreate{}
	for i, temp := range temps {
		id := pid.New()
		info, err := temp.Stat()
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("failed to get temp file stats: %s", err)
		}
		imgCreate := tx.ProcessedImage.Create().
			SetDimentions(imgs[i].Bounds().Dx()).
			SetID(id).
			SetSizeBits(uint32(info.Size())).
			SetSourceID(source).
			SetType(processedimage.TypeWEBP)
		imgCreates = append(imgCreates, imgCreate)
		err = temp.Close()
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("failed to close temp file: %s", err)
		}
	}

	dbImgs, err := tx.ProcessedImage.CreateBulk(imgCreates...).Save(ctx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: %v", err, rerr)
		}
		return nil, err
	}

	for i, temp := range temps {
		err = os.Rename(temp.Name(), path.Join(d.Conf.DataLocation, IMG_DIR, dbImgs[i].ID.String()))
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: %v", err, rerr)
			}
			return nil, fmt.Errorf("failed to move file: %s", err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	// End transaction ===================================

	for _, dbImg := range dbImgs {
		dbImg.Unwrap()
	}
	return dbImgs, nil
}

func ParseImageType(s string) (entImage.Type, error) {
	switch strings.ToUpper(s) {
	case "WEBP":
		return entImage.TypeWEBP, nil
	case "PNG":
		return entImage.TypePNG, nil
	case "JPG", "JPEG":
		return entImage.TypeJPG, nil
	default:
		return "", fmt.Errorf("invalid ImageType: %q", s)
	}
}

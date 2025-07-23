package database

import (
	"fmt"
	"strings"

	entImage "github.com/Pineapple217/cvrs/pkg/ent/image"
	"github.com/Pineapple217/cvrs/pkg/ent/release"
)

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

func ParseReleaseType(s string) (release.Type, error) {
	switch strings.ToLower(s) {
	case "album":
		return release.TypeAlbum, nil
	case "compilation":
		return release.TypeCompilation, nil
	case "ep":
		return release.TypeEP, nil
	case "single":
		return release.TypeSingle, nil
	case "unknown":
		return release.TypeUnknown, nil
	default:
		return "", fmt.Errorf("invalid ReleaseType: %q", s)
	}
}

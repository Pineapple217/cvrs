package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/Pineapple217/cvrs/pkg/pid"
	"github.com/Pineapple217/cvrs/pkg/users"
	"github.com/labstack/echo/v4"
)

type ReleaseAddRequest struct {
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	ReleaseDate time.Time `json:"releaseDate"`
	Artists     []pid.ID  `json:"artists"`
}

func (h *Handler) ReleaseAdd(c echo.Context) error {
	f, err := c.MultipartForm()
	if err != nil {
		return err
	}
	var data ReleaseAddRequest
	_, ok := f.Value["json"]
	if !ok || len(f.Value["json"]) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "no json found")
	}
	err = json.Unmarshal([]byte(f.Value["json"][0]), &data)
	if err != nil {
		return err
	}
	_, ok = f.File["img"]
	if !ok || len(f.File["img"]) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "to img found")
	}
	img := f.File["img"][0]

	t, err := database.ParseReleaseType(data.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "bad release type")
	}

	_, claims := users.IsAuth(c)
	DBimg, err := h.DB.SaveImg(c.Request().Context(), img, claims.UserId)
	if err != nil {
		return err
	}

	_, err = h.DB.Client.Release.Create().
		SetName(strings.TrimSpace(data.Name)).
		SetImage(DBimg).
		SetType(t).
		SetReleaseDate(data.ReleaseDate).
		Save(c.Request().Context())
	if err != nil {
		h.DB.HardDeleteImg(c.Request().Context(), DBimg.ID)
		return err
	}

	return nil
}

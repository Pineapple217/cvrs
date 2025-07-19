package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/Pineapple217/cvrs/pkg/ent/artist"
	"github.com/Pineapple217/cvrs/pkg/pid"
	"github.com/Pineapple217/cvrs/pkg/users"
	"github.com/labstack/echo/v4"
)

type ArtistsAddRequest struct {
	Name string `json:"name"`
}

func (h *Handler) ArtistsAdd(c echo.Context) error {
	f, err := c.MultipartForm()
	if err != nil {
		return err
	}
	var data ArtistsAddRequest
	_, ok := f.Value["json"]
	if !ok || len(f.Value["json"]) == 0 {
		return echo.ErrBadRequest
	}
	err = json.Unmarshal([]byte(f.Value["json"][0]), &data)
	if err != nil {
		return err
	}
	_, ok = f.File["img"]
	if !ok || len(f.File["img"]) == 0 {
		return echo.ErrBadRequest
	}
	img := f.File["img"][0]

	_, claims := users.IsAuth(c)
	DBimg, err := h.DB.SaveImg(c.Request().Context(), img, claims.UserId)
	if err != nil {
		return err
	}

	_, err = h.DB.Client.Artist.Create().
		SetName(strings.TrimSpace(data.Name)).
		SetImage(DBimg).
		Save(c.Request().Context())
	if err != nil {
		h.DB.HardDeleteImg(c.Request().Context(), DBimg.ID)
		return err
	}
	return c.NoContent(http.StatusOK)
}

type ArtistsPage struct {
	Limit  int           `json:"limit"`
	Offset int           `json:"offset"`
	Artist []*ent.Artist `josn:"artists"`
}

func (h *Handler) ArtistsGet(c echo.Context) error {
	var err error
	var offset, limit int
	offsetStr := c.QueryParam("offset")
	limitStr := c.QueryParam("limit")

	offset, err = strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}
	if limit > 200 {
		limit = 200
	}

	as, err := h.DB.Client.Artist.Query().
		Order(ent.Desc(artist.FieldUpdatedAt)).
		Offset(offset).
		Limit(limit).
		WithImage(func(iq *ent.ImageQuery) {
			iq.Select("id").WithProccesedImage()
		}).
		All(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ArtistsPage{
		Limit:  limit,
		Offset: offset,
		Artist: as,
	})
}

func (h *Handler) ArtistGetId(c echo.Context) error {
	idStr := c.Param("id")
	id, err := pid.DecodeBase32(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ID format provided")
	}

	a, err := h.DB.Client.Artist.Query().
		Where(artist.IDEQ(id)).
		WithImage(func(iq *ent.ImageQuery) {
			iq.WithProccesedImage()
		}).
		Only(c.Request().Context())
	if ent.IsNotFound(err) {
		return echo.NewHTTPError(http.StatusNotFound)
	}
	if ent.IsNotSingular(err) {
		slog.Warn("not singular", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "not singular")
	}

	return c.JSON(http.StatusOK, a)
}

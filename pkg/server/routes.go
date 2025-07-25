package server

import (
	"log/slog"
	"net/http"

	"github.com/Pineapple217/cvrs/pkg/handler"
	"github.com/Pineapple217/cvrs/pkg/static"
	"github.com/Pineapple217/cvrs/pkg/users"
	"github.com/labstack/echo/v4"
)

func (server *Server) RegisterRoutes(hdlr *handler.Handler, imgDir string) {
	slog.Info("Registering routes")
	e := server.e

	// backend
	api := e.Group("/api")

	img := api.Group("/i")
	img.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Add("Cache-Control", "public, max-age=31536000, immutable")
			return next(c)
		}
	})
	img.Static("/", imgDir)

	api.POST("/auth/login", hdlr.Login)
	api.GET("/auth/users", users.CheckAdmin(hdlr.Users))

	api.POST("/artists/add", users.CheckAuth(hdlr.ArtistsAdd))
	api.GET("/artist/:id", hdlr.ArtistGetId)
	api.GET("/artists", hdlr.ArtistsGet)

	api.POST("/releases/add", users.CheckAdmin(hdlr.ReleaseAdd))

	// frontend
	frontend := static.GetFrontend()
	e.StaticFS("", frontend)
	e.GET("/", func(c echo.Context) error {
		file, err := frontend.Open("index.html")
		if err != nil {
			panic("index.html not found, likely due to bad build")
		}
		stat, _ := file.Stat()
		c.Response().Header().Set("Cache-Control", "no-cache, max-age=0")
		return c.Stream(http.StatusOK, "text/html", http.MaxBytesReader(c.Response().Writer, file, stat.Size()))
	})
}

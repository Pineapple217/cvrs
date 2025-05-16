package server

import (
	"log/slog"
	"net/http"

	"github.com/Pineapple217/cvrs/pkg/handler"
	"github.com/Pineapple217/cvrs/pkg/static"
	"github.com/Pineapple217/cvrs/pkg/users"
	"github.com/labstack/echo/v4"
)

func (server *Server) RegisterRoutes(hdlr *handler.Handler) {
	slog.Info("Registering routes")
	e := server.e

	// backend
	api := e.Group("api")
	api.POST("/auth/login", hdlr.Login)
	api.GET("/auth/users", users.CheckAdmin(hdlr.Users))

	// Temp: del me
	api.GET("/admin", users.CheckAdmin(func(c echo.Context) error { return c.String(200, "admin cool") }))
	api.GET("/a", users.CheckAuth(func(c echo.Context) error { return c.String(200, "auth cool") }))

	// frontend
	frontend := static.GetFrontend()
	e.StaticFS("", frontend)
	e.GET("/", func(c echo.Context) error {
		file, err := frontend.Open("index.html")
		if err != nil {
			panic("index.html not found, likely do to bad build")
		}
		stat, _ := file.Stat()
		c.Response().Header().Set("Cache-Control", "no-cache, max-age=0")
		return c.Stream(http.StatusOK, "text/html", http.MaxBytesReader(c.Response().Writer, file, stat.Size()))
	})
}

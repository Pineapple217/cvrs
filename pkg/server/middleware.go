package server

import (
	"log/slog"
	"strings"
	"time"

	"github.com/Pineapple217/cvrs/pkg/users"
	"github.com/labstack/echo/v4"
	echoMw "github.com/labstack/echo/v4/middleware"
)

func (s *Server) ApplyMiddleware(dev bool) {
	slog.Info("Applying middlewares")
	s.e.Use(echoMw.RateLimiterWithConfig(echoMw.RateLimiterConfig{
		Store: echoMw.NewRateLimiterMemoryStoreWithConfig(
			echoMw.RateLimiterMemoryStoreConfig{Rate: 30, Burst: 60, ExpiresIn: 3 * time.Minute},
		),
	}))
	s.e.Use(echoMw.RequestLoggerWithConfig(echoMw.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v echoMw.RequestLoggerValues) error {
			slog.Info("request",
				"method", v.Method,
				"status", v.Status,
				"latency", v.Latency,
				"path", v.URI,
			)
			return nil

		},
	}))

	if dev {
		s.e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
			AllowOrigins: []string{"http://localhost:5173"},
		}))
	}

	s.e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqPath := c.Request().URL.Path

			if strings.HasPrefix(reqPath, "/assets/") {
				c.Response().Header().Set("Cache-Control", "public, max-age=31536000, immutable")
			}

			return next(c)
		}
	})
	s.e.Use(users.Auth([]byte("adsjfkaweijrfsdjfkla")))

	s.e.Use(echoMw.GzipWithConfig(echoMw.GzipConfig{
		Level: 5,
	}))
}

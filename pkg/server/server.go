package server

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	listen = "0.0.0.0"
	// listen = "127.0.0.1"
	port = ":3000"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Debug = true
	NewServer := &Server{
		e: e,
	}

	return NewServer
}

// Starts the server in a new routine
func (s *Server) Start() {
	flag.Parse()
	slog.Info("Starting server")
	go func() {
		if err := s.e.Start(listen + port); err != nil && err != http.ErrServerClosed {
			slog.Error("Shutting down the server", "error", err.Error())
		}
	}()
	slog.Info("Server started", "bind", listen, "port", port)
}

// Tries to the stops the server gracefully
func (s *Server) Stop() {
	slog.Info("Stopping server")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.e.Shutdown(ctx); err != nil {
		slog.Error(err.Error())
	}
}

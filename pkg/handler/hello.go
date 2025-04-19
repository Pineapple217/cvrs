package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Hello(c echo.Context) error {
	return c.JSON(200, "heyhey")
}

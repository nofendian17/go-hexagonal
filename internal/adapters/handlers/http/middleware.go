package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterAppMiddleware(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
}

package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	internalMiddleware "user-svc/internal/middleware"
	"user-svc/internal/shared/logger"
)

func RegisterAppMiddleware(e *echo.Echo, logger *logger.LogWrapper) {
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	loggingMiddleware := internalMiddleware.NewLoggingMiddleware(logger)
	e.Use(loggingMiddleware.LogRequestAndResponse)
}

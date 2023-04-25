package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
	"user-svc/internal/adapters/repository/postgres"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
	"user-svc/internal/shared/config"
	appError "user-svc/internal/shared/error"
	validatorHelper "user-svc/internal/shared/validator"
)

const shutdownTimeout = 10 * time.Second

func Start() {
	e := echo.New()
	cfg := config.New()
	// hex can switch storage
	// with implement interface
	repo := postgres.NewRepository(cfg)
	userService := services.NewUserService(repo)
	roleService := services.NewRoleService(repo)
	RegisterRoutes(e, *userService, *roleService)
	RegisterMiddleware(e)
	e.Debug = true
	e.Validator = &validatorHelper.CustomValidator{
		Validator: validator.New(),
	}
	e.HTTPErrorHandler = errorHandler
	startServer(e, cfg.App.Port)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	shutdownServer(e)
}

func errorHandler(err error, c echo.Context) {
	var appErr *appError.AppError
	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, domain.Response{
			Code:    appErr.Code,
			Message: appErr.Message,
		})
		return
	}

	var report *echo.HTTPError
	if errors.As(err, &report) {
		switch report.Code {
		case http.StatusNotFound:
			c.JSON(http.StatusNotFound, domain.Response{
				Code:    http.StatusNotFound,
				Message: "Resource not found",
			})
		case http.StatusInternalServerError:
			c.JSON(http.StatusInternalServerError, domain.Response{
				Code:    http.StatusInternalServerError,
				Message: "Internal server error",
			})
		default:
			c.JSON(report.Code, domain.Response{
				Code:    report.Code,
				Message: report.Message.(string),
			})
		}
		return
	}

	switch err.(type) {
	case validator.ValidationErrors:
		var messages []string
		for _, v := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s: invalid value '%v'", v.Field(), v.Value())
			messages = append(messages, message)
		}
		c.JSON(http.StatusBadRequest, domain.Response{
			Code:    http.StatusBadRequest,
			Message: strings.Join(messages, ", "),
		})
	default:
		c.JSON(http.StatusInternalServerError, domain.Response{
			Code:    http.StatusInternalServerError,
			Message: "Internal server error",
		})
	}
}

func startServer(e *echo.Echo, port int) {
	go func() {
		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()
}

func shutdownServer(e *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

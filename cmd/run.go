package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	"user-svc/internal/core/domain"
	appError "user-svc/internal/shared/error"
	validatorHelper "user-svc/internal/shared/validator"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"user-svc/internal/adapters/handler"
	"user-svc/internal/adapters/repository/postgres"
	"user-svc/internal/core/services"
	"user-svc/internal/shared/config"
)

const shutdownTimeout = 10 * time.Second

func Run() {
	cfg := config.New()
	repo := postgres.NewRepository(cfg)
	userService := services.NewUserService(repo)
	roleService := services.NewRoleService(repo)
	InitRoutes(cfg.App.Port, userService, roleService)
}

func InitRoutes(port int32, userService *services.UserService, roleService *services.RoleService) {
	e := echo.New()
	e.Validator = &validatorHelper.CustomValidator{
		Validator: validator.New(),
	}

	h := handler.NewHttpHandler(*userService, *roleService)

	v1 := e.Group("/api/v1")
	// user route
	v1.POST("/user", h.CreateUser)
	v1.PUT("/user", h.UpdateUser)
	v1.DELETE("/user/:id", h.DeleteUser)
	v1.GET("/user/:id", h.User)
	v1.GET("/users", h.Users)

	// role route
	v1.POST("/role", h.CreateRole)
	v1.PUT("/role", h.UpdateRole)
	v1.DELETE("/role/:id", h.DeleteRole)
	v1.GET("/role/:id", h.Role)
	v1.GET("/roles", h.Roles)

	e.HTTPErrorHandler = errorHandler

	startServer(e, port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	shutdownServer(e)
}

func errorHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	if appErr, ok := err.(*appError.AppError); ok {
		report.Code = appErr.Code
		c.Logger().Error(appErr)
		c.JSON(appErr.Code, domain.Response{
			Code:    appErr.Code,
			Message: appErr.Message,
		})
		return
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not a valid email", err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s value must be greater than %s", err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s value must be lower than %s", err.Field(), err.Param())
			case "uuid":
				report.Message = fmt.Sprintf("%s is not a valid uuid", err.Field())
			}

			break
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, domain.Response{
		Code:    report.Code,
		Message: report.Message.(string),
	})
	return
}

func startServer(e *echo.Echo, port int32) {
	go func() {
		if err := e.Start(":" + strconv.Itoa(int(port))); err != nil && err != http.ErrServerClosed {
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

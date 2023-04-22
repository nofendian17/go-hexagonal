package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
)

type HttpHandler struct {
	userService services.UserService
}

func NewHttpHandler(userService services.UserService) *HttpHandler {
	return &HttpHandler{
		userService: userService,
	}
}

func (h *HttpHandler) CreateUser(c echo.Context) error {
	var user domain.CreateUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.Create(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *HttpHandler) UpdateUser(c echo.Context) error {
	var user domain.UpdateUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.Update(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) DeleteUser(c echo.Context) error {
	var user domain.DeleteUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.Delete(user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) Users(c echo.Context) error {
	result, err := h.userService.Users()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) User(c echo.Context) error {
	var user domain.GetUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.User(user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

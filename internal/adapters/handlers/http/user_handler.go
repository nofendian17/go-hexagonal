package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var user domain.CreateUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.CreateUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var user domain.UpdateUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.UpdateUser(&user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	var user domain.DeleteUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.DeleteUser(user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *UserHandler) Users(c echo.Context) error {
	result, err := h.userService.GetUsers()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *UserHandler) User(c echo.Context) error {
	var user domain.GetUserRequest
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := c.Validate(&user); err != nil {
		return err
	}

	result, err := h.userService.GetUser(user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

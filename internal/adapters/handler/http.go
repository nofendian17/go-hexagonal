package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
)

type HttpHandler struct {
	userService services.UserService
	roleService services.RoleService
}

func NewHttpHandler(userService services.UserService, roleService services.RoleService) *HttpHandler {
	return &HttpHandler{
		userService: userService,
		roleService: roleService,
	}
}

// User

func (h *HttpHandler) CreateUser(c echo.Context) error {
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

func (h *HttpHandler) UpdateUser(c echo.Context) error {
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

func (h *HttpHandler) DeleteUser(c echo.Context) error {
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

func (h *HttpHandler) Users(c echo.Context) error {
	result, err := h.userService.GetUsers()
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

	result, err := h.userService.GetUser(user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

// Role

func (h *HttpHandler) CreateRole(c echo.Context) error {
	var role domain.CreateRoleRequest
	if err := c.Bind(&role); err != nil {
		return err
	}

	if err := c.Validate(&role); err != nil {
		return err
	}
	result, err := h.roleService.CreateRole(&role)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, result)
}

func (h *HttpHandler) UpdateRole(c echo.Context) error {
	var role domain.UpdateRoleRequest
	if err := c.Bind(&role); err != nil {
		return err
	}

	if err := c.Validate(&role); err != nil {
		return err
	}
	result, err := h.roleService.UpdateRole(&role)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, result)
}

func (h *HttpHandler) DeleteRole(c echo.Context) error {
	var role domain.DeleteRoleRequest
	if err := c.Bind(&role); err != nil {
		return err
	}

	if err := c.Validate(&role); err != nil {
		return err
	}

	result, err := h.roleService.DeleteRole(role.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) Roles(c echo.Context) error {
	result, err := h.roleService.GetRoles()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *HttpHandler) Role(c echo.Context) error {
	var role domain.GetRoleRequest
	if err := c.Bind(&role); err != nil {
		return err
	}

	if err := c.Validate(&role); err != nil {
		return err
	}

	result, err := h.roleService.GetRole(role.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

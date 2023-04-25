package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
)

type RoleHandler struct {
	roleService services.RoleService
}

func NewRoleHandler(roleService services.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) CreateRole(c echo.Context) error {
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

func (h *RoleHandler) UpdateRole(c echo.Context) error {
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

func (h *RoleHandler) DeleteRole(c echo.Context) error {
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

func (h *RoleHandler) Roles(c echo.Context) error {
	result, err := h.roleService.GetRoles()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *RoleHandler) Role(c echo.Context) error {
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

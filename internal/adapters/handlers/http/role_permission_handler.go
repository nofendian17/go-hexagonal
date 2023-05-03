package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
)

type RolePermissionHandler struct {
	rolePermissionService services.RolePermissionService
}

func NewRolePermissionHandler(rolePermissionService services.RolePermissionService) *RolePermissionHandler {
	return &RolePermissionHandler{
		rolePermissionService: rolePermissionService,
	}
}

func (h *RolePermissionHandler) AssignPermissionsToRole(c echo.Context) error {
	var rolePermission domain.AssignPermissionToRoleRequest
	if err := c.Bind(&rolePermission); err != nil {
		return err
	}

	if err := c.Validate(&rolePermission); err != nil {
		return err
	}

	result, err := h.rolePermissionService.AssignPermissionsToRole(&rolePermission)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *RolePermissionHandler) GetRolePermissions(c echo.Context) error {
	var rolePermission domain.GetRolePermissionRequest
	if err := c.Bind(&rolePermission); err != nil {
		return err
	}

	if err := c.Validate(&rolePermission); err != nil {
		return err
	}

	result, err := h.rolePermissionService.GetRolePermissions(&rolePermission)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *RolePermissionHandler) RemovePermissionsFromRole(c echo.Context) error {
	var rolePermission domain.RemovePermissionFromRoleRequest
	if err := c.Bind(&rolePermission); err != nil {
		return err
	}

	if err := c.Validate(&rolePermission); err != nil {
		return err
	}

	result, err := h.rolePermissionService.RemovePermissionsFromRole(&rolePermission)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

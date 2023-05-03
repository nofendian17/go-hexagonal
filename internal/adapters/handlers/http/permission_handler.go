package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
)

type PermissionHandler struct {
	permissionService services.PermissionService
}

func NewPermissionHandler(permissionService services.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
}

func (h *PermissionHandler) CreatePermission(c echo.Context) error {
	var permission domain.CreatePermissionRequest
	if err := c.Bind(&permission); err != nil {
		return err
	}

	if err := c.Validate(&permission); err != nil {
		return err
	}
	result, err := h.permissionService.CreatePermission(&permission)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, result)
}

func (h *PermissionHandler) UpdatePermission(c echo.Context) error {
	var permission domain.UpdatePermissionRequest
	if err := c.Bind(&permission); err != nil {
		return err
	}

	if err := c.Validate(&permission); err != nil {
		return err
	}
	result, err := h.permissionService.UpdatePermission(&permission)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, result)
}

func (h *PermissionHandler) DeletePermission(c echo.Context) error {
	var permission domain.DeletePermissionRequest
	if err := c.Bind(&permission); err != nil {
		return err
	}

	if err := c.Validate(&permission); err != nil {
		return err
	}

	result, err := h.permissionService.DeletePermission(permission.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *PermissionHandler) Permissions(c echo.Context) error {
	result, err := h.permissionService.GetPermissions()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *PermissionHandler) Permission(c echo.Context) error {
	var permission domain.GetPermissionRequest
	if err := c.Bind(&permission); err != nil {
		return err
	}

	if err := c.Validate(&permission); err != nil {
		return err
	}

	result, err := h.permissionService.GetPermission(permission.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

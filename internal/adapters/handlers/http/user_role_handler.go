package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
)

type UserRoleHandler struct {
	userRoleService services.UserRoleService
}

func NewUserRoleHandler(userRoleService services.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{
		userRoleService: userRoleService,
	}
}

func (h *UserRoleHandler) AssignRolesToUser(c echo.Context) error {
	var userRole domain.AssignRolesToUserRequest
	if err := c.Bind(&userRole); err != nil {
		return err
	}

	if err := c.Validate(&userRole); err != nil {
		return err
	}

	result, err := h.userRoleService.AssignRolesToUser(&userRole)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *UserRoleHandler) GetUserRoles(c echo.Context) error {
	var userRole domain.GetUserRolesRequest
	if err := c.Bind(&userRole); err != nil {
		return err
	}

	if err := c.Validate(&userRole); err != nil {
		return err
	}

	result, err := h.userRoleService.GetUserRoles(&userRole)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func (h *UserRoleHandler) RemoveRolesFromUser(c echo.Context) error {
	var userRole domain.RemoveRolesFromUserRequest
	if err := c.Bind(&userRole); err != nil {
		return err
	}

	if err := c.Validate(&userRole); err != nil {
		return err
	}

	result, err := h.userRoleService.RemoveRolesFromUser(&userRole)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

package http

import (
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
	"user-svc/internal/shared/constants"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Authenticate(c echo.Context) error {
	var auth domain.GetTokenRequest
	if err := c.Bind(&auth); err != nil {
		return err
	}

	if err := c.Validate(&auth); err != nil {
		return err
	}
	result, err := h.authService.Authenticate(&auth)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	var auth domain.RefreshTokenRequest
	if err := c.Bind(&auth); err != nil {
		return err
	}

	if err := c.Validate(&auth); err != nil {
		return err
	}
	result, err := h.authService.Refresh(&auth)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func (h *AuthHandler) Logout(c echo.Context) error {
	authID := c.Get(constants.KeyAuthID).(string)
	result, err := h.authService.Logout(authID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

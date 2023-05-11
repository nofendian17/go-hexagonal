package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/services"
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

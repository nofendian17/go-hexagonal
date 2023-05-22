package middleware

import (
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	"user-svc/internal/core/services"
	"user-svc/internal/shared/constants"
	appError "user-svc/internal/shared/error"

	"github.com/labstack/echo/v4"
)

type PermissionChecker interface {
	Check(c echo.Context, requiredPermission string) (bool, error)
}

type PermissionCheckerImpl struct {
	AuthRepository        ports.AuthRepository
	RolePermissionService services.RolePermissionService
}

func (p *PermissionCheckerImpl) Check(c echo.Context, requiredPermission string) (bool, error) {
	authID := c.Get(constants.KeyAuthID).(string)
	tokenInfo, err := p.AuthRepository.GetToken(authID)
	if err != nil {
		return false, err
	}
	for _, role := range tokenInfo.Roles {
		d := &domain.GetRolePermissionRequest{
			RoleId: role.Id,
		}
		permissions, err := p.RolePermissionService.GetRolePermissions(d)
		if err != nil {
			return false, err
		}
		for _, permission := range permissions.Data.([]*domain.Permission) {
			if permission.Name == requiredPermission {
				c.Set(constants.KeyUserID, tokenInfo.UserID)
				return true, nil
			}
		}
	}

	return false, nil
}

type PermissionMiddleware struct {
	Checker PermissionChecker
}

func (m *PermissionMiddleware) Handle(requiredPermission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if hasPermission, err := m.Checker.Check(c, requiredPermission); err != nil {
				return err
			} else if !hasPermission {
				return c.JSON(http.StatusForbidden, &appError.AppError{
					Code:    http.StatusForbidden,
					Message: "forbidden",
				})
			}
			return next(c)
		}
	}
}

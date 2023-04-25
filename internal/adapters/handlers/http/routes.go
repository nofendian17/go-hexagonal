package http

import (
	"github.com/labstack/echo/v4"
	"user-svc/internal/core/services"
)

func RegisterRoutes(e *echo.Echo, userService services.UserService, roleService services.RoleService) {
	// Create user handler
	userHandler := NewUserHandler(userService)

	// Create role handler
	roleHandler := NewRoleHandler(roleService)

	v1 := e.Group("/api/v1")

	// Register user endpoints
	v1.POST("/user", userHandler.CreateUser)
	v1.PUT("/user", userHandler.UpdateUser)
	v1.DELETE("/user/:id", userHandler.DeleteUser)
	v1.GET("/user/:id", userHandler.User)
	v1.GET("/users", userHandler.Users)

	// Register role endpoints
	v1.POST("/role", roleHandler.CreateRole)
	v1.PUT("/role", roleHandler.UpdateRole)
	v1.DELETE("/role/:id", roleHandler.DeleteRole)
	v1.GET("/role/:id", roleHandler.Role)
	v1.GET("/roles", roleHandler.Roles)
}

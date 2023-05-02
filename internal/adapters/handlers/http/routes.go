package http

import (
	"github.com/labstack/echo/v4"
	"user-svc/internal/core/services"
)

func RegisterRoutes(
	e *echo.Echo,
	userService services.UserService,
	roleService services.RoleService,
	permissionService services.PermissionService,
	userRoleService services.UserRoleService,
) {
	// Create user handler
	userHandler := NewUserHandler(userService)
	// Create user role handler
	userRoleHandler := NewUserRoleHandler(userRoleService)
	// Create role handler
	roleHandler := NewRoleHandler(roleService)
	// Create permission handler
	permissionHandler := NewPermissionHandler(permissionService)

	v1 := e.Group("/api/v1")

	// Register user endpoints
	v1.POST("/user", userHandler.CreateUser)
	v1.PUT("/user/:id", userHandler.UpdateUser)
	v1.DELETE("/user/:id", userHandler.DeleteUser)
	v1.GET("/user/:id", userHandler.User)
	v1.GET("/users", userHandler.Users)

	// Register user role endpoints
	v1.GET("/user/:user_id/roles", userRoleHandler.GetUserRoles)
	v1.POST("/user/:user_id/roles/assign", userRoleHandler.AssignRolesToUser)
	v1.DELETE("/user/:user_id/roles/revoke", userRoleHandler.RemoveRolesFromUser)

	// Register role endpoints
	v1.POST("/role", roleHandler.CreateRole)
	v1.PUT("/role/:id", roleHandler.UpdateRole)
	v1.DELETE("/role/:id", roleHandler.DeleteRole)
	v1.GET("/role/:id", roleHandler.Role)
	v1.GET("/roles", roleHandler.Roles)

	// Register permission endpoints
	v1.POST("/permission", permissionHandler.CreatePermission)
	v1.PUT("/permission/:id", permissionHandler.UpdatePermission)
	v1.DELETE("/permission/:id", permissionHandler.DeletePermission)
	v1.GET("/permission/:id", permissionHandler.Permission)
	v1.GET("/permissions", permissionHandler.Permissions)
}

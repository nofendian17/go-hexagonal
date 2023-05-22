package http

import (
	"github.com/labstack/echo/v4"
	"user-svc/internal/core/ports"
	"user-svc/internal/core/services"
	"user-svc/internal/middleware"
	"user-svc/internal/shared/config"
)

const (
	apiPrefix           = "/api/v1"
	usersPath           = "/users"
	rolesPath           = "/roles"
	permissionsPath     = "/permissions"
	userRolesPath       = "/user/:user_id/roles"
	rolePermissionsPath = "/role/:role_id/permissions"
)

func RegisterHTTPRoutes(
	e *echo.Echo,
	cfg *config.Config,
	authRepository ports.AuthRepository,
	userService services.UserService,
	roleService services.RoleService,
	permissionService services.PermissionService,
	userRoleService services.UserRoleService,
	rolePermissionService services.RolePermissionService,
	authService services.AuthService,
) {
	// Create user handler
	userHandler := NewUserHandler(userService)
	// Create user role handler
	userRoleHandler := NewUserRoleHandler(userRoleService)
	// Create role handler
	roleHandler := NewRoleHandler(roleService)
	// Create permission handler
	permissionHandler := NewPermissionHandler(permissionService)
	// Create role permission handler
	rolePermissionHandler := NewRolePermissionHandler(rolePermissionService)
	// Create auth handler
	authHandler := NewAuthHandler(authService)

	// Register JWT Middleware for routes
	authenticator := &middleware.JWTAuthenticatorImpl{
		SecretKey: []byte(cfg.App.Auth.AccessKey),
	}

	jwtMiddleware := &middleware.JWTMiddleware{
		Authenticator:  authenticator,
		AuthRepository: authRepository,
	}

	// create a new instance of the permission middleware
	checker := &middleware.PermissionCheckerImpl{
		AuthRepository:        authRepository,
		RolePermissionService: rolePermissionService,
	}
	permissionMiddleware := &middleware.PermissionMiddleware{
		Checker: checker,
	}

	v1 := e.Group(apiPrefix)

	// Register auth endpoint
	authGroup := v1.Group("/auth")
	authGroup.POST("/login", authHandler.Authenticate)
	authGroup.POST("/refresh", authHandler.Refresh)
	authGroup.DELETE("/logout", authHandler.Logout, jwtMiddleware.Handle)

	// Register user endpoints
	userGroup := v1.Group(usersPath, jwtMiddleware.Handle)
	userGroup.POST("", userHandler.CreateUser, permissionMiddleware.Handle("Create-User"))
	userGroup.PUT("/:id", userHandler.UpdateUser, permissionMiddleware.Handle("Update-User"))
	userGroup.DELETE("/:id", userHandler.DeleteUser, permissionMiddleware.Handle("Delete-User"))
	userGroup.GET("/:id", userHandler.User, permissionMiddleware.Handle("View-User"))
	userGroup.GET("", userHandler.Users, permissionMiddleware.Handle("List-User"))

	// Register user role endpoints
	userRoleGroup := v1.Group(userRolesPath, jwtMiddleware.Handle)
	userRoleGroup.GET("", userRoleHandler.GetUserRoles, permissionMiddleware.Handle("View-Role"))
	userRoleGroup.POST("/assign", userRoleHandler.AssignRolesToUser, permissionMiddleware.Handle("Update-Role"))
	userRoleGroup.DELETE("/revoke", userRoleHandler.RemoveRolesFromUser, permissionMiddleware.Handle("Update-Role"))

	// Register role endpoints
	roleGroup := v1.Group(rolesPath, jwtMiddleware.Handle)
	roleGroup.POST("", roleHandler.CreateRole, permissionMiddleware.Handle("Create-Role"))
	roleGroup.PUT("/:id", roleHandler.UpdateRole, permissionMiddleware.Handle("Update-Role"))
	roleGroup.DELETE("/:id", roleHandler.DeleteRole, permissionMiddleware.Handle("Delete-Role"))
	roleGroup.GET("/:id", roleHandler.Role, permissionMiddleware.Handle("View-Role"))
	roleGroup.GET("", roleHandler.Roles, permissionMiddleware.Handle("List-Role"))

	// Register role permission endpoints
	rolePermissionGroup := v1.Group(rolePermissionsPath, jwtMiddleware.Handle)
	rolePermissionGroup.GET("", rolePermissionHandler.GetRolePermissions, permissionMiddleware.Handle("View-Permission"))
	rolePermissionGroup.POST("/assign", rolePermissionHandler.AssignPermissionsToRole, permissionMiddleware.Handle("Update-Permission"))
	rolePermissionGroup.DELETE("/revoke", rolePermissionHandler.RemovePermissionsFromRole, permissionMiddleware.Handle("Update-Permission"))

	// Register permission endpoints
	permissionGroup := v1.Group(permissionsPath, jwtMiddleware.Handle)
	permissionGroup.POST("", permissionHandler.CreatePermission, permissionMiddleware.Handle("Create-Permission"))
	permissionGroup.PUT("/:id", permissionHandler.UpdatePermission, permissionMiddleware.Handle("Update-Permission"))
	permissionGroup.DELETE("/:id", permissionHandler.DeletePermission, permissionMiddleware.Handle("Delete-Permission"))
	permissionGroup.GET("/:id", permissionHandler.Permission, permissionMiddleware.Handle("View-Permission"))
	permissionGroup.GET("", permissionHandler.Permissions, permissionMiddleware.Handle("List-Permission"))
}

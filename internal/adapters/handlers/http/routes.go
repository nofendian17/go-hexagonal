package http

import (
	"github.com/labstack/echo/v4"
	"user-svc/internal/core/ports"
	"user-svc/internal/core/services"
	"user-svc/internal/middleware"
	"user-svc/internal/shared/config"
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

	// define func permission example: route GET "/users" only can access by user has permission "list-user"
	// common defined in db
	// - list-{module-name}
	// - view-{module-name}
	// - create-{module-name}
	// - update-{module-name}
	// - delete-{module-name}

	getUsersPermissionFunc := func(next echo.HandlerFunc) echo.HandlerFunc {
		return permissionMiddleware.Handle(next, "list-user")
	}

	v1 := e.Group("/api/v1")

	// Register auth endpoint
	auth := v1.Group("/auth")
	auth.POST("/login", authHandler.Authenticate)
	auth.POST("/refresh", authHandler.Refresh)
	auth.DELETE("/logout", authHandler.Logout, jwtMiddleware.Handle)

	// Register user endpoints
	v1.POST("/user", userHandler.CreateUser, jwtMiddleware.Handle)
	v1.PUT("/user/:id", userHandler.UpdateUser, jwtMiddleware.Handle)
	v1.DELETE("/user/:id", userHandler.DeleteUser, jwtMiddleware.Handle)
	v1.GET("/user/:id", userHandler.User, jwtMiddleware.Handle)
	v1.GET("/users", userHandler.Users, jwtMiddleware.Handle, getUsersPermissionFunc)

	// Register user role endpoints
	v1.GET("/user/:user_id/roles", userRoleHandler.GetUserRoles, jwtMiddleware.Handle)
	v1.POST("/user/:user_id/roles/assign", userRoleHandler.AssignRolesToUser, jwtMiddleware.Handle)
	v1.DELETE("/user/:user_id/roles/revoke", userRoleHandler.RemoveRolesFromUser, jwtMiddleware.Handle)

	// Register role endpoints
	v1.POST("/role", roleHandler.CreateRole, jwtMiddleware.Handle)
	v1.PUT("/role/:id", roleHandler.UpdateRole, jwtMiddleware.Handle)
	v1.DELETE("/role/:id", roleHandler.DeleteRole, jwtMiddleware.Handle)
	v1.GET("/role/:id", roleHandler.Role, jwtMiddleware.Handle)
	v1.GET("/roles", roleHandler.Roles, jwtMiddleware.Handle)

	// Register role permission endpoints
	v1.GET("/role/:role_id/permissions", rolePermissionHandler.GetRolePermissions, jwtMiddleware.Handle)
	v1.POST("/role/:role_id/permissions/assign", rolePermissionHandler.AssignPermissionsToRole, jwtMiddleware.Handle)
	v1.DELETE("/role/:role_id/permissions/revoke", rolePermissionHandler.RemovePermissionsFromRole, jwtMiddleware.Handle)

	// Register permission endpoints
	v1.POST("/permission", permissionHandler.CreatePermission, jwtMiddleware.Handle)
	v1.PUT("/permission/:id", permissionHandler.UpdatePermission, jwtMiddleware.Handle)
	v1.DELETE("/permission/:id", permissionHandler.DeletePermission, jwtMiddleware.Handle)
	v1.GET("/permission/:id", permissionHandler.Permission, jwtMiddleware.Handle)
	v1.GET("/permissions", permissionHandler.Permissions, jwtMiddleware.Handle)
}

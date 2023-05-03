package ports

import "user-svc/internal/core/domain"

type RolePermissionService interface {
	GetRolePermissions(request *domain.GetRolePermissionRequest) (*domain.Response, error)
	AssignPermissionsToRole(request *domain.AssignPermissionToRoleRequest) (*domain.Response, error)
	RemovePermissionsFromRole(request *domain.RemovePermissionFromRoleRequest) (*domain.Response, error)
}

type RolePermissionRepository interface {
	GetRolePermissions(roleId string) ([]*domain.Permission, error)
	AddRolePermissions(roleId string, permissions []string) error
	RemoveRolePermissions(roleId string, permissions []string) error
}

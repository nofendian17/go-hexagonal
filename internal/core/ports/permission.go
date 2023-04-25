package ports

import "user-svc/internal/core/domain"

type PermissionService interface {
	CreatePermission(request *domain.CreatePermissionRequest) (*domain.Response, error)
	UpdatePermission(request *domain.UpdatePermissionRequest) (*domain.Response, error)
	DeletePermission(id string) (*domain.Response, error)
	GetPermissions() (*domain.Response, error)
	GetPermission(id string) (*domain.Response, error)
}

type PermissionRepository interface {
	CreatePermission(role *domain.Permission) error
	UpdatePermission(role *domain.Permission) error
	DeletePermission(id string) error
	GetAllPermission() ([]*domain.Permission, error)
	GetPermissionByID(id string) (*domain.Permission, error)
	GetPermissionByName(name string) (*domain.Permission, error)
	PermissionIsExist(name string) (bool, error)
}

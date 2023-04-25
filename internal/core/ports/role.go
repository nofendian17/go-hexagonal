package ports

import "user-svc/internal/core/domain"

type RoleService interface {
	CreateRole(request *domain.CreateRoleRequest) (*domain.Response, error)
	UpdateRole(request *domain.UpdateRoleRequest) (*domain.Response, error)
	DeleteRole(id string) (*domain.Response, error)
	GetRoles() (*domain.Response, error)
	GetRole(id string) (*domain.Response, error)
}

type RoleRepository interface {
	CreateRole(role *domain.Role) error
	UpdateRole(role *domain.Role) error
	DeleteRole(id string) error
	GetAllRole() ([]*domain.Role, error)
	GetRoleByID(id string) (*domain.Role, error)
	GetRoleByName(name string) (*domain.Role, error)
	RoleIsExist(name string) (bool, error)
}

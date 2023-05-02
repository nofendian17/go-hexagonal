package ports

import "user-svc/internal/core/domain"

type UserRoleService interface {
	GetUserRoles(request *domain.GetUserRolesRequest) (*domain.Response, error)
	AssignRolesToUser(request *domain.AssignRolesToUserRequest) (*domain.Response, error)
	RemoveRolesFromUser(request *domain.RemoveRolesFromUserRequest) (*domain.Response, error)
}

type UserRoleRepository interface {
	GetUserRoles(userID string) ([]*domain.Role, error)
	AddUserRoles(userID string, roles []string) error
	RemoveUserRoles(userID string, roles []string) error
}

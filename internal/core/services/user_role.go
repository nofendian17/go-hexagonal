package services

import (
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	appError "user-svc/internal/shared/error"
)

type UserRoleService struct {
	userRoleRepository ports.UserRoleRepository
	userService        ports.UserService
	roleService        ports.RoleService
}

func NewUserRoleService(userRoleRepository ports.UserRoleRepository, userService ports.UserService, roleService ports.RoleService) *UserRoleService {
	return &UserRoleService{
		userRoleRepository: userRoleRepository,
		userService:        userService,
		roleService:        roleService,
	}
}

func (s *UserRoleService) GetUserRoles(request *domain.GetUserRolesRequest) (*domain.Response, error) {
	user, err := s.userService.GetUser(request.UserId)
	if err != nil && user == nil {
		return nil, err
	}

	result, err := s.userRoleRepository.GetUserRoles(request.UserId)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

func (s *UserRoleService) AssignRolesToUser(request *domain.AssignRolesToUserRequest) (*domain.Response, error) {
	user, err := s.userService.GetUser(request.UserId)
	if err != nil && user == nil {
		return nil, err
	}

	for _, roleID := range request.RolesId {
		role, err := s.roleService.GetRole(roleID)
		if err != nil && role == nil {
			return nil, err
		}
	}

	err = s.userRoleRepository.AddUserRoles(request.UserId, request.RolesId)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    nil,
	}, nil
}

func (s *UserRoleService) RemoveRolesFromUser(request *domain.RemoveRolesFromUserRequest) (*domain.Response, error) {
	user, err := s.userService.GetUser(request.UserId)
	if err != nil && user == nil {
		return nil, err
	}

	for _, roleID := range request.RolesId {
		role, err := s.roleService.GetRole(roleID)
		if err != nil && role == nil {
			return nil, err
		}
	}

	err = s.userRoleRepository.RemoveUserRoles(request.UserId, request.RolesId)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

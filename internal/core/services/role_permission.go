package services

import (
	"net/http"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	appError "user-svc/internal/shared/error"
)

type RolePermissionService struct {
	rolePermissionRepository ports.RolePermissionRepository
	roleService              ports.RoleService
	permissionService        ports.PermissionService
}

func NewRolePermissionService(rolePermissionRepository ports.RolePermissionRepository, roleService ports.RoleService, permissionService ports.PermissionService) *RolePermissionService {
	return &RolePermissionService{
		rolePermissionRepository: rolePermissionRepository,
		roleService:              roleService,
		permissionService:        permissionService,
	}
}

func (s *RolePermissionService) GetRolePermissions(request *domain.GetRolePermissionRequest) (*domain.Response, error) {
	role, err := s.roleService.GetRole(request.RoleId)
	if err != nil && role == nil {
		return nil, err
	}

	result, err := s.rolePermissionRepository.GetRolePermissions(request.RoleId)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

func (s *RolePermissionService) AssignPermissionsToRole(request *domain.AssignPermissionToRoleRequest) (*domain.Response, error) {
	role, err := s.roleService.GetRole(request.RoleId)
	if err != nil && role == nil {
		return nil, err
	}

	for _, permissionID := range request.PermissionsId {
		permission, err := s.permissionService.GetPermission(permissionID)
		if err != nil && permission == nil {
			return nil, err
		}
	}

	err = s.rolePermissionRepository.AddRolePermissions(request.RoleId, request.PermissionsId)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    nil,
	}, nil
}

func (s *RolePermissionService) RemovePermissionsFromRole(request *domain.RemovePermissionFromRoleRequest) (*domain.Response, error) {
	role, err := s.roleService.GetRole(request.RoleId)
	if err != nil && role == nil {
		return nil, err
	}

	for _, permissionID := range request.PermissionsId {
		role, err := s.roleService.GetRole(permissionID)
		if err != nil && role == nil {
			return nil, err
		}
	}

	err = s.rolePermissionRepository.RemoveRolePermissions(request.RoleId, request.PermissionsId)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

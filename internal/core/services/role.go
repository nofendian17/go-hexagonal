package services

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	appError "user-svc/internal/shared/error"
)

type RoleService struct {
	roleRepository ports.RoleRepository
}

func NewRoleService(roleRepository ports.RoleRepository) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (r *RoleService) CreateRole(request *domain.CreateRoleRequest) (*domain.Response, error) {
	if exist, err := r.roleRepository.RoleIsExist(request.Name); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	} else if exist {
		return nil, &appError.AppError{Code: http.StatusConflict, Message: fmt.Sprintf("role %s already exist", request.Name)}
	}

	role := &domain.Role{
		Id:        uuid.New().String(),
		Name:      request.Name,
		Active:    *request.Active,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.roleRepository.CreateRole(role); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    nil,
	}, nil
}

func (r *RoleService) UpdateRole(request *domain.UpdateRoleRequest) (*domain.Response, error) {
	role, err := r.roleRepository.GetRoleByID(request.Id)
	if err != nil && role == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("role with id %s not exist", request.Id)}
	}

	check, _ := r.roleRepository.GetRoleByName(request.Name)
	if check != nil && check.Id != role.Id {
		return nil, &appError.AppError{Code: http.StatusConflict, Message: fmt.Sprintf("role with name %s already exist", request.Name)}
	}

	role.Name = request.Name
	role.Active = *request.Active
	role.UpdatedAt = time.Now()

	err = r.roleRepository.UpdateRole(role)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (r *RoleService) DeleteRole(id string) (*domain.Response, error) {
	role, err := r.roleRepository.GetRoleByID(id)
	if err != nil && role == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("role with id %s not exist", id)}
	}

	err = r.roleRepository.DeleteRole(role.Id)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (r *RoleService) GetRoles() (*domain.Response, error) {
	result, err := r.roleRepository.GetAllRole()
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

func (r *RoleService) GetRole(id string) (*domain.Response, error) {
	result, err := r.roleRepository.GetRoleByID(id)
	if err != nil && result == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("role with id %s not exist", id)}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

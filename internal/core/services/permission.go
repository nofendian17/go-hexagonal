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

type PermissionService struct {
	permissionRepository ports.PermissionRepository
}

func NewPermissionService(permissionRepository ports.PermissionRepository) *PermissionService {
	return &PermissionService{
		permissionRepository: permissionRepository,
	}
}

func (r *PermissionService) CreatePermission(request *domain.CreatePermissionRequest) (*domain.Response, error) {
	if exist, err := r.permissionRepository.PermissionIsExist(request.Name); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	} else if exist {
		return nil, &appError.AppError{Code: http.StatusConflict, Message: fmt.Sprintf("permission %s already exist", request.Name)}
	}

	permission := &domain.Permission{
		Id:        uuid.New().String(),
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := r.permissionRepository.CreatePermission(permission); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    nil,
	}, nil
}

func (r *PermissionService) UpdatePermission(request *domain.UpdatePermissionRequest) (*domain.Response, error) {
	permission, err := r.permissionRepository.GetPermissionByID(request.Id)
	if err != nil && permission == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("permission with id %s not exist", request.Id)}
	}

	check, _ := r.permissionRepository.GetPermissionByName(request.Name)
	if check != nil && check.Id != permission.Id {
		return nil, &appError.AppError{Code: http.StatusConflict, Message: fmt.Sprintf("permission with name %s already exist", request.Name)}
	}

	permission.Name = request.Name
	permission.UpdatedAt = time.Now()

	err = r.permissionRepository.UpdatePermission(permission)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (r *PermissionService) DeletePermission(id string) (*domain.Response, error) {
	permission, err := r.permissionRepository.GetPermissionByID(id)
	if err != nil && permission == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("permission with id %s not exist", id)}
	}

	err = r.permissionRepository.DeletePermission(permission.Id)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (r *PermissionService) GetPermissions() (*domain.Response, error) {
	result, err := r.permissionRepository.GetAllPermission()
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

func (r *PermissionService) GetPermission(id string) (*domain.Response, error) {
	result, err := r.permissionRepository.GetPermissionByID(id)
	if err != nil && result == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("permission with id %s not exist", id)}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

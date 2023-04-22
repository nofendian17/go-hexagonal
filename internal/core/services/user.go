package services

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	appError "user-svc/internal/shared/error"
	"user-svc/internal/shared/hash"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) Create(request *domain.CreateUserRequest) (*domain.Response, error) {
	if exist, err := u.repo.Exist(request.Email); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	} else if exist {
		return nil, &appError.AppError{Code: http.StatusConflict, Message: fmt.Sprintf("request with %s already exist", request.Email)}
	}

	hashedPassword, salt, err := hash.HashPassword(request.Password)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	user := &domain.User{
		Id:        uuid.New().String(),
		Name:      request.Name,
		Email:     request.Email,
		Salt:      salt,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.repo.Create(user); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusCreated,
		Message: "created",
		Data:    nil,
	}, nil
}

func (u *UserService) Update(request *domain.UpdateUserRequest) (*domain.Response, error) {
	user, err := u.repo.UserByEmail(request.Email)
	if err != nil && user == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("user with %s not exist", request.Email)}
	}

	if request.Password != "" {
		hashedPassword, salt, err := hash.HashPassword(request.Password)
		if err != nil {
			return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		}
		user.Password = hashedPassword
		user.Salt = salt
	}

	user.Name = request.Name
	user.Active = *request.Active
	user.UpdatedAt = time.Now()

	err = u.repo.Update(user)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    nil,
	}, nil
}

func (u *UserService) Delete(id string) (*domain.Response, error) {
	user, err := u.repo.UserByID(id)
	if err != nil && user == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("user with id %s not exist", id)}
	}

	err = u.repo.Delete(user.Id)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    nil,
	}, nil
}

func (u *UserService) Users() (*domain.Response, error) {
	result, err := u.repo.Users()
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    result,
	}, nil
}

func (u *UserService) User(id string) (*domain.Response, error) {
	result, err := u.repo.UserByID(id)
	if err != nil && result == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("user with id %s not exist", id)}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    result,
	}, nil
}

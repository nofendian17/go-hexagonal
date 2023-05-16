package services

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"time"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	appError "user-svc/internal/shared/error"
	"user-svc/internal/shared/hash"
	"user-svc/internal/shared/logger"
)

type UserService struct {
	userRepository ports.UserRepository
	hasher         hash.Hasher
	logger         logger.Logger
}

func NewUserService(userRepository ports.UserRepository, hasher hash.Hasher, logger logger.Logger) *UserService {
	return &UserService{
		userRepository: userRepository,
		hasher:         hasher,
		logger:         logger,
	}
}

func (u *UserService) CreateUser(request *domain.CreateUserRequest) (*domain.Response, error) {
	if exist, err := u.userRepository.UserIsExist(request.Email); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	} else if exist {
		return nil, &appError.AppError{Code: http.StatusConflict, Message: fmt.Sprintf("user %s already exist", request.Email)}
	}

	salt, err := u.hasher.GenerateRandomSalt()
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	hashedPassword := u.hasher.HashPassword(request.Password, salt)

	user := &domain.User{
		Id:        uuid.New().String(),
		Name:      request.Name,
		Email:     request.Email,
		Salt:      base64.URLEncoding.EncodeToString(salt),
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := u.userRepository.CreateUser(user); err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
		Data:    nil,
	}, nil
}

func (u *UserService) UpdateUser(request *domain.UpdateUserRequest) (*domain.Response, error) {
	user, err := u.userRepository.GetUserByID(request.Id)
	if err != nil && user == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("user with id %s not exist", request.Id)}
	}

	check, _ := u.userRepository.GetUserByEmail(request.Email)
	if check != nil && check.Id != user.Id {
		return nil, &appError.AppError{Code: http.StatusConflict, Message: fmt.Sprintf("user with email %s already exist", request.Email)}
	}

	if request.Password != "" {
		salt, err := u.hasher.GenerateRandomSalt()
		if err != nil {
			return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		hashedPassword := u.hasher.HashPassword(request.Password, salt)
		user.Password = hashedPassword
		user.Salt = base64.URLEncoding.EncodeToString(salt)
	}

	user.Name = request.Name
	user.Email = request.Email
	user.Active = *request.Active
	user.UpdatedAt = time.Now()

	err = u.userRepository.UpdateUser(user)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (u *UserService) DeleteUser(id string) (*domain.Response, error) {
	user, err := u.userRepository.GetUserByID(id)
	if err != nil && user == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("user with id %s not exist", id)}
	}

	err = u.userRepository.DeleteUser(user.Id)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (u *UserService) GetUsers() (*domain.Response, error) {
	result, err := u.userRepository.GetAllUsers()
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

func (u *UserService) GetUser(id string) (*domain.Response, error) {
	result, err := u.userRepository.GetUserByID(id)
	if err != nil && result == nil {
		return nil, &appError.AppError{Code: http.StatusNotFound, Message: fmt.Sprintf("user with id %s not exist", id)}
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    result,
	}, nil
}

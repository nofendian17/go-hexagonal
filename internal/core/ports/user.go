package ports

import "user-svc/internal/core/domain"

type UserService interface {
	CreateUser(request *domain.CreateUserRequest) (*domain.Response, error)
	UpdateUser(request *domain.UpdateUserRequest) (*domain.Response, error)
	DeleteUser(id string) (*domain.Response, error)
	GetUsers() (*domain.Response, error)
	GetUser(id string) (*domain.Response, error)
}

type UserRepository interface {
	CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
	GetAllUsers() ([]*domain.User, error)
	GetUserByID(id string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	UserIsExist(email string) (bool, error)
}

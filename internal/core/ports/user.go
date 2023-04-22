package ports

import "user-svc/internal/core/domain"

type UserService interface {
	Create(request *domain.CreateUserRequest) (*domain.Response, error)
	Update(request *domain.UpdateUserRequest) (*domain.Response, error)
	Delete(id string) (*domain.Response, error)
	Users() (*domain.Response, error)
	User(id string) (*domain.Response, error)
}

type UserRepository interface {
	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id string) error
	Users() ([]*domain.User, error)
	UserByID(id string) (*domain.User, error)
	UserByEmail(email string) (*domain.User, error)
	Exist(email string) (bool, error)
}

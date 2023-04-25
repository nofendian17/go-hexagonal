package domain

import (
	"time"
)

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	Salt      string    `json:"-"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Id       string `param:"id" validate:"required,uuid"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Active   *bool  `json:"active" validate:"required"`
	Password string `json:"password"`
}

type DeleteUserRequest struct {
	Id string `param:"id" validate:"required,uuid"`
}

type GetUserRequest struct {
	Id string `param:"id" validate:"required,uuid"`
}

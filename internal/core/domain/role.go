package domain

import "time"

type Role struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreateRoleRequest struct {
	Name   string `json:"name" validate:"required"`
	Active *bool  `json:"active" validate:"required"`
}

type UpdateRoleRequest struct {
	Id     string `param:"id" validate:"required,uuid"`
	Name   string `json:"name" validate:"required"`
	Active *bool  `json:"active" validate:"required"`
}

type DeleteRoleRequest struct {
	Id string `param:"id" validate:"required,uuid"`
}

type GetRoleRequest struct {
	Id string `param:"id" validate:"required,uuid"`
}

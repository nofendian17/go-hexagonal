package domain

import "time"

type Permission struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type CreatePermissionRequest struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePermissionRequest struct {
	Id   string `param:"id" validate:"required,uuid"`
	Name string `json:"name" validate:"required"`
}

type DeletePermissionRequest struct {
	Id string `param:"id" validate:"required,uuid"`
}

type GetPermissionRequest struct {
	Id string `param:"id" validate:"required,uuid"`
}

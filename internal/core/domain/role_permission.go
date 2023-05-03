package domain

type GetRolePermissionRequest struct {
	RoleId string `param:"role_id" validate:"required,uuid"`
}

type AssignPermissionToRoleRequest struct {
	RoleId        string   `param:"role_id" validate:"required,uuid"`
	PermissionsId []string `json:"permissions_id" validate:"required,min=1,dive,uuid"`
}

type RemovePermissionFromRoleRequest struct {
	RoleId        string   `param:"role_id" validate:"required,uuid"`
	PermissionsId []string `json:"permissions_id" validate:"required,min=1,dive,uuid"`
}

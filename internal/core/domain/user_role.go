package domain

type GetUserRolesRequest struct {
	UserId string `param:"user_id" validate:"required,uuid"`
}

type AssignRolesToUserRequest struct {
	UserId  string   `param:"user_id" validate:"required,uuid"`
	RolesId []string `json:"roles_id" validate:"required,min=1,dive,uuid"`
}

type RemoveRolesFromUserRequest struct {
	UserId  string   `param:"user_id" validate:"required,uuid"`
	RolesId []string `json:"roles_id" validate:"dive,required,min=1,uuid"`
}

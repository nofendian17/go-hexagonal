package domain

import "time"

type Token struct {
	AccessToken      string        `json:"access_token"`
	AccessExpiresIn  time.Duration `json:"access_expires_in"`
	RefreshToken     string        `json:"refresh_token"`
	RefreshExpiresIn time.Duration `json:"refresh_expires_in"`
	CreatedDate      time.Time     `json:"created_date"`
}

type TokenInfo struct {
	UserID          string            `json:"user_id"`
	Roles           []*Role           `json:"roles"`
	AdditionalField map[string]string `json:"additional_field"`
}

type GetTokenRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type GetRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type GetDestroyTokenRequest struct {
	AccessToken string `header:"Authorization" validate:"required"`
}

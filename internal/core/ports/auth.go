package ports

import (
	"time"
	"user-svc/internal/core/domain"
)

type AuthService interface {
	Authenticate(request *domain.GetTokenRequest) (*domain.Response, error)
	Refresh(request *domain.RefreshTokenRequest) (*domain.Response, error)
	Logout(authID string) (*domain.Response, error)
}

type AuthRepository interface {
	SaveToken(key string, tokenDetail *domain.TokenInfo, expiration time.Duration) error
	GetToken(key string) (*domain.TokenInfo, error)
	DeleteToken(key string) error
}

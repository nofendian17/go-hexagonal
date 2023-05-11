package ports

import (
	"time"
	"user-svc/internal/core/domain"
)

type AuthService interface {
	Authenticate(request *domain.GetTokenRequest) (*domain.Response, error)
	Refresh(request *domain.GetRefreshTokenRequest) (*domain.Response, error)
	Logout(request *domain.GetDestroyTokenRequest) (*domain.Response, error)
}

type AuthRepository interface {
	SaveToken(key string, tokenDetail *domain.TokenInfo, expiration time.Duration) error
}

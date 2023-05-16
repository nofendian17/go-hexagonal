package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"user-svc/internal/core/ports"
	"user-svc/internal/shared/constants"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	appError "user-svc/internal/shared/error"
)

type JWTAuthenticator interface {
	Authenticate(c echo.Context) (*jwt.Token, error)
}

type JWTAuthenticatorImpl struct {
	SecretKey []byte
}

type Middleware interface {
	Handle(next echo.HandlerFunc) echo.HandlerFunc
}

type JWTMiddleware struct {
	Authenticator  JWTAuthenticator
	AuthRepository ports.AuthRepository
}

func (a *JWTAuthenticatorImpl) Authenticate(c echo.Context) (*jwt.Token, error) {
	authHeader := c.Request().Header.Get(constants.KeyAuthorization)
	tokenString := strings.Replace(authHeader, constants.KeyBearer+" ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return a.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (m *JWTMiddleware) Handle(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := m.Authenticator.Authenticate(c)
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, &appError.AppError{
				Code:    http.StatusUnauthorized,
				Message: "invalid authorization token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, &appError.AppError{
				Code:    http.StatusUnauthorized,
				Message: "invalid authorization token",
			})
		}

		authID := claims[constants.KeyAuthID].(string)

		isExist, err := m.AuthRepository.TokenExist(authID)
		if err != nil || !isExist {
			return c.JSON(http.StatusUnauthorized, &appError.AppError{
				Code:    http.StatusUnauthorized,
				Message: "invalid authorization token",
			})
		}

		c.Set(constants.KeyAuthID, authID)

		return next(c)
	}
}

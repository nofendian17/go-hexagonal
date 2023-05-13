package services

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	"user-svc/internal/shared/config"
	"user-svc/internal/shared/constants"
	appError "user-svc/internal/shared/error"
	"user-svc/internal/shared/hash"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AuthService struct {
	config          *config.Config
	userRepository  ports.UserRepository
	authRepository  ports.AuthRepository
	userRoleService ports.UserRoleService
	hasher          hash.Hasher
}

func NewAuthService(config *config.Config, userRepository ports.UserRepository, authRepository ports.AuthRepository, userRoleService ports.UserRoleService, hasher hash.Hasher) *AuthService {
	return &AuthService{
		config:          config,
		userRepository:  userRepository,
		authRepository:  authRepository,
		userRoleService: userRoleService,
		hasher:          hasher,
	}
}

func (s *AuthService) Authenticate(request *domain.GetTokenRequest) (*domain.Response, error) {
	user, err := s.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &appError.AppError{Code: http.StatusUnauthorized, Message: fmt.Sprintf("user with email %s not found", request.Email)}
		}
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	if user == nil {
		return nil, &appError.AppError{Code: http.StatusUnauthorized, Message: fmt.Sprintf("user with email %s not found", request.Email)}
	}
	if !user.Active {
		return nil, &appError.AppError{Code: http.StatusUnauthorized, Message: fmt.Sprintf("user with email %s is blocked", request.Email)}
	}

	salt, err := base64.URLEncoding.DecodeString(user.Salt)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	isMatch := s.hasher.CheckPassword(user.Password, request.Password, salt)

	if !isMatch {
		return nil, &appError.AppError{Code: http.StatusUnauthorized, Message: "incorrect password"}
	}

	roles, err := s.getUserRoles(user.Id)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	tokenInfo := &domain.TokenInfo{
		UserID: user.Id,
		Roles:  roles,
		AdditionalField: map[string]string{
			"x_device_id": "iphone",
		},
	}

	token := jwt.New(jwt.SigningMethodHS256)
	accessUUID, generateTime, accessToken, authTokenExpiredIn, err := s.crateAccessToken(token)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	refreshUUID, refreshToken, refreshTokenExpiredIn, err := s.createRefreshToken(token, generateTime)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	accessExpiresIn := time.Duration(authTokenExpiredIn) * time.Millisecond
	refreshExpiresIn := time.Duration(refreshTokenExpiredIn) * time.Millisecond

	saveTokenErrs := make(chan error, 2)
	go func() {
		saveTokenErrs <- s.authRepository.SaveToken(accessUUID, tokenInfo, accessExpiresIn)
	}()

	go func() {
		saveTokenErrs <- s.authRepository.SaveToken(refreshUUID, tokenInfo, refreshExpiresIn)
	}()

	for i := 0; i < 2; i++ {
		if err := <-saveTokenErrs; err != nil {
			return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		}
	}

	authToken := domain.Token{
		AccessToken:      accessToken,
		AccessExpiresIn:  accessExpiresIn,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: refreshExpiresIn,
		CreatedDate:      time.Unix(generateTime, 0),
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    authToken,
	}, nil
}

func (s *AuthService) Refresh(request *domain.RefreshTokenRequest) (*domain.Response, error) {
	secretKey := s.config.App.Auth.RefreshKey
	token, err := s.parseToken(request.RefreshToken, secretKey)
	if err != nil || !token.Valid {
		return nil, &appError.AppError{Code: http.StatusUnauthorized, Message: "invalid refresh token"}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &appError.AppError{Code: http.StatusUnauthorized, Message: "invalid refresh token"}
	}

	authID := claims[constants.KeyAuthID].(string)

	tokenInfo, err := s.authRepository.GetToken(authID)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusUnauthorized, Message: "invalid refresh token"}
	}

	regenerateToken := jwt.New(jwt.SigningMethodHS256)
	accessUUID, generateTime, accessToken, authTokenExpiredIn, err := s.crateAccessToken(regenerateToken)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	refreshUUID, refreshToken, refreshTokenExpiredIn, err := s.createRefreshToken(regenerateToken, generateTime)
	if err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	accessExpiresIn := time.Duration(authTokenExpiredIn) * time.Millisecond
	refreshExpiresIn := time.Duration(refreshTokenExpiredIn) * time.Millisecond

	saveTokenErrs := make(chan error, 2)
	deleteTokenErr := make(chan error, 1)
	go func() {
		saveTokenErrs <- s.authRepository.SaveToken(accessUUID, tokenInfo, accessExpiresIn)
	}()

	go func() {
		saveTokenErrs <- s.authRepository.SaveToken(refreshUUID, tokenInfo, refreshExpiresIn)
	}()

	for i := 0; i < 2; i++ {
		if err := <-saveTokenErrs; err != nil {
			return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
		}
	}

	go func() {
		deleteTokenErr <- s.authRepository.DeleteToken(authID)
	}()

	if err := <-deleteTokenErr; err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}

	authToken := domain.Token{
		AccessToken:      accessToken,
		AccessExpiresIn:  accessExpiresIn,
		RefreshToken:     refreshToken,
		RefreshExpiresIn: refreshExpiresIn,
		CreatedDate:      time.Unix(generateTime, 0),
	}

	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    authToken,
	}, nil
}

func (s *AuthService) Logout(authID string) (*domain.Response, error) {
	deleteTokenErr := make(chan error, 1)
	go func() {
		deleteTokenErr <- s.authRepository.DeleteToken(authID)
	}()

	if err := <-deleteTokenErr; err != nil {
		return nil, &appError.AppError{Code: http.StatusInternalServerError, Message: err.Error()}
	}
	return &domain.Response{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    nil,
	}, nil
}

func (s *AuthService) crateAccessToken(token *jwt.Token) (accessUUID string, generateTime int64, tokenString string, expiredIn int64, err error) {
	secretKey := s.config.App.Auth.AccessKey
	accessUUID = uuid.New().String()
	expiredAtTime := time.Now().Add(time.Minute * time.Duration(s.config.App.Auth.AccessLifeTime))
	expiredIn = expiredAtTime.Sub(time.Now()).Milliseconds()
	generateTime = time.Now().Unix()

	token.Claims = jwt.MapClaims{
		constants.KeyTokenType:    "access",
		constants.KeyGenerateTime: generateTime,
		constants.KeyAuthID:       accessUUID,
		constants.KeyExp:          expiredAtTime.Unix(),
	}

	tokenString, err = token.SignedString([]byte(secretKey))
	return
}

func (s *AuthService) createRefreshToken(token *jwt.Token, generateTime int64) (refreshUUID string, refreshToken string, expiredIn int64, err error) {
	secretKey := s.config.App.Auth.RefreshKey
	refreshUUID = uuid.New().String()
	expiredAtTime := time.Now().Add(time.Minute * time.Duration(s.config.App.Auth.RefreshLifeTime))
	expiredIn = expiredAtTime.Sub(time.Now()).Milliseconds()
	token.Claims = jwt.MapClaims{
		constants.KeyTokenType:    "refresh",
		constants.KeyGenerateTime: generateTime,
		constants.KeyAuthID:       refreshUUID,
		constants.KeyExp:          expiredAtTime.Unix(),
	}
	refreshToken, err = token.SignedString([]byte(secretKey))
	return
}

func (s *AuthService) getUserRoles(userID string) ([]*domain.Role, error) {
	userRoles := domain.GetUserRolesRequest{
		UserId: userID,
	}

	roles, err := s.userRoleService.GetUserRoles(&userRoles)
	if err != nil {
		return nil, err
	}

	return roles.Data.([]*domain.Role), nil
}

func (s *AuthService) parseToken(tokenString, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

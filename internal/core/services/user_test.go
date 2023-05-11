package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"testing"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	mocksUserRepository "user-svc/internal/mocks/core/ports"
	mocksHasher "user-svc/internal/mocks/shared/hash"
	"user-svc/internal/shared/hash"
)

func TestNewUserService(t *testing.T) {
	mockUserRepository := mocksUserRepository.UserRepository{}
	mockCacheRepository := mocksUserRepository.CacheRepository{}
	mockHasher := mocksHasher.Hasher{}
	type args struct {
		repo  ports.UserRepository
		cache ports.CacheRepository
		hash  hash.Hasher
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		{
			name: "success",
			args: args{
				repo: &mockUserRepository,
				hash: &mockHasher,
			},
			want: NewUserService(
				&mockUserRepository,
				&mockCacheRepository,
				&mockHasher,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repo, tt.args.cache, tt.args.hash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	type args struct {
		user *domain.CreateUserRequest
	}
	tests := []struct {
		name         string
		args         args
		existResult  []interface{}
		createResult error
		hasherResult []interface{}
		want         *domain.Response
		wantErr      bool
	}{
		{
			name: "success - created",
			args: args{user: &domain.CreateUserRequest{
				Name:     "test",
				Email:    "test@mail.com",
				Password: "secret",
			}},
			existResult: []interface{}{
				false, nil,
			},
			createResult: nil,
			hasherResult: []interface{}{
				"hashedPassword", "salt", nil,
			},
			want: &domain.Response{
				Code:    http.StatusCreated,
				Message: http.StatusText(http.StatusCreated),
				Data:    nil,
			},
			wantErr: false,
		},
		{
			name: "failed - check user with email error",
			args: args{user: &domain.CreateUserRequest{
				Name:     "test",
				Email:    "test@mail.com",
				Password: "secret",
			}},
			existResult: []interface{}{
				false, errors.New("error"),
			},
			hasherResult: []interface{}{
				"hashedPassword", "salt", nil,
			},
			createResult: nil,
			want:         nil,
			wantErr:      true,
		},
		{
			name: "failed - check user with email already exist",
			args: args{user: &domain.CreateUserRequest{
				Name:     "test",
				Email:    "test@mail.com",
				Password: "secret",
			}},
			existResult: []interface{}{
				true, nil,
			},
			createResult: nil,
			hasherResult: []interface{}{
				"hashedPassword", "salt", nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - hash password error",
			args: args{user: &domain.CreateUserRequest{
				Name:     "test",
				Email:    "test@mail.com",
				Password: "secret",
			}},
			existResult: []interface{}{
				false, nil,
			},
			createResult: nil,
			hasherResult: []interface{}{
				"", "", errors.New("hash password error"),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "failed - unable to save user data",
			args: args{user: &domain.CreateUserRequest{
				Name:     "test",
				Email:    "test@mail.com",
				Password: "secret",
			}},
			existResult: []interface{}{
				false, nil,
			},
			hasherResult: []interface{}{
				"hashedPassword", "salt", nil,
			},
			createResult: errors.New("has error"),
			want:         nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := mocksUserRepository.UserRepository{}
			mockUserRepository.On("UserIsExist", mock.Anything).Return(tt.existResult...)
			mockUserRepository.On("CreateUser", mock.Anything).Return(tt.createResult)

			mockHasher := mocksHasher.Hasher{}
			mockHasher.On("HashPassword", mock.Anything).Return(tt.hasherResult...)
			u := UserService{
				userRepository: &mockUserRepository,
				hasher:         &mockHasher,
			}
			got, err := u.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

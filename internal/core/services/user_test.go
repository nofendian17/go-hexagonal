package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"testing"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	mockCore "user-svc/internal/mocks/core/ports"
	mockShared "user-svc/internal/mocks/shared/hash"
	"user-svc/internal/shared/hash"
)

func TestNewUserService(t *testing.T) {
	mockUserRepository := mockCore.UserRepository{}
	mockHasher := mockShared.Hasher{}
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
				&mockHasher,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repo, tt.args.hash); !reflect.DeepEqual(got, tt.want) {
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
		saltResult   []interface{}
		hasherResult string
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
			saltResult: []interface{}{
				[]uint8("salt"), nil,
			},
			hasherResult: "secret",
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
			saltResult: []interface{}{
				[]uint8("salt"), nil,
			},
			existResult: []interface{}{
				false, errors.New("error"),
			},
			hasherResult: "12345",
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
			saltResult: []interface{}{
				[]uint8("salt"), nil,
			},
			existResult: []interface{}{
				true, nil,
			},
			createResult: nil,
			hasherResult: "secret",
			want:         nil,
			wantErr:      true,
		},
		{
			name: "failed - hash password error",
			args: args{user: &domain.CreateUserRequest{
				Name:     "test",
				Email:    "test@mail.com",
				Password: "secret",
			}},
			saltResult: []interface{}{
				nil, errors.New("error"),
			},
			existResult: []interface{}{
				false, nil,
			},
			createResult: nil,
			hasherResult: "secret",
			want:         nil,
			wantErr:      true,
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
			saltResult: []interface{}{
				[]uint8("salt"), nil,
			},
			hasherResult: "secret",
			createResult: errors.New("has error"),
			want:         nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := mockCore.UserRepository{}
			mockUserRepository.On("UserIsExist", mock.Anything).Return(tt.existResult...)
			mockUserRepository.On("CreateUser", mock.Anything).Return(tt.createResult)

			mockHasher := mockShared.Hasher{}
			mockHasher.On("GenerateRandomSalt").Return(tt.saltResult...)
			mockHasher.On("HashPassword", mock.Anything, mock.Anything).Return(tt.hasherResult)
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

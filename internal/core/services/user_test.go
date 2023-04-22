package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"net/http"
	"reflect"
	"testing"
	"user-svc/internal/core/domain"
	"user-svc/internal/core/ports"
	mocks "user-svc/internal/mocks/core/ports"
)

func TestNewUserService(t *testing.T) {
	mockUserRepository := mocks.UserRepository{}
	type args struct {
		repo ports.UserRepository
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
		{
			name: "success",
			args: args{repo: &mockUserRepository},
			want: NewUserService(&mockUserRepository),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
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
			want: &domain.Response{
				Code:    http.StatusCreated,
				Message: "created",
				Data:    nil,
			},
			wantErr: false,
		},
		{
			name: "failed - user already exist",
			args: args{user: &domain.CreateUserRequest{
				Name:     "test",
				Email:    "test@mail.com",
				Password: "secret",
			}},
			existResult: []interface{}{
				true, errors.New("user already exist"),
			},
			createResult: nil,
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
			createResult: errors.New("has error"),
			want:         nil,
			wantErr:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := mocks.UserRepository{}
			mockUserRepository.On("Exist", mock.Anything).Return(tt.existResult...)
			mockUserRepository.On("Create", mock.Anything).Return(tt.createResult)

			u := UserService{
				repo: &mockUserRepository,
			}
			got, err := u.Create(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

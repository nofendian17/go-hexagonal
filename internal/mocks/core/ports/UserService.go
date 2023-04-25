// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	domain "user-svc/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: request
func (_m *UserService) CreateUser(request *domain.CreateUserRequest) (*domain.Response, error) {
	ret := _m.Called(request)

	var r0 *domain.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.CreateUserRequest) (*domain.Response, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*domain.CreateUserRequest) *domain.Response); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.CreateUserRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: id
func (_m *UserService) DeleteUser(id string) (*domain.Response, error) {
	ret := _m.Called(id)

	var r0 *domain.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Response, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Response); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: id
func (_m *UserService) GetUser(id string) (*domain.Response, error) {
	ret := _m.Called(id)

	var r0 *domain.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Response, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Response); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields:
func (_m *UserService) GetUsers() (*domain.Response, error) {
	ret := _m.Called()

	var r0 *domain.Response
	var r1 error
	if rf, ok := ret.Get(0).(func() (*domain.Response, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *domain.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: request
func (_m *UserService) UpdateUser(request *domain.UpdateUserRequest) (*domain.Response, error) {
	ret := _m.Called(request)

	var r0 *domain.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.UpdateUserRequest) (*domain.Response, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(*domain.UpdateUserRequest) *domain.Response); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.UpdateUserRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

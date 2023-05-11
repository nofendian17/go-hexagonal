// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	domain "user-svc/internal/core/domain"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// SaveToken provides a mock function with given fields: key, tokenDetail, expiration
func (_m *AuthRepository) SaveToken(key string, tokenDetail *domain.TokenInfo, expiration time.Duration) error {
	ret := _m.Called(key, tokenDetail, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *domain.TokenInfo, time.Duration) error); ok {
		r0 = rf(key, tokenDetail, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAuthRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthRepository(t mockConstructorTestingTNewAuthRepository) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

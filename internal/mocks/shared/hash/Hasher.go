// Code generated by mockery v2.26.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Hasher is an autogenerated mock type for the Hasher type
type Hasher struct {
	mock.Mock
}

// CheckPassword provides a mock function with given fields: password, hashedPassword, salt
func (_m *Hasher) CheckPassword(password string, hashedPassword string, salt string) (bool, error) {
	ret := _m.Called(password, hashedPassword, salt)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (bool, error)); ok {
		return rf(password, hashedPassword, salt)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) bool); ok {
		r0 = rf(password, hashedPassword, salt)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(password, hashedPassword, salt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashPassword provides a mock function with given fields: password
func (_m *Hasher) HashPassword(password string) (string, string, error) {
	ret := _m.Called(password)

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (string, string, error)); ok {
		return rf(password)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewHasher interface {
	mock.TestingT
	Cleanup(func())
}

// NewHasher creates a new instance of Hasher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHasher(t mockConstructorTestingTNewHasher) *Hasher {
	mock := &Hasher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
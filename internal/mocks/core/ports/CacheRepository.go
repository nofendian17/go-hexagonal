// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// CacheRepository is an autogenerated mock type for the CacheRepository type
type CacheRepository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *CacheRepository) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: key
func (_m *CacheRepository) Delete(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: key
func (_m *CacheRepository) Exists(key string) (bool, error) {
	ret := _m.Called(key)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key
func (_m *CacheRepository) Get(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields:
func (_m *CacheRepository) Ping() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Set provides a mock function with given fields: key, value, expiration
func (_m *CacheRepository) Set(key string, value interface{}, expiration time.Duration) error {
	ret := _m.Called(key, value, expiration)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) error); ok {
		r0 = rf(key, value, expiration)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCacheRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheRepository creates a new instance of CacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheRepository(t mockConstructorTestingTNewCacheRepository) *CacheRepository {
	mock := &CacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
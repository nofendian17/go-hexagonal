// Code generated by mockery v2.26.1. DO NOT EDIT.

package mocks

import (
	domain "user-svc/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// RolePermissionRepository is an autogenerated mock type for the RolePermissionRepository type
type RolePermissionRepository struct {
	mock.Mock
}

// AddRolePermissions provides a mock function with given fields: roleId, permissions
func (_m *RolePermissionRepository) AddRolePermissions(roleId string, permissions []string) error {
	ret := _m.Called(roleId, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(roleId, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRolePermissions provides a mock function with given fields: roleId
func (_m *RolePermissionRepository) GetRolePermissions(roleId string) ([]*domain.Permission, error) {
	ret := _m.Called(roleId)

	var r0 []*domain.Permission
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*domain.Permission, error)); ok {
		return rf(roleId)
	}
	if rf, ok := ret.Get(0).(func(string) []*domain.Permission); ok {
		r0 = rf(roleId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Permission)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(roleId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveRolePermissions provides a mock function with given fields: roleId, permissions
func (_m *RolePermissionRepository) RemoveRolePermissions(roleId string, permissions []string) error {
	ret := _m.Called(roleId, permissions)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(roleId, permissions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRolePermissionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRolePermissionRepository creates a new instance of RolePermissionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRolePermissionRepository(t mockConstructorTestingTNewRolePermissionRepository) *RolePermissionRepository {
	mock := &RolePermissionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

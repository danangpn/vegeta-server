// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import io "io"
import mock "github.com/stretchr/testify/mock"
import models "vegeta-server/models"
import time "time"

// ITaskGetter is an autogenerated mock type for the ITaskGetter type
type ITaskGetter struct {
	mock.Mock
}

// CreatedAt provides a mock function with given fields:
func (_m *ITaskGetter) CreatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *ITaskGetter) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Params provides a mock function with given fields:
func (_m *ITaskGetter) Params() models.AttackParams {
	ret := _m.Called()

	var r0 models.AttackParams
	if rf, ok := ret.Get(0).(func() models.AttackParams); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.AttackParams)
	}

	return r0
}

// Result provides a mock function with given fields:
func (_m *ITaskGetter) Result() io.Reader {
	ret := _m.Called()

	var r0 io.Reader
	if rf, ok := ret.Get(0).(func() io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Reader)
		}
	}

	return r0
}

// Status provides a mock function with given fields:
func (_m *ITaskGetter) Status() models.AttackStatus {
	ret := _m.Called()

	var r0 models.AttackStatus
	if rf, ok := ret.Get(0).(func() models.AttackStatus); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.AttackStatus)
	}

	return r0
}

// UpdatedAt provides a mock function with given fields:
func (_m *ITaskGetter) UpdatedAt() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}
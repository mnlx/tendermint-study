// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Source is an autogenerated mock type for the Source type
type Source struct {
	mock.Mock
}

// Now provides a mock function with given fields:
func (_m *Source) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	goka "github.com/lovoo/goka"
	mock "github.com/stretchr/testify/mock"
)

// GokaEventHandler is an autogenerated mock type for the GokaEventHandler type
type GokaEventHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, message
func (_m *GokaEventHandler) Handle(ctx goka.Context, message interface{}) {
	_m.Called(ctx, message)
}

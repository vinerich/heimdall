// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	heimdall "github.com/dadrus/heimdall/internal/heimdall"
	errorhandlers "github.com/dadrus/heimdall/internal/rules/mechanisms/errorhandlers"

	mock "github.com/stretchr/testify/mock"
)

// ErrorHandlerMock is an autogenerated mock type for the ErrorHandler type
type ErrorHandlerMock struct {
	mock.Mock
}

type ErrorHandlerMock_Expecter struct {
	mock *mock.Mock
}

func (_m *ErrorHandlerMock) EXPECT() *ErrorHandlerMock_Expecter {
	return &ErrorHandlerMock_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, err
func (_m *ErrorHandlerMock) Execute(ctx heimdall.Context, err error) (bool, error) {
	ret := _m.Called(ctx, err)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(heimdall.Context, error) (bool, error)); ok {
		return rf(ctx, err)
	}
	if rf, ok := ret.Get(0).(func(heimdall.Context, error) bool); ok {
		r0 = rf(ctx, err)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(heimdall.Context, error) error); ok {
		r1 = rf(ctx, err)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ErrorHandlerMock_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ErrorHandlerMock_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx heimdall.Context
//   - err error
func (_e *ErrorHandlerMock_Expecter) Execute(ctx interface{}, err interface{}) *ErrorHandlerMock_Execute_Call {
	return &ErrorHandlerMock_Execute_Call{Call: _e.mock.On("Execute", ctx, err)}
}

func (_c *ErrorHandlerMock_Execute_Call) Run(run func(ctx heimdall.Context, err error)) *ErrorHandlerMock_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(heimdall.Context), args[1].(error))
	})
	return _c
}

func (_c *ErrorHandlerMock_Execute_Call) Return(_a0 bool, _a1 error) *ErrorHandlerMock_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ErrorHandlerMock_Execute_Call) RunAndReturn(run func(heimdall.Context, error) (bool, error)) *ErrorHandlerMock_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// WithConfig provides a mock function with given fields: config
func (_m *ErrorHandlerMock) WithConfig(config map[string]interface{}) (errorhandlers.ErrorHandler, error) {
	ret := _m.Called(config)

	var r0 errorhandlers.ErrorHandler
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) (errorhandlers.ErrorHandler, error)); ok {
		return rf(config)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) errorhandlers.ErrorHandler); ok {
		r0 = rf(config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errorhandlers.ErrorHandler)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ErrorHandlerMock_WithConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithConfig'
type ErrorHandlerMock_WithConfig_Call struct {
	*mock.Call
}

// WithConfig is a helper method to define mock.On call
//   - config map[string]interface{}
func (_e *ErrorHandlerMock_Expecter) WithConfig(config interface{}) *ErrorHandlerMock_WithConfig_Call {
	return &ErrorHandlerMock_WithConfig_Call{Call: _e.mock.On("WithConfig", config)}
}

func (_c *ErrorHandlerMock_WithConfig_Call) Run(run func(config map[string]interface{})) *ErrorHandlerMock_WithConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(map[string]interface{}))
	})
	return _c
}

func (_c *ErrorHandlerMock_WithConfig_Call) Return(_a0 errorhandlers.ErrorHandler, _a1 error) *ErrorHandlerMock_WithConfig_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ErrorHandlerMock_WithConfig_Call) RunAndReturn(run func(map[string]interface{}) (errorhandlers.ErrorHandler, error)) *ErrorHandlerMock_WithConfig_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewErrorHandlerMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewErrorHandlerMock creates a new instance of ErrorHandlerMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewErrorHandlerMock(t mockConstructorTestingTNewErrorHandlerMock) *ErrorHandlerMock {
	mock := &ErrorHandlerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

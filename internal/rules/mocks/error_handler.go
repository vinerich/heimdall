// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	heimdall "github.com/dadrus/heimdall/internal/heimdall"
	mock "github.com/stretchr/testify/mock"
)

// ErrorHandlerMock is an autogenerated mock type for the errorHandler type
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

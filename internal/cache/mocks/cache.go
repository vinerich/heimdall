// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// CacheMock is an autogenerated mock type for the Cache type
type CacheMock struct {
	mock.Mock
}

type CacheMock_Expecter struct {
	mock *mock.Mock
}

func (_m *CacheMock) EXPECT() *CacheMock_Expecter {
	return &CacheMock_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: key
func (_m *CacheMock) Delete(key string) {
	_m.Called(key)
}

// CacheMock_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type CacheMock_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - key string
func (_e *CacheMock_Expecter) Delete(key interface{}) *CacheMock_Delete_Call {
	return &CacheMock_Delete_Call{Call: _e.mock.On("Delete", key)}
}

func (_c *CacheMock_Delete_Call) Run(run func(key string)) *CacheMock_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *CacheMock_Delete_Call) Return() *CacheMock_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *CacheMock_Delete_Call) RunAndReturn(run func(string)) *CacheMock_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *CacheMock) Get(key string) interface{} {
	ret := _m.Called(key)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// CacheMock_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CacheMock_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *CacheMock_Expecter) Get(key interface{}) *CacheMock_Get_Call {
	return &CacheMock_Get_Call{Call: _e.mock.On("Get", key)}
}

func (_c *CacheMock_Get_Call) Run(run func(key string)) *CacheMock_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *CacheMock_Get_Call) Return(_a0 interface{}) *CacheMock_Get_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheMock_Get_Call) RunAndReturn(run func(string) interface{}) *CacheMock_Get_Call {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value, ttl
func (_m *CacheMock) Set(key string, value interface{}, ttl time.Duration) {
	_m.Called(key, value, ttl)
}

// CacheMock_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type CacheMock_Set_Call struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - value interface{}
//   - ttl time.Duration
func (_e *CacheMock_Expecter) Set(key interface{}, value interface{}, ttl interface{}) *CacheMock_Set_Call {
	return &CacheMock_Set_Call{Call: _e.mock.On("Set", key, value, ttl)}
}

func (_c *CacheMock_Set_Call) Run(run func(key string, value interface{}, ttl time.Duration)) *CacheMock_Set_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(interface{}), args[2].(time.Duration))
	})
	return _c
}

func (_c *CacheMock_Set_Call) Return() *CacheMock_Set_Call {
	_c.Call.Return()
	return _c
}

func (_c *CacheMock_Set_Call) RunAndReturn(run func(string, interface{}, time.Duration)) *CacheMock_Set_Call {
	_c.Call.Return(run)
	return _c
}

// Start provides a mock function with given fields: ctx
func (_m *CacheMock) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheMock_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type CacheMock_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CacheMock_Expecter) Start(ctx interface{}) *CacheMock_Start_Call {
	return &CacheMock_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *CacheMock_Start_Call) Run(run func(ctx context.Context)) *CacheMock_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CacheMock_Start_Call) Return(_a0 error) *CacheMock_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheMock_Start_Call) RunAndReturn(run func(context.Context) error) *CacheMock_Start_Call {
	_c.Call.Return(run)
	return _c
}

// Stop provides a mock function with given fields: ctx
func (_m *CacheMock) Stop(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheMock_Stop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Stop'
type CacheMock_Stop_Call struct {
	*mock.Call
}

// Stop is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CacheMock_Expecter) Stop(ctx interface{}) *CacheMock_Stop_Call {
	return &CacheMock_Stop_Call{Call: _e.mock.On("Stop", ctx)}
}

func (_c *CacheMock_Stop_Call) Run(run func(ctx context.Context)) *CacheMock_Stop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CacheMock_Stop_Call) Return(_a0 error) *CacheMock_Stop_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheMock_Stop_Call) RunAndReturn(run func(context.Context) error) *CacheMock_Stop_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCacheMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheMock creates a new instance of CacheMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheMock(t mockConstructorTestingTNewCacheMock) *CacheMock {
	mock := &CacheMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

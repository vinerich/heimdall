// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	heimdall "github.com/dadrus/heimdall/internal/heimdall"
	mock "github.com/stretchr/testify/mock"

	rule "github.com/dadrus/heimdall/internal/rules/rule"

	url "net/url"
)

// RuleMock is an autogenerated mock type for the Rule type
type RuleMock struct {
	mock.Mock
}

type RuleMock_Expecter struct {
	mock *mock.Mock
}

func (_m *RuleMock) EXPECT() *RuleMock_Expecter {
	return &RuleMock_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *RuleMock) Execute(_a0 heimdall.Context) (rule.URIMutator, error) {
	ret := _m.Called(_a0)

	var r0 rule.URIMutator
	var r1 error
	if rf, ok := ret.Get(0).(func(heimdall.Context) (rule.URIMutator, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(heimdall.Context) rule.URIMutator); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rule.URIMutator)
		}
	}

	if rf, ok := ret.Get(1).(func(heimdall.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RuleMock_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type RuleMock_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 heimdall.Context
func (_e *RuleMock_Expecter) Execute(_a0 interface{}) *RuleMock_Execute_Call {
	return &RuleMock_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *RuleMock_Execute_Call) Run(run func(_a0 heimdall.Context)) *RuleMock_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(heimdall.Context))
	})
	return _c
}

func (_c *RuleMock_Execute_Call) Return(_a0 rule.URIMutator, _a1 error) *RuleMock_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RuleMock_Execute_Call) RunAndReturn(run func(heimdall.Context) (rule.URIMutator, error)) *RuleMock_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ID provides a mock function with given fields:
func (_m *RuleMock) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RuleMock_ID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ID'
type RuleMock_ID_Call struct {
	*mock.Call
}

// ID is a helper method to define mock.On call
func (_e *RuleMock_Expecter) ID() *RuleMock_ID_Call {
	return &RuleMock_ID_Call{Call: _e.mock.On("ID")}
}

func (_c *RuleMock_ID_Call) Run(run func()) *RuleMock_ID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RuleMock_ID_Call) Return(_a0 string) *RuleMock_ID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuleMock_ID_Call) RunAndReturn(run func() string) *RuleMock_ID_Call {
	_c.Call.Return(run)
	return _c
}

// MatchesMethod provides a mock function with given fields: _a0
func (_m *RuleMock) MatchesMethod(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RuleMock_MatchesMethod_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MatchesMethod'
type RuleMock_MatchesMethod_Call struct {
	*mock.Call
}

// MatchesMethod is a helper method to define mock.On call
//   - _a0 string
func (_e *RuleMock_Expecter) MatchesMethod(_a0 interface{}) *RuleMock_MatchesMethod_Call {
	return &RuleMock_MatchesMethod_Call{Call: _e.mock.On("MatchesMethod", _a0)}
}

func (_c *RuleMock_MatchesMethod_Call) Run(run func(_a0 string)) *RuleMock_MatchesMethod_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *RuleMock_MatchesMethod_Call) Return(_a0 bool) *RuleMock_MatchesMethod_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuleMock_MatchesMethod_Call) RunAndReturn(run func(string) bool) *RuleMock_MatchesMethod_Call {
	_c.Call.Return(run)
	return _c
}

// MatchesURL provides a mock function with given fields: _a0
func (_m *RuleMock) MatchesURL(_a0 *url.URL) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*url.URL) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RuleMock_MatchesURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MatchesURL'
type RuleMock_MatchesURL_Call struct {
	*mock.Call
}

// MatchesURL is a helper method to define mock.On call
//   - _a0 *url.URL
func (_e *RuleMock_Expecter) MatchesURL(_a0 interface{}) *RuleMock_MatchesURL_Call {
	return &RuleMock_MatchesURL_Call{Call: _e.mock.On("MatchesURL", _a0)}
}

func (_c *RuleMock_MatchesURL_Call) Run(run func(_a0 *url.URL)) *RuleMock_MatchesURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*url.URL))
	})
	return _c
}

func (_c *RuleMock_MatchesURL_Call) Return(_a0 bool) *RuleMock_MatchesURL_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuleMock_MatchesURL_Call) RunAndReturn(run func(*url.URL) bool) *RuleMock_MatchesURL_Call {
	_c.Call.Return(run)
	return _c
}

// SrcID provides a mock function with given fields:
func (_m *RuleMock) SrcID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RuleMock_SrcID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SrcID'
type RuleMock_SrcID_Call struct {
	*mock.Call
}

// SrcID is a helper method to define mock.On call
func (_e *RuleMock_Expecter) SrcID() *RuleMock_SrcID_Call {
	return &RuleMock_SrcID_Call{Call: _e.mock.On("SrcID")}
}

func (_c *RuleMock_SrcID_Call) Run(run func()) *RuleMock_SrcID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RuleMock_SrcID_Call) Return(_a0 string) *RuleMock_SrcID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RuleMock_SrcID_Call) RunAndReturn(run func() string) *RuleMock_SrcID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewRuleMock interface {
	mock.TestingT
	Cleanup(func())
}

// NewRuleMock creates a new instance of RuleMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRuleMock(t mockConstructorTestingTNewRuleMock) *RuleMock {
	mock := &RuleMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

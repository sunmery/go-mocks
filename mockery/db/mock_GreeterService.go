// Code generated by mockery v2.40.3. DO NOT EDIT.

package db

import mock "github.com/stretchr/testify/mock"

// MockGreeterService is an autogenerated mock type for the GreeterService type
type MockGreeterService struct {
	mock.Mock
}

type MockGreeterService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGreeterService) EXPECT() *MockGreeterService_Expecter {
	return &MockGreeterService_Expecter{mock: &_m.Mock}
}

// Greet provides a mock function with given fields:
func (_m *MockGreeterService) Greet() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Greet")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockGreeterService_Greet_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Greet'
type MockGreeterService_Greet_Call struct {
	*mock.Call
}

// Greet is a helper method to define mock.On call
func (_e *MockGreeterService_Expecter) Greet() *MockGreeterService_Greet_Call {
	return &MockGreeterService_Greet_Call{Call: _e.mock.On("Greet")}
}

func (_c *MockGreeterService_Greet_Call) Run(run func()) *MockGreeterService_Greet_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGreeterService_Greet_Call) Return(_a0 string) *MockGreeterService_Greet_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGreeterService_Greet_Call) RunAndReturn(run func() string) *MockGreeterService_Greet_Call {
	_c.Call.Return(run)
	return _c
}

// GreetInDefaultMsg provides a mock function with given fields:
func (_m *MockGreeterService) GreetInDefaultMsg() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GreetInDefaultMsg")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockGreeterService_GreetInDefaultMsg_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GreetInDefaultMsg'
type MockGreeterService_GreetInDefaultMsg_Call struct {
	*mock.Call
}

// GreetInDefaultMsg is a helper method to define mock.On call
func (_e *MockGreeterService_Expecter) GreetInDefaultMsg() *MockGreeterService_GreetInDefaultMsg_Call {
	return &MockGreeterService_GreetInDefaultMsg_Call{Call: _e.mock.On("GreetInDefaultMsg")}
}

func (_c *MockGreeterService_GreetInDefaultMsg_Call) Run(run func()) *MockGreeterService_GreetInDefaultMsg_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGreeterService_GreetInDefaultMsg_Call) Return(_a0 string) *MockGreeterService_GreetInDefaultMsg_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGreeterService_GreetInDefaultMsg_Call) RunAndReturn(run func() string) *MockGreeterService_GreetInDefaultMsg_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGreeterService creates a new instance of MockGreeterService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGreeterService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGreeterService {
	mock := &MockGreeterService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

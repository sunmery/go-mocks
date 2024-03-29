// Code generated by mockery v2.40.3. DO NOT EDIT.

package db

import mock "github.com/stretchr/testify/mock"

// MockDB is an autogenerated mock type for the DB type
type MockDB struct {
	mock.Mock
}

type MockDB_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDB) EXPECT() *MockDB_Expecter {
	return &MockDB_Expecter{mock: &_m.Mock}
}

// FetchDefaultMessage provides a mock function with given fields:
func (_m *MockDB) FetchDefaultMessage() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FetchDefaultMessage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_FetchDefaultMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchDefaultMessage'
type MockDB_FetchDefaultMessage_Call struct {
	*mock.Call
}

// FetchDefaultMessage is a helper method to define mock.On call
func (_e *MockDB_Expecter) FetchDefaultMessage() *MockDB_FetchDefaultMessage_Call {
	return &MockDB_FetchDefaultMessage_Call{Call: _e.mock.On("FetchDefaultMessage")}
}

func (_c *MockDB_FetchDefaultMessage_Call) Run(run func()) *MockDB_FetchDefaultMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockDB_FetchDefaultMessage_Call) Return(_a0 string, _a1 error) *MockDB_FetchDefaultMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_FetchDefaultMessage_Call) RunAndReturn(run func() (string, error)) *MockDB_FetchDefaultMessage_Call {
	_c.Call.Return(run)
	return _c
}

// FetchMessage provides a mock function with given fields: lang
func (_m *MockDB) FetchMessage(lang string) (string, error) {
	ret := _m.Called(lang)

	if len(ret) == 0 {
		panic("no return value specified for FetchMessage")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(lang)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(lang)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(lang)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDB_FetchMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchMessage'
type MockDB_FetchMessage_Call struct {
	*mock.Call
}

// FetchMessage is a helper method to define mock.On call
//   - lang string
func (_e *MockDB_Expecter) FetchMessage(lang interface{}) *MockDB_FetchMessage_Call {
	return &MockDB_FetchMessage_Call{Call: _e.mock.On("FetchMessage", lang)}
}

func (_c *MockDB_FetchMessage_Call) Run(run func(lang string)) *MockDB_FetchMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockDB_FetchMessage_Call) Return(_a0 string, _a1 error) *MockDB_FetchMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDB_FetchMessage_Call) RunAndReturn(run func(string) (string, error)) *MockDB_FetchMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDB creates a new instance of MockDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDB {
	mock := &MockDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

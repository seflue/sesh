// Code generated by mockery v2.46.3. DO NOT EDIT.

package home

import mock "github.com/stretchr/testify/mock"

// MockHome is an autogenerated mock type for the Home type
type MockHome struct {
	mock.Mock
}

type MockHome_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHome) EXPECT() *MockHome_Expecter {
	return &MockHome_Expecter{mock: &_m.Mock}
}

// ExpandHome provides a mock function with given fields: path
func (_m *MockHome) ExpandHome(path string) (string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for ExpandHome")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHome_ExpandHome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExpandHome'
type MockHome_ExpandHome_Call struct {
	*mock.Call
}

// ExpandHome is a helper method to define mock.On call
//   - path string
func (_e *MockHome_Expecter) ExpandHome(path interface{}) *MockHome_ExpandHome_Call {
	return &MockHome_ExpandHome_Call{Call: _e.mock.On("ExpandHome", path)}
}

func (_c *MockHome_ExpandHome_Call) Run(run func(path string)) *MockHome_ExpandHome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockHome_ExpandHome_Call) Return(_a0 string, _a1 error) *MockHome_ExpandHome_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHome_ExpandHome_Call) RunAndReturn(run func(string) (string, error)) *MockHome_ExpandHome_Call {
	_c.Call.Return(run)
	return _c
}

// ShortenHome provides a mock function with given fields: path
func (_m *MockHome) ShortenHome(path string) (string, error) {
	ret := _m.Called(path)

	if len(ret) == 0 {
		panic("no return value specified for ShortenHome")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(path)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockHome_ShortenHome_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ShortenHome'
type MockHome_ShortenHome_Call struct {
	*mock.Call
}

// ShortenHome is a helper method to define mock.On call
//   - path string
func (_e *MockHome_Expecter) ShortenHome(path interface{}) *MockHome_ShortenHome_Call {
	return &MockHome_ShortenHome_Call{Call: _e.mock.On("ShortenHome", path)}
}

func (_c *MockHome_ShortenHome_Call) Run(run func(path string)) *MockHome_ShortenHome_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockHome_ShortenHome_Call) Return(_a0 string, _a1 error) *MockHome_ShortenHome_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockHome_ShortenHome_Call) RunAndReturn(run func(string) (string, error)) *MockHome_ShortenHome_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHome creates a new instance of MockHome. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHome(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHome {
	mock := &MockHome{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

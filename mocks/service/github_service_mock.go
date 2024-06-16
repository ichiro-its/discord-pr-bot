// Code generated by mockery v2.43.2. DO NOT EDIT.

package service

import mock "github.com/stretchr/testify/mock"

// GithubServiceMock is an autogenerated mock type for the GithubService type
type GithubServiceMock struct {
	mock.Mock
}

type GithubServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *GithubServiceMock) EXPECT() *GithubServiceMock_Expecter {
	return &GithubServiceMock_Expecter{mock: &_m.Mock}
}

// GetOpenPullRequestUrls provides a mock function with given fields:
func (_m *GithubServiceMock) GetOpenPullRequestUrls() ([]string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetOpenPullRequestUrls")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GithubServiceMock_GetOpenPullRequestUrls_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOpenPullRequestUrls'
type GithubServiceMock_GetOpenPullRequestUrls_Call struct {
	*mock.Call
}

// GetOpenPullRequestUrls is a helper method to define mock.On call
func (_e *GithubServiceMock_Expecter) GetOpenPullRequestUrls() *GithubServiceMock_GetOpenPullRequestUrls_Call {
	return &GithubServiceMock_GetOpenPullRequestUrls_Call{Call: _e.mock.On("GetOpenPullRequestUrls")}
}

func (_c *GithubServiceMock_GetOpenPullRequestUrls_Call) Run(run func()) *GithubServiceMock_GetOpenPullRequestUrls_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GithubServiceMock_GetOpenPullRequestUrls_Call) Return(pullRequestUrls []string, err error) *GithubServiceMock_GetOpenPullRequestUrls_Call {
	_c.Call.Return(pullRequestUrls, err)
	return _c
}

func (_c *GithubServiceMock_GetOpenPullRequestUrls_Call) RunAndReturn(run func() ([]string, error)) *GithubServiceMock_GetOpenPullRequestUrls_Call {
	_c.Call.Return(run)
	return _c
}

// NewGithubServiceMock creates a new instance of GithubServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGithubServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *GithubServiceMock {
	mock := &GithubServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
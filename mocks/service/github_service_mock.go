// Code generated by mockery v2.43.2. DO NOT EDIT.

package service

import (
	entity "github.com/ichiro-its/discord-pr-bot/entity"
	mock "github.com/stretchr/testify/mock"
)

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

// GetOpenPullRequests provides a mock function with given fields: org, repo
func (_m *GithubServiceMock) GetOpenPullRequests(org string, repo string) ([]*entity.PullRequest, error) {
	ret := _m.Called(org, repo)

	if len(ret) == 0 {
		panic("no return value specified for GetOpenPullRequests")
	}

	var r0 []*entity.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]*entity.PullRequest, error)); ok {
		return rf(org, repo)
	}
	if rf, ok := ret.Get(0).(func(string, string) []*entity.PullRequest); ok {
		r0 = rf(org, repo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(org, repo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GithubServiceMock_GetOpenPullRequests_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetOpenPullRequests'
type GithubServiceMock_GetOpenPullRequests_Call struct {
	*mock.Call
}

// GetOpenPullRequests is a helper method to define mock.On call
//   - org string
//   - repo string
func (_e *GithubServiceMock_Expecter) GetOpenPullRequests(org interface{}, repo interface{}) *GithubServiceMock_GetOpenPullRequests_Call {
	return &GithubServiceMock_GetOpenPullRequests_Call{Call: _e.mock.On("GetOpenPullRequests", org, repo)}
}

func (_c *GithubServiceMock_GetOpenPullRequests_Call) Run(run func(org string, repo string)) *GithubServiceMock_GetOpenPullRequests_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *GithubServiceMock_GetOpenPullRequests_Call) Return(_a0 []*entity.PullRequest, _a1 error) *GithubServiceMock_GetOpenPullRequests_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *GithubServiceMock_GetOpenPullRequests_Call) RunAndReturn(run func(string, string) ([]*entity.PullRequest, error)) *GithubServiceMock_GetOpenPullRequests_Call {
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

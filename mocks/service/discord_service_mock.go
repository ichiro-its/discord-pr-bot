// Code generated by mockery v2.43.2. DO NOT EDIT.

package service

import (
	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"
)

// DiscordServiceMock is an autogenerated mock type for the DiscordService type
type DiscordServiceMock struct {
	mock.Mock
}

type DiscordServiceMock_Expecter struct {
	mock *mock.Mock
}

func (_m *DiscordServiceMock) EXPECT() *DiscordServiceMock_Expecter {
	return &DiscordServiceMock_Expecter{mock: &_m.Mock}
}

// DeleteMessage provides a mock function with given fields: channelId, messageId
func (_m *DiscordServiceMock) DeleteMessage(channelId string, messageId string) error {
	ret := _m.Called(channelId, messageId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(channelId, messageId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscordServiceMock_DeleteMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteMessage'
type DiscordServiceMock_DeleteMessage_Call struct {
	*mock.Call
}

// DeleteMessage is a helper method to define mock.On call
//   - channelId string
//   - messageId string
func (_e *DiscordServiceMock_Expecter) DeleteMessage(channelId interface{}, messageId interface{}) *DiscordServiceMock_DeleteMessage_Call {
	return &DiscordServiceMock_DeleteMessage_Call{Call: _e.mock.On("DeleteMessage", channelId, messageId)}
}

func (_c *DiscordServiceMock_DeleteMessage_Call) Run(run func(channelId string, messageId string)) *DiscordServiceMock_DeleteMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *DiscordServiceMock_DeleteMessage_Call) Return(_a0 error) *DiscordServiceMock_DeleteMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DiscordServiceMock_DeleteMessage_Call) RunAndReturn(run func(string, string) error) *DiscordServiceMock_DeleteMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessages provides a mock function with given fields: channelId
func (_m *DiscordServiceMock) GetMessages(channelId string) ([]*discordgo.Message, error) {
	ret := _m.Called(channelId)

	if len(ret) == 0 {
		panic("no return value specified for GetMessages")
	}

	var r0 []*discordgo.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*discordgo.Message, error)); ok {
		return rf(channelId)
	}
	if rf, ok := ret.Get(0).(func(string) []*discordgo.Message); ok {
		r0 = rf(channelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DiscordServiceMock_GetMessages_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessages'
type DiscordServiceMock_GetMessages_Call struct {
	*mock.Call
}

// GetMessages is a helper method to define mock.On call
//   - channelId string
func (_e *DiscordServiceMock_Expecter) GetMessages(channelId interface{}) *DiscordServiceMock_GetMessages_Call {
	return &DiscordServiceMock_GetMessages_Call{Call: _e.mock.On("GetMessages", channelId)}
}

func (_c *DiscordServiceMock_GetMessages_Call) Run(run func(channelId string)) *DiscordServiceMock_GetMessages_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *DiscordServiceMock_GetMessages_Call) Return(_a0 []*discordgo.Message, _a1 error) *DiscordServiceMock_GetMessages_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DiscordServiceMock_GetMessages_Call) RunAndReturn(run func(string) ([]*discordgo.Message, error)) *DiscordServiceMock_GetMessages_Call {
	_c.Call.Return(run)
	return _c
}

// SendMessage provides a mock function with given fields: channelId, content
func (_m *DiscordServiceMock) SendMessage(channelId string, content string) error {
	ret := _m.Called(channelId, content)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(channelId, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscordServiceMock_SendMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendMessage'
type DiscordServiceMock_SendMessage_Call struct {
	*mock.Call
}

// SendMessage is a helper method to define mock.On call
//   - channelId string
//   - content string
func (_e *DiscordServiceMock_Expecter) SendMessage(channelId interface{}, content interface{}) *DiscordServiceMock_SendMessage_Call {
	return &DiscordServiceMock_SendMessage_Call{Call: _e.mock.On("SendMessage", channelId, content)}
}

func (_c *DiscordServiceMock_SendMessage_Call) Run(run func(channelId string, content string)) *DiscordServiceMock_SendMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *DiscordServiceMock_SendMessage_Call) Return(_a0 error) *DiscordServiceMock_SendMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DiscordServiceMock_SendMessage_Call) RunAndReturn(run func(string, string) error) *DiscordServiceMock_SendMessage_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateMessage provides a mock function with given fields: channelId, messageId, content
func (_m *DiscordServiceMock) UpdateMessage(channelId string, messageId string, content string) error {
	ret := _m.Called(channelId, messageId, content)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(channelId, messageId, content)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DiscordServiceMock_UpdateMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateMessage'
type DiscordServiceMock_UpdateMessage_Call struct {
	*mock.Call
}

// UpdateMessage is a helper method to define mock.On call
//   - channelId string
//   - messageId string
//   - content string
func (_e *DiscordServiceMock_Expecter) UpdateMessage(channelId interface{}, messageId interface{}, content interface{}) *DiscordServiceMock_UpdateMessage_Call {
	return &DiscordServiceMock_UpdateMessage_Call{Call: _e.mock.On("UpdateMessage", channelId, messageId, content)}
}

func (_c *DiscordServiceMock_UpdateMessage_Call) Run(run func(channelId string, messageId string, content string)) *DiscordServiceMock_UpdateMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *DiscordServiceMock_UpdateMessage_Call) Return(_a0 error) *DiscordServiceMock_UpdateMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DiscordServiceMock_UpdateMessage_Call) RunAndReturn(run func(string, string, string) error) *DiscordServiceMock_UpdateMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewDiscordServiceMock creates a new instance of DiscordServiceMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDiscordServiceMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *DiscordServiceMock {
	mock := &DiscordServiceMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

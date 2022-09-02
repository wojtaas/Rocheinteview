// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	rocheinteview "rocheinteview"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Ping provides a mock function with given fields: message
func (_m *Service) Ping(message string) rocheinteview.PingResponse {
	ret := _m.Called(message)

	var r0 rocheinteview.PingResponse
	if rf, ok := ret.Get(0).(func(string) rocheinteview.PingResponse); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Get(0).(rocheinteview.PingResponse)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

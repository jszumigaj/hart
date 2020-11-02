// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jszumigaj/hart (interfaces: FrameSender)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFrameSender is a mock of FrameSender interface.
type MockFrameSender struct {
	ctrl     *gomock.Controller
	recorder *MockFrameSenderMockRecorder
}

// MockFrameSenderMockRecorder is the mock recorder for MockFrameSender.
type MockFrameSenderMockRecorder struct {
	mock *MockFrameSender
}

// NewMockFrameSender creates a new mock instance.
func NewMockFrameSender(ctrl *gomock.Controller) *MockFrameSender {
	mock := &MockFrameSender{ctrl: ctrl}
	mock.recorder = &MockFrameSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrameSender) EXPECT() *MockFrameSenderMockRecorder {
	return m.recorder
}

// SendFrame mocks base method.
func (m *MockFrameSender) SendFrame(arg0, arg1 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendFrame", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendFrame indicates an expected call of SendFrame.
func (mr *MockFrameSenderMockRecorder) SendFrame(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendFrame", reflect.TypeOf((*MockFrameSender)(nil).SendFrame), arg0, arg1)
}
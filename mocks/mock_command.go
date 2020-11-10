// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jszumigaj/hart (interfaces: Command)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	hart "github.com/jszumigaj/hart"
)

// MockCommand is a mock of Command interface.
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand.
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance.
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// Data mocks base method.
func (m *MockCommand) Data() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Data")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Data indicates an expected call of Data.
func (mr *MockCommandMockRecorder) Data() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Data", reflect.TypeOf((*MockCommand)(nil).Data))
}

// Description mocks base method.
func (m *MockCommand) Description() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Description")
	ret0, _ := ret[0].(string)
	return ret0
}

// Description indicates an expected call of Description.
func (mr *MockCommandMockRecorder) Description() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Description", reflect.TypeOf((*MockCommand)(nil).Description))
}

// No mocks base method.
func (m *MockCommand) No() byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "No")
	ret0, _ := ret[0].(byte)
	return ret0
}

// No indicates an expected call of No.
func (mr *MockCommandMockRecorder) No() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "No", reflect.TypeOf((*MockCommand)(nil).No))
}

// SetData mocks base method.
func (m *MockCommand) SetData(arg0 []byte, arg1 hart.CommandStatus) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetData", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetData indicates an expected call of SetData.
func (mr *MockCommandMockRecorder) SetData(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetData", reflect.TypeOf((*MockCommand)(nil).SetData), arg0, arg1)
}

// Status mocks base method.
func (m *MockCommand) Status() hart.CommandStatus {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(hart.CommandStatus)
	return ret0
}

// Status indicates an expected call of Status.
func (mr *MockCommandMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCommand)(nil).Status))
}
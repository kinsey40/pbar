// Code generated by MockGen. DO NOT EDIT.
// Source: tqdm.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTqdmInterface is a mock of TqdmInterface interface
type MockTqdmInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTqdmInterfaceMockRecorder
}

// MockTqdmInterfaceMockRecorder is the mock recorder for MockTqdmInterface
type MockTqdmInterfaceMockRecorder struct {
	mock *MockTqdmInterface
}

// NewMockTqdmInterface creates a new mock instance
func NewMockTqdmInterface(ctrl *gomock.Controller) *MockTqdmInterface {
	mock := &MockTqdmInterface{ctrl: ctrl}
	mock.recorder = &MockTqdmInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTqdmInterface) EXPECT() *MockTqdmInterfaceMockRecorder {
	return m.recorder
}

// Update mocks base method
func (m *MockTqdmInterface) Update() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Update")
}

// Update indicates an expected call of Update
func (mr *MockTqdmInterfaceMockRecorder) Update() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockTqdmInterface)(nil).Update))
}

// SetDescription mocks base method
func (m *MockTqdmInterface) SetDescription(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetDescription", arg0)
}

// SetDescription indicates an expected call of SetDescription
func (mr *MockTqdmInterfaceMockRecorder) SetDescription(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDescription", reflect.TypeOf((*MockTqdmInterface)(nil).SetDescription), arg0)
}

// GetDescription mocks base method
func (m *MockTqdmInterface) GetDescription() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDescription indicates an expected call of GetDescription
func (mr *MockTqdmInterfaceMockRecorder) GetDescription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDescription", reflect.TypeOf((*MockTqdmInterface)(nil).GetDescription))
}

// SetRetain mocks base method
func (m *MockTqdmInterface) SetRetain(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetRetain", arg0)
}

// SetRetain indicates an expected call of SetRetain
func (mr *MockTqdmInterfaceMockRecorder) SetRetain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRetain", reflect.TypeOf((*MockTqdmInterface)(nil).SetRetain), arg0)
}

// GetRetain mocks base method
func (m *MockTqdmInterface) GetRetain() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRetain")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetRetain indicates an expected call of GetRetain
func (mr *MockTqdmInterfaceMockRecorder) GetRetain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRetain", reflect.TypeOf((*MockTqdmInterface)(nil).GetRetain))
}

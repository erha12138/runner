// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cox96de/runner/testtool (interfaces: TestingT)
//
// Generated by this command:
//
//	mockgen -destination mock/mockgen_tb.go -package mock . TestingT
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTestingT is a mock of TestingT interface.
type MockTestingT struct {
	ctrl     *gomock.Controller
	recorder *MockTestingTMockRecorder
}

// MockTestingTMockRecorder is the mock recorder for MockTestingT.
type MockTestingTMockRecorder struct {
	mock *MockTestingT
}

// NewMockTestingT creates a new mock instance.
func NewMockTestingT(ctrl *gomock.Controller) *MockTestingT {
	mock := &MockTestingT{ctrl: ctrl}
	mock.recorder = &MockTestingTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestingT) EXPECT() *MockTestingTMockRecorder {
	return m.recorder
}

// FailNow mocks base method.
func (m *MockTestingT) FailNow() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FailNow")
}

// FailNow indicates an expected call of FailNow.
func (mr *MockTestingTMockRecorder) FailNow() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailNow", reflect.TypeOf((*MockTestingT)(nil).FailNow))
}

// Helper mocks base method.
func (m *MockTestingT) Helper() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Helper")
}

// Helper indicates an expected call of Helper.
func (mr *MockTestingTMockRecorder) Helper() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Helper", reflect.TypeOf((*MockTestingT)(nil).Helper))
}

// Logf mocks base method.
func (m *MockTestingT) Logf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Logf", varargs...)
}

// Logf indicates an expected call of Logf.
func (mr *MockTestingTMockRecorder) Logf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logf", reflect.TypeOf((*MockTestingT)(nil).Logf), varargs...)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/kubelogin/pkg/token (interfaces: ExecCredentialWriter)

// Package mock_token is a generated GoMock package.
package mock_token

import (
	io "io"
	reflect "reflect"

	adal "github.com/Azure/go-autorest/autorest/adal"
	gomock "github.com/golang/mock/gomock"
)

// MockExecCredentialWriter is a mock of ExecCredentialWriter interface.
type MockExecCredentialWriter struct {
	ctrl     *gomock.Controller
	recorder *MockExecCredentialWriterMockRecorder
}

// MockExecCredentialWriterMockRecorder is the mock recorder for MockExecCredentialWriter.
type MockExecCredentialWriterMockRecorder struct {
	mock *MockExecCredentialWriter
}

// NewMockExecCredentialWriter creates a new mock instance.
func NewMockExecCredentialWriter(ctrl *gomock.Controller) *MockExecCredentialWriter {
	mock := &MockExecCredentialWriter{ctrl: ctrl}
	mock.recorder = &MockExecCredentialWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecCredentialWriter) EXPECT() *MockExecCredentialWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockExecCredentialWriter) Write(arg0 adal.Token, arg1 io.Writer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockExecCredentialWriterMockRecorder) Write(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockExecCredentialWriter)(nil).Write), arg0, arg1)
}

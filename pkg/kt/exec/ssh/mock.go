// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/kt/exec/ssh/types.go

// Package ssh is a generated GoMock package.
package ssh

import (
	exec "os/exec"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCliInterface is a mock of CliInterface interface.
type MockCliInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCliInterfaceMockRecorder
}

// MockCliInterfaceMockRecorder is the mock recorder for MockCliInterface.
type MockCliInterfaceMockRecorder struct {
	mock *MockCliInterface
}

// NewMockCliInterface creates a new mock instance.
func NewMockCliInterface(ctrl *gomock.Controller) *MockCliInterface {
	mock := &MockCliInterface{ctrl: ctrl}
	mock.recorder = &MockCliInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCliInterface) EXPECT() *MockCliInterfaceMockRecorder {
	return m.recorder
}

// TunnelToRemote mocks base method.
func (m *MockCliInterface) TunnelToRemote(localTun int, remoteHost, privateKeyPath string, remoteSSHPort int) *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TunnelToRemote", localTun, remoteHost, privateKeyPath, remoteSSHPort)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// TunnelToRemote indicates an expected call of TunnelToRemote.
func (mr *MockCliInterfaceMockRecorder) TunnelToRemote(localTun, remoteHost, privateKeyPath, remoteSSHPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TunnelToRemote", reflect.TypeOf((*MockCliInterface)(nil).TunnelToRemote), localTun, remoteHost, privateKeyPath, remoteSSHPort)
}

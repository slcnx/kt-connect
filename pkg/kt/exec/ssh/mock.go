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

// DynamicForwardLocalRequestToRemote mocks base method.
func (m *MockCliInterface) DynamicForwardLocalRequestToRemote(remoteHost, privateKeyPath string, remoteSSHPort, proxyPort int) *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DynamicForwardLocalRequestToRemote", remoteHost, privateKeyPath, remoteSSHPort, proxyPort)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// DynamicForwardLocalRequestToRemote indicates an expected call of DynamicForwardLocalRequestToRemote.
func (mr *MockCliInterfaceMockRecorder) DynamicForwardLocalRequestToRemote(remoteHost, privateKeyPath, remoteSSHPort, proxyPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DynamicForwardLocalRequestToRemote", reflect.TypeOf((*MockCliInterface)(nil).DynamicForwardLocalRequestToRemote), remoteHost, privateKeyPath, remoteSSHPort, proxyPort)
}

// ForwardRemoteRequestToLocal mocks base method.
func (m *MockCliInterface) ForwardRemoteRequestToLocal(localPort, remoteHost, remotePort, privateKeyPath string, remoteSSHPort int) *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForwardRemoteRequestToLocal", localPort, remoteHost, remotePort, privateKeyPath, remoteSSHPort)
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// ForwardRemoteRequestToLocal indicates an expected call of ForwardRemoteRequestToLocal.
func (mr *MockCliInterfaceMockRecorder) ForwardRemoteRequestToLocal(localPort, remoteHost, remotePort, privateKeyPath, remoteSSHPort interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardRemoteRequestToLocal", reflect.TypeOf((*MockCliInterface)(nil).ForwardRemoteRequestToLocal), localPort, remoteHost, remotePort, privateKeyPath, remoteSSHPort)
}

// Version mocks base method.
func (m *MockCliInterface) Version() *exec.Cmd {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*exec.Cmd)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockCliInterfaceMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockCliInterface)(nil).Version))
}

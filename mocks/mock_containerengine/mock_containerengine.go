// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/nitrictech/cli/pkg/containerengine (interfaces: ContainerEngine)

// Package mock_containerengine is a generated GoMock package.
package mock_containerengine

import (
	io "io"
	reflect "reflect"
	time "time"

	types "github.com/docker/docker/api/types"
	container "github.com/docker/docker/api/types/container"
	network "github.com/docker/docker/api/types/network"
	gomock "github.com/golang/mock/gomock"
	containerengine "github.com/nitrictech/cli/pkg/containerengine"
)

// MockContainerEngine is a mock of ContainerEngine interface.
type MockContainerEngine struct {
	ctrl     *gomock.Controller
	recorder *MockContainerEngineMockRecorder
}

// MockContainerEngineMockRecorder is the mock recorder for MockContainerEngine.
type MockContainerEngineMockRecorder struct {
	mock *MockContainerEngine
}

// NewMockContainerEngine creates a new mock instance.
func NewMockContainerEngine(ctrl *gomock.Controller) *MockContainerEngine {
	mock := &MockContainerEngine{ctrl: ctrl}
	mock.recorder = &MockContainerEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainerEngine) EXPECT() *MockContainerEngineMockRecorder {
	return m.recorder
}

// Build mocks base method.
func (m *MockContainerEngine) Build(arg0, arg1, arg2 string, arg3 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Build", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Build indicates an expected call of Build.
func (mr *MockContainerEngineMockRecorder) Build(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Build", reflect.TypeOf((*MockContainerEngine)(nil).Build), arg0, arg1, arg2, arg3)
}

// ContainerCreate mocks base method.
func (m *MockContainerEngine) ContainerCreate(arg0 *container.Config, arg1 *container.HostConfig, arg2 *network.NetworkingConfig, arg3 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerCreate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerCreate indicates an expected call of ContainerCreate.
func (mr *MockContainerEngineMockRecorder) ContainerCreate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerCreate", reflect.TypeOf((*MockContainerEngine)(nil).ContainerCreate), arg0, arg1, arg2, arg3)
}

// ContainerExec mocks base method.
func (m *MockContainerEngine) ContainerExec(arg0 string, arg1 []string, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerExec", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ContainerExec indicates an expected call of ContainerExec.
func (mr *MockContainerEngineMockRecorder) ContainerExec(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerExec", reflect.TypeOf((*MockContainerEngine)(nil).ContainerExec), arg0, arg1, arg2)
}

// ContainerLogs mocks base method.
func (m *MockContainerEngine) ContainerLogs(arg0 string, arg1 types.ContainerLogsOptions) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerLogs", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerLogs indicates an expected call of ContainerLogs.
func (mr *MockContainerEngineMockRecorder) ContainerLogs(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerLogs", reflect.TypeOf((*MockContainerEngine)(nil).ContainerLogs), arg0, arg1)
}

// ContainerWait mocks base method.
func (m *MockContainerEngine) ContainerWait(arg0 string, arg1 container.WaitCondition) (<-chan container.ContainerWaitOKBody, <-chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerWait", arg0, arg1)
	ret0, _ := ret[0].(<-chan container.ContainerWaitOKBody)
	ret1, _ := ret[1].(<-chan error)
	return ret0, ret1
}

// ContainerWait indicates an expected call of ContainerWait.
func (mr *MockContainerEngineMockRecorder) ContainerWait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerWait", reflect.TypeOf((*MockContainerEngine)(nil).ContainerWait), arg0, arg1)
}

// ContainersListByLabel mocks base method.
func (m *MockContainerEngine) ContainersListByLabel(arg0 map[string]string) ([]types.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainersListByLabel", arg0)
	ret0, _ := ret[0].([]types.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainersListByLabel indicates an expected call of ContainersListByLabel.
func (mr *MockContainerEngineMockRecorder) ContainersListByLabel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainersListByLabel", reflect.TypeOf((*MockContainerEngine)(nil).ContainersListByLabel), arg0)
}

// CopyFromArchive mocks base method.
func (m *MockContainerEngine) CopyFromArchive(arg0, arg1 string, arg2 io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyFromArchive", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyFromArchive indicates an expected call of CopyFromArchive.
func (mr *MockContainerEngineMockRecorder) CopyFromArchive(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyFromArchive", reflect.TypeOf((*MockContainerEngine)(nil).CopyFromArchive), arg0, arg1, arg2)
}

// ImagePull mocks base method.
func (m *MockContainerEngine) ImagePull(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImagePull", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImagePull indicates an expected call of ImagePull.
func (mr *MockContainerEngineMockRecorder) ImagePull(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImagePull", reflect.TypeOf((*MockContainerEngine)(nil).ImagePull), arg0)
}

// ListImages mocks base method.
func (m *MockContainerEngine) ListImages(arg0, arg1 string) ([]containerengine.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].([]containerengine.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages.
func (mr *MockContainerEngineMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockContainerEngine)(nil).ListImages), arg0, arg1)
}

// Logger mocks base method.
func (m *MockContainerEngine) Logger(arg0 string) containerengine.ContainerLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger", arg0)
	ret0, _ := ret[0].(containerengine.ContainerLogger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockContainerEngineMockRecorder) Logger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockContainerEngine)(nil).Logger), arg0)
}

// NetworkCreate mocks base method.
func (m *MockContainerEngine) NetworkCreate(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkCreate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetworkCreate indicates an expected call of NetworkCreate.
func (mr *MockContainerEngineMockRecorder) NetworkCreate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkCreate", reflect.TypeOf((*MockContainerEngine)(nil).NetworkCreate), arg0)
}

// RemoveByLabel mocks base method.
func (m *MockContainerEngine) RemoveByLabel(arg0 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveByLabel", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveByLabel indicates an expected call of RemoveByLabel.
func (mr *MockContainerEngineMockRecorder) RemoveByLabel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByLabel", reflect.TypeOf((*MockContainerEngine)(nil).RemoveByLabel), arg0)
}

// Start mocks base method.
func (m *MockContainerEngine) Start(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockContainerEngineMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockContainerEngine)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockContainerEngine) Stop(arg0 string, arg1 *time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockContainerEngineMockRecorder) Stop(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockContainerEngine)(nil).Stop), arg0, arg1)
}

// Type mocks base method.
func (m *MockContainerEngine) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockContainerEngineMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockContainerEngine)(nil).Type))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/vmware-tanzu/octant/pkg/plugin/api (interfaces: Service)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	action "github.com/vmware-tanzu/octant/pkg/action"
	event "github.com/vmware-tanzu/octant/pkg/event"
	api "github.com/vmware-tanzu/octant/pkg/plugin/api"
	store "github.com/vmware-tanzu/octant/pkg/store"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CancelPortForward mocks base method.
func (m *MockService) CancelPortForward(arg0 context.Context, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelPortForward", arg0, arg1)
}

// CancelPortForward indicates an expected call of CancelPortForward.
func (mr *MockServiceMockRecorder) CancelPortForward(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelPortForward", reflect.TypeOf((*MockService)(nil).CancelPortForward), arg0, arg1)
}

// Create mocks base method.
func (m *MockService) Create(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), arg0, arg1)
}

// CreateLink mocks base method.
func (m *MockService) CreateLink(arg0 context.Context, arg1 store.Key) (api.LinkResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLink", arg0, arg1)
	ret0, _ := ret[0].(api.LinkResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateLink indicates an expected call of CreateLink.
func (mr *MockServiceMockRecorder) CreateLink(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLink", reflect.TypeOf((*MockService)(nil).CreateLink), arg0, arg1)
}

// Delete mocks base method.
func (m *MockService) Delete(arg0 context.Context, arg1 store.Key) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), arg0, arg1)
}

// ForceFrontendUpdate mocks base method.
func (m *MockService) ForceFrontendUpdate(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForceFrontendUpdate", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForceFrontendUpdate indicates an expected call of ForceFrontendUpdate.
func (mr *MockServiceMockRecorder) ForceFrontendUpdate(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForceFrontendUpdate", reflect.TypeOf((*MockService)(nil).ForceFrontendUpdate), arg0)
}

// Get mocks base method.
func (m *MockService) Get(arg0 context.Context, arg1 store.Key) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockService)(nil).Get), arg0, arg1)
}

// List mocks base method.
func (m *MockService) List(arg0 context.Context, arg1 store.Key) (*unstructured.UnstructuredList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockService)(nil).List), arg0, arg1)
}

// ListNamespaces mocks base method.
func (m *MockService) ListNamespaces(arg0 context.Context) (api.NamespacesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNamespaces", arg0)
	ret0, _ := ret[0].(api.NamespacesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNamespaces indicates an expected call of ListNamespaces.
func (mr *MockServiceMockRecorder) ListNamespaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNamespaces", reflect.TypeOf((*MockService)(nil).ListNamespaces), arg0)
}

// PortForward mocks base method.
func (m *MockService) PortForward(arg0 context.Context, arg1 api.PortForwardRequest) (api.PortForwardResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PortForward", arg0, arg1)
	ret0, _ := ret[0].(api.PortForwardResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PortForward indicates an expected call of PortForward.
func (mr *MockServiceMockRecorder) PortForward(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PortForward", reflect.TypeOf((*MockService)(nil).PortForward), arg0, arg1)
}

// SendAlert mocks base method.
func (m *MockService) SendAlert(arg0 context.Context, arg1 string, arg2 action.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAlert", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAlert indicates an expected call of SendAlert.
func (mr *MockServiceMockRecorder) SendAlert(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAlert", reflect.TypeOf((*MockService)(nil).SendAlert), arg0, arg1, arg2)
}

// SendEvent mocks base method.
func (m *MockService) SendEvent(arg0 context.Context, arg1 string, arg2 event.EventType, arg3 action.Payload) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendEvent", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendEvent indicates an expected call of SendEvent.
func (mr *MockServiceMockRecorder) SendEvent(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendEvent", reflect.TypeOf((*MockService)(nil).SendEvent), arg0, arg1, arg2, arg3)
}

// Update mocks base method.
func (m *MockService) Update(arg0 context.Context, arg1 *unstructured.Unstructured) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), arg0, arg1)
}

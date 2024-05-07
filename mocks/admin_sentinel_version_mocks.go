// Code generated by MockGen. DO NOT EDIT.
// Source: admin_sentinel_version.go
//
// Generated by this command:
//
//	mockgen -source=admin_sentinel_version.go -destination=mocks/admin_sentinel_version_mocks.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	tfe "github.com/hashicorp/go-tfe"
	gomock "go.uber.org/mock/gomock"
)

// MockAdminSentinelVersions is a mock of AdminSentinelVersions interface.
type MockAdminSentinelVersions struct {
	ctrl     *gomock.Controller
	recorder *MockAdminSentinelVersionsMockRecorder
}

// MockAdminSentinelVersionsMockRecorder is the mock recorder for MockAdminSentinelVersions.
type MockAdminSentinelVersionsMockRecorder struct {
	mock *MockAdminSentinelVersions
}

// NewMockAdminSentinelVersions creates a new mock instance.
func NewMockAdminSentinelVersions(ctrl *gomock.Controller) *MockAdminSentinelVersions {
	mock := &MockAdminSentinelVersions{ctrl: ctrl}
	mock.recorder = &MockAdminSentinelVersionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminSentinelVersions) EXPECT() *MockAdminSentinelVersionsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAdminSentinelVersions) Create(ctx context.Context, options tfe.AdminSentinelVersionCreateOptions) (*tfe.AdminSentinelVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminSentinelVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdminSentinelVersionsMockRecorder) Create(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdminSentinelVersions)(nil).Create), ctx, options)
}

// Delete mocks base method.
func (m *MockAdminSentinelVersions) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminSentinelVersionsMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminSentinelVersions)(nil).Delete), ctx, id)
}

// List mocks base method.
func (m *MockAdminSentinelVersions) List(ctx context.Context, options *tfe.AdminSentinelVersionsListOptions) (*tfe.AdminSentinelVersionsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].(*tfe.AdminSentinelVersionsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAdminSentinelVersionsMockRecorder) List(ctx, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAdminSentinelVersions)(nil).List), ctx, options)
}

// Read mocks base method.
func (m *MockAdminSentinelVersions) Read(ctx context.Context, id string) (*tfe.AdminSentinelVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, id)
	ret0, _ := ret[0].(*tfe.AdminSentinelVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockAdminSentinelVersionsMockRecorder) Read(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockAdminSentinelVersions)(nil).Read), ctx, id)
}

// Update mocks base method.
func (m *MockAdminSentinelVersions) Update(ctx context.Context, id string, options tfe.AdminSentinelVersionUpdateOptions) (*tfe.AdminSentinelVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, options)
	ret0, _ := ret[0].(*tfe.AdminSentinelVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAdminSentinelVersionsMockRecorder) Update(ctx, id, options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminSentinelVersions)(nil).Update), ctx, id, options)
}

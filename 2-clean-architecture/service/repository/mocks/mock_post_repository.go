// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/service (interfaces: PostRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/wnfrx/unit-testing-sharing-session/2-clean-architecture/models"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockPostRepository) Add(arg0 context.Context, arg1 models.Post) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockPostRepositoryMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPostRepository)(nil).Add), arg0, arg1)
}

// Delete mocks base method.
func (m *MockPostRepository) Delete(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPostRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostRepository)(nil).Delete), arg0, arg1)
}

// Get mocks base method.
func (m *MockPostRepository) Get(arg0 context.Context, arg1 int64) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPostRepositoryMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostRepository)(nil).Get), arg0, arg1)
}

// GetByAuthor mocks base method.
func (m *MockPostRepository) GetByAuthor(arg0 context.Context, arg1 string) (models.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAuthor", arg0, arg1)
	ret0, _ := ret[0].(models.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAuthor indicates an expected call of GetByAuthor.
func (mr *MockPostRepositoryMockRecorder) GetByAuthor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAuthor", reflect.TypeOf((*MockPostRepository)(nil).GetByAuthor), arg0, arg1)
}

// Update mocks base method.
func (m *MockPostRepository) Update(arg0 context.Context, arg1 models.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPostRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPostRepository)(nil).Update), arg0, arg1)
}

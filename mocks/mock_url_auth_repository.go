// Code generated by MockGen. DO NOT EDIT.
// Source: koc-digital-case/services (interfaces: URLAuthRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "koc-digital-case/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockURLAuthRepository is a mock of URLAuthRepository interface.
type MockURLAuthRepository struct {
	ctrl     *gomock.Controller
	recorder *MockURLAuthRepositoryMockRecorder
}

// MockURLAuthRepositoryMockRecorder is the mock recorder for MockURLAuthRepository.
type MockURLAuthRepositoryMockRecorder struct {
	mock *MockURLAuthRepository
}

// NewMockURLAuthRepository creates a new mock instance.
func NewMockURLAuthRepository(ctrl *gomock.Controller) *MockURLAuthRepository {
	mock := &MockURLAuthRepository{ctrl: ctrl}
	mock.recorder = &MockURLAuthRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLAuthRepository) EXPECT() *MockURLAuthRepositoryMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockURLAuthRepository) GetUser(arg0 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockURLAuthRepositoryMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockURLAuthRepository)(nil).GetUser), arg0)
}
// Code generated by MockGen. DO NOT EDIT.
// Source: koc-digital-case/services (interfaces: URLRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "koc-digital-case/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockURLRepository is a mock of URLRepository interface.
type MockURLRepository struct {
	ctrl     *gomock.Controller
	recorder *MockURLRepositoryMockRecorder
}

// MockURLRepositoryMockRecorder is the mock recorder for MockURLRepository.
type MockURLRepositoryMockRecorder struct {
	mock *MockURLRepository
}

// NewMockURLRepository creates a new mock instance.
func NewMockURLRepository(ctrl *gomock.Controller) *MockURLRepository {
	mock := &MockURLRepository{ctrl: ctrl}
	mock.recorder = &MockURLRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockURLRepository) EXPECT() *MockURLRepositoryMockRecorder {
	return m.recorder
}

// DeleteURL mocks base method.
func (m *MockURLRepository) DeleteURL(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteURL", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteURL indicates an expected call of DeleteURL.
func (mr *MockURLRepositoryMockRecorder) DeleteURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteURL", reflect.TypeOf((*MockURLRepository)(nil).DeleteURL), arg0)
}

// GetShortenedURL mocks base method.
func (m *MockURLRepository) GetShortenedURL(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortenedURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShortenedURL indicates an expected call of GetShortenedURL.
func (mr *MockURLRepositoryMockRecorder) GetShortenedURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortenedURL", reflect.TypeOf((*MockURLRepository)(nil).GetShortenedURL), arg0)
}

// GetURL mocks base method.
func (m *MockURLRepository) GetURL(arg0 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURL", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURL indicates an expected call of GetURL.
func (mr *MockURLRepositoryMockRecorder) GetURL(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURL", reflect.TypeOf((*MockURLRepository)(nil).GetURL), arg0)
}

// GetURLs mocks base method.
func (m *MockURLRepository) GetURLs() ([]models.URLMapping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetURLs")
	ret0, _ := ret[0].([]models.URLMapping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetURLs indicates an expected call of GetURLs.
func (mr *MockURLRepositoryMockRecorder) GetURLs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetURLs", reflect.TypeOf((*MockURLRepository)(nil).GetURLs))
}

// GetUserRemainingBenefits mocks base method.
func (m *MockURLRepository) GetUserRemainingBenefits(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserRemainingBenefits", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserRemainingBenefits indicates an expected call of GetUserRemainingBenefits.
func (mr *MockURLRepositoryMockRecorder) GetUserRemainingBenefits(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserRemainingBenefits", reflect.TypeOf((*MockURLRepository)(nil).GetUserRemainingBenefits), arg0)
}

// StoreURLMapping mocks base method.
func (m *MockURLRepository) StoreURLMapping(arg0 int, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreURLMapping", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreURLMapping indicates an expected call of StoreURLMapping.
func (mr *MockURLRepositoryMockRecorder) StoreURLMapping(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreURLMapping", reflect.TypeOf((*MockURLRepository)(nil).StoreURLMapping), arg0, arg1, arg2, arg3)
}

// UpdateUserUsage mocks base method.
func (m *MockURLRepository) UpdateUserUsage(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserUsage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserUsage indicates an expected call of UpdateUserUsage.
func (mr *MockURLRepositoryMockRecorder) UpdateUserUsage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserUsage", reflect.TypeOf((*MockURLRepository)(nil).UpdateUserUsage), arg0, arg1)
}

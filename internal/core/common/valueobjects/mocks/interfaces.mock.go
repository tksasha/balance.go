// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/common/valueobjects/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/core/common/valueobjects/interfaces.go -package mocks -destination internal/core/common/valueobjects/mocks/interfaces.mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCurrentDateProvider is a mock of CurrentDateProvider interface.
type MockCurrentDateProvider struct {
	ctrl     *gomock.Controller
	recorder *MockCurrentDateProviderMockRecorder
	isgomock struct{}
}

// MockCurrentDateProviderMockRecorder is the mock recorder for MockCurrentDateProvider.
type MockCurrentDateProviderMockRecorder struct {
	mock *MockCurrentDateProvider
}

// NewMockCurrentDateProvider creates a new mock instance.
func NewMockCurrentDateProvider(ctrl *gomock.Controller) *MockCurrentDateProvider {
	mock := &MockCurrentDateProvider{ctrl: ctrl}
	mock.recorder = &MockCurrentDateProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrentDateProvider) EXPECT() *MockCurrentDateProviderMockRecorder {
	return m.recorder
}

// CurrentMonth mocks base method.
func (m *MockCurrentDateProvider) CurrentMonth() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentMonth")
	ret0, _ := ret[0].(int)
	return ret0
}

// CurrentMonth indicates an expected call of CurrentMonth.
func (mr *MockCurrentDateProviderMockRecorder) CurrentMonth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentMonth", reflect.TypeOf((*MockCurrentDateProvider)(nil).CurrentMonth))
}

// CurrentYear mocks base method.
func (m *MockCurrentDateProvider) CurrentYear() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentYear")
	ret0, _ := ret[0].(int)
	return ret0
}

// CurrentYear indicates an expected call of CurrentYear.
func (mr *MockCurrentDateProviderMockRecorder) CurrentYear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentYear", reflect.TypeOf((*MockCurrentDateProvider)(nil).CurrentYear))
}

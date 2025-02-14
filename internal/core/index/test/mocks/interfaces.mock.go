// Code generated by MockGen. DO NOT EDIT.
// Source: internal/core/index/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/core/index/interfaces.go -package mocks -destination internal/core/index/test/mocks/interfaces.mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
	isgomock struct{}
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Expense mocks base method.
func (m *MockRepository) Expense(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expense", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Expense indicates an expected call of Expense.
func (mr *MockRepositoryMockRecorder) Expense(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expense", reflect.TypeOf((*MockRepository)(nil).Expense), ctx)
}

// Income mocks base method.
func (m *MockRepository) Income(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Income", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Income indicates an expected call of Income.
func (mr *MockRepositoryMockRecorder) Income(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Income", reflect.TypeOf((*MockRepository)(nil).Income), ctx)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
	isgomock struct{}
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

// Balance mocks base method.
func (m *MockService) Balance(ctx context.Context) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balance", ctx)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balance indicates an expected call of Balance.
func (mr *MockServiceMockRecorder) Balance(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockService)(nil).Balance), ctx)
}

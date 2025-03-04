// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/categoryreport/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/app/categoryreport/interfaces.go -package mocks -destination internal/app/categoryreport/test/mocks/interfaces.mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	categoryreport "github.com/tksasha/balance/internal/app/categoryreport"
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

// Group mocks base method.
func (m *MockRepository) Group(ctx context.Context, filters categoryreport.Filters) (categoryreport.Entities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", ctx, filters)
	ret0, _ := ret[0].(categoryreport.Entities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Group indicates an expected call of Group.
func (mr *MockRepositoryMockRecorder) Group(ctx, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockRepository)(nil).Group), ctx, filters)
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

// Report mocks base method.
func (m *MockService) Report(ctx context.Context, request categoryreport.Request) (categoryreport.Entities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Report", ctx, request)
	ret0, _ := ret[0].(categoryreport.Entities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Report indicates an expected call of Report.
func (mr *MockServiceMockRecorder) Report(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Report", reflect.TypeOf((*MockService)(nil).Report), ctx, request)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/category_repository.go
//
// Generated by this command:
//
//	mockgen -source internal/repositories/category_repository.go -package mockedrepositories -destination mocks/repositories/category_repository.mock.go
//

// Package mockedrepositories is a generated GoMock package.
package mockedrepositories

import (
	context "context"
	reflect "reflect"

	models "github.com/tksasha/balance/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
	isgomock struct{}
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockCategoryRepository) GetCategories(ctx context.Context) (models.Categories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", ctx)
	ret0, _ := ret[0].(models.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockCategoryRepositoryMockRecorder) GetCategories(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategories), ctx)
}
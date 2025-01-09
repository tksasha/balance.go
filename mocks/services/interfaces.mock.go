// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/services/interfaces.go -package mocksforservices -destination mocks/services/interfaces.mock.go
//

// Package mocksforservices is a generated GoMock package.
package mocksforservices

import (
	context "context"
	reflect "reflect"

	models "github.com/tksasha/balance/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockItemRepository is a mock of ItemRepository interface.
type MockItemRepository struct {
	ctrl     *gomock.Controller
	recorder *MockItemRepositoryMockRecorder
	isgomock struct{}
}

// MockItemRepositoryMockRecorder is the mock recorder for MockItemRepository.
type MockItemRepositoryMockRecorder struct {
	mock *MockItemRepository
}

// NewMockItemRepository creates a new mock instance.
func NewMockItemRepository(ctrl *gomock.Controller) *MockItemRepository {
	mock := &MockItemRepository{ctrl: ctrl}
	mock.recorder = &MockItemRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemRepository) EXPECT() *MockItemRepositoryMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockItemRepository) CreateItem(ctx context.Context, item *models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockItemRepositoryMockRecorder) CreateItem(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockItemRepository)(nil).CreateItem), ctx, item)
}

// DeleteItem mocks base method.
func (m *MockItemRepository) DeleteItem(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockItemRepositoryMockRecorder) DeleteItem(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockItemRepository)(nil).DeleteItem), ctx, id)
}

// GetItem mocks base method.
func (m *MockItemRepository) GetItem(ctx context.Context, id int) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, id)
	ret0, _ := ret[0].(*models.Item)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockItemRepositoryMockRecorder) GetItem(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockItemRepository)(nil).GetItem), ctx, id)
}

// GetItems mocks base method.
func (m *MockItemRepository) GetItems(ctx context.Context) (models.Items, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ctx)
	ret0, _ := ret[0].(models.Items)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockItemRepositoryMockRecorder) GetItems(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockItemRepository)(nil).GetItems), ctx)
}

// UpdateItem mocks base method.
func (m *MockItemRepository) UpdateItem(ctx context.Context, item *models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockItemRepositoryMockRecorder) UpdateItem(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockItemRepository)(nil).UpdateItem), ctx, item)
}

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

// Create mocks base method.
func (m *MockCategoryRepository) Create(ctx context.Context, category *models.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCategoryRepositoryMockRecorder) Create(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategoryRepository)(nil).Create), ctx, category)
}

// FindByID mocks base method.
func (m *MockCategoryRepository) FindByID(ctx context.Context, id int) (*models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockCategoryRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockCategoryRepository)(nil).FindByID), ctx, id)
}

// FindByName mocks base method.
func (m *MockCategoryRepository) FindByName(ctx context.Context, name string) (*models.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", ctx, name)
	ret0, _ := ret[0].(*models.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockCategoryRepositoryMockRecorder) FindByName(ctx, name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockCategoryRepository)(nil).FindByName), ctx, name)
}

// GetAll mocks base method.
func (m *MockCategoryRepository) GetAll(ctx context.Context) (models.Categories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].(models.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockCategoryRepositoryMockRecorder) GetAll(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockCategoryRepository)(nil).GetAll), ctx)
}

// Update mocks base method.
func (m *MockCategoryRepository) Update(ctx context.Context, category *models.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCategoryRepositoryMockRecorder) Update(ctx, category any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategoryRepository)(nil).Update), ctx, category)
}

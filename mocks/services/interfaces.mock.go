// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/services/interfaces.go -package mockedservices -destination mocks/services/interfaces.mock.go
//

// Package mockedservices is a generated GoMock package.
package mockedservices

import (
	context "context"
	reflect "reflect"

	decorators "github.com/tksasha/balance/internal/decorators"
	models "github.com/tksasha/balance/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockItemsGetter is a mock of ItemsGetter interface.
type MockItemsGetter struct {
	ctrl     *gomock.Controller
	recorder *MockItemsGetterMockRecorder
}

// MockItemsGetterMockRecorder is the mock recorder for MockItemsGetter.
type MockItemsGetterMockRecorder struct {
	mock *MockItemsGetter
}

// NewMockItemsGetter creates a new mock instance.
func NewMockItemsGetter(ctrl *gomock.Controller) *MockItemsGetter {
	mock := &MockItemsGetter{ctrl: ctrl}
	mock.recorder = &MockItemsGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemsGetter) EXPECT() *MockItemsGetterMockRecorder {
	return m.recorder
}

// GetItems mocks base method.
func (m *MockItemsGetter) GetItems(ctx context.Context, currency models.Currency) (*decorators.ItemsDecorator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItems", ctx, currency)
	ret0, _ := ret[0].(*decorators.ItemsDecorator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItems indicates an expected call of GetItems.
func (mr *MockItemsGetterMockRecorder) GetItems(ctx, currency any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItems", reflect.TypeOf((*MockItemsGetter)(nil).GetItems), ctx, currency)
}

// MockItemGetter is a mock of ItemGetter interface.
type MockItemGetter struct {
	ctrl     *gomock.Controller
	recorder *MockItemGetterMockRecorder
}

// MockItemGetterMockRecorder is the mock recorder for MockItemGetter.
type MockItemGetterMockRecorder struct {
	mock *MockItemGetter
}

// NewMockItemGetter creates a new mock instance.
func NewMockItemGetter(ctrl *gomock.Controller) *MockItemGetter {
	mock := &MockItemGetter{ctrl: ctrl}
	mock.recorder = &MockItemGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemGetter) EXPECT() *MockItemGetterMockRecorder {
	return m.recorder
}

// GetItem mocks base method.
func (m *MockItemGetter) GetItem(ctx context.Context, id string) (*decorators.ItemDecorator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, id)
	ret0, _ := ret[0].(*decorators.ItemDecorator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockItemGetterMockRecorder) GetItem(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockItemGetter)(nil).GetItem), ctx, id)
}

// MockItemUpdater is a mock of ItemUpdater interface.
type MockItemUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockItemUpdaterMockRecorder
}

// MockItemUpdaterMockRecorder is the mock recorder for MockItemUpdater.
type MockItemUpdaterMockRecorder struct {
	mock *MockItemUpdater
}

// NewMockItemUpdater creates a new mock instance.
func NewMockItemUpdater(ctrl *gomock.Controller) *MockItemUpdater {
	mock := &MockItemUpdater{ctrl: ctrl}
	mock.recorder = &MockItemUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemUpdater) EXPECT() *MockItemUpdaterMockRecorder {
	return m.recorder
}

// UpdateItem mocks base method.
func (m *MockItemUpdater) UpdateItem(ctx context.Context, item *models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateItem", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockItemUpdaterMockRecorder) UpdateItem(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockItemUpdater)(nil).UpdateItem), ctx, item)
}

// MockItemDeleter is a mock of ItemDeleter interface.
type MockItemDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockItemDeleterMockRecorder
}

// MockItemDeleterMockRecorder is the mock recorder for MockItemDeleter.
type MockItemDeleterMockRecorder struct {
	mock *MockItemDeleter
}

// NewMockItemDeleter creates a new mock instance.
func NewMockItemDeleter(ctrl *gomock.Controller) *MockItemDeleter {
	mock := &MockItemDeleter{ctrl: ctrl}
	mock.recorder = &MockItemDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemDeleter) EXPECT() *MockItemDeleterMockRecorder {
	return m.recorder
}

// DeleteItem mocks base method.
func (m *MockItemDeleter) DeleteItem(ctx context.Context, item *models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteItem", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockItemDeleterMockRecorder) DeleteItem(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockItemDeleter)(nil).DeleteItem), ctx, item)
}

// MockCategoriesGetter is a mock of CategoriesGetter interface.
type MockCategoriesGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCategoriesGetterMockRecorder
}

// MockCategoriesGetterMockRecorder is the mock recorder for MockCategoriesGetter.
type MockCategoriesGetterMockRecorder struct {
	mock *MockCategoriesGetter
}

// NewMockCategoriesGetter creates a new mock instance.
func NewMockCategoriesGetter(ctrl *gomock.Controller) *MockCategoriesGetter {
	mock := &MockCategoriesGetter{ctrl: ctrl}
	mock.recorder = &MockCategoriesGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoriesGetter) EXPECT() *MockCategoriesGetterMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockCategoriesGetter) GetCategories(ctx context.Context, currency int) (*decorators.CategoriesDecorator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", ctx, currency)
	ret0, _ := ret[0].(*decorators.CategoriesDecorator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockCategoriesGetterMockRecorder) GetCategories(ctx, currency any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategoriesGetter)(nil).GetCategories), ctx, currency)
}

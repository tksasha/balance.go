// Code generated by MockGen. DO NOT EDIT.
// Source: internal/interfaces/interfaces.go
//
// Generated by this command:
//
//	mockgen -source internal/interfaces/interfaces.go -package mocks -destination test/mocks/interfaces.mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/tksasha/balance/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockItemCreator is a mock of ItemCreator interface.
type MockItemCreator struct {
	ctrl     *gomock.Controller
	recorder *MockItemCreatorMockRecorder
}

// MockItemCreatorMockRecorder is the mock recorder for MockItemCreator.
type MockItemCreatorMockRecorder struct {
	mock *MockItemCreator
}

// NewMockItemCreator creates a new mock instance.
func NewMockItemCreator(ctrl *gomock.Controller) *MockItemCreator {
	mock := &MockItemCreator{ctrl: ctrl}
	mock.recorder = &MockItemCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockItemCreator) EXPECT() *MockItemCreatorMockRecorder {
	return m.recorder
}

// CreateItem mocks base method.
func (m *MockItemCreator) CreateItem(ctx context.Context, item *models.Item) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateItem", ctx, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateItem indicates an expected call of CreateItem.
func (mr *MockItemCreatorMockRecorder) CreateItem(ctx, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateItem", reflect.TypeOf((*MockItemCreator)(nil).CreateItem), ctx, item)
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
func (m *MockItemGetter) GetItem(ctx context.Context, id int) (*models.Item, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetItem", ctx, id)
	ret0, _ := ret[0].(*models.Item)
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

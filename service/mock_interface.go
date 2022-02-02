// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ridhdhish-desai-zs/product-gofr/service (interfaces: Product)

// Package service is a generated GoMock package.
package service

import (
	gofr "developer.zopsmart.com/go/gofr/pkg/gofr"
	gomock "github.com/golang/mock/gomock"
	models "github.com/ridhdhish-desai-zs/product-gofr/models"
	reflect "reflect"
)

// MockProduct is a mock of Product interface
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// GetById mocks base method
func (m *MockProduct) GetById(arg0 *gofr.Context, arg1 string) (*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockProductMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockProduct)(nil).GetById), arg0, arg1)
}

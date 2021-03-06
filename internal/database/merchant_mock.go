// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ferruvich/curve-prepaid-card/internal/database (interfaces: Merchant)

// Package database is a generated GoMock package.
package database

import (
	model "github.com/ferruvich/curve-prepaid-card/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMerchant is a mock of Merchant interface
type MockMerchant struct {
	ctrl     *gomock.Controller
	recorder *MockMerchantMockRecorder
}

// MockMerchantMockRecorder is the mock recorder for MockMerchant
type MockMerchantMockRecorder struct {
	mock *MockMerchant
}

// NewMockMerchant creates a new mock instance
func NewMockMerchant(ctrl *gomock.Controller) *MockMerchant {
	mock := &MockMerchant{ctrl: ctrl}
	mock.recorder = &MockMerchantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMerchant) EXPECT() *MockMerchantMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockMerchant) Write(arg0 *model.Merchant) error {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockMerchantMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockMerchant)(nil).Write), arg0)
}

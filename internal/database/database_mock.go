// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ferruvich/curve-prepaid-card/internal/database (interfaces: DataBase)

// Package database is a generated GoMock package.
package database

import (
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDataBase is a mock of DataBase interface
type MockDataBase struct {
	ctrl     *gomock.Controller
	recorder *MockDataBaseMockRecorder
}

// MockDataBaseMockRecorder is the mock recorder for MockDataBase
type MockDataBaseMockRecorder struct {
	mock *MockDataBase
}

// NewMockDataBase creates a new mock instance
func NewMockDataBase(ctrl *gomock.Controller) *MockDataBase {
	mock := &MockDataBase{ctrl: ctrl}
	mock.recorder = &MockDataBaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataBase) EXPECT() *MockDataBaseMockRecorder {
	return m.recorder
}

// AuthorizationRequest mocks base method
func (m *MockDataBase) AuthorizationRequest() AuthorizationRequest {
	ret := m.ctrl.Call(m, "AuthorizationRequest")
	ret0, _ := ret[0].(AuthorizationRequest)
	return ret0
}

// AuthorizationRequest indicates an expected call of AuthorizationRequest
func (mr *MockDataBaseMockRecorder) AuthorizationRequest() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationRequest", reflect.TypeOf((*MockDataBase)(nil).AuthorizationRequest))
}

// Card mocks base method
func (m *MockDataBase) Card() Card {
	ret := m.ctrl.Call(m, "Card")
	ret0, _ := ret[0].(Card)
	return ret0
}

// Card indicates an expected call of Card
func (mr *MockDataBaseMockRecorder) Card() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Card", reflect.TypeOf((*MockDataBase)(nil).Card))
}

// Merchant mocks base method
func (m *MockDataBase) Merchant() Merchant {
	ret := m.ctrl.Call(m, "Merchant")
	ret0, _ := ret[0].(Merchant)
	return ret0
}

// Merchant indicates an expected call of Merchant
func (mr *MockDataBaseMockRecorder) Merchant() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merchant", reflect.TypeOf((*MockDataBase)(nil).Merchant))
}

// User mocks base method
func (m *MockDataBase) User() User {
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(User)
	return ret0
}

// User indicates an expected call of User
func (mr *MockDataBaseMockRecorder) User() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockDataBase)(nil).User))
}

// newPipelineStmt mocks base method
func (m *MockDataBase) newPipelineStmt(arg0 string, arg1 ...interface{}) *pipelineStmt {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "newPipelineStmt", varargs...)
	ret0, _ := ret[0].(*pipelineStmt)
	return ret0
}

// newPipelineStmt indicates an expected call of newPipelineStmt
func (mr *MockDataBaseMockRecorder) newPipelineStmt(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "newPipelineStmt", reflect.TypeOf((*MockDataBase)(nil).newPipelineStmt), varargs...)
}

// runPipeline mocks base method
func (m *MockDataBase) runPipeline(arg0 transaction, arg1 ...*pipelineStmt) (*sql.Rows, error) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "runPipeline", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// runPipeline indicates an expected call of runPipeline
func (mr *MockDataBaseMockRecorder) runPipeline(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "runPipeline", reflect.TypeOf((*MockDataBase)(nil).runPipeline), varargs...)
}

// withTransaction mocks base method
func (m *MockDataBase) withTransaction(arg0 *sql.DB, arg1 func(transaction) (*sql.Rows, error)) (*sql.Rows, error) {
	ret := m.ctrl.Call(m, "withTransaction", arg0, arg1)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// withTransaction indicates an expected call of withTransaction
func (mr *MockDataBaseMockRecorder) withTransaction(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "withTransaction", reflect.TypeOf((*MockDataBase)(nil).withTransaction), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package balance is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	"test-example/balance"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetAllMovements mocks base method
func (m *MockClient) GetAllMovements(userId string) ([]*balance.Movement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllMovements", userId)
	ret0, _ := ret[0].([]*balance.Movement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllMovements indicates an expected call of GetAllMovements
func (mr *MockClientMockRecorder) GetAllMovements(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllMovements", reflect.TypeOf((*MockClient)(nil).GetAllMovements), userId)
}

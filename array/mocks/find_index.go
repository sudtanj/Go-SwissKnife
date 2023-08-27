// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces/find_index.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIFindIndex is a mock of IFindIndex interface.
type MockIFindIndex[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockIFindIndexMockRecorder[T]
}

// MockIFindIndexMockRecorder is the mock recorder for MockIFindIndex.
type MockIFindIndexMockRecorder[T any] struct {
	mock *MockIFindIndex[T]
}

// NewMockIFindIndex creates a new mock instance.
func NewMockIFindIndex[T any](ctrl *gomock.Controller) *MockIFindIndex[T] {
	mock := &MockIFindIndex[T]{ctrl: ctrl}
	mock.recorder = &MockIFindIndexMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFindIndex[T]) EXPECT() *MockIFindIndexMockRecorder[T] {
	return m.recorder
}

// FindIndex mocks base method.
func (m *MockIFindIndex[T]) FindIndex(values []T, compFunc func(T) bool) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIndex", values, compFunc)
	ret0, _ := ret[0].(int)
	return ret0
}

// FindIndex indicates an expected call of FindIndex.
func (mr *MockIFindIndexMockRecorder[T]) FindIndex(values, compFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIndex", reflect.TypeOf((*MockIFindIndex[T])(nil).FindIndex), values, compFunc)
}

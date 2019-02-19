// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpack/pack (interfaces: Cache)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCache is a mock of Cache interface
type MockCache struct {
	ctrl     *gomock.Controller
	recorder *MockCacheMockRecorder
}

// MockCacheMockRecorder is the mock recorder for MockCache
type MockCacheMockRecorder struct {
	mock *MockCache
}

// NewMockCache creates a new mock instance
func NewMockCache(ctrl *gomock.Controller) *MockCache {
	mock := &MockCache{ctrl: ctrl}
	mock.recorder = &MockCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCache) EXPECT() *MockCacheMockRecorder {
	return m.recorder
}

// Clear mocks base method
func (m *MockCache) Clear(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Clear", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Clear indicates an expected call of Clear
func (mr *MockCacheMockRecorder) Clear(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clear", reflect.TypeOf((*MockCache)(nil).Clear), arg0)
}

// Volume mocks base method
func (m *MockCache) Volume() string {
	ret := m.ctrl.Call(m, "Volume")
	ret0, _ := ret[0].(string)
	return ret0
}

// Volume indicates an expected call of Volume
func (mr *MockCacheMockRecorder) Volume() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Volume", reflect.TypeOf((*MockCache)(nil).Volume))
}

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain (interfaces: ShortenerService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/AssylzhanZharzhanov/go-shortener/internal/shortener/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockShortenerService is a mock of ShortenerService interface.
type MockShortenerService struct {
	ctrl     *gomock.Controller
	recorder *MockShortenerServiceMockRecorder
}

// MockShortenerServiceMockRecorder is the mock recorder for MockShortenerService.
type MockShortenerServiceMockRecorder struct {
	mock *MockShortenerService
}

// NewMockShortenerService creates a new mock instance.
func NewMockShortenerService(ctrl *gomock.Controller) *MockShortenerService {
	mock := &MockShortenerService{ctrl: ctrl}
	mock.recorder = &MockShortenerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShortenerService) EXPECT() *MockShortenerServiceMockRecorder {
	return m.recorder
}

// CreateShortenURL mocks base method.
func (m *MockShortenerService) CreateShortenURL(arg0 context.Context, arg1 *domain.Link) (*domain.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortenURL", arg0, arg1)
	ret0, _ := ret[0].(*domain.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShortenURL indicates an expected call of CreateShortenURL.
func (mr *MockShortenerServiceMockRecorder) CreateShortenURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortenURL", reflect.TypeOf((*MockShortenerService)(nil).CreateShortenURL), arg0, arg1)
}

// Get mocks base method.
func (m *MockShortenerService) Get(arg0 context.Context, arg1 string) (*domain.Link, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*domain.Link)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShortenerServiceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShortenerService)(nil).Get), arg0, arg1)
}

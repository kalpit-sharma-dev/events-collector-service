// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "github.com/kalpit-sharma-dev/kalpit.cool2006-gmail.com/src/models"
	reflect "reflect"
)

// MockEventCollectorService is a mock of EventCollectorService interface
type MockEventCollectorService struct {
	ctrl     *gomock.Controller
	recorder *MockEventCollectorServiceMockRecorder
}

// MockEventCollectorServiceMockRecorder is the mock recorder for MockEventCollectorService
type MockEventCollectorServiceMockRecorder struct {
	mock *MockEventCollectorService
}

// NewMockEventCollectorService creates a new mock instance
func NewMockEventCollectorService(ctrl *gomock.Controller) *MockEventCollectorService {
	mock := &MockEventCollectorService{ctrl: ctrl}
	mock.recorder = &MockEventCollectorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventCollectorService) EXPECT() *MockEventCollectorServiceMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method
func (m *MockEventCollectorService) CreateEvent(ctx context.Context, eventReq models.Event) (models.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", ctx, eventReq)
	ret0, _ := ret[0].(models.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent
func (mr *MockEventCollectorServiceMockRecorder) CreateEvent(ctx, eventReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockEventCollectorService)(nil).CreateEvent), ctx, eventReq)
}

// GetAllEvents mocks base method
func (m *MockEventCollectorService) GetAllEvents(ctx context.Context) (models.EventList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvents", ctx)
	ret0, _ := ret[0].(models.EventList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEvents indicates an expected call of GetAllEvents
func (mr *MockEventCollectorServiceMockRecorder) GetAllEvents(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvents", reflect.TypeOf((*MockEventCollectorService)(nil).GetAllEvents), ctx)
}
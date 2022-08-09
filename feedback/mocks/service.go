// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_feedback is a generated GoMock package.
package mock_feedback

import (
	feedback "product-feedback/feedback"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFeedbackService is a mock of FeedbackService interface.
type MockFeedbackService struct {
	ctrl     *gomock.Controller
	recorder *MockFeedbackServiceMockRecorder
}

// MockFeedbackServiceMockRecorder is the mock recorder for MockFeedbackService.
type MockFeedbackServiceMockRecorder struct {
	mock *MockFeedbackService
}

// NewMockFeedbackService creates a new mock instance.
func NewMockFeedbackService(ctrl *gomock.Controller) *MockFeedbackService {
	mock := &MockFeedbackService{ctrl: ctrl}
	mock.recorder = &MockFeedbackServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeedbackService) EXPECT() *MockFeedbackServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFeedbackService) Create(userId int, f feedback.CreateFeedbackInput) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, f)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFeedbackServiceMockRecorder) Create(userId, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFeedbackService)(nil).Create), userId, f)
}

// Delete mocks base method.
func (m *MockFeedbackService) Delete(userId, feedbackId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, feedbackId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockFeedbackServiceMockRecorder) Delete(userId, feedbackId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockFeedbackService)(nil).Delete), userId, feedbackId)
}

// GetAll mocks base method.
func (m *MockFeedbackService) GetAll() ([]feedback.Feedback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]feedback.Feedback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockFeedbackServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockFeedbackService)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockFeedbackService) GetById(feedbackId int) (feedback.Feedback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", feedbackId)
	ret0, _ := ret[0].(feedback.Feedback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockFeedbackServiceMockRecorder) GetById(feedbackId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockFeedbackService)(nil).GetById), feedbackId)
}

// Update mocks base method.
func (m *MockFeedbackService) Update(userId, feedbackId int, f feedback.UpdateFeedbackInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, feedbackId, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockFeedbackServiceMockRecorder) Update(userId, feedbackId, f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockFeedbackService)(nil).Update), userId, feedbackId, f)
}
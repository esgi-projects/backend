// Code generated by MockGen. DO NOT EDIT.
// Source: internal/note/usecase.go
//
// Generated by this command:
//
//	mockgen -source=internal/note/usecase.go -destination=internal/note/mock/usecase_mock.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	models "github.com/esgi-challenge/backend/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUseCase) Create(note *models.Note) (*models.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", note)
	ret0, _ := ret[0].(*models.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(note any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), note)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), id)
}

// GetAllByUser mocks base method.
func (m *MockUseCase) GetAllByUser(user *models.User) (*[]models.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUser", user)
	ret0, _ := ret[0].(*[]models.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUser indicates an expected call of GetAllByUser.
func (mr *MockUseCaseMockRecorder) GetAllByUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUser", reflect.TypeOf((*MockUseCase)(nil).GetAllByUser), user)
}

// GetById mocks base method.
func (m *MockUseCase) GetById(id uint) (*models.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUseCaseMockRecorder) GetById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUseCase)(nil).GetById), id)
}

// Update mocks base method.
func (m *MockUseCase) Update(id uint, updatedNote *models.Note) (*models.Note, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, updatedNote)
	ret0, _ := ret[0].(*models.Note)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUseCaseMockRecorder) Update(id, updatedNote any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), id, updatedNote)
}
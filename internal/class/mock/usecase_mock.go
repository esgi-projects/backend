// Code generated by MockGen. DO NOT EDIT.
// Source: internal/class/usecase.go
//
// Generated by this command:
//
//	mockgen -source=internal/class/usecase.go -destination=internal/class/mock/usecase_mock.go -package=mock
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

// Add mocks base method.
func (m *MockUseCase) Add(id uint, class *models.ClassAdd) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", id, class)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockUseCaseMockRecorder) Add(id, class any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockUseCase)(nil).Add), id, class)
}

// Create mocks base method.
func (m *MockUseCase) Create(user *models.User, class *models.Class) (*models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user, class)
	ret0, _ := ret[0].(*models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUseCaseMockRecorder) Create(user, class any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), user, class)
}

// Delete mocks base method.
func (m *MockUseCase) Delete(user *models.User, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", user, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUseCaseMockRecorder) Delete(user, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), user, id)
}

// GetAll mocks base method.
func (m *MockUseCase) GetAll() (*[]models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUseCaseMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUseCase)(nil).GetAll))
}

// GetAllBySchoolId mocks base method.
func (m *MockUseCase) GetAllBySchoolId(schoolId uint) (*[]models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBySchoolId", schoolId)
	ret0, _ := ret[0].(*[]models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllBySchoolId indicates an expected call of GetAllBySchoolId.
func (mr *MockUseCaseMockRecorder) GetAllBySchoolId(schoolId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBySchoolId", reflect.TypeOf((*MockUseCase)(nil).GetAllBySchoolId), schoolId)
}

// GetById mocks base method.
func (m *MockUseCase) GetById(id uint) (*models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockUseCaseMockRecorder) GetById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockUseCase)(nil).GetById), id)
}

// GetClassLessStudents mocks base method.
func (m *MockUseCase) GetClassLessStudents(schoolId uint) (*[]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClassLessStudents", schoolId)
	ret0, _ := ret[0].(*[]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassLessStudents indicates an expected call of GetClassLessStudents.
func (mr *MockUseCaseMockRecorder) GetClassLessStudents(schoolId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassLessStudents", reflect.TypeOf((*MockUseCase)(nil).GetClassLessStudents), schoolId)
}

// Remove mocks base method.
func (m *MockUseCase) Remove(id uint, class *models.ClassRemove) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", id, class)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Remove indicates an expected call of Remove.
func (mr *MockUseCaseMockRecorder) Remove(id, class any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockUseCase)(nil).Remove), id, class)
}

// Update mocks base method.
func (m *MockUseCase) Update(id uint, updatedClass *models.Class) (*models.Class, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, updatedClass)
	ret0, _ := ret[0].(*models.Class)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUseCaseMockRecorder) Update(id, updatedClass any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), id, updatedClass)
}
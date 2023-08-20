// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	generated "github.com/SawitProRecruitment/UserService/generated"
	gomock "github.com/golang/mock/gomock"
	v4 "github.com/labstack/echo/v4"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Generated mocks base method.
func (m *MockRepositoryInterface) Generated(ctx context.Context, id int, phoneNumber string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generated", ctx, id, phoneNumber)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generated indicates an expected call of Generated.
func (mr *MockRepositoryInterfaceMockRecorder) Generated(ctx, id, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generated", reflect.TypeOf((*MockRepositoryInterface)(nil).Generated), ctx, id, phoneNumber)
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}

// GetUserByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhoneNumber", ctx, phoneNumber)
	ret0, _ := ret[0].(Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhoneNumber indicates an expected call of GetUserByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) GetUserByPhoneNumber(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).GetUserByPhoneNumber), ctx, phoneNumber)
}

// InsertUsers mocks base method.
func (m *MockRepositoryInterface) InsertUsers(ctx v4.Context, input generated.RegistrationRequest) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUsers", ctx, input)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUsers indicates an expected call of InsertUsers.
func (mr *MockRepositoryInterfaceMockRecorder) InsertUsers(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUsers", reflect.TypeOf((*MockRepositoryInterface)(nil).InsertUsers), ctx, input)
}

// UpdateCounterLogin mocks base method.
func (m *MockRepositoryInterface) UpdateCounterLogin(ctx context.Context, loginCounter int, phoneNumber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCounterLogin", ctx, loginCounter, phoneNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCounterLogin indicates an expected call of UpdateCounterLogin.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateCounterLogin(ctx, loginCounter, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCounterLogin", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateCounterLogin), ctx, loginCounter, phoneNumber)
}

// UpdateUserByPhoneNumber mocks base method.
func (m *MockRepositoryInterface) UpdateUserByPhoneNumber(ctx context.Context, req generated.ProfileRequest, phoneNumber string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByPhoneNumber", ctx, req, phoneNumber)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByPhoneNumber indicates an expected call of UpdateUserByPhoneNumber.
func (mr *MockRepositoryInterfaceMockRecorder) UpdateUserByPhoneNumber(ctx, req, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByPhoneNumber", reflect.TypeOf((*MockRepositoryInterface)(nil).UpdateUserByPhoneNumber), ctx, req, phoneNumber)
}

// Verify mocks base method.
func (m *MockRepositoryInterface) Verify(authToken string) (*JwtCustomClaims, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", authToken)
	ret0, _ := ret[0].(*JwtCustomClaims)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockRepositoryInterfaceMockRecorder) Verify(authToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockRepositoryInterface)(nil).Verify), authToken)
}

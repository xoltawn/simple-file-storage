// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xoltawn/simple-file-storage/domain (interfaces: FileUsecase)

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/xoltawn/simple-file-storage/domain"
)

// MockFileUsecase is a mock of FileUsecase interface.
type MockFileUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockFileUsecaseMockRecorder
}

// MockFileUsecaseMockRecorder is the mock recorder for MockFileUsecase.
type MockFileUsecaseMockRecorder struct {
	mock *MockFileUsecase
}

// NewMockFileUsecase creates a new mock instance.
func NewMockFileUsecase(ctrl *gomock.Controller) *MockFileUsecase {
	mock := &MockFileUsecase{ctrl: ctrl}
	mock.recorder = &MockFileUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileUsecase) EXPECT() *MockFileUsecaseMockRecorder {
	return m.recorder
}

// DownloadFromTextFile mocks base method.
func (m *MockFileUsecase) DownloadFromTextFile(arg0 context.Context, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadFromTextFile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadFromTextFile indicates an expected call of DownloadFromTextFile.
func (mr *MockFileUsecaseMockRecorder) DownloadFromTextFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadFromTextFile", reflect.TypeOf((*MockFileUsecase)(nil).DownloadFromTextFile), arg0, arg1)
}

// FetchFiles mocks base method.
func (m *MockFileUsecase) FetchFiles(arg0 context.Context, arg1, arg2 int) ([]domain.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchFiles", arg0, arg1, arg2)
	ret0, _ := ret[0].([]domain.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFiles indicates an expected call of FetchFiles.
func (mr *MockFileUsecaseMockRecorder) FetchFiles(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFiles", reflect.TypeOf((*MockFileUsecase)(nil).FetchFiles), arg0, arg1, arg2)
}

// UploadFile mocks base method.
func (m *MockFileUsecase) UploadFile(arg0 context.Context, arg1 []byte) (domain.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", arg0, arg1)
	ret0, _ := ret[0].(domain.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockFileUsecaseMockRecorder) UploadFile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockFileUsecase)(nil).UploadFile), arg0, arg1)
}
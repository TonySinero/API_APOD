// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/apod/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockImageRepo is a mock of ImageRepo interface.
type MockImageRepo struct {
	ctrl     *gomock.Controller
	recorder *MockImageRepoMockRecorder
}

// MockImageRepoMockRecorder is the mock recorder for MockImageRepo.
type MockImageRepoMockRecorder struct {
	mock *MockImageRepo
}

// NewMockImageRepo creates a new mock instance.
func NewMockImageRepo(ctrl *gomock.Controller) *MockImageRepo {
	mock := &MockImageRepo{ctrl: ctrl}
	mock.recorder = &MockImageRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageRepo) EXPECT() *MockImageRepoMockRecorder {
	return m.recorder
}

// CreateAlbum mocks base method.
func (m *MockImageRepo) CreateAlbum(im *model.Nasa) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAlbum", im)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAlbum indicates an expected call of CreateAlbum.
func (mr *MockImageRepoMockRecorder) CreateAlbum(im interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAlbum", reflect.TypeOf((*MockImageRepo)(nil).CreateAlbum), im)
}

// GetAll mocks base method.
func (m *MockImageRepo) GetAll() ([]model.Nasa, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]model.Nasa)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockImageRepoMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockImageRepo)(nil).GetAll))
}

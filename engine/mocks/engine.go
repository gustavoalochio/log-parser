// Code generated by MockGen. DO NOT EDIT.
// Source: engine.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "log-parser/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// BuildDeadReasons mocks base method.
func (m *MockEngine) BuildDeadReasons(game *entity.Game, gameCounter int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildDeadReasons", game, gameCounter)
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDeadReasons indicates an expected call of BuildDeadReasons.
func (mr *MockEngineMockRecorder) BuildDeadReasons(game, gameCounter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildDeadReasons", reflect.TypeOf((*MockEngine)(nil).BuildDeadReasons), game, gameCounter)
}

// BuildResultGame mocks base method.
func (m *MockEngine) BuildResultGame(result entity.ResultGameInformation, gameCounter int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildResultGame", result, gameCounter)
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildResultGame indicates an expected call of BuildResultGame.
func (mr *MockEngineMockRecorder) BuildResultGame(result, gameCounter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildResultGame", reflect.TypeOf((*MockEngine)(nil).BuildResultGame), result, gameCounter)
}

// EndGame mocks base method.
func (m *MockEngine) EndGame(game *entity.Game) entity.ResultGameInformation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndGame", game)
	ret0, _ := ret[0].(entity.ResultGameInformation)
	return ret0
}

// EndGame indicates an expected call of EndGame.
func (mr *MockEngineMockRecorder) EndGame(game interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndGame", reflect.TypeOf((*MockEngine)(nil).EndGame), game)
}

// InitGame mocks base method.
func (m *MockEngine) InitGame() *engine {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGame")
	ret0, _ := ret[0].(*engine)
	return ret0
}

// InitGame indicates an expected call of InitGame.
func (mr *MockEngineMockRecorder) InitGame() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGame", reflect.TypeOf((*MockEngine)(nil).InitGame))
}

// Kill mocks base method.
func (m *MockEngine) Kill(game *entity.Game, playerKilledInformation entity.PlayerKilledInformation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill", game, playerKilledInformation)
}

// Kill indicates an expected call of Kill.
func (mr *MockEngineMockRecorder) Kill(game, playerKilledInformation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockEngine)(nil).Kill), game, playerKilledInformation)
}

// Start mocks base method.
func (m *MockEngine) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockEngineMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEngine)(nil).Start))
}

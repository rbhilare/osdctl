// Code generated by MockGen. DO NOT EDIT.
// Source: jira.go
//
// Generated by this command:
//
//	mockgen -source=jira.go -destination=./mocks/jira_mock.go -package=utils
//

// Package utils is a generated GoMock package.
package utils

import (
	reflect "reflect"

	jira "github.com/andygrunwald/go-jira"
	gomock "go.uber.org/mock/gomock"
)

// MockJiraClientInterface is a mock of JiraClientInterface interface.
type MockJiraClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJiraClientInterfaceMockRecorder
	isgomock struct{}
}

// MockJiraClientInterfaceMockRecorder is the mock recorder for MockJiraClientInterface.
type MockJiraClientInterfaceMockRecorder struct {
	mock *MockJiraClientInterface
}

// NewMockJiraClientInterface creates a new mock instance.
func NewMockJiraClientInterface(ctrl *gomock.Controller) *MockJiraClientInterface {
	mock := &MockJiraClientInterface{ctrl: ctrl}
	mock.recorder = &MockJiraClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJiraClientInterface) EXPECT() *MockJiraClientInterfaceMockRecorder {
	return m.recorder
}

// Board mocks base method.
func (m *MockJiraClientInterface) Board() *jira.BoardService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Board")
	ret0, _ := ret[0].(*jira.BoardService)
	return ret0
}

// Board indicates an expected call of Board.
func (mr *MockJiraClientInterfaceMockRecorder) Board() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Board", reflect.TypeOf((*MockJiraClientInterface)(nil).Board))
}

// CreateIssue mocks base method.
func (m *MockJiraClientInterface) CreateIssue(issue *jira.Issue) (*jira.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIssue", issue)
	ret0, _ := ret[0].(*jira.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIssue indicates an expected call of CreateIssue.
func (mr *MockJiraClientInterfaceMockRecorder) CreateIssue(issue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIssue", reflect.TypeOf((*MockJiraClientInterface)(nil).CreateIssue), issue)
}

// CreateVersion mocks base method.
func (m *MockJiraClientInterface) CreateVersion(version *jira.Version) (*jira.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVersion", version)
	ret0, _ := ret[0].(*jira.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateVersion indicates an expected call of CreateVersion.
func (mr *MockJiraClientInterfaceMockRecorder) CreateVersion(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVersion", reflect.TypeOf((*MockJiraClientInterface)(nil).CreateVersion), version)
}

// Issue mocks base method.
func (m *MockJiraClientInterface) Issue() *jira.IssueService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Issue")
	ret0, _ := ret[0].(*jira.IssueService)
	return ret0
}

// Issue indicates an expected call of Issue.
func (mr *MockJiraClientInterfaceMockRecorder) Issue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Issue", reflect.TypeOf((*MockJiraClientInterface)(nil).Issue))
}

// SearchIssues mocks base method.
func (m *MockJiraClientInterface) SearchIssues(jql string) ([]jira.Issue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchIssues", jql)
	ret0, _ := ret[0].([]jira.Issue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchIssues indicates an expected call of SearchIssues.
func (mr *MockJiraClientInterfaceMockRecorder) SearchIssues(jql any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchIssues", reflect.TypeOf((*MockJiraClientInterface)(nil).SearchIssues), jql)
}

// Sprint mocks base method.
func (m *MockJiraClientInterface) Sprint() *jira.SprintService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sprint")
	ret0, _ := ret[0].(*jira.SprintService)
	return ret0
}

// Sprint indicates an expected call of Sprint.
func (mr *MockJiraClientInterfaceMockRecorder) Sprint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sprint", reflect.TypeOf((*MockJiraClientInterface)(nil).Sprint))
}

// User mocks base method.
func (m *MockJiraClientInterface) User() *jira.UserService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(*jira.UserService)
	return ret0
}

// User indicates an expected call of User.
func (mr *MockJiraClientInterfaceMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockJiraClientInterface)(nil).User))
}

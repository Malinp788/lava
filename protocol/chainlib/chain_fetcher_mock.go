// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/chainlib/chain_fetcher.go
//
// Generated by this command:
//
//	mockgen -source=protocol/chainlib/chain_fetcher.go -destination protocol/chainlib/chain_fetcher_mock.go -package chainlib
//
// Package chainlib is a generated GoMock package.
package chainlib

import (
	context "context"
	reflect "reflect"

	lavasession "github.com/lavanet/lava/protocol/lavasession"
	gomock "go.uber.org/mock/gomock"
)

// MockChainFetcherIf is a mock of ChainFetcherIf interface.
type MockChainFetcherIf struct {
	ctrl     *gomock.Controller
	recorder *MockChainFetcherIfMockRecorder
}

// MockChainFetcherIfMockRecorder is the mock recorder for MockChainFetcherIf.
type MockChainFetcherIfMockRecorder struct {
	mock *MockChainFetcherIf
}

// NewMockChainFetcherIf creates a new mock instance.
func NewMockChainFetcherIf(ctrl *gomock.Controller) *MockChainFetcherIf {
	mock := &MockChainFetcherIf{ctrl: ctrl}
	mock.recorder = &MockChainFetcherIfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainFetcherIf) EXPECT() *MockChainFetcherIfMockRecorder {
	return m.recorder
}

// FetchBlockHashByNum mocks base method.
func (m *MockChainFetcherIf) FetchBlockHashByNum(ctx context.Context, blockNum int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBlockHashByNum", ctx, blockNum)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBlockHashByNum indicates an expected call of FetchBlockHashByNum.
func (mr *MockChainFetcherIfMockRecorder) FetchBlockHashByNum(ctx, blockNum any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBlockHashByNum", reflect.TypeOf((*MockChainFetcherIf)(nil).FetchBlockHashByNum), ctx, blockNum)
}

// FetchEndpoint mocks base method.
func (m *MockChainFetcherIf) FetchEndpoint() lavasession.RPCProviderEndpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchEndpoint")
	ret0, _ := ret[0].(lavasession.RPCProviderEndpoint)
	return ret0
}

// FetchEndpoint indicates an expected call of FetchEndpoint.
func (mr *MockChainFetcherIfMockRecorder) FetchEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchEndpoint", reflect.TypeOf((*MockChainFetcherIf)(nil).FetchEndpoint))
}

// FetchLatestBlockNum mocks base method.
func (m *MockChainFetcherIf) FetchLatestBlockNum(ctx context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchLatestBlockNum", ctx)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchLatestBlockNum indicates an expected call of FetchLatestBlockNum.
func (mr *MockChainFetcherIfMockRecorder) FetchLatestBlockNum(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchLatestBlockNum", reflect.TypeOf((*MockChainFetcherIf)(nil).FetchLatestBlockNum), ctx)
}

// Validate mocks base method.
func (m *MockChainFetcherIf) Validate(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validate", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Validate indicates an expected call of Validate.
func (mr *MockChainFetcherIfMockRecorder) Validate(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validate", reflect.TypeOf((*MockChainFetcherIf)(nil).Validate), ctx)
}
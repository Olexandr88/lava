// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/rpcconsumer/rpcconsumer.go
//
// Generated by this command:
//
//	mockgen -source=protocol/rpcconsumer/rpcconsumer.go -destination protocol/rpcconsumer/consumer_state_tracker_mock.go -package rpcconsumer
//

// Package rpcconsumer is a generated GoMock package.
package rpcconsumer

import (
	context "context"
	reflect "reflect"

	common "github.com/lavanet/lava/v4/protocol/common"
	finalizationconsensus "github.com/lavanet/lava/v4/protocol/lavaprotocol/finalizationconsensus"
	lavasession "github.com/lavanet/lava/v4/protocol/lavasession"
	updaters "github.com/lavanet/lava/v4/protocol/statetracker/updaters"
	types "github.com/lavanet/lava/v4/x/conflict/types"
	types0 "github.com/lavanet/lava/v4/x/plans/types"
	types1 "github.com/lavanet/lava/v4/x/protocol/types"
	gomock "go.uber.org/mock/gomock"
)

// MockConsumerStateTrackerInf is a mock of ConsumerStateTrackerInf interface.
type MockConsumerStateTrackerInf struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerStateTrackerInfMockRecorder
}

// MockConsumerStateTrackerInfMockRecorder is the mock recorder for MockConsumerStateTrackerInf.
type MockConsumerStateTrackerInfMockRecorder struct {
	mock *MockConsumerStateTrackerInf
}

// NewMockConsumerStateTrackerInf creates a new mock instance.
func NewMockConsumerStateTrackerInf(ctrl *gomock.Controller) *MockConsumerStateTrackerInf {
	mock := &MockConsumerStateTrackerInf{ctrl: ctrl}
	mock.recorder = &MockConsumerStateTrackerInfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumerStateTrackerInf) EXPECT() *MockConsumerStateTrackerInfMockRecorder {
	return m.recorder
}

// GetConsumerPolicy mocks base method.
func (m *MockConsumerStateTrackerInf) GetConsumerPolicy(ctx context.Context, consumerAddress, chainID string) (*types0.Policy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConsumerPolicy", ctx, consumerAddress, chainID)
	ret0, _ := ret[0].(*types0.Policy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConsumerPolicy indicates an expected call of GetConsumerPolicy.
func (mr *MockConsumerStateTrackerInfMockRecorder) GetConsumerPolicy(ctx, consumerAddress, chainID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsumerPolicy", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).GetConsumerPolicy), ctx, consumerAddress, chainID)
}

// GetLatestVirtualEpoch mocks base method.
func (m *MockConsumerStateTrackerInf) GetLatestVirtualEpoch() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestVirtualEpoch")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetLatestVirtualEpoch indicates an expected call of GetLatestVirtualEpoch.
func (mr *MockConsumerStateTrackerInfMockRecorder) GetLatestVirtualEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestVirtualEpoch", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).GetLatestVirtualEpoch))
}

// GetProtocolVersion mocks base method.
func (m *MockConsumerStateTrackerInf) GetProtocolVersion(ctx context.Context) (*updaters.ProtocolVersionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtocolVersion", ctx)
	ret0, _ := ret[0].(*updaters.ProtocolVersionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProtocolVersion indicates an expected call of GetProtocolVersion.
func (mr *MockConsumerStateTrackerInfMockRecorder) GetProtocolVersion(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtocolVersion", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).GetProtocolVersion), ctx)
}

// RegisterConsumerSessionManagerForPairingUpdates mocks base method.
func (m *MockConsumerStateTrackerInf) RegisterConsumerSessionManagerForPairingUpdates(ctx context.Context, consumerSessionManager *lavasession.ConsumerSessionManager, staticProviders []*lavasession.RPCProviderEndpoint) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterConsumerSessionManagerForPairingUpdates", ctx, consumerSessionManager)
}

// RegisterConsumerSessionManagerForPairingUpdates indicates an expected call of RegisterConsumerSessionManagerForPairingUpdates.
func (mr *MockConsumerStateTrackerInfMockRecorder) RegisterConsumerSessionManagerForPairingUpdates(ctx, consumerSessionManager any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterConsumerSessionManagerForPairingUpdates", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).RegisterConsumerSessionManagerForPairingUpdates), ctx, consumerSessionManager)
}

// RegisterFinalizationConsensusForUpdates mocks base method.
func (m *MockConsumerStateTrackerInf) RegisterFinalizationConsensusForUpdates(arg0 context.Context, arg1 *finalizationconsensus.FinalizationConsensus) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterFinalizationConsensusForUpdates", arg0, arg1)
}

// RegisterFinalizationConsensusForUpdates indicates an expected call of RegisterFinalizationConsensusForUpdates.
func (mr *MockConsumerStateTrackerInfMockRecorder) RegisterFinalizationConsensusForUpdates(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFinalizationConsensusForUpdates", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).RegisterFinalizationConsensusForUpdates), arg0, arg1)
}

// RegisterForDowntimeParamsUpdates mocks base method.
func (m *MockConsumerStateTrackerInf) RegisterForDowntimeParamsUpdates(ctx context.Context, downtimeParamsUpdatable updaters.DowntimeParamsUpdatable) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterForDowntimeParamsUpdates", ctx, downtimeParamsUpdatable)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterForDowntimeParamsUpdates indicates an expected call of RegisterForDowntimeParamsUpdates.
func (mr *MockConsumerStateTrackerInfMockRecorder) RegisterForDowntimeParamsUpdates(ctx, downtimeParamsUpdatable any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterForDowntimeParamsUpdates", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).RegisterForDowntimeParamsUpdates), ctx, downtimeParamsUpdatable)
}

// RegisterForSpecUpdates mocks base method.
func (m *MockConsumerStateTrackerInf) RegisterForSpecUpdates(ctx context.Context, specUpdatable updaters.SpecUpdatable, endpoint lavasession.RPCEndpoint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterForSpecUpdates", ctx, specUpdatable, endpoint)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterForSpecUpdates indicates an expected call of RegisterForSpecUpdates.
func (mr *MockConsumerStateTrackerInfMockRecorder) RegisterForSpecUpdates(ctx, specUpdatable, endpoint any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterForSpecUpdates", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).RegisterForSpecUpdates), ctx, specUpdatable, endpoint)
}

// RegisterForVersionUpdates mocks base method.
func (m *MockConsumerStateTrackerInf) RegisterForVersionUpdates(ctx context.Context, version *types1.Version, versionValidator updaters.VersionValidationInf) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterForVersionUpdates", ctx, version, versionValidator)
}

// RegisterForVersionUpdates indicates an expected call of RegisterForVersionUpdates.
func (mr *MockConsumerStateTrackerInfMockRecorder) RegisterForVersionUpdates(ctx, version, versionValidator any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterForVersionUpdates", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).RegisterForVersionUpdates), ctx, version, versionValidator)
}

// TxConflictDetection mocks base method.
func (m *MockConsumerStateTrackerInf) TxConflictDetection(ctx context.Context, finalizationConflict *types.FinalizationConflict, responseConflict *types.ResponseConflict, conflictHandler common.ConflictHandlerInterface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TxConflictDetection", ctx, finalizationConflict, responseConflict, conflictHandler)
	ret0, _ := ret[0].(error)
	return ret0
}

// TxConflictDetection indicates an expected call of TxConflictDetection.
func (mr *MockConsumerStateTrackerInfMockRecorder) TxConflictDetection(ctx, finalizationConflict, responseConflict, conflictHandler any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TxConflictDetection", reflect.TypeOf((*MockConsumerStateTrackerInf)(nil).TxConflictDetection), ctx, finalizationConflict, responseConflict, conflictHandler)
}

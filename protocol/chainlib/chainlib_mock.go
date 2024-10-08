// Code generated by MockGen. DO NOT EDIT.
// Source: protocol/chainlib/chainlib.go

// Package chainlib is a generated GoMock package.
package chainlib

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	rpcInterfaceMessages "github.com/lavanet/lava/v3/protocol/chainlib/chainproxy/rpcInterfaceMessages"
	rpcclient "github.com/lavanet/lava/v3/protocol/chainlib/chainproxy/rpcclient"
	extensionslib "github.com/lavanet/lava/v3/protocol/chainlib/extensionslib"
	common "github.com/lavanet/lava/v3/protocol/common"
	metrics "github.com/lavanet/lava/v3/protocol/metrics"
	types "github.com/lavanet/lava/v3/x/pairing/types"
	types0 "github.com/lavanet/lava/v3/x/spec/types"
)

// MockChainParser is a mock of ChainParser interface.
type MockChainParser struct {
	ctrl     *gomock.Controller
	recorder *MockChainParserMockRecorder
}

// MockChainParserMockRecorder is the mock recorder for MockChainParser.
type MockChainParserMockRecorder struct {
	mock *MockChainParser
}

// NewMockChainParser creates a new mock instance.
func NewMockChainParser(ctrl *gomock.Controller) *MockChainParser {
	mock := &MockChainParser{ctrl: ctrl}
	mock.recorder = &MockChainParserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainParser) EXPECT() *MockChainParserMockRecorder {
	return m.recorder
}

// Activate mocks base method.
func (m *MockChainParser) Activate() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Activate")
}

// Activate indicates an expected call of Activate.
func (mr *MockChainParserMockRecorder) Activate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activate", reflect.TypeOf((*MockChainParser)(nil).Activate))
}

// Active mocks base method.
func (m *MockChainParser) Active() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Active")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Active indicates an expected call of Active.
func (mr *MockChainParserMockRecorder) Active() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Active", reflect.TypeOf((*MockChainParser)(nil).Active))
}

// ChainBlockStats mocks base method.
func (m *MockChainParser) ChainBlockStats() (int64, time.Duration, uint32, uint32) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainBlockStats")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(uint32)
	ret3, _ := ret[3].(uint32)
	return ret0, ret1, ret2, ret3
}

// ChainBlockStats indicates an expected call of ChainBlockStats.
func (mr *MockChainParserMockRecorder) ChainBlockStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainBlockStats", reflect.TypeOf((*MockChainParser)(nil).ChainBlockStats))
}

// CraftMessage mocks base method.
func (m *MockChainParser) CraftMessage(parser *types0.ParseDirective, connectionType string, craftData *CraftData, metadata []types.Metadata) (ChainMessageForSend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CraftMessage", parser, connectionType, craftData, metadata)
	ret0, _ := ret[0].(ChainMessageForSend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CraftMessage indicates an expected call of CraftMessage.
func (mr *MockChainParserMockRecorder) CraftMessage(parser, connectionType, craftData, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CraftMessage", reflect.TypeOf((*MockChainParser)(nil).CraftMessage), parser, connectionType, craftData, metadata)
}

// DataReliabilityParams mocks base method.
func (m *MockChainParser) DataReliabilityParams() (bool, uint32) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DataReliabilityParams")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(uint32)
	return ret0, ret1
}

// DataReliabilityParams indicates an expected call of DataReliabilityParams.
func (mr *MockChainParserMockRecorder) DataReliabilityParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataReliabilityParams", reflect.TypeOf((*MockChainParser)(nil).DataReliabilityParams))
}

// ExtensionsParser mocks base method.
func (m *MockChainParser) ExtensionsParser() *extensionslib.ExtensionParser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtensionsParser")
	ret0, _ := ret[0].(*extensionslib.ExtensionParser)
	return ret0
}

// ExtensionsParser indicates an expected call of ExtensionsParser.
func (mr *MockChainParserMockRecorder) ExtensionsParser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionsParser", reflect.TypeOf((*MockChainParser)(nil).ExtensionsParser))
}

// GetParsingByTag mocks base method.
func (m *MockChainParser) GetParsingByTag(tag types0.FUNCTION_TAG) (*types0.ParseDirective, *types0.ApiCollection, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParsingByTag", tag)
	ret0, _ := ret[0].(*types0.ParseDirective)
	ret1, _ := ret[1].(*types0.ApiCollection)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// GetParsingByTag indicates an expected call of GetParsingByTag.
func (mr *MockChainParserMockRecorder) GetParsingByTag(tag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParsingByTag", reflect.TypeOf((*MockChainParser)(nil).GetParsingByTag), tag)
}

// GetUniqueName mocks base method.
func (m *MockChainParser) GetUniqueName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUniqueName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetUniqueName indicates an expected call of GetUniqueName.
func (mr *MockChainParserMockRecorder) GetUniqueName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUniqueName", reflect.TypeOf((*MockChainParser)(nil).GetUniqueName))
}

// GetVerifications mocks base method.
func (m *MockChainParser) GetVerifications(supported []string) ([]VerificationContainer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVerifications", supported)
	ret0, _ := ret[0].([]VerificationContainer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVerifications indicates an expected call of GetVerifications.
func (mr *MockChainParserMockRecorder) GetVerifications(supported interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVerifications", reflect.TypeOf((*MockChainParser)(nil).GetVerifications), supported)
}

// HandleHeaders mocks base method.
func (m *MockChainParser) HandleHeaders(metadata []types.Metadata, apiCollection *types0.ApiCollection, headersDirection types0.Header_HeaderType) ([]types.Metadata, string, []types.Metadata) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleHeaders", metadata, apiCollection, headersDirection)
	ret0, _ := ret[0].([]types.Metadata)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].([]types.Metadata)
	return ret0, ret1, ret2
}

// HandleHeaders indicates an expected call of HandleHeaders.
func (mr *MockChainParserMockRecorder) HandleHeaders(metadata, apiCollection, headersDirection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleHeaders", reflect.TypeOf((*MockChainParser)(nil).HandleHeaders), metadata, apiCollection, headersDirection)
}

// ParseMsg mocks base method.
func (m *MockChainParser) ParseMsg(url string, data []byte, connectionType string, metadata []types.Metadata, extensionInfo extensionslib.ExtensionInfo) (ChainMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseMsg", url, data, connectionType, metadata, extensionInfo)
	ret0, _ := ret[0].(ChainMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseMsg indicates an expected call of ParseMsg.
func (mr *MockChainParserMockRecorder) ParseMsg(url, data, connectionType, metadata, extensionInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseMsg", reflect.TypeOf((*MockChainParser)(nil).ParseMsg), url, data, connectionType, metadata, extensionInfo)
}

// SeparateAddonsExtensions mocks base method.
func (m *MockChainParser) SeparateAddonsExtensions(supported []string) ([]string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SeparateAddonsExtensions", supported)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SeparateAddonsExtensions indicates an expected call of SeparateAddonsExtensions.
func (mr *MockChainParserMockRecorder) SeparateAddonsExtensions(supported interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SeparateAddonsExtensions", reflect.TypeOf((*MockChainParser)(nil).SeparateAddonsExtensions), supported)
}

// SetPolicy mocks base method.
func (m *MockChainParser) SetPolicy(policy PolicyInf, chainId, apiInterface string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPolicy", policy, chainId, apiInterface)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPolicy indicates an expected call of SetPolicy.
func (mr *MockChainParserMockRecorder) SetPolicy(policy, chainId, apiInterface interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPolicy", reflect.TypeOf((*MockChainParser)(nil).SetPolicy), policy, chainId, apiInterface)
}

// SetSpec mocks base method.
func (m *MockChainParser) SetSpec(spec types0.Spec) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSpec", spec)
}

// SetSpec indicates an expected call of SetSpec.
func (mr *MockChainParserMockRecorder) SetSpec(spec interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSpec", reflect.TypeOf((*MockChainParser)(nil).SetSpec), spec)
}

// UpdateBlockTime mocks base method.
func (m *MockChainParser) UpdateBlockTime(newBlockTime time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateBlockTime", newBlockTime)
}

// UpdateBlockTime indicates an expected call of UpdateBlockTime.
func (mr *MockChainParserMockRecorder) UpdateBlockTime(newBlockTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBlockTime", reflect.TypeOf((*MockChainParser)(nil).UpdateBlockTime), newBlockTime)
}

// MockChainMessage is a mock of ChainMessage interface.
type MockChainMessage struct {
	ctrl     *gomock.Controller
	recorder *MockChainMessageMockRecorder
}

// MockChainMessageMockRecorder is the mock recorder for MockChainMessage.
type MockChainMessageMockRecorder struct {
	mock *MockChainMessage
}

// NewMockChainMessage creates a new mock instance.
func NewMockChainMessage(ctrl *gomock.Controller) *MockChainMessage {
	mock := &MockChainMessage{ctrl: ctrl}
	mock.recorder = &MockChainMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainMessage) EXPECT() *MockChainMessageMockRecorder {
	return m.recorder
}

// AppendHeader mocks base method.
func (m *MockChainMessage) GetRequestedBlocksHashes() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRequestedBlocksHashes")
	ret0, _ := ret[0].([]string)
	return ret0
}

// AppendHeader mocks base method.
func (m *MockChainMessage) AppendHeader(metadata []types.Metadata) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AppendHeader", metadata)
}

// AppendHeader indicates an expected call of AppendHeader.
func (mr *MockChainMessageMockRecorder) AppendHeader(metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendHeader", reflect.TypeOf((*MockChainMessage)(nil).AppendHeader), metadata)
}

// CheckResponseError mocks base method.
func (m *MockChainMessage) CheckResponseError(data []byte, httpStatusCode int) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckResponseError", data, httpStatusCode)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// CheckResponseError indicates an expected call of CheckResponseError.
func (mr *MockChainMessageMockRecorder) CheckResponseError(data, httpStatusCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckResponseError", reflect.TypeOf((*MockChainMessage)(nil).CheckResponseError), data, httpStatusCode)
}

// DisableErrorHandling mocks base method.
func (m *MockChainMessage) DisableErrorHandling() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DisableErrorHandling")
}

// DisableErrorHandling indicates an expected call of DisableErrorHandling.
func (mr *MockChainMessageMockRecorder) DisableErrorHandling() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableErrorHandling", reflect.TypeOf((*MockChainMessage)(nil).DisableErrorHandling))
}

// GetApi mocks base method.
func (m *MockChainMessage) GetApi() *types0.Api {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApi")
	ret0, _ := ret[0].(*types0.Api)
	return ret0
}

// GetApi indicates an expected call of GetApi.
func (mr *MockChainMessageMockRecorder) GetApi() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApi", reflect.TypeOf((*MockChainMessage)(nil).GetApi))
}

// GetApiCollection mocks base method.
func (m *MockChainMessage) GetApiCollection() *types0.ApiCollection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiCollection")
	ret0, _ := ret[0].(*types0.ApiCollection)
	return ret0
}

// GetApiCollection indicates an expected call of GetApiCollection.
func (mr *MockChainMessageMockRecorder) GetApiCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiCollection", reflect.TypeOf((*MockChainMessage)(nil).GetApiCollection))
}

// GetExtensions mocks base method.
func (m *MockChainMessage) GetExtensions() []*types0.Extension {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExtensions")
	ret0, _ := ret[0].([]*types0.Extension)
	return ret0
}

// GetExtensions indicates an expected call of GetExtensions.
func (mr *MockChainMessageMockRecorder) GetExtensions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExtensions", reflect.TypeOf((*MockChainMessage)(nil).GetExtensions))
}

// GetForceCacheRefresh mocks base method.
func (m *MockChainMessage) GetForceCacheRefresh() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetForceCacheRefresh")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GetForceCacheRefresh indicates an expected call of GetForceCacheRefresh.
func (mr *MockChainMessageMockRecorder) GetForceCacheRefresh() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetForceCacheRefresh", reflect.TypeOf((*MockChainMessage)(nil).GetForceCacheRefresh))
}

// GetParseDirective mocks base method.
func (m *MockChainMessage) GetParseDirective() *types0.ParseDirective {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParseDirective")
	ret0, _ := ret[0].(*types0.ParseDirective)
	return ret0
}

// GetParseDirective indicates an expected call of GetParseDirective.
func (mr *MockChainMessageMockRecorder) GetParseDirective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParseDirective", reflect.TypeOf((*MockChainMessage)(nil).GetParseDirective))
}

// GetRPCMessage mocks base method.
func (m *MockChainMessage) GetRPCMessage() rpcInterfaceMessages.GenericMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRPCMessage")
	ret0, _ := ret[0].(rpcInterfaceMessages.GenericMessage)
	return ret0
}

// GetRPCMessage indicates an expected call of GetRPCMessage.
func (mr *MockChainMessageMockRecorder) GetRPCMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRPCMessage", reflect.TypeOf((*MockChainMessage)(nil).GetRPCMessage))
}

// GetRawRequestHash mocks base method.
func (m *MockChainMessage) GetRawRequestHash() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawRequestHash")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawRequestHash indicates an expected call of GetRawRequestHash.
func (mr *MockChainMessageMockRecorder) GetRawRequestHash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawRequestHash", reflect.TypeOf((*MockChainMessage)(nil).GetRawRequestHash))
}

// OverrideExtensions mocks base method.
func (m *MockChainMessage) OverrideExtensions(extensionNames []string, extensionParser *extensionslib.ExtensionParser) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OverrideExtensions", extensionNames, extensionParser)
}

// OverrideExtensions indicates an expected call of OverrideExtensions.
func (mr *MockChainMessageMockRecorder) OverrideExtensions(extensionNames, extensionParser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverrideExtensions", reflect.TypeOf((*MockChainMessage)(nil).OverrideExtensions), extensionNames, extensionParser)
}

// RequestedBlock mocks base method.
func (m *MockChainMessage) RequestedBlock() (int64, int64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestedBlock")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	return ret0, ret1
}

// RequestedBlock indicates an expected call of RequestedBlock.
func (mr *MockChainMessageMockRecorder) RequestedBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestedBlock", reflect.TypeOf((*MockChainMessage)(nil).RequestedBlock))
}

// SetForceCacheRefresh mocks base method.
func (m *MockChainMessage) SetForceCacheRefresh(force bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetForceCacheRefresh", force)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SetForceCacheRefresh indicates an expected call of SetForceCacheRefresh.
func (mr *MockChainMessageMockRecorder) SetForceCacheRefresh(force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetForceCacheRefresh", reflect.TypeOf((*MockChainMessage)(nil).SetForceCacheRefresh), force)
}

// SubscriptionIdExtractor mocks base method.
func (m *MockChainMessage) SubscriptionIdExtractor(reply *rpcclient.JsonrpcMessage) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionIdExtractor", reply)
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionIdExtractor indicates an expected call of SubscriptionIdExtractor.
func (mr *MockChainMessageMockRecorder) SubscriptionIdExtractor(reply interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionIdExtractor", reflect.TypeOf((*MockChainMessage)(nil).SubscriptionIdExtractor), reply)
}

// TimeoutOverride mocks base method.
func (m *MockChainMessage) TimeoutOverride(arg0 ...time.Duration) time.Duration {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TimeoutOverride", varargs...)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeoutOverride indicates an expected call of TimeoutOverride.
func (mr *MockChainMessageMockRecorder) TimeoutOverride(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeoutOverride", reflect.TypeOf((*MockChainMessage)(nil).TimeoutOverride), arg0...)
}

// UpdateLatestBlockInMessage mocks base method.
func (m *MockChainMessage) UpdateLatestBlockInMessage(latestBlock int64, modifyContent bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLatestBlockInMessage", latestBlock, modifyContent)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateLatestBlockInMessage indicates an expected call of UpdateLatestBlockInMessage.
func (mr *MockChainMessageMockRecorder) UpdateLatestBlockInMessage(latestBlock, modifyContent interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLatestBlockInMessage", reflect.TypeOf((*MockChainMessage)(nil).UpdateLatestBlockInMessage), latestBlock, modifyContent)
}

// MockChainMessageForSend is a mock of ChainMessageForSend interface.
type MockChainMessageForSend struct {
	ctrl     *gomock.Controller
	recorder *MockChainMessageForSendMockRecorder
}

// MockChainMessageForSendMockRecorder is the mock recorder for MockChainMessageForSend.
type MockChainMessageForSendMockRecorder struct {
	mock *MockChainMessageForSend
}

// NewMockChainMessageForSend creates a new mock instance.
func NewMockChainMessageForSend(ctrl *gomock.Controller) *MockChainMessageForSend {
	mock := &MockChainMessageForSend{ctrl: ctrl}
	mock.recorder = &MockChainMessageForSendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainMessageForSend) EXPECT() *MockChainMessageForSendMockRecorder {
	return m.recorder
}

// CheckResponseError mocks base method.
func (m *MockChainMessageForSend) CheckResponseError(data []byte, httpStatusCode int) (bool, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckResponseError", data, httpStatusCode)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// CheckResponseError indicates an expected call of CheckResponseError.
func (mr *MockChainMessageForSendMockRecorder) CheckResponseError(data, httpStatusCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckResponseError", reflect.TypeOf((*MockChainMessageForSend)(nil).CheckResponseError), data, httpStatusCode)
}

// GetApi mocks base method.
func (m *MockChainMessageForSend) GetApi() *types0.Api {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApi")
	ret0, _ := ret[0].(*types0.Api)
	return ret0
}

// GetApi indicates an expected call of GetApi.
func (mr *MockChainMessageForSendMockRecorder) GetApi() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApi", reflect.TypeOf((*MockChainMessageForSend)(nil).GetApi))
}

// GetApiCollection mocks base method.
func (m *MockChainMessageForSend) GetApiCollection() *types0.ApiCollection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApiCollection")
	ret0, _ := ret[0].(*types0.ApiCollection)
	return ret0
}

// GetApiCollection indicates an expected call of GetApiCollection.
func (mr *MockChainMessageForSendMockRecorder) GetApiCollection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApiCollection", reflect.TypeOf((*MockChainMessageForSend)(nil).GetApiCollection))
}

// GetParseDirective mocks base method.
func (m *MockChainMessageForSend) GetParseDirective() *types0.ParseDirective {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParseDirective")
	ret0, _ := ret[0].(*types0.ParseDirective)
	return ret0
}

// GetParseDirective indicates an expected call of GetParseDirective.
func (mr *MockChainMessageForSendMockRecorder) GetParseDirective() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParseDirective", reflect.TypeOf((*MockChainMessageForSend)(nil).GetParseDirective))
}

// GetRPCMessage mocks base method.
func (m *MockChainMessageForSend) GetRPCMessage() rpcInterfaceMessages.GenericMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRPCMessage")
	ret0, _ := ret[0].(rpcInterfaceMessages.GenericMessage)
	return ret0
}

// GetRPCMessage indicates an expected call of GetRPCMessage.
func (mr *MockChainMessageForSendMockRecorder) GetRPCMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRPCMessage", reflect.TypeOf((*MockChainMessageForSend)(nil).GetRPCMessage))
}

// TimeoutOverride mocks base method.
func (m *MockChainMessageForSend) TimeoutOverride(arg0 ...time.Duration) time.Duration {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TimeoutOverride", varargs...)
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// TimeoutOverride indicates an expected call of TimeoutOverride.
func (mr *MockChainMessageForSendMockRecorder) TimeoutOverride(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeoutOverride", reflect.TypeOf((*MockChainMessageForSend)(nil).TimeoutOverride), arg0...)
}

// MockHealthReporter is a mock of HealthReporter interface.
type MockHealthReporter struct {
	ctrl     *gomock.Controller
	recorder *MockHealthReporterMockRecorder
}

// MockHealthReporterMockRecorder is the mock recorder for MockHealthReporter.
type MockHealthReporterMockRecorder struct {
	mock *MockHealthReporter
}

// NewMockHealthReporter creates a new mock instance.
func NewMockHealthReporter(ctrl *gomock.Controller) *MockHealthReporter {
	mock := &MockHealthReporter{ctrl: ctrl}
	mock.recorder = &MockHealthReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHealthReporter) EXPECT() *MockHealthReporterMockRecorder {
	return m.recorder
}

// IsHealthy mocks base method.
func (m *MockHealthReporter) IsHealthy() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHealthy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsHealthy indicates an expected call of IsHealthy.
func (mr *MockHealthReporterMockRecorder) IsHealthy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealthy", reflect.TypeOf((*MockHealthReporter)(nil).IsHealthy))
}

// MockRelaySender is a mock of RelaySender interface.
type MockRelaySender struct {
	ctrl     *gomock.Controller
	recorder *MockRelaySenderMockRecorder
}

// MockRelaySenderMockRecorder is the mock recorder for MockRelaySender.
type MockRelaySenderMockRecorder struct {
	mock *MockRelaySender
}

// NewMockRelaySender creates a new mock instance.
func NewMockRelaySender(ctrl *gomock.Controller) *MockRelaySender {
	mock := &MockRelaySender{ctrl: ctrl}
	mock.recorder = &MockRelaySenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRelaySender) EXPECT() *MockRelaySenderMockRecorder {
	return m.recorder
}

// CancelSubscriptionContext mocks base method.
func (m *MockRelaySender) CancelSubscriptionContext(subscriptionKey string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CancelSubscriptionContext", subscriptionKey)
}

// CancelSubscriptionContext indicates an expected call of CancelSubscriptionContext.
func (mr *MockRelaySenderMockRecorder) CancelSubscriptionContext(subscriptionKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelSubscriptionContext", reflect.TypeOf((*MockRelaySender)(nil).CancelSubscriptionContext), subscriptionKey)
}

// CreateDappKey mocks base method.
func (m *MockRelaySender) CreateDappKey(userData common.UserData) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDappKey", userData)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateDappKey indicates an expected call of CreateDappKey.
func (mr *MockRelaySenderMockRecorder) CreateDappKey(userData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDappKey", reflect.TypeOf((*MockRelaySender)(nil).CreateDappKey), userData)
}

// ParseRelay mocks base method.
func (m *MockRelaySender) ParseRelay(ctx context.Context, url, req, connectionType, dappID, consumerIp string, metadata []types.Metadata) (ProtocolMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseRelay", ctx, url, req, connectionType, dappID, consumerIp, metadata)
	ret0, _ := ret[0].(ProtocolMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseRelay indicates an expected call of ParseRelay.
func (mr *MockRelaySenderMockRecorder) ParseRelay(ctx, url, req, connectionType, dappID, consumerIp, metadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseRelay", reflect.TypeOf((*MockRelaySender)(nil).ParseRelay), ctx, url, req, connectionType, dappID, consumerIp, metadata)
}

// SendParsedRelay mocks base method.
func (m *MockRelaySender) SendParsedRelay(ctx context.Context, analytics *metrics.RelayMetrics, protocolMessage ProtocolMessage) (*common.RelayResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendParsedRelay", ctx, analytics, protocolMessage)
	ret0, _ := ret[0].(*common.RelayResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendParsedRelay indicates an expected call of SendParsedRelay.
func (mr *MockRelaySenderMockRecorder) SendParsedRelay(ctx, analytics, protocolMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendParsedRelay", reflect.TypeOf((*MockRelaySender)(nil).SendParsedRelay), ctx, analytics, protocolMessage)
}

// SendRelay mocks base method.
func (m *MockRelaySender) SendRelay(ctx context.Context, url, req, connectionType, dappID, consumerIp string, analytics *metrics.RelayMetrics, metadataValues []types.Metadata) (*common.RelayResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRelay", ctx, url, req, connectionType, dappID, consumerIp, analytics, metadataValues)
	ret0, _ := ret[0].(*common.RelayResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendRelay indicates an expected call of SendRelay.
func (mr *MockRelaySenderMockRecorder) SendRelay(ctx, url, req, connectionType, dappID, consumerIp, analytics, metadataValues interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRelay", reflect.TypeOf((*MockRelaySender)(nil).SendRelay), ctx, url, req, connectionType, dappID, consumerIp, analytics, metadataValues)
}

// SetConsistencySeenBlock mocks base method.
func (m *MockRelaySender) SetConsistencySeenBlock(blockSeen int64, key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetConsistencySeenBlock", blockSeen, key)
}

// SetConsistencySeenBlock indicates an expected call of SetConsistencySeenBlock.
func (mr *MockRelaySenderMockRecorder) SetConsistencySeenBlock(blockSeen, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConsistencySeenBlock", reflect.TypeOf((*MockRelaySender)(nil).SetConsistencySeenBlock), blockSeen, key)
}

// MockChainListener is a mock of ChainListener interface.
type MockChainListener struct {
	ctrl     *gomock.Controller
	recorder *MockChainListenerMockRecorder
}

// MockChainListenerMockRecorder is the mock recorder for MockChainListener.
type MockChainListenerMockRecorder struct {
	mock *MockChainListener
}

// NewMockChainListener creates a new mock instance.
func NewMockChainListener(ctrl *gomock.Controller) *MockChainListener {
	mock := &MockChainListener{ctrl: ctrl}
	mock.recorder = &MockChainListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainListener) EXPECT() *MockChainListenerMockRecorder {
	return m.recorder
}

// GetListeningAddress mocks base method.
func (m *MockChainListener) GetListeningAddress() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListeningAddress")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetListeningAddress indicates an expected call of GetListeningAddress.
func (mr *MockChainListenerMockRecorder) GetListeningAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListeningAddress", reflect.TypeOf((*MockChainListener)(nil).GetListeningAddress))
}

// Serve mocks base method.
func (m *MockChainListener) Serve(ctx context.Context, cmdFlags common.ConsumerCmdFlags) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Serve", ctx, cmdFlags)
}

// Serve indicates an expected call of Serve.
func (mr *MockChainListenerMockRecorder) Serve(ctx, cmdFlags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockChainListener)(nil).Serve), ctx, cmdFlags)
}

// MockChainRouter is a mock of ChainRouter interface.
type MockChainRouter struct {
	ctrl     *gomock.Controller
	recorder *MockChainRouterMockRecorder
}

// MockChainRouterMockRecorder is the mock recorder for MockChainRouter.
type MockChainRouterMockRecorder struct {
	mock *MockChainRouter
}

// NewMockChainRouter creates a new mock instance.
func NewMockChainRouter(ctrl *gomock.Controller) *MockChainRouter {
	mock := &MockChainRouter{ctrl: ctrl}
	mock.recorder = &MockChainRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainRouter) EXPECT() *MockChainRouterMockRecorder {
	return m.recorder
}

// ExtensionsSupported mocks base method.
func (m *MockChainRouter) ExtensionsSupported(arg0 []string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtensionsSupported", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ExtensionsSupported indicates an expected call of ExtensionsSupported.
func (mr *MockChainRouterMockRecorder) ExtensionsSupported(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtensionsSupported", reflect.TypeOf((*MockChainRouter)(nil).ExtensionsSupported), arg0)
}

// SendNodeMsg mocks base method.
func (m *MockChainRouter) SendNodeMsg(ctx context.Context, ch chan interface{}, chainMessage ChainMessageForSend, extensions []string) (*RelayReplyWrapper, string, *rpcclient.ClientSubscription, common.NodeUrl, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendNodeMsg", ctx, ch, chainMessage, extensions)
	ret0, _ := ret[0].(*RelayReplyWrapper)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*rpcclient.ClientSubscription)
	ret3, _ := ret[3].(common.NodeUrl)
	ret4, _ := ret[4].(string)
	ret5, _ := ret[5].(error)
	return ret0, ret1, ret2, ret3, ret4, ret5
}

// SendNodeMsg indicates an expected call of SendNodeMsg.
func (mr *MockChainRouterMockRecorder) SendNodeMsg(ctx, ch, chainMessage, extensions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNodeMsg", reflect.TypeOf((*MockChainRouter)(nil).SendNodeMsg), ctx, ch, chainMessage, extensions)
}

// MockChainProxy is a mock of ChainProxy interface.
type MockChainProxy struct {
	ctrl     *gomock.Controller
	recorder *MockChainProxyMockRecorder
}

// MockChainProxyMockRecorder is the mock recorder for MockChainProxy.
type MockChainProxyMockRecorder struct {
	mock *MockChainProxy
}

// NewMockChainProxy creates a new mock instance.
func NewMockChainProxy(ctrl *gomock.Controller) *MockChainProxy {
	mock := &MockChainProxy{ctrl: ctrl}
	mock.recorder = &MockChainProxyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainProxy) EXPECT() *MockChainProxyMockRecorder {
	return m.recorder
}

// GetChainProxyInformation mocks base method.
func (m *MockChainProxy) GetChainProxyInformation() (common.NodeUrl, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainProxyInformation")
	ret0, _ := ret[0].(common.NodeUrl)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// GetChainProxyInformation indicates an expected call of GetChainProxyInformation.
func (mr *MockChainProxyMockRecorder) GetChainProxyInformation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainProxyInformation", reflect.TypeOf((*MockChainProxy)(nil).GetChainProxyInformation))
}

// SendNodeMsg mocks base method.
func (m *MockChainProxy) SendNodeMsg(ctx context.Context, ch chan interface{}, chainMessage ChainMessageForSend) (*RelayReplyWrapper, string, *rpcclient.ClientSubscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendNodeMsg", ctx, ch, chainMessage)
	ret0, _ := ret[0].(*RelayReplyWrapper)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*rpcclient.ClientSubscription)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SendNodeMsg indicates an expected call of SendNodeMsg.
func (mr *MockChainProxyMockRecorder) SendNodeMsg(ctx, ch, chainMessage interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendNodeMsg", reflect.TypeOf((*MockChainProxy)(nil).SendNodeMsg), ctx, ch, chainMessage)
}

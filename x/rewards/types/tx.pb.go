// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/rewards/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSetIprpcData struct {
	Authority          string     `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	MinIprpcCost       types.Coin `protobuf:"bytes,2,opt,name=min_iprpc_cost,json=minIprpcCost,proto3" json:"min_iprpc_cost"`
	IprpcSubscriptions []string   `protobuf:"bytes,3,rep,name=iprpc_subscriptions,json=iprpcSubscriptions,proto3" json:"iprpc_subscriptions,omitempty"`
}

func (m *MsgSetIprpcData) Reset()         { *m = MsgSetIprpcData{} }
func (m *MsgSetIprpcData) String() string { return proto.CompactTextString(m) }
func (*MsgSetIprpcData) ProtoMessage()    {}
func (*MsgSetIprpcData) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4c66e189226d78, []int{0}
}
func (m *MsgSetIprpcData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetIprpcData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetIprpcData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetIprpcData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetIprpcData.Merge(m, src)
}
func (m *MsgSetIprpcData) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetIprpcData) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetIprpcData.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetIprpcData proto.InternalMessageInfo

func (m *MsgSetIprpcData) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSetIprpcData) GetMinIprpcCost() types.Coin {
	if m != nil {
		return m.MinIprpcCost
	}
	return types.Coin{}
}

func (m *MsgSetIprpcData) GetIprpcSubscriptions() []string {
	if m != nil {
		return m.IprpcSubscriptions
	}
	return nil
}

type MsgSetIprpcDataResponse struct {
}

func (m *MsgSetIprpcDataResponse) Reset()         { *m = MsgSetIprpcDataResponse{} }
func (m *MsgSetIprpcDataResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetIprpcDataResponse) ProtoMessage()    {}
func (*MsgSetIprpcDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4c66e189226d78, []int{1}
}
func (m *MsgSetIprpcDataResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetIprpcDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetIprpcDataResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetIprpcDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetIprpcDataResponse.Merge(m, src)
}
func (m *MsgSetIprpcDataResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetIprpcDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetIprpcDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetIprpcDataResponse proto.InternalMessageInfo

type MsgFundIprpc struct {
	Creator  string                                   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Duration uint64                                   `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Amounts  github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=amounts,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amounts"`
	Spec     string                                   `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (m *MsgFundIprpc) Reset()         { *m = MsgFundIprpc{} }
func (m *MsgFundIprpc) String() string { return proto.CompactTextString(m) }
func (*MsgFundIprpc) ProtoMessage()    {}
func (*MsgFundIprpc) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4c66e189226d78, []int{2}
}
func (m *MsgFundIprpc) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundIprpc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundIprpc.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundIprpc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundIprpc.Merge(m, src)
}
func (m *MsgFundIprpc) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundIprpc) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundIprpc.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundIprpc proto.InternalMessageInfo

func (m *MsgFundIprpc) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgFundIprpc) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *MsgFundIprpc) GetAmounts() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amounts
	}
	return nil
}

func (m *MsgFundIprpc) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

type MsgFundIprpcResponse struct {
}

func (m *MsgFundIprpcResponse) Reset()         { *m = MsgFundIprpcResponse{} }
func (m *MsgFundIprpcResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFundIprpcResponse) ProtoMessage()    {}
func (*MsgFundIprpcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4c66e189226d78, []int{3}
}
func (m *MsgFundIprpcResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFundIprpcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFundIprpcResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFundIprpcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFundIprpcResponse.Merge(m, src)
}
func (m *MsgFundIprpcResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFundIprpcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFundIprpcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFundIprpcResponse proto.InternalMessageInfo

type MsgCoverIbcIprpcFundCost struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Index   uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *MsgCoverIbcIprpcFundCost) Reset()         { *m = MsgCoverIbcIprpcFundCost{} }
func (m *MsgCoverIbcIprpcFundCost) String() string { return proto.CompactTextString(m) }
func (*MsgCoverIbcIprpcFundCost) ProtoMessage()    {}
func (*MsgCoverIbcIprpcFundCost) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4c66e189226d78, []int{4}
}
func (m *MsgCoverIbcIprpcFundCost) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoverIbcIprpcFundCost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoverIbcIprpcFundCost.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoverIbcIprpcFundCost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoverIbcIprpcFundCost.Merge(m, src)
}
func (m *MsgCoverIbcIprpcFundCost) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoverIbcIprpcFundCost) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoverIbcIprpcFundCost.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoverIbcIprpcFundCost proto.InternalMessageInfo

func (m *MsgCoverIbcIprpcFundCost) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCoverIbcIprpcFundCost) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type MsgCoverIbcIprpcFundCostResponse struct {
}

func (m *MsgCoverIbcIprpcFundCostResponse) Reset()         { *m = MsgCoverIbcIprpcFundCostResponse{} }
func (m *MsgCoverIbcIprpcFundCostResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCoverIbcIprpcFundCostResponse) ProtoMessage()    {}
func (*MsgCoverIbcIprpcFundCostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4c66e189226d78, []int{5}
}
func (m *MsgCoverIbcIprpcFundCostResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoverIbcIprpcFundCostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoverIbcIprpcFundCostResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoverIbcIprpcFundCostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoverIbcIprpcFundCostResponse.Merge(m, src)
}
func (m *MsgCoverIbcIprpcFundCostResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoverIbcIprpcFundCostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoverIbcIprpcFundCostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoverIbcIprpcFundCostResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetIprpcData)(nil), "lavanet.lava.rewards.MsgSetIprpcData")
	proto.RegisterType((*MsgSetIprpcDataResponse)(nil), "lavanet.lava.rewards.MsgSetIprpcDataResponse")
	proto.RegisterType((*MsgFundIprpc)(nil), "lavanet.lava.rewards.MsgFundIprpc")
	proto.RegisterType((*MsgFundIprpcResponse)(nil), "lavanet.lava.rewards.MsgFundIprpcResponse")
	proto.RegisterType((*MsgCoverIbcIprpcFundCost)(nil), "lavanet.lava.rewards.MsgCoverIbcIprpcFundCost")
	proto.RegisterType((*MsgCoverIbcIprpcFundCostResponse)(nil), "lavanet.lava.rewards.MsgCoverIbcIprpcFundCostResponse")
}

func init() { proto.RegisterFile("lavanet/lava/rewards/tx.proto", fileDescriptor_6a4c66e189226d78) }

var fileDescriptor_6a4c66e189226d78 = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xd6, 0xc2, 0xa8, 0xa9, 0x40, 0x98, 0xc2, 0xd2, 0x68, 0x64, 0x55, 0x24, 0xa4, 0xaa,
	0xd2, 0x6c, 0xb6, 0x49, 0x1c, 0x76, 0x5c, 0x07, 0xd2, 0x90, 0x7a, 0xc9, 0x6e, 0x70, 0x98, 0x9c,
	0xc4, 0xca, 0x2c, 0x88, 0x1d, 0xc5, 0x4e, 0xe9, 0x4e, 0xdc, 0x39, 0xf1, 0x77, 0x70, 0xe2, 0x9f,
	0x40, 0x1a, 0xb7, 0x1d, 0x11, 0x07, 0x40, 0xed, 0x81, 0x7f, 0x03, 0xd9, 0x49, 0xda, 0xae, 0x6a,
	0xf9, 0x71, 0x89, 0xfd, 0xfc, 0x7d, 0x7e, 0xef, 0xfb, 0xec, 0x17, 0x83, 0x47, 0x6f, 0xc8, 0x88,
	0x70, 0xaa, 0xb0, 0x1e, 0x71, 0x46, 0xdf, 0x92, 0x2c, 0x92, 0x58, 0x8d, 0x51, 0x9a, 0x09, 0x25,
	0x60, 0xbb, 0x84, 0x91, 0x1e, 0x51, 0x09, 0x3b, 0xf7, 0x48, 0xc2, 0xb8, 0xc0, 0xe6, 0x5b, 0x10,
	0x1d, 0x37, 0x14, 0x32, 0x11, 0x12, 0x07, 0x44, 0x52, 0x3c, 0xda, 0x0b, 0xa8, 0x22, 0x7b, 0x38,
	0x14, 0x8c, 0x97, 0x78, 0x3b, 0x16, 0xb1, 0x30, 0x53, 0xac, 0x67, 0xc5, 0xaa, 0xf7, 0xd9, 0x02,
	0x77, 0x87, 0x32, 0x3e, 0xa5, 0xea, 0x24, 0xcd, 0xd2, 0xf0, 0x98, 0x28, 0x02, 0xb7, 0x41, 0x93,
	0xe4, 0xea, 0x5c, 0x64, 0x4c, 0x5d, 0xd8, 0x56, 0xd7, 0xea, 0x35, 0xfd, 0xf9, 0x02, 0x7c, 0x06,
	0xee, 0x24, 0x8c, 0x9f, 0x31, 0x4d, 0x3f, 0x0b, 0x85, 0x54, 0xf6, 0x46, 0xd7, 0xea, 0xdd, 0xde,
	0xef, 0xa0, 0x42, 0x00, 0xd2, 0x02, 0x50, 0x29, 0x00, 0x0d, 0x04, 0xe3, 0x47, 0x8d, 0xcb, 0xef,
	0x3b, 0x35, 0xbf, 0x95, 0x30, 0x6e, 0x8a, 0x0c, 0x84, 0x54, 0x10, 0x83, 0xfb, 0x45, 0x0a, 0x99,
	0x07, 0x32, 0xcc, 0x58, 0xaa, 0x98, 0xe0, 0xd2, 0xae, 0x77, 0xeb, 0xbd, 0xa6, 0x0f, 0x0d, 0x74,
	0xba, 0x88, 0x1c, 0x6e, 0xbf, 0xff, 0xf5, 0xa9, 0xbf, 0x55, 0x9d, 0xcf, 0x92, 0x66, 0xaf, 0x03,
	0xb6, 0x96, 0x96, 0x7c, 0x2a, 0x53, 0xc1, 0x25, 0xf5, 0xbe, 0x59, 0xa0, 0x35, 0x94, 0xf1, 0xf3,
	0x9c, 0x47, 0x06, 0x84, 0x36, 0xd8, 0x0c, 0x33, 0x4a, 0x94, 0xc8, 0x4a, 0x77, 0x55, 0x08, 0x1d,
	0x70, 0x2b, 0xca, 0x33, 0xa2, 0x0b, 0x1a, 0x57, 0x0d, 0x7f, 0x16, 0x43, 0x0a, 0x36, 0x49, 0x22,
	0x72, 0xae, 0x0a, 0x91, 0x7f, 0x34, 0xfc, 0x44, 0x1b, 0xfe, 0xf8, 0x63, 0xa7, 0x17, 0x33, 0x75,
	0x9e, 0x07, 0x28, 0x14, 0x09, 0x2e, 0xaf, 0xa7, 0x18, 0x76, 0x65, 0xf4, 0x1a, 0xab, 0x8b, 0x94,
	0x4a, 0xb3, 0x41, 0xfa, 0x55, 0x6e, 0x08, 0x41, 0x43, 0xa6, 0x34, 0xb4, 0x1b, 0x46, 0x99, 0x99,
	0x1f, 0x76, 0xb4, 0xf5, 0xf6, 0x82, 0xf5, 0x99, 0x17, 0xef, 0x21, 0x68, 0x2f, 0xc6, 0x33, 0xd3,
	0x2f, 0x80, 0x3d, 0x94, 0xf1, 0x40, 0x8c, 0x68, 0x76, 0x12, 0x84, 0x06, 0xd3, 0x24, 0x73, 0xf4,
	0xeb, 0xfd, 0xb7, 0xc1, 0x0d, 0xc6, 0x23, 0x3a, 0x2e, 0xcd, 0x17, 0x81, 0xe7, 0x81, 0xee, 0xba,
	0x5c, 0x55, 0xbd, 0xfd, 0x2f, 0x1b, 0xa0, 0x3e, 0x94, 0x31, 0x8c, 0x40, 0xeb, 0x5a, 0x2f, 0x3d,
	0x46, 0xab, 0xfa, 0x17, 0x2d, 0xdd, 0x95, 0xb3, 0xfb, 0x4f, 0xb4, 0xaa, 0x1a, 0x7c, 0x05, 0x9a,
	0xf3, 0xeb, 0xf4, 0xd6, 0xee, 0x9d, 0x71, 0x9c, 0xfe, 0xdf, 0x39, 0xb3, 0xe4, 0xef, 0xc0, 0x83,
	0xd5, 0xe7, 0x86, 0xd6, 0x26, 0x59, 0xc9, 0x77, 0x9e, 0xfe, 0x1f, 0xbf, 0x12, 0x70, 0x74, 0x7c,
	0x39, 0x71, 0xad, 0xab, 0x89, 0x6b, 0xfd, 0x9c, 0xb8, 0xd6, 0x87, 0xa9, 0x5b, 0xbb, 0x9a, 0xba,
	0xb5, 0xaf, 0x53, 0xb7, 0xf6, 0xb2, 0xbf, 0xd0, 0x4f, 0xd7, 0x9e, 0x8d, 0xd1, 0x01, 0x1e, 0xcf,
	0xdf, 0x0e, 0xdd, 0x57, 0xc1, 0x4d, 0xf3, 0x83, 0x1f, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x26, 0xcd, 0x02, 0x60, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SetIprpcData(ctx context.Context, in *MsgSetIprpcData, opts ...grpc.CallOption) (*MsgSetIprpcDataResponse, error)
	FundIprpc(ctx context.Context, in *MsgFundIprpc, opts ...grpc.CallOption) (*MsgFundIprpcResponse, error)
	CoverIbcIprpcFundCost(ctx context.Context, in *MsgCoverIbcIprpcFundCost, opts ...grpc.CallOption) (*MsgCoverIbcIprpcFundCostResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetIprpcData(ctx context.Context, in *MsgSetIprpcData, opts ...grpc.CallOption) (*MsgSetIprpcDataResponse, error) {
	out := new(MsgSetIprpcDataResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.rewards.Msg/SetIprpcData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) FundIprpc(ctx context.Context, in *MsgFundIprpc, opts ...grpc.CallOption) (*MsgFundIprpcResponse, error) {
	out := new(MsgFundIprpcResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.rewards.Msg/FundIprpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CoverIbcIprpcFundCost(ctx context.Context, in *MsgCoverIbcIprpcFundCost, opts ...grpc.CallOption) (*MsgCoverIbcIprpcFundCostResponse, error) {
	out := new(MsgCoverIbcIprpcFundCostResponse)
	err := c.cc.Invoke(ctx, "/lavanet.lava.rewards.Msg/CoverIbcIprpcFundCost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SetIprpcData(context.Context, *MsgSetIprpcData) (*MsgSetIprpcDataResponse, error)
	FundIprpc(context.Context, *MsgFundIprpc) (*MsgFundIprpcResponse, error)
	CoverIbcIprpcFundCost(context.Context, *MsgCoverIbcIprpcFundCost) (*MsgCoverIbcIprpcFundCostResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetIprpcData(ctx context.Context, req *MsgSetIprpcData) (*MsgSetIprpcDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetIprpcData not implemented")
}
func (*UnimplementedMsgServer) FundIprpc(ctx context.Context, req *MsgFundIprpc) (*MsgFundIprpcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FundIprpc not implemented")
}
func (*UnimplementedMsgServer) CoverIbcIprpcFundCost(ctx context.Context, req *MsgCoverIbcIprpcFundCost) (*MsgCoverIbcIprpcFundCostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoverIbcIprpcFundCost not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetIprpcData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetIprpcData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetIprpcData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.rewards.Msg/SetIprpcData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetIprpcData(ctx, req.(*MsgSetIprpcData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_FundIprpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFundIprpc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FundIprpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.rewards.Msg/FundIprpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FundIprpc(ctx, req.(*MsgFundIprpc))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CoverIbcIprpcFundCost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCoverIbcIprpcFundCost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CoverIbcIprpcFundCost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.rewards.Msg/CoverIbcIprpcFundCost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CoverIbcIprpcFundCost(ctx, req.(*MsgCoverIbcIprpcFundCost))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lavanet.lava.rewards.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetIprpcData",
			Handler:    _Msg_SetIprpcData_Handler,
		},
		{
			MethodName: "FundIprpc",
			Handler:    _Msg_FundIprpc_Handler,
		},
		{
			MethodName: "CoverIbcIprpcFundCost",
			Handler:    _Msg_CoverIbcIprpcFundCost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lavanet/lava/rewards/tx.proto",
}

func (m *MsgSetIprpcData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetIprpcData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetIprpcData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IprpcSubscriptions) > 0 {
		for iNdEx := len(m.IprpcSubscriptions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IprpcSubscriptions[iNdEx])
			copy(dAtA[i:], m.IprpcSubscriptions[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.IprpcSubscriptions[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.MinIprpcCost.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetIprpcDataResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetIprpcDataResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetIprpcDataResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgFundIprpc) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundIprpc) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundIprpc) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Spec) > 0 {
		i -= len(m.Spec)
		copy(dAtA[i:], m.Spec)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Spec)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Amounts) > 0 {
		for iNdEx := len(m.Amounts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amounts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Duration != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFundIprpcResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFundIprpcResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFundIprpcResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCoverIbcIprpcFundCost) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoverIbcIprpcFundCost) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoverIbcIprpcFundCost) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Index != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCoverIbcIprpcFundCostResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoverIbcIprpcFundCostResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoverIbcIprpcFundCostResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetIprpcData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.MinIprpcCost.Size()
	n += 1 + l + sovTx(uint64(l))
	if len(m.IprpcSubscriptions) > 0 {
		for _, s := range m.IprpcSubscriptions {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSetIprpcDataResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgFundIprpc) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovTx(uint64(m.Duration))
	}
	if len(m.Amounts) > 0 {
		for _, e := range m.Amounts {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Spec)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFundIprpcResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCoverIbcIprpcFundCost) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovTx(uint64(m.Index))
	}
	return n
}

func (m *MsgCoverIbcIprpcFundCostResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetIprpcData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetIprpcData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetIprpcData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinIprpcCost", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinIprpcCost.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IprpcSubscriptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IprpcSubscriptions = append(m.IprpcSubscriptions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetIprpcDataResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetIprpcDataResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetIprpcDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgFundIprpc) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgFundIprpc: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundIprpc: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amounts = append(m.Amounts, types.Coin{})
			if err := m.Amounts[len(m.Amounts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgFundIprpcResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgFundIprpcResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFundIprpcResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCoverIbcIprpcFundCost) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCoverIbcIprpcFundCost: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoverIbcIprpcFundCost: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCoverIbcIprpcFundCostResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCoverIbcIprpcFundCostResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoverIbcIprpcFundCostResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

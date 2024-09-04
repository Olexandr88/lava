// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/rewards/iprpc.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// object that holds the list for iprpc funcs for a specific month id
type IprpcReward struct {
	Id        uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SpecFunds []Specfund `protobuf:"bytes,2,rep,name=spec_funds,json=specFunds,proto3" json:"spec_funds"`
}

func (m *IprpcReward) Reset()         { *m = IprpcReward{} }
func (m *IprpcReward) String() string { return proto.CompactTextString(m) }
func (*IprpcReward) ProtoMessage()    {}
func (*IprpcReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_1293618a311573f7, []int{0}
}
func (m *IprpcReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IprpcReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IprpcReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IprpcReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IprpcReward.Merge(m, src)
}
func (m *IprpcReward) XXX_Size() int {
	return m.Size()
}
func (m *IprpcReward) XXX_DiscardUnknown() {
	xxx_messageInfo_IprpcReward.DiscardUnknown(m)
}

var xxx_messageInfo_IprpcReward proto.InternalMessageInfo

func (m *IprpcReward) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IprpcReward) GetSpecFunds() []Specfund {
	if m != nil {
		return m.SpecFunds
	}
	return nil
}

type Specfund struct {
	Spec string                                   `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Fund github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=fund,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fund"`
}

func (m *Specfund) Reset()         { *m = Specfund{} }
func (m *Specfund) String() string { return proto.CompactTextString(m) }
func (*Specfund) ProtoMessage()    {}
func (*Specfund) Descriptor() ([]byte, []int) {
	return fileDescriptor_1293618a311573f7, []int{1}
}
func (m *Specfund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Specfund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Specfund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Specfund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specfund.Merge(m, src)
}
func (m *Specfund) XXX_Size() int {
	return m.Size()
}
func (m *Specfund) XXX_DiscardUnknown() {
	xxx_messageInfo_Specfund.DiscardUnknown(m)
}

var xxx_messageInfo_Specfund proto.InternalMessageInfo

func (m *Specfund) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *Specfund) GetFund() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Fund
	}
	return nil
}

type IprpcMemo struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Spec     string `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Duration uint64 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (m *IprpcMemo) Reset()         { *m = IprpcMemo{} }
func (m *IprpcMemo) String() string { return proto.CompactTextString(m) }
func (*IprpcMemo) ProtoMessage()    {}
func (*IprpcMemo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1293618a311573f7, []int{2}
}
func (m *IprpcMemo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IprpcMemo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IprpcMemo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IprpcMemo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IprpcMemo.Merge(m, src)
}
func (m *IprpcMemo) XXX_Size() int {
	return m.Size()
}
func (m *IprpcMemo) XXX_DiscardUnknown() {
	xxx_messageInfo_IprpcMemo.DiscardUnknown(m)
}

var xxx_messageInfo_IprpcMemo proto.InternalMessageInfo

func (m *IprpcMemo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *IprpcMemo) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *IprpcMemo) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type PendingIbcIprpcFund struct {
	Index    uint64     `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator  string     `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Spec     string     `protobuf:"bytes,3,opt,name=spec,proto3" json:"spec,omitempty"`
	Duration uint64     `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Fund     types.Coin `protobuf:"bytes,5,opt,name=fund,proto3" json:"fund"`
	Expiry   uint64     `protobuf:"varint,6,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (m *PendingIbcIprpcFund) Reset()         { *m = PendingIbcIprpcFund{} }
func (m *PendingIbcIprpcFund) String() string { return proto.CompactTextString(m) }
func (*PendingIbcIprpcFund) ProtoMessage()    {}
func (*PendingIbcIprpcFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_1293618a311573f7, []int{3}
}
func (m *PendingIbcIprpcFund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingIbcIprpcFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingIbcIprpcFund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingIbcIprpcFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingIbcIprpcFund.Merge(m, src)
}
func (m *PendingIbcIprpcFund) XXX_Size() int {
	return m.Size()
}
func (m *PendingIbcIprpcFund) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingIbcIprpcFund.DiscardUnknown(m)
}

var xxx_messageInfo_PendingIbcIprpcFund proto.InternalMessageInfo

func (m *PendingIbcIprpcFund) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PendingIbcIprpcFund) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *PendingIbcIprpcFund) GetSpec() string {
	if m != nil {
		return m.Spec
	}
	return ""
}

func (m *PendingIbcIprpcFund) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *PendingIbcIprpcFund) GetFund() types.Coin {
	if m != nil {
		return m.Fund
	}
	return types.Coin{}
}

func (m *PendingIbcIprpcFund) GetExpiry() uint64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func init() {
	proto.RegisterType((*IprpcReward)(nil), "lavanet.lava.rewards.IprpcReward")
	proto.RegisterType((*Specfund)(nil), "lavanet.lava.rewards.Specfund")
	proto.RegisterType((*IprpcMemo)(nil), "lavanet.lava.rewards.IprpcMemo")
	proto.RegisterType((*PendingIbcIprpcFund)(nil), "lavanet.lava.rewards.PendingIbcIprpcFund")
}

func init() { proto.RegisterFile("lavanet/lava/rewards/iprpc.proto", fileDescriptor_1293618a311573f7) }

var fileDescriptor_1293618a311573f7 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x4d, 0xd2, 0xbc, 0xf2, 0xea, 0x27, 0x31, 0x98, 0x0a, 0x85, 0x0e, 0x7e, 0x51, 0xa6, 0x0a,
	0x09, 0x9b, 0xd2, 0x3f, 0x68, 0x11, 0x52, 0x07, 0x24, 0x14, 0xc4, 0xc2, 0x52, 0x25, 0xb6, 0x09,
	0x16, 0xd4, 0x8e, 0xec, 0xb4, 0xb4, 0x13, 0xbf, 0xc0, 0x77, 0xf0, 0x19, 0x4c, 0x1d, 0x3b, 0x32,
	0x01, 0x6a, 0x7f, 0x04, 0xd9, 0x71, 0xab, 0x56, 0x2a, 0x4c, 0xce, 0xd5, 0x3d, 0xf7, 0x9c, 0x73,
	0x4f, 0x2e, 0x48, 0x3f, 0x17, 0xab, 0x42, 0xf2, 0x86, 0xd8, 0x97, 0x68, 0xfe, 0xa5, 0xd0, 0xcc,
	0x10, 0x51, 0xeb, 0x9a, 0xe2, 0x5a, 0xab, 0x46, 0xc1, 0xbe, 0x47, 0x60, 0xfb, 0x62, 0x8f, 0x18,
	0xf4, 0x2b, 0x55, 0x29, 0x07, 0x20, 0xf6, 0xab, 0xc5, 0x0e, 0x10, 0x55, 0x66, 0xa1, 0x0c, 0x29,
	0x0b, 0xc3, 0xc9, 0x6a, 0x54, 0xf2, 0xa6, 0x18, 0x11, 0xaa, 0x84, 0x6c, 0xfb, 0x59, 0x09, 0xee,
	0x66, 0x96, 0x3a, 0x77, 0x2c, 0xf0, 0x21, 0x88, 0x04, 0x4b, 0xc2, 0x34, 0x1c, 0xc6, 0x79, 0x24,
	0x18, 0x9c, 0x02, 0x60, 0x6a, 0x4e, 0xe7, 0x1f, 0x96, 0x92, 0x99, 0x24, 0x4a, 0x3b, 0xc3, 0xbb,
	0x17, 0x08, 0x5f, 0xd3, 0xc7, 0x6f, 0x6b, 0x4e, 0x2d, 0x6c, 0x12, 0x6f, 0x7f, 0xdd, 0x07, 0x79,
	0xcf, 0xce, 0xbd, 0xb2, 0x63, 0xd9, 0x57, 0x70, 0x7b, 0x6c, 0x42, 0x08, 0x62, 0xdb, 0x70, 0x12,
	0xbd, 0xdc, 0x7d, 0xc3, 0x39, 0x88, 0x6d, 0xcf, 0xd3, 0x3f, 0xc1, 0xad, 0x65, 0x6c, 0x2d, 0x63,
	0x6f, 0x19, 0x4f, 0x95, 0x90, 0x93, 0xe7, 0x96, 0xf9, 0xfb, 0xef, 0xfb, 0x61, 0x25, 0x9a, 0x8f,
	0xcb, 0x12, 0x53, 0xb5, 0x20, 0x7e, 0xbf, 0xf6, 0x79, 0x66, 0xd8, 0x27, 0xd2, 0x6c, 0x6a, 0x6e,
	0xdc, 0x80, 0xc9, 0x1d, 0x71, 0xf6, 0x0e, 0xf4, 0xdc, 0x92, 0xaf, 0xf9, 0x42, 0xc1, 0x04, 0x3c,
	0xa0, 0x9a, 0x17, 0x8d, 0xd2, 0xde, 0xc4, 0xb1, 0x3c, 0x79, 0x8b, 0xce, 0xbc, 0x0d, 0xc0, 0x2d,
	0x5b, 0xea, 0xa2, 0x11, 0x4a, 0x26, 0x1d, 0x17, 0xcb, 0xa9, 0xce, 0x7e, 0x84, 0xe0, 0xd1, 0x1b,
	0x2e, 0x99, 0x90, 0xd5, 0xac, 0xa4, 0x4e, 0xc1, 0x2e, 0x0c, 0xfb, 0xe0, 0x46, 0x48, 0xc6, 0xd7,
	0x3e, 0xc7, 0xb6, 0x38, 0xd7, 0x8d, 0xae, 0xeb, 0x76, 0xfe, 0xa1, 0x1b, 0x5f, 0xea, 0xc2, 0xb1,
	0xcf, 0xeb, 0x26, 0x0d, 0xff, 0x9f, 0x57, 0xfb, 0x27, 0x1c, 0x18, 0x3e, 0x06, 0x5d, 0xbe, 0xae,
	0x85, 0xde, 0x24, 0x5d, 0x47, 0xe7, 0xab, 0xc9, 0xcb, 0xed, 0x1e, 0x85, 0xbb, 0x3d, 0x0a, 0xff,
	0xec, 0x51, 0xf8, 0xed, 0x80, 0x82, 0xdd, 0x01, 0x05, 0x3f, 0x0f, 0x28, 0x78, 0xff, 0xf4, 0x2c,
	0xe5, 0x8b, 0x9b, 0x5c, 0x8d, 0xc9, 0xfa, 0x74, 0x98, 0x2e, 0xed, 0xb2, 0xeb, 0xae, 0x69, 0xfc,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x17, 0x17, 0x4e, 0xbd, 0x02, 0x00, 0x00,
}

func (m *IprpcReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IprpcReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IprpcReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpecFunds) > 0 {
		for iNdEx := len(m.SpecFunds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpecFunds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIprpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Id != 0 {
		i = encodeVarintIprpc(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Specfund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Specfund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Specfund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fund) > 0 {
		for iNdEx := len(m.Fund) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fund[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIprpc(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Spec) > 0 {
		i -= len(m.Spec)
		copy(dAtA[i:], m.Spec)
		i = encodeVarintIprpc(dAtA, i, uint64(len(m.Spec)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IprpcMemo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IprpcMemo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IprpcMemo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Duration != 0 {
		i = encodeVarintIprpc(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Spec) > 0 {
		i -= len(m.Spec)
		copy(dAtA[i:], m.Spec)
		i = encodeVarintIprpc(dAtA, i, uint64(len(m.Spec)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIprpc(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PendingIbcIprpcFund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingIbcIprpcFund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingIbcIprpcFund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiry != 0 {
		i = encodeVarintIprpc(dAtA, i, uint64(m.Expiry))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.Fund.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIprpc(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Duration != 0 {
		i = encodeVarintIprpc(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Spec) > 0 {
		i -= len(m.Spec)
		copy(dAtA[i:], m.Spec)
		i = encodeVarintIprpc(dAtA, i, uint64(len(m.Spec)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintIprpc(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintIprpc(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIprpc(dAtA []byte, offset int, v uint64) int {
	offset -= sovIprpc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *IprpcReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovIprpc(uint64(m.Id))
	}
	if len(m.SpecFunds) > 0 {
		for _, e := range m.SpecFunds {
			l = e.Size()
			n += 1 + l + sovIprpc(uint64(l))
		}
	}
	return n
}

func (m *Specfund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Spec)
	if l > 0 {
		n += 1 + l + sovIprpc(uint64(l))
	}
	if len(m.Fund) > 0 {
		for _, e := range m.Fund {
			l = e.Size()
			n += 1 + l + sovIprpc(uint64(l))
		}
	}
	return n
}

func (m *IprpcMemo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIprpc(uint64(l))
	}
	l = len(m.Spec)
	if l > 0 {
		n += 1 + l + sovIprpc(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovIprpc(uint64(m.Duration))
	}
	return n
}

func (m *PendingIbcIprpcFund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovIprpc(uint64(m.Index))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovIprpc(uint64(l))
	}
	l = len(m.Spec)
	if l > 0 {
		n += 1 + l + sovIprpc(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovIprpc(uint64(m.Duration))
	}
	l = m.Fund.Size()
	n += 1 + l + sovIprpc(uint64(l))
	if m.Expiry != 0 {
		n += 1 + sovIprpc(uint64(m.Expiry))
	}
	return n
}

func sovIprpc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIprpc(x uint64) (n int) {
	return sovIprpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IprpcReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIprpc
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
			return fmt.Errorf("proto: IprpcReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IprpcReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpecFunds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpecFunds = append(m.SpecFunds, Specfund{})
			if err := m.SpecFunds[len(m.SpecFunds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIprpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIprpc
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
func (m *Specfund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIprpc
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
			return fmt.Errorf("proto: Specfund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Specfund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fund", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fund = append(m.Fund, types.Coin{})
			if err := m.Fund[len(m.Fund)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIprpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIprpc
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
func (m *IprpcMemo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIprpc
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
			return fmt.Errorf("proto: IprpcMemo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IprpcMemo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
		default:
			iNdEx = preIndex
			skippy, err := skipIprpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIprpc
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
func (m *PendingIbcIprpcFund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIprpc
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
			return fmt.Errorf("proto: PendingIbcIprpcFund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingIbcIprpcFund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Spec = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fund", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
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
				return ErrInvalidLengthIprpc
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIprpc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fund.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIprpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiry |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIprpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIprpc
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
func skipIprpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIprpc
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
					return 0, ErrIntOverflowIprpc
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
					return 0, ErrIntOverflowIprpc
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
				return 0, ErrInvalidLengthIprpc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIprpc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIprpc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIprpc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIprpc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIprpc = fmt.Errorf("proto: unexpected end of group")
)

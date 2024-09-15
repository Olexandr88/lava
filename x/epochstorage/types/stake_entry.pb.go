// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/stake_entry.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	types1 "github.com/cosmos/cosmos-sdk/x/staking/types"
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

type StakeEntry struct {
	Stake              types.Coin         `protobuf:"bytes,1,opt,name=stake,proto3" json:"stake"`
	Address            string             `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	StakeAppliedBlock  uint64             `protobuf:"varint,3,opt,name=stake_applied_block,json=stakeAppliedBlock,proto3" json:"stake_applied_block,omitempty"`
	Endpoints          []Endpoint         `protobuf:"bytes,4,rep,name=endpoints,proto3" json:"endpoints"`
	Geolocation        int32              `protobuf:"varint,5,opt,name=geolocation,proto3" json:"geolocation,omitempty"`
	Chain              string             `protobuf:"bytes,6,opt,name=chain,proto3" json:"chain,omitempty"`
	Moniker            string             `protobuf:"bytes,8,opt,name=moniker,proto3" json:"moniker,omitempty"`
	DelegateTotal      types.Coin         `protobuf:"bytes,9,opt,name=delegate_total,json=delegateTotal,proto3" json:"delegate_total"`
	DelegateCommission uint64             `protobuf:"varint,11,opt,name=delegate_commission,json=delegateCommission,proto3" json:"delegate_commission,omitempty"`
	LastChange         uint64             `protobuf:"varint,12,opt,name=last_change,json=lastChange,proto3" json:"last_change,omitempty"`
	BlockReport        *BlockReport       `protobuf:"bytes,13,opt,name=block_report,json=blockReport,proto3" json:"block_report,omitempty"`
	Vault              string             `protobuf:"bytes,14,opt,name=vault,proto3" json:"vault,omitempty"`
	Description        types1.Description `protobuf:"bytes,15,opt,name=description,proto3" json:"description"`
	Jails              uint64             `protobuf:"varint,16,opt,name=jails,proto3" json:"jails,omitempty"`
	JailEndTime        int64              `protobuf:"varint,17,opt,name=jail_end_time,json=jailEndTime,proto3" json:"jail_end_time,omitempty"`
}

func (m *StakeEntry) Reset()         { *m = StakeEntry{} }
func (m *StakeEntry) String() string { return proto.CompactTextString(m) }
func (*StakeEntry) ProtoMessage()    {}
func (*StakeEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6302d6b53c056e, []int{0}
}
func (m *StakeEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StakeEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StakeEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StakeEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeEntry.Merge(m, src)
}
func (m *StakeEntry) XXX_Size() int {
	return m.Size()
}
func (m *StakeEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StakeEntry proto.InternalMessageInfo

func (m *StakeEntry) GetStake() types.Coin {
	if m != nil {
		return m.Stake
	}
	return types.Coin{}
}

func (m *StakeEntry) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *StakeEntry) GetStakeAppliedBlock() uint64 {
	if m != nil {
		return m.StakeAppliedBlock
	}
	return 0
}

func (m *StakeEntry) GetEndpoints() []Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func (m *StakeEntry) GetGeolocation() int32 {
	if m != nil {
		return m.Geolocation
	}
	return 0
}

func (m *StakeEntry) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *StakeEntry) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *StakeEntry) GetDelegateTotal() types.Coin {
	if m != nil {
		return m.DelegateTotal
	}
	return types.Coin{}
}

func (m *StakeEntry) GetDelegateCommission() uint64 {
	if m != nil {
		return m.DelegateCommission
	}
	return 0
}

func (m *StakeEntry) GetLastChange() uint64 {
	if m != nil {
		return m.LastChange
	}
	return 0
}

func (m *StakeEntry) GetBlockReport() *BlockReport {
	if m != nil {
		return m.BlockReport
	}
	return nil
}

func (m *StakeEntry) GetVault() string {
	if m != nil {
		return m.Vault
	}
	return ""
}

func (m *StakeEntry) GetDescription() types1.Description {
	if m != nil {
		return m.Description
	}
	return types1.Description{}
}

func (m *StakeEntry) GetJails() uint64 {
	if m != nil {
		return m.Jails
	}
	return 0
}

func (m *StakeEntry) GetJailEndTime() int64 {
	if m != nil {
		return m.JailEndTime
	}
	return 0
}

// BlockReport holds the most up-to-date info regarding blocks of the provider
// It is set in the relay payment TX logic
// used by the consumer to calculate the provider's sync score
type BlockReport struct {
	Epoch       uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	LatestBlock uint64 `protobuf:"varint,2,opt,name=latest_block,json=latestBlock,proto3" json:"latest_block,omitempty"`
}

func (m *BlockReport) Reset()         { *m = BlockReport{} }
func (m *BlockReport) String() string { return proto.CompactTextString(m) }
func (*BlockReport) ProtoMessage()    {}
func (*BlockReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6302d6b53c056e, []int{1}
}
func (m *BlockReport) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockReport.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockReport.Merge(m, src)
}
func (m *BlockReport) XXX_Size() int {
	return m.Size()
}
func (m *BlockReport) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockReport.DiscardUnknown(m)
}

var xxx_messageInfo_BlockReport proto.InternalMessageInfo

func (m *BlockReport) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *BlockReport) GetLatestBlock() uint64 {
	if m != nil {
		return m.LatestBlock
	}
	return 0
}

func init() {
	proto.RegisterType((*StakeEntry)(nil), "lavanet.lava.epochstorage.StakeEntry")
	proto.RegisterType((*BlockReport)(nil), "lavanet.lava.epochstorage.BlockReport")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/stake_entry.proto", fileDescriptor_df6302d6b53c056e)
}

var fileDescriptor_df6302d6b53c056e = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xbd, 0x4e, 0xdd, 0x30,
	0x18, 0xbd, 0x81, 0x5c, 0xb8, 0xd7, 0x01, 0x0a, 0x86, 0xc1, 0x30, 0x84, 0x14, 0xaa, 0x2a, 0x6a,
	0xa5, 0x58, 0x80, 0xfa, 0x00, 0xbd, 0x14, 0x2a, 0x98, 0xaa, 0x94, 0xa9, 0x4b, 0xe4, 0x24, 0x56,
	0xae, 0x4b, 0x62, 0x47, 0xb1, 0xb9, 0x2a, 0x6f, 0xd1, 0xc7, 0xe8, 0xd8, 0xbd, 0x2f, 0xc0, 0xc8,
	0xd8, 0xa9, 0xaa, 0x60, 0xe8, 0x6b, 0x54, 0xb6, 0x93, 0xfb, 0x33, 0x50, 0x75, 0x49, 0x7c, 0xbe,
	0x73, 0xfc, 0xf9, 0x3b, 0x47, 0x36, 0x78, 0x5d, 0x92, 0x09, 0xe1, 0x54, 0x61, 0xfd, 0xc7, 0xb4,
	0x16, 0xd9, 0x58, 0x2a, 0xd1, 0x90, 0x82, 0x62, 0xa9, 0xc8, 0x35, 0x4d, 0x28, 0x57, 0xcd, 0x6d,
	0x54, 0x37, 0x42, 0x09, 0xb8, 0xdb, 0x8a, 0x23, 0xfd, 0x8f, 0xe6, 0xc5, 0x7b, 0xe1, 0xd3, 0x7d,
	0x28, 0xcf, 0x6b, 0xc1, 0xb8, 0xb2, 0x4d, 0xf6, 0x76, 0x0a, 0x51, 0x08, 0xb3, 0xc4, 0x7a, 0xd5,
	0x56, 0xfd, 0x4c, 0xc8, 0x4a, 0x48, 0x9c, 0x12, 0x49, 0xf1, 0xe4, 0x28, 0xa5, 0x8a, 0x1c, 0xe1,
	0x4c, 0x30, 0xde, 0xf2, 0x2f, 0x5a, 0x5e, 0x0f, 0xc5, 0x78, 0x31, 0x95, 0xb4, 0xb8, 0x55, 0x6d,
	0x91, 0x8a, 0x71, 0x81, 0xcd, 0xd7, 0x96, 0x0e, 0x7e, 0xf4, 0x01, 0xf8, 0xa8, 0x9d, 0x9c, 0x69,
	0x23, 0xf0, 0x0d, 0xe8, 0x1b, 0x5f, 0xc8, 0x09, 0x9c, 0xd0, 0x3b, 0xde, 0x8d, 0x6c, 0xdf, 0x48,
	0x9f, 0x1b, 0xb5, 0x4d, 0xa3, 0x53, 0xc1, 0xf8, 0xc8, 0xbd, 0xfb, 0xb5, 0xdf, 0x8b, 0xad, 0x1a,
	0x22, 0xb0, 0x4a, 0xf2, 0xbc, 0xa1, 0x52, 0xa2, 0xa5, 0xc0, 0x09, 0x87, 0x71, 0x07, 0x61, 0x04,
	0xb6, 0x6d, 0x50, 0xa4, 0xae, 0x4b, 0x46, 0xf3, 0x24, 0x2d, 0x45, 0x76, 0x8d, 0x96, 0x03, 0x27,
	0x74, 0xe3, 0x2d, 0x43, 0xbd, 0xb5, 0xcc, 0x48, 0x13, 0xf0, 0x3d, 0x18, 0x76, 0x81, 0x48, 0xe4,
	0x06, 0xcb, 0xa1, 0x77, 0x7c, 0x18, 0x3d, 0x99, 0x6b, 0x74, 0xd6, 0x6a, 0xdb, 0x71, 0x66, 0x7b,
	0x61, 0x00, 0xbc, 0x82, 0x8a, 0x52, 0x64, 0x44, 0x31, 0xc1, 0x51, 0x3f, 0x70, 0xc2, 0x7e, 0x3c,
	0x5f, 0x82, 0x3b, 0xa0, 0x9f, 0x8d, 0x09, 0xe3, 0x68, 0xc5, 0x8c, 0x6c, 0x81, 0xb6, 0x52, 0x09,
	0xce, 0xae, 0x69, 0x83, 0x06, 0xd6, 0x4a, 0x0b, 0xe1, 0x39, 0xd8, 0xc8, 0x69, 0x49, 0x0b, 0xa2,
	0x68, 0xa2, 0x84, 0x22, 0x25, 0x1a, 0xfe, 0x5f, 0x48, 0xeb, 0xdd, 0xb6, 0x2b, 0xbd, 0x0b, 0x62,
	0xb0, 0x3d, 0xed, 0x93, 0x89, 0xaa, 0x62, 0x52, 0xea, 0x09, 0x3d, 0x13, 0x09, 0xec, 0xa8, 0xd3,
	0x29, 0x03, 0xf7, 0x81, 0x57, 0x12, 0xa9, 0x92, 0x6c, 0x4c, 0x78, 0x41, 0xd1, 0x9a, 0x11, 0x02,
	0x5d, 0x3a, 0x35, 0x15, 0x78, 0x01, 0xd6, 0x4c, 0xac, 0x49, 0x43, 0x6b, 0xd1, 0x28, 0xb4, 0x6e,
	0xe6, 0x7a, 0xf9, 0x8f, 0xdc, 0x4c, 0xd8, 0xb1, 0x51, 0xc7, 0x5e, 0x3a, 0x03, 0x3a, 0x94, 0x09,
	0xb9, 0x29, 0x15, 0xda, 0xb0, 0xa1, 0x18, 0x00, 0x3f, 0x00, 0x2f, 0xa7, 0x32, 0x6b, 0x58, 0x6d,
	0xc2, 0x7c, 0x66, 0xfa, 0x1f, 0x76, 0xbe, 0xbb, 0x4b, 0xd6, 0x59, 0x7f, 0x37, 0x93, 0x8e, 0x86,
	0x3a, 0x81, 0x6f, 0x7f, 0xbe, 0xbf, 0x72, 0xe2, 0xf9, 0x16, 0xfa, 0x9c, 0xcf, 0x84, 0x95, 0x12,
	0x6d, 0x1a, 0x37, 0x16, 0xc0, 0x03, 0xb0, 0xae, 0x17, 0x09, 0xe5, 0x79, 0xa2, 0x58, 0x45, 0xd1,
	0x56, 0xe0, 0x84, 0xcb, 0xb1, 0xa7, 0x8b, 0x67, 0x3c, 0xbf, 0x62, 0x15, 0xbd, 0x74, 0x07, 0xab,
	0x9b, 0x83, 0x4b, 0x77, 0x00, 0x36, 0xbd, 0x83, 0x73, 0xe0, 0x8d, 0x16, 0x87, 0x37, 0x2e, 0xcd,
	0xed, 0x75, 0x63, 0x0b, 0xe0, 0x73, 0xb0, 0x56, 0x12, 0x45, 0xa5, 0x6a, 0xef, 0xde, 0x92, 0x21,
	0x3d, 0x5b, 0x33, 0xdb, 0x47, 0x17, 0x77, 0x0f, 0xbe, 0x73, 0xff, 0xe0, 0x3b, 0xbf, 0x1f, 0x7c,
	0xe7, 0xeb, 0xa3, 0xdf, 0xbb, 0x7f, 0xf4, 0x7b, 0x3f, 0x1f, 0xfd, 0xde, 0x27, 0x5c, 0x30, 0x35,
	0xbe, 0x49, 0xa3, 0x4c, 0x54, 0x78, 0xe1, 0x0d, 0x4f, 0x4e, 0xf0, 0x97, 0xc5, 0x87, 0xac, 0x6e,
	0x6b, 0x2a, 0xd3, 0x15, 0xf3, 0xae, 0x4e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x28, 0x02, 0x10,
	0xfc, 0x3a, 0x04, 0x00, 0x00,
}

func (m *StakeEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StakeEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StakeEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.JailEndTime != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.JailEndTime))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if m.Jails != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.Jails))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x7a
	if len(m.Vault) > 0 {
		i -= len(m.Vault)
		copy(dAtA[i:], m.Vault)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Vault)))
		i--
		dAtA[i] = 0x72
	}
	if m.BlockReport != nil {
		{
			size, err := m.BlockReport.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintStakeEntry(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x6a
	}
	if m.LastChange != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.LastChange))
		i--
		dAtA[i] = 0x60
	}
	if m.DelegateCommission != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.DelegateCommission))
		i--
		dAtA[i] = 0x58
	}
	{
		size, err := m.DelegateTotal.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x32
	}
	if m.Geolocation != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.Geolocation))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintStakeEntry(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.StakeAppliedBlock != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.StakeAppliedBlock))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintStakeEntry(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintStakeEntry(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BlockReport) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockReport) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockReport) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LatestBlock != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.LatestBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.Epoch != 0 {
		i = encodeVarintStakeEntry(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintStakeEntry(dAtA []byte, offset int, v uint64) int {
	offset -= sovStakeEntry(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StakeEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Stake.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	if m.StakeAppliedBlock != 0 {
		n += 1 + sovStakeEntry(uint64(m.StakeAppliedBlock))
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovStakeEntry(uint64(l))
		}
	}
	if m.Geolocation != 0 {
		n += 1 + sovStakeEntry(uint64(m.Geolocation))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = m.DelegateTotal.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	if m.DelegateCommission != 0 {
		n += 1 + sovStakeEntry(uint64(m.DelegateCommission))
	}
	if m.LastChange != 0 {
		n += 1 + sovStakeEntry(uint64(m.LastChange))
	}
	if m.BlockReport != nil {
		l = m.BlockReport.Size()
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = len(m.Vault)
	if l > 0 {
		n += 1 + l + sovStakeEntry(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovStakeEntry(uint64(l))
	if m.Jails != 0 {
		n += 2 + sovStakeEntry(uint64(m.Jails))
	}
	if m.JailEndTime != 0 {
		n += 2 + sovStakeEntry(uint64(m.JailEndTime))
	}
	return n
}

func (m *BlockReport) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovStakeEntry(uint64(m.Epoch))
	}
	if m.LatestBlock != 0 {
		n += 1 + sovStakeEntry(uint64(m.LatestBlock))
	}
	return n
}

func sovStakeEntry(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStakeEntry(x uint64) (n int) {
	return sovStakeEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StakeEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakeEntry
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
			return fmt.Errorf("proto: StakeEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StakeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeAppliedBlock", wireType)
			}
			m.StakeAppliedBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StakeAppliedBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Geolocation", wireType)
			}
			m.Geolocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Geolocation |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateTotal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DelegateTotal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegateCommission", wireType)
			}
			m.DelegateCommission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelegateCommission |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastChange", wireType)
			}
			m.LastChange = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastChange |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockReport", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockReport == nil {
				m.BlockReport = &BlockReport{}
			}
			if err := m.BlockReport.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vault", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vault = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
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
				return ErrInvalidLengthStakeEntry
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthStakeEntry
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jails", wireType)
			}
			m.Jails = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Jails |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailEndTime", wireType)
			}
			m.JailEndTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JailEndTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStakeEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakeEntry
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
func (m *BlockReport) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStakeEntry
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
			return fmt.Errorf("proto: BlockReport: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockReport: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestBlock", wireType)
			}
			m.LatestBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStakeEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStakeEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStakeEntry
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
func skipStakeEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStakeEntry
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
					return 0, ErrIntOverflowStakeEntry
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
					return 0, ErrIntOverflowStakeEntry
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
				return 0, ErrInvalidLengthStakeEntry
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStakeEntry
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStakeEntry
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStakeEntry        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStakeEntry          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStakeEntry = fmt.Errorf("proto: unexpected end of group")
)

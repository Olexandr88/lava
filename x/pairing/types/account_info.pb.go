// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/pairing/account_info.proto

package types

import (
	fmt "fmt"
	types3 "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types4 "github.com/lavanet/lava/v4/x/dualstaking/types"
	types "github.com/lavanet/lava/v4/x/epochstorage/types"
	types2 "github.com/lavanet/lava/v4/x/projects/types"
	types1 "github.com/lavanet/lava/v4/x/subscription/types"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
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

type QueryAccountInfoResponse struct {
	Provider              []types.StakeEntry          `protobuf:"bytes,1,rep,name=provider,proto3" json:"provider"`
	Frozen                []types.StakeEntry          `protobuf:"bytes,2,rep,name=frozen,proto3" json:"frozen"`
	Unstaked              []types.StakeEntry          `protobuf:"bytes,4,rep,name=unstaked,proto3" json:"unstaked"`
	Subscription          *types1.Subscription        `protobuf:"bytes,5,opt,name=subscription,proto3" json:"subscription,omitempty"`
	Project               *types2.Project             `protobuf:"bytes,6,opt,name=project,proto3" json:"project,omitempty"`
	DelegationsValidators []types3.DelegationResponse `protobuf:"bytes,7,rep,name=delegations_validators,json=delegationsValidators,proto3" json:"delegations_validators"`
	DelegationsProviders  []types4.Delegation         `protobuf:"bytes,8,rep,name=delegations_providers,json=delegationsProviders,proto3" json:"delegations_providers"`
	TotalDelegations      uint64                      `protobuf:"varint,9,opt,name=total_delegations,json=totalDelegations,proto3" json:"total_delegations,omitempty"`
}

func (m *QueryAccountInfoResponse) Reset()         { *m = QueryAccountInfoResponse{} }
func (m *QueryAccountInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAccountInfoResponse) ProtoMessage()    {}
func (*QueryAccountInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5569ac5696c70068, []int{0}
}
func (m *QueryAccountInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAccountInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAccountInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAccountInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAccountInfoResponse.Merge(m, src)
}
func (m *QueryAccountInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAccountInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAccountInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAccountInfoResponse proto.InternalMessageInfo

func (m *QueryAccountInfoResponse) GetProvider() []types.StakeEntry {
	if m != nil {
		return m.Provider
	}
	return nil
}

func (m *QueryAccountInfoResponse) GetFrozen() []types.StakeEntry {
	if m != nil {
		return m.Frozen
	}
	return nil
}

func (m *QueryAccountInfoResponse) GetUnstaked() []types.StakeEntry {
	if m != nil {
		return m.Unstaked
	}
	return nil
}

func (m *QueryAccountInfoResponse) GetSubscription() *types1.Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (m *QueryAccountInfoResponse) GetProject() *types2.Project {
	if m != nil {
		return m.Project
	}
	return nil
}

func (m *QueryAccountInfoResponse) GetDelegationsValidators() []types3.DelegationResponse {
	if m != nil {
		return m.DelegationsValidators
	}
	return nil
}

func (m *QueryAccountInfoResponse) GetDelegationsProviders() []types4.Delegation {
	if m != nil {
		return m.DelegationsProviders
	}
	return nil
}

func (m *QueryAccountInfoResponse) GetTotalDelegations() uint64 {
	if m != nil {
		return m.TotalDelegations
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryAccountInfoResponse)(nil), "lavanet.lava.pairing.QueryAccountInfoResponse")
}

func init() {
	proto.RegisterFile("lavanet/lava/pairing/account_info.proto", fileDescriptor_5569ac5696c70068)
}

var fileDescriptor_5569ac5696c70068 = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x21, 0xa4, 0xc5, 0xe5, 0x00, 0x56, 0x40, 0x56, 0x0f, 0x26, 0x82, 0xa2, 0x46, 0x2d,
	0xda, 0x55, 0x0b, 0x67, 0x24, 0x4a, 0x11, 0xe2, 0xd6, 0xa6, 0x12, 0x07, 0x2e, 0xd6, 0xc6, 0x9e,
	0xb8, 0x06, 0x77, 0xc7, 0xda, 0x5d, 0x1b, 0xc2, 0x57, 0xf0, 0x59, 0x3d, 0xf6, 0xc8, 0x09, 0xa1,
	0xe4, 0x33, 0xb8, 0x20, 0xaf, 0x77, 0xc3, 0xee, 0x11, 0x4e, 0x9e, 0xd1, 0xbc, 0x79, 0x33, 0xf3,
	0xf6, 0x39, 0xdc, 0xaf, 0x58, 0xcb, 0x38, 0x28, 0xda, 0x7d, 0x69, 0xcd, 0x4a, 0x51, 0xf2, 0x82,
	0xb2, 0x2c, 0xc3, 0x86, 0xab, 0xb4, 0xe4, 0x0b, 0x24, 0xb5, 0x40, 0x85, 0xd1, 0xd8, 0x00, 0x49,
	0xf7, 0x25, 0x06, 0xb8, 0x3b, 0x2e, 0xb0, 0x40, 0x0d, 0xa0, 0x5d, 0xd4, 0x63, 0x77, 0x93, 0x02,
	0xb1, 0xa8, 0x80, 0xea, 0x6c, 0xde, 0x2c, 0xe8, 0x17, 0xc1, 0xea, 0x1a, 0x84, 0x34, 0xf5, 0x43,
	0x6f, 0x28, 0xd4, 0x98, 0x5d, 0x4a, 0x85, 0x82, 0x15, 0x40, 0xa5, 0x62, 0x9f, 0x21, 0x05, 0xae,
	0xc4, 0xd2, 0x80, 0xfd, 0x0d, 0xf3, 0x86, 0x55, 0x1d, 0xa6, 0xdb, 0x32, 0x87, 0x0a, 0x0a, 0xa6,
	0xc0, 0x00, 0xf7, 0x32, 0x94, 0x57, 0x28, 0xa9, 0x2d, 0xb7, 0x47, 0x73, 0x50, 0xec, 0xc8, 0xe6,
	0x06, 0xf5, 0xdc, 0xa3, 0x93, 0xcd, 0x5c, 0x66, 0xa2, 0xac, 0x55, 0x89, 0xdc, 0x4b, 0x0c, 0xfa,
	0xa9, 0x2f, 0x8f, 0xc0, 0x4f, 0x90, 0x29, 0x69, 0x83, 0x1e, 0xf4, 0xe4, 0xf7, 0x30, 0x8c, 0xcf,
	0x1b, 0x10, 0xcb, 0xd7, 0xbd, 0x6c, 0xef, 0xf9, 0x02, 0x67, 0x20, 0x6b, 0xe4, 0x12, 0xa2, 0x77,
	0xe1, 0x76, 0x2d, 0xb0, 0x2d, 0x73, 0x10, 0x71, 0x30, 0xb9, 0x3d, 0xdd, 0x39, 0x7e, 0x46, 0x3c,
	0x29, 0xdd, 0xf3, 0xc9, 0x45, 0x77, 0xfe, 0xdb, 0xee, 0xfa, 0x93, 0xe1, 0xf5, 0xcf, 0xc7, 0x83,
	0xd9, 0xa6, 0x39, 0x7a, 0x13, 0x8e, 0x16, 0x02, 0xbf, 0x01, 0x8f, 0x6f, 0xfd, 0x3b, 0x8d, 0x69,
	0xed, 0xb6, 0x69, 0xb8, 0xd6, 0x38, 0x8f, 0x87, 0xff, 0xb1, 0x8d, 0x6d, 0x8e, 0xce, 0xc3, 0x7b,
	0xae, 0x5c, 0xf1, 0x9d, 0x49, 0x30, 0xdd, 0x39, 0xde, 0xf7, 0xc9, 0x3c, 0x41, 0x2f, 0x9c, 0x44,
	0xd3, 0x05, 0x33, 0x8f, 0x22, 0x7a, 0x15, 0x6e, 0x19, 0x5d, 0xe3, 0x91, 0x66, 0x4b, 0x7c, 0x36,
	0xab, 0x3e, 0x39, 0xeb, 0x03, 0x43, 0x62, 0x9b, 0xa2, 0x22, 0x7c, 0x64, 0x1c, 0x51, 0x22, 0x97,
	0x69, 0xcb, 0xaa, 0x32, 0x67, 0x0a, 0x85, 0x8c, 0xb7, 0xf4, 0xa5, 0x07, 0xa4, 0x37, 0x08, 0xb1,
	0x86, 0x30, 0x06, 0x21, 0xa7, 0x9b, 0x2e, 0xfb, 0x6a, 0xe6, 0xdc, 0x87, 0x0e, 0xdf, 0x87, 0x0d,
	0x5d, 0x94, 0x86, 0x6e, 0x21, 0xb5, 0x2f, 0x24, 0xe3, 0x6d, 0x3d, 0x67, 0xcf, 0x5f, 0xdb, 0x71,
	0xac, 0x33, 0xc9, 0x4c, 0x18, 0x3b, 0x44, 0x67, 0x96, 0x27, 0x3a, 0x0c, 0x1f, 0x28, 0x54, 0xac,
	0x4a, 0x9d, 0x6a, 0x7c, 0x77, 0x12, 0x4c, 0x87, 0xb3, 0xfb, 0xba, 0xf0, 0x97, 0x47, 0x9e, 0x9c,
	0x5e, 0xaf, 0x92, 0xe0, 0x66, 0x95, 0x04, 0xbf, 0x56, 0x49, 0xf0, 0x7d, 0x9d, 0x0c, 0x6e, 0xd6,
	0xc9, 0xe0, 0xc7, 0x3a, 0x19, 0x7c, 0x3c, 0x28, 0x4a, 0x75, 0xd9, 0xcc, 0x49, 0x86, 0x57, 0xd4,
	0xf3, 0x71, 0xfb, 0x92, 0x7e, 0xdd, 0xfc, 0xeb, 0x6a, 0x59, 0x83, 0x9c, 0x8f, 0xb4, 0x95, 0x5f,
	0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0xab, 0x04, 0x7b, 0x26, 0x10, 0x04, 0x00, 0x00,
}

func (m *QueryAccountInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAccountInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAccountInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalDelegations != 0 {
		i = encodeVarintAccountInfo(dAtA, i, uint64(m.TotalDelegations))
		i--
		dAtA[i] = 0x48
	}
	if len(m.DelegationsProviders) > 0 {
		for iNdEx := len(m.DelegationsProviders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationsProviders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.DelegationsValidators) > 0 {
		for iNdEx := len(m.DelegationsValidators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationsValidators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.Project != nil {
		{
			size, err := m.Project.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccountInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Subscription != nil {
		{
			size, err := m.Subscription.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAccountInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Unstaked) > 0 {
		for iNdEx := len(m.Unstaked) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Unstaked[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Frozen) > 0 {
		for iNdEx := len(m.Frozen) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Frozen[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Provider) > 0 {
		for iNdEx := len(m.Provider) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Provider[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAccountInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Provider) > 0 {
		for _, e := range m.Provider {
			l = e.Size()
			n += 1 + l + sovAccountInfo(uint64(l))
		}
	}
	if len(m.Frozen) > 0 {
		for _, e := range m.Frozen {
			l = e.Size()
			n += 1 + l + sovAccountInfo(uint64(l))
		}
	}
	if len(m.Unstaked) > 0 {
		for _, e := range m.Unstaked {
			l = e.Size()
			n += 1 + l + sovAccountInfo(uint64(l))
		}
	}
	if m.Subscription != nil {
		l = m.Subscription.Size()
		n += 1 + l + sovAccountInfo(uint64(l))
	}
	if m.Project != nil {
		l = m.Project.Size()
		n += 1 + l + sovAccountInfo(uint64(l))
	}
	if len(m.DelegationsValidators) > 0 {
		for _, e := range m.DelegationsValidators {
			l = e.Size()
			n += 1 + l + sovAccountInfo(uint64(l))
		}
	}
	if len(m.DelegationsProviders) > 0 {
		for _, e := range m.DelegationsProviders {
			l = e.Size()
			n += 1 + l + sovAccountInfo(uint64(l))
		}
	}
	if m.TotalDelegations != 0 {
		n += 1 + sovAccountInfo(uint64(m.TotalDelegations))
	}
	return n
}

func sovAccountInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountInfo(x uint64) (n int) {
	return sovAccountInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAccountInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountInfo
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
			return fmt.Errorf("proto: QueryAccountInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAccountInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
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
				return ErrInvalidLengthAccountInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = append(m.Provider, types.StakeEntry{})
			if err := m.Provider[len(m.Provider)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
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
				return ErrInvalidLengthAccountInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Frozen = append(m.Frozen, types.StakeEntry{})
			if err := m.Frozen[len(m.Frozen)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unstaked", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
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
				return ErrInvalidLengthAccountInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unstaked = append(m.Unstaked, types.StakeEntry{})
			if err := m.Unstaked[len(m.Unstaked)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subscription", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
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
				return ErrInvalidLengthAccountInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Subscription == nil {
				m.Subscription = &types1.Subscription{}
			}
			if err := m.Subscription.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Project", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
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
				return ErrInvalidLengthAccountInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Project == nil {
				m.Project = &types2.Project{}
			}
			if err := m.Project.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationsValidators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
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
				return ErrInvalidLengthAccountInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationsValidators = append(m.DelegationsValidators, types3.DelegationResponse{})
			if err := m.DelegationsValidators[len(m.DelegationsValidators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationsProviders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
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
				return ErrInvalidLengthAccountInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegationsProviders = append(m.DelegationsProviders, types4.Delegation{})
			if err := m.DelegationsProviders[len(m.DelegationsProviders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDelegations", wireType)
			}
			m.TotalDelegations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalDelegations |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAccountInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountInfo
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
func skipAccountInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountInfo
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
					return 0, ErrIntOverflowAccountInfo
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
					return 0, ErrIntOverflowAccountInfo
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
				return 0, ErrInvalidLengthAccountInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountInfo = fmt.Errorf("proto: unexpected end of group")
)

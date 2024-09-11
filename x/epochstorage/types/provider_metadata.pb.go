// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/provider_metadata.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type ProviderMetadata struct {
	Provider         string     `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	SelfDelegation   types.Coin `protobuf:"bytes,2,opt,name=SelfDelegation,proto3" json:"SelfDelegation"`
	TotalDelegations types.Coin `protobuf:"bytes,3,opt,name=TotalDelegations,proto3" json:"TotalDelegations"`
	Chains           []string   `protobuf:"bytes,4,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (m *ProviderMetadata) Reset()         { *m = ProviderMetadata{} }
func (m *ProviderMetadata) String() string { return proto.CompactTextString(m) }
func (*ProviderMetadata) ProtoMessage()    {}
func (*ProviderMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_eea88d8dbfad23fd, []int{0}
}
func (m *ProviderMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProviderMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProviderMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProviderMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProviderMetadata.Merge(m, src)
}
func (m *ProviderMetadata) XXX_Size() int {
	return m.Size()
}
func (m *ProviderMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ProviderMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ProviderMetadata proto.InternalMessageInfo

func (m *ProviderMetadata) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *ProviderMetadata) GetSelfDelegation() types.Coin {
	if m != nil {
		return m.SelfDelegation
	}
	return types.Coin{}
}

func (m *ProviderMetadata) GetTotalDelegations() types.Coin {
	if m != nil {
		return m.TotalDelegations
	}
	return types.Coin{}
}

func (m *ProviderMetadata) GetChains() []string {
	if m != nil {
		return m.Chains
	}
	return nil
}

func init() {
	proto.RegisterType((*ProviderMetadata)(nil), "lavanet.lava.epochstorage.ProviderMetadata")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/provider_metadata.proto", fileDescriptor_eea88d8dbfad23fd)
}

var fileDescriptor_eea88d8dbfad23fd = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x14, 0xc0, 0x1b, 0x37, 0x86, 0x8b, 0x20, 0xb3, 0x88, 0x74, 0x3b, 0xc4, 0xe1, 0x69, 0xa7, 0x84,
	0xb9, 0x6f, 0x30, 0x05, 0x11, 0x11, 0x64, 0x7a, 0xf2, 0x22, 0xaf, 0x5d, 0xec, 0x02, 0x6d, 0x5e,
	0x69, 0x62, 0xd1, 0x6f, 0xe1, 0xc7, 0xda, 0x71, 0x47, 0x0f, 0x22, 0xb2, 0x7e, 0x11, 0x69, 0x1b,
	0xff, 0x4c, 0x2f, 0x5e, 0xf2, 0xf2, 0x5e, 0xf2, 0xcb, 0x2f, 0xbc, 0x47, 0xc7, 0x09, 0x14, 0xa0,
	0xa5, 0x15, 0x55, 0x14, 0x32, 0xc3, 0x68, 0x61, 0x2c, 0xe6, 0x10, 0x4b, 0x91, 0xe5, 0x58, 0xa8,
	0xb9, 0xcc, 0xef, 0x52, 0x69, 0x61, 0x0e, 0x16, 0x78, 0x96, 0xa3, 0x45, 0xbf, 0xef, 0x10, 0x5e,
	0x45, 0xfe, 0x13, 0x19, 0xec, 0xc7, 0x18, 0x63, 0x7d, 0x4b, 0x54, 0xbb, 0x06, 0x18, 0xb0, 0x08,
	0x4d, 0x8a, 0x46, 0x84, 0x60, 0xa4, 0x28, 0xc6, 0xa1, 0xb4, 0x30, 0x16, 0x11, 0x2a, 0xed, 0xce,
	0xf7, 0x20, 0x55, 0x1a, 0x45, 0xbd, 0x36, 0xa5, 0xa3, 0x57, 0x42, 0x7b, 0x57, 0xce, 0x7f, 0xe9,
	0xf4, 0xfe, 0x80, 0x6e, 0x7f, 0xfe, 0x29, 0x20, 0x43, 0x32, 0xea, 0xce, 0xbe, 0x72, 0xff, 0x8c,
	0xee, 0x5e, 0xcb, 0xe4, 0xfe, 0x54, 0x26, 0x32, 0x06, 0xab, 0x50, 0x07, 0x5b, 0x43, 0x32, 0xda,
	0x39, 0xee, 0xf3, 0x46, 0xce, 0x2b, 0x39, 0x77, 0x72, 0x7e, 0x82, 0x4a, 0x4f, 0xdb, 0xcb, 0xb7,
	0x43, 0x6f, 0xf6, 0x0b, 0xf3, 0x2f, 0x68, 0xef, 0x06, 0x2d, 0x24, 0xdf, 0x25, 0x13, 0xb4, 0xfe,
	0xf7, 0xd4, 0x1f, 0xd0, 0x3f, 0xa0, 0x9d, 0x68, 0x01, 0x4a, 0x9b, 0xa0, 0x3d, 0x6c, 0x8d, 0xba,
	0x33, 0x97, 0x4d, 0xcf, 0x97, 0x6b, 0x46, 0x56, 0x6b, 0x46, 0xde, 0xd7, 0x8c, 0x3c, 0x97, 0xcc,
	0x5b, 0x95, 0xcc, 0x7b, 0x29, 0x99, 0x77, 0x2b, 0x62, 0x65, 0x17, 0x0f, 0x21, 0x8f, 0x30, 0x15,
	0x1b, 0xa3, 0x29, 0x26, 0xe2, 0x71, 0x73, 0x3e, 0xf6, 0x29, 0x93, 0x26, 0xec, 0xd4, 0x0d, 0x9b,
	0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x2e, 0xe7, 0x8c, 0xc9, 0x01, 0x00, 0x00,
}

func (m *ProviderMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProviderMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProviderMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintProviderMetadata(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.TotalDelegations.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProviderMetadata(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SelfDelegation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProviderMetadata(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintProviderMetadata(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProviderMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovProviderMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProviderMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovProviderMetadata(uint64(l))
	}
	l = m.SelfDelegation.Size()
	n += 1 + l + sovProviderMetadata(uint64(l))
	l = m.TotalDelegations.Size()
	n += 1 + l + sovProviderMetadata(uint64(l))
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovProviderMetadata(uint64(l))
		}
	}
	return n
}

func sovProviderMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProviderMetadata(x uint64) (n int) {
	return sovProviderMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProviderMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProviderMetadata
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
			return fmt.Errorf("proto: ProviderMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProviderMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderMetadata
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
				return ErrInvalidLengthProviderMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SelfDelegation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderMetadata
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
				return ErrInvalidLengthProviderMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProviderMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SelfDelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalDelegations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderMetadata
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
				return ErrInvalidLengthProviderMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProviderMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalDelegations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProviderMetadata
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
				return ErrInvalidLengthProviderMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProviderMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProviderMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProviderMetadata
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
func skipProviderMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProviderMetadata
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
					return 0, ErrIntOverflowProviderMetadata
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
					return 0, ErrIntOverflowProviderMetadata
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
				return 0, ErrInvalidLengthProviderMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProviderMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProviderMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProviderMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProviderMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProviderMetadata = fmt.Errorf("proto: unexpected end of group")
)

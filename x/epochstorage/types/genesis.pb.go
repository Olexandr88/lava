// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the epochstorage module's genesis state.
type GenesisState struct {
	Params            Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	StakeStorageList  []StakeStorage  `protobuf:"bytes,2,rep,name=stakeStorageList,proto3" json:"stakeStorageList"`
	EpochDetails      *EpochDetails   `protobuf:"bytes,3,opt,name=epochDetails,proto3" json:"epochDetails,omitempty"`
	FixatedParamsList []FixatedParams `protobuf:"bytes,4,rep,name=fixatedParamsList,proto3" json:"fixatedParamsList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b4c1a9bc1f09c0e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetStakeStorageList() []StakeStorage {
	if m != nil {
		return m.StakeStorageList
	}
	return nil
}

func (m *GenesisState) GetEpochDetails() *EpochDetails {
	if m != nil {
		return m.EpochDetails
	}
	return nil
}

func (m *GenesisState) GetFixatedParamsList() []FixatedParams {
	if m != nil {
		return m.FixatedParamsList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.epochstorage.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/genesis.proto", fileDescriptor_6b4c1a9bc1f09c0e)
}

var fileDescriptor_6b4c1a9bc1f09c0e = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcf, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0xa9, 0x05, 0xf9, 0xc9, 0x19, 0xc5, 0x25, 0xf9, 0x45,
	0x89, 0xe9, 0xa9, 0xfa, 0xe9, 0xa9, 0x79, 0xa9, 0xc5, 0x99, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0x92, 0x50, 0x85, 0x7a, 0x20, 0x5a, 0x0f, 0x59, 0xa1, 0x94, 0x48, 0x7a, 0x7e, 0x7a,
	0x3e, 0x58, 0x95, 0x3e, 0x88, 0x05, 0xd1, 0x20, 0xa5, 0x86, 0xdb, 0xe4, 0x82, 0xc4, 0xa2, 0xc4,
	0x5c, 0xa8, 0xc1, 0x52, 0xba, 0xb8, 0xd5, 0x15, 0x97, 0x24, 0x66, 0xa7, 0xc6, 0x43, 0x79, 0x84,
	0x95, 0x83, 0x39, 0xf1, 0x29, 0xa9, 0x25, 0x89, 0x99, 0x39, 0x30, 0xd3, 0xf5, 0x70, 0x2b, 0x4f,
	0xcb, 0xac, 0x48, 0x2c, 0x49, 0x4d, 0x89, 0x47, 0x76, 0x8d, 0xd2, 0x55, 0x26, 0x2e, 0x1e, 0x77,
	0x88, 0xc7, 0x83, 0x4b, 0x12, 0x4b, 0x52, 0x85, 0xec, 0xb9, 0xd8, 0x20, 0x0a, 0x24, 0x18, 0x15,
	0x18, 0x35, 0xb8, 0x8d, 0x14, 0xf5, 0x70, 0x06, 0x84, 0x5e, 0x00, 0x58, 0xa1, 0x13, 0xcb, 0x89,
	0x7b, 0xf2, 0x0c, 0x41, 0x50, 0x6d, 0x42, 0x91, 0x5c, 0x02, 0x60, 0x7f, 0x04, 0x43, 0x14, 0xf9,
	0x64, 0x16, 0x97, 0x48, 0x30, 0x29, 0x30, 0x6b, 0x70, 0x1b, 0xa9, 0xe3, 0x31, 0x2a, 0x18, 0x49,
	0x0b, 0xd4, 0x40, 0x0c, 0x63, 0x84, 0xbc, 0xb9, 0x78, 0xc0, 0x9a, 0x5c, 0x20, 0x5e, 0x96, 0x60,
	0x06, 0xbb, 0x10, 0x9f, 0xb1, 0xae, 0x48, 0xca, 0x83, 0x50, 0x34, 0x0b, 0xc5, 0x70, 0x09, 0x42,
	0x43, 0x04, 0xe2, 0x0d, 0xb0, 0x43, 0x59, 0xc0, 0x0e, 0xd5, 0xc0, 0x63, 0xa2, 0x1b, 0xb2, 0x1e,
	0xa8, 0x4b, 0x31, 0x0d, 0x72, 0xf2, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86,
	0x28, 0xfd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x94, 0xc8, 0x2a,
	0x33, 0xd6, 0xaf, 0x40, 0x8d, 0xb1, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x4c, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x9c, 0x7e, 0xcc, 0xbb, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FixatedParamsList) > 0 {
		for iNdEx := len(m.FixatedParamsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FixatedParamsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.EpochDetails != nil {
		{
			size, err := m.EpochDetails.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StakeStorageList) > 0 {
		for iNdEx := len(m.StakeStorageList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.StakeStorageList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.StakeStorageList) > 0 {
		for _, e := range m.StakeStorageList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.EpochDetails != nil {
		l = m.EpochDetails.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.FixatedParamsList) > 0 {
		for _, e := range m.FixatedParamsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StakeStorageList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StakeStorageList = append(m.StakeStorageList, StakeStorage{})
			if err := m.StakeStorageList[len(m.StakeStorageList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochDetails", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EpochDetails == nil {
				m.EpochDetails = &EpochDetails{}
			}
			if err := m.EpochDetails.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FixatedParamsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FixatedParamsList = append(m.FixatedParamsList, FixatedParams{})
			if err := m.FixatedParamsList[len(m.FixatedParamsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)

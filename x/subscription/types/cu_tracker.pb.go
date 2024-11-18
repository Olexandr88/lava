// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/subscription/cu_tracker.proto

package types

import (
	fmt "fmt"
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

type TrackedCu struct {
	Cu uint64 `protobuf:"varint,1,opt,name=cu,proto3" json:"cu,omitempty"`
}

func (m *TrackedCu) Reset()         { *m = TrackedCu{} }
func (m *TrackedCu) String() string { return proto.CompactTextString(m) }
func (*TrackedCu) ProtoMessage()    {}
func (*TrackedCu) Descriptor() ([]byte, []int) {
	return fileDescriptor_5974e118ddf7c543, []int{0}
}
func (m *TrackedCu) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrackedCu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrackedCu.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrackedCu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackedCu.Merge(m, src)
}
func (m *TrackedCu) XXX_Size() int {
	return m.Size()
}
func (m *TrackedCu) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackedCu.DiscardUnknown(m)
}

var xxx_messageInfo_TrackedCu proto.InternalMessageInfo

func (m *TrackedCu) GetCu() uint64 {
	if m != nil {
		return m.Cu
	}
	return 0
}

type CuTrackerTimerData struct {
	Block  uint64     `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	Credit types.Coin `protobuf:"bytes,2,opt,name=credit,proto3" json:"credit"`
}

func (m *CuTrackerTimerData) Reset()         { *m = CuTrackerTimerData{} }
func (m *CuTrackerTimerData) String() string { return proto.CompactTextString(m) }
func (*CuTrackerTimerData) ProtoMessage()    {}
func (*CuTrackerTimerData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5974e118ddf7c543, []int{1}
}
func (m *CuTrackerTimerData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CuTrackerTimerData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CuTrackerTimerData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CuTrackerTimerData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CuTrackerTimerData.Merge(m, src)
}
func (m *CuTrackerTimerData) XXX_Size() int {
	return m.Size()
}
func (m *CuTrackerTimerData) XXX_DiscardUnknown() {
	xxx_messageInfo_CuTrackerTimerData.DiscardUnknown(m)
}

var xxx_messageInfo_CuTrackerTimerData proto.InternalMessageInfo

func (m *CuTrackerTimerData) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *CuTrackerTimerData) GetCredit() types.Coin {
	if m != nil {
		return m.Credit
	}
	return types.Coin{}
}

type TrackedCuInfo struct {
	Provider  string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	ChainID   string `protobuf:"bytes,2,opt,name=chainID,proto3" json:"chainID,omitempty"`
	TrackedCu uint64 `protobuf:"varint,3,opt,name=trackedCu,proto3" json:"trackedCu,omitempty"`
	Block     uint64 `protobuf:"varint,4,opt,name=block,proto3" json:"block,omitempty"`
}

func (m *TrackedCuInfo) Reset()         { *m = TrackedCuInfo{} }
func (m *TrackedCuInfo) String() string { return proto.CompactTextString(m) }
func (*TrackedCuInfo) ProtoMessage()    {}
func (*TrackedCuInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5974e118ddf7c543, []int{2}
}
func (m *TrackedCuInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TrackedCuInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TrackedCuInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TrackedCuInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrackedCuInfo.Merge(m, src)
}
func (m *TrackedCuInfo) XXX_Size() int {
	return m.Size()
}
func (m *TrackedCuInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TrackedCuInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TrackedCuInfo proto.InternalMessageInfo

func (m *TrackedCuInfo) GetProvider() string {
	if m != nil {
		return m.Provider
	}
	return ""
}

func (m *TrackedCuInfo) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *TrackedCuInfo) GetTrackedCu() uint64 {
	if m != nil {
		return m.TrackedCu
	}
	return 0
}

func (m *TrackedCuInfo) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func init() {
	proto.RegisterType((*TrackedCu)(nil), "lavanet.lava.subscription.TrackedCu")
	proto.RegisterType((*CuTrackerTimerData)(nil), "lavanet.lava.subscription.CuTrackerTimerData")
	proto.RegisterType((*TrackedCuInfo)(nil), "lavanet.lava.subscription.TrackedCuInfo")
}

func init() {
	proto.RegisterFile("lavanet/lava/subscription/cu_tracker.proto", fileDescriptor_5974e118ddf7c543)
}

var fileDescriptor_5974e118ddf7c543 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0x6e, 0xc2, 0x30,
	0x14, 0x85, 0x13, 0x4a, 0x69, 0xe3, 0xaa, 0x1d, 0x2c, 0x86, 0x40, 0x2b, 0x17, 0x31, 0xa1, 0x0e,
	0xb6, 0x68, 0x2b, 0x75, 0x07, 0x16, 0xd6, 0x88, 0xa9, 0x4b, 0xe5, 0x18, 0x17, 0x2c, 0x20, 0x37,
	0x72, 0xec, 0xa8, 0xbc, 0x45, 0x1f, 0x8b, 0x91, 0xb1, 0x53, 0x55, 0x91, 0x17, 0xa9, 0xf2, 0xc3,
	0xdf, 0x74, 0xef, 0xb9, 0x3e, 0xf6, 0xf9, 0xe4, 0x8b, 0x9e, 0x96, 0x3c, 0xe5, 0x91, 0x34, 0x2c,
	0xaf, 0x2c, 0xb1, 0x61, 0x22, 0xb4, 0x8a, 0x8d, 0x82, 0x88, 0x09, 0xfb, 0x61, 0x34, 0x17, 0x0b,
	0xa9, 0x69, 0xac, 0xc1, 0x00, 0x6e, 0x55, 0x5e, 0x9a, 0x57, 0x7a, 0xea, 0x6d, 0x13, 0x01, 0xc9,
	0x0a, 0x12, 0x16, 0xf2, 0x44, 0xb2, 0xb4, 0x1f, 0x4a, 0xc3, 0xfb, 0x4c, 0x80, 0x8a, 0xca, 0xab,
	0xed, 0xe6, 0x0c, 0x66, 0x50, 0xb4, 0x2c, 0xef, 0xca, 0x69, 0xf7, 0x1e, 0x79, 0x93, 0x22, 0x61,
	0x3a, 0xb4, 0xf8, 0x0e, 0xd5, 0x84, 0xf5, 0xdd, 0x8e, 0xdb, 0xab, 0x07, 0x35, 0x61, 0xbb, 0x02,
	0xe1, 0xa1, 0x2d, 0x8f, 0xf5, 0x44, 0xad, 0xa4, 0x1e, 0x71, 0xc3, 0x71, 0x13, 0x5d, 0x86, 0x4b,
	0x10, 0x8b, 0xca, 0x58, 0x0a, 0xfc, 0x86, 0x1a, 0x42, 0xcb, 0xa9, 0x32, 0x7e, 0xad, 0xe3, 0xf6,
	0x6e, 0x9e, 0x5b, 0xb4, 0xe4, 0xa1, 0x39, 0x0f, 0xad, 0x78, 0xe8, 0x10, 0x54, 0x34, 0xa8, 0x6f,
	0x7e, 0x1f, 0x9d, 0xa0, 0xb2, 0x77, 0xd7, 0xe8, 0xf6, 0x40, 0x30, 0x8e, 0x3e, 0x01, 0xb7, 0xd1,
	0x75, 0xac, 0x21, 0x55, 0x53, 0xa9, 0x8b, 0x08, 0x2f, 0x38, 0x68, 0xec, 0xa3, 0x2b, 0x31, 0xe7,
	0x2a, 0x1a, 0x8f, 0x8a, 0x18, 0x2f, 0xd8, 0x4b, 0xfc, 0x80, 0x3c, 0xb3, 0x7f, 0xc6, 0xbf, 0x28,
	0xc8, 0x8e, 0x83, 0x23, 0x73, 0xfd, 0x84, 0x79, 0x30, 0xde, 0xec, 0x88, 0xbb, 0xdd, 0x11, 0xf7,
	0x6f, 0x47, 0xdc, 0xef, 0x8c, 0x38, 0xdb, 0x8c, 0x38, 0x3f, 0x19, 0x71, 0xde, 0xd9, 0x4c, 0x99,
	0xb9, 0x0d, 0xa9, 0x80, 0x15, 0x3b, 0x5b, 0x4f, 0xfa, 0xca, 0xbe, 0xce, 0x77, 0x64, 0xd6, 0xb1,
	0x4c, 0xc2, 0x46, 0xf1, 0x9d, 0x2f, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x17, 0xcc, 0x38,
	0xcd, 0x01, 0x00, 0x00,
}

func (m *TrackedCu) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrackedCu) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrackedCu) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Cu != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.Cu))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CuTrackerTimerData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CuTrackerTimerData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CuTrackerTimerData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Credit.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCuTracker(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Block != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TrackedCuInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TrackedCuInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TrackedCuInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Block != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x20
	}
	if m.TrackedCu != 0 {
		i = encodeVarintCuTracker(dAtA, i, uint64(m.TrackedCu))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintCuTracker(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Provider) > 0 {
		i -= len(m.Provider)
		copy(dAtA[i:], m.Provider)
		i = encodeVarintCuTracker(dAtA, i, uint64(len(m.Provider)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCuTracker(dAtA []byte, offset int, v uint64) int {
	offset -= sovCuTracker(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TrackedCu) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Cu != 0 {
		n += 1 + sovCuTracker(uint64(m.Cu))
	}
	return n
}

func (m *CuTrackerTimerData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Block != 0 {
		n += 1 + sovCuTracker(uint64(m.Block))
	}
	l = m.Credit.Size()
	n += 1 + l + sovCuTracker(uint64(l))
	return n
}

func (m *TrackedCuInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Provider)
	if l > 0 {
		n += 1 + l + sovCuTracker(uint64(l))
	}
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovCuTracker(uint64(l))
	}
	if m.TrackedCu != 0 {
		n += 1 + sovCuTracker(uint64(m.TrackedCu))
	}
	if m.Block != 0 {
		n += 1 + sovCuTracker(uint64(m.Block))
	}
	return n
}

func sovCuTracker(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCuTracker(x uint64) (n int) {
	return sovCuTracker(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TrackedCu) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCuTracker
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
			return fmt.Errorf("proto: TrackedCu: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrackedCu: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cu", wireType)
			}
			m.Cu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Cu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCuTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCuTracker
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
func (m *CuTrackerTimerData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCuTracker
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
			return fmt.Errorf("proto: CuTrackerTimerData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CuTrackerTimerData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
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
				return ErrInvalidLengthCuTracker
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCuTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Credit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCuTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCuTracker
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
func (m *TrackedCuInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCuTracker
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
			return fmt.Errorf("proto: TrackedCuInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TrackedCuInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Provider", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
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
				return ErrInvalidLengthCuTracker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCuTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Provider = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
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
				return ErrInvalidLengthCuTracker
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCuTracker
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackedCu", wireType)
			}
			m.TrackedCu = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrackedCu |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCuTracker
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCuTracker(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCuTracker
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
func skipCuTracker(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCuTracker
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
					return 0, ErrIntOverflowCuTracker
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
					return 0, ErrIntOverflowCuTracker
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
				return 0, ErrInvalidLengthCuTracker
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCuTracker
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCuTracker
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCuTracker        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCuTracker          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCuTracker = fmt.Errorf("proto: unexpected end of group")
)

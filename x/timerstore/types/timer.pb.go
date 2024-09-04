// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/timerstore/timer.proto

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

type GenesisState struct {
	Version         uint64              `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	NextBlockHeight uint64              `protobuf:"varint,2,opt,name=next_block_height,json=nextBlockHeight,proto3" json:"next_block_height,omitempty"`
	NextBlockTime   uint64              `protobuf:"varint,3,opt,name=next_block_time,json=nextBlockTime,proto3" json:"next_block_time,omitempty"`
	TimeEntries     []GenesisTimerEntry `protobuf:"bytes,4,rep,name=time_entries,json=timeEntries,proto3" json:"time_entries"`
	BlockEntries    []GenesisTimerEntry `protobuf:"bytes,5,rep,name=block_entries,json=blockEntries,proto3" json:"block_entries"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bfa2b3f0d19f843, []int{0}
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

func (m *GenesisState) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GenesisState) GetNextBlockHeight() uint64 {
	if m != nil {
		return m.NextBlockHeight
	}
	return 0
}

func (m *GenesisState) GetNextBlockTime() uint64 {
	if m != nil {
		return m.NextBlockTime
	}
	return 0
}

func (m *GenesisState) GetTimeEntries() []GenesisTimerEntry {
	if m != nil {
		return m.TimeEntries
	}
	return nil
}

func (m *GenesisState) GetBlockEntries() []GenesisTimerEntry {
	if m != nil {
		return m.BlockEntries
	}
	return nil
}

type GenesisTimerEntry struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	Data  []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *GenesisTimerEntry) Reset()         { *m = GenesisTimerEntry{} }
func (m *GenesisTimerEntry) String() string { return proto.CompactTextString(m) }
func (*GenesisTimerEntry) ProtoMessage()    {}
func (*GenesisTimerEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bfa2b3f0d19f843, []int{1}
}
func (m *GenesisTimerEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisTimerEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisTimerEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisTimerEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisTimerEntry.Merge(m, src)
}
func (m *GenesisTimerEntry) XXX_Size() int {
	return m.Size()
}
func (m *GenesisTimerEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisTimerEntry.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisTimerEntry proto.InternalMessageInfo

func (m *GenesisTimerEntry) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GenesisTimerEntry) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *GenesisTimerEntry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.timerstore.GenesisState")
	proto.RegisterType((*GenesisTimerEntry)(nil), "lavanet.lava.timerstore.GenesisTimerEntry")
}

func init() {
	proto.RegisterFile("lavanet/lava/timerstore/timer.proto", fileDescriptor_0bfa2b3f0d19f843)
}

var fileDescriptor_0bfa2b3f0d19f843 = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x6a, 0xf2, 0x40,
	0x10, 0xc7, 0x13, 0x8d, 0xdf, 0x47, 0xd7, 0x88, 0x75, 0x11, 0x1a, 0x7a, 0x48, 0xc5, 0x42, 0x11,
	0xa1, 0x09, 0xd4, 0x37, 0x10, 0xc4, 0xde, 0x0a, 0xb1, 0xbd, 0xf4, 0x22, 0xd1, 0x0e, 0x71, 0x51,
	0x77, 0x65, 0x77, 0x0c, 0xfa, 0x16, 0x7d, 0x8c, 0x3e, 0x8a, 0x47, 0x8f, 0x3d, 0x95, 0xa2, 0x2f,
	0x52, 0x76, 0x13, 0xdb, 0x4a, 0xe9, 0xa5, 0xa7, 0xf9, 0xcf, 0xcc, 0x2f, 0xff, 0xcc, 0xce, 0x90,
	0xcb, 0x59, 0x9c, 0xc6, 0x1c, 0x30, 0xd4, 0x31, 0x44, 0x36, 0x07, 0xa9, 0x50, 0x48, 0xc8, 0x64,
	0xb0, 0x90, 0x02, 0x05, 0x3d, 0xcb, 0xa1, 0x40, 0xc7, 0xe0, 0x0b, 0x3a, 0xaf, 0x27, 0x22, 0x11,
	0x86, 0x09, 0xb5, 0xca, 0xf0, 0xe6, 0x4b, 0x81, 0xb8, 0x7d, 0xe0, 0xa0, 0x98, 0x1a, 0x60, 0x8c,
	0x40, 0x3d, 0xf2, 0x3f, 0x05, 0xa9, 0x98, 0xe0, 0x9e, 0xdd, 0xb0, 0x5b, 0x4e, 0x74, 0x48, 0x69,
	0x9b, 0xd4, 0x38, 0xac, 0x70, 0x38, 0x9a, 0x89, 0xf1, 0x74, 0x38, 0x01, 0x96, 0x4c, 0xd0, 0x2b,
	0x18, 0xa6, 0xaa, 0x1b, 0x5d, 0x5d, 0xbf, 0x35, 0x65, 0x7a, 0x45, 0xaa, 0xdf, 0x58, 0x3d, 0x85,
	0x57, 0x34, 0x64, 0xe5, 0x93, 0xbc, 0x67, 0x73, 0xa0, 0x03, 0xe2, 0xea, 0xe6, 0x10, 0x38, 0x4a,
	0x06, 0xca, 0x73, 0x1a, 0xc5, 0x56, 0xf9, 0xa6, 0x1d, 0xfc, 0xf2, 0x88, 0x20, 0x1f, 0x55, 0x7f,
	0x2b, 0x7b, 0x1c, 0xe5, 0xba, 0xeb, 0x6c, 0xde, 0x2e, 0xac, 0xa8, 0xac, 0x99, 0x5e, 0x66, 0x42,
	0x1f, 0x48, 0x25, 0xfb, 0xef, 0xc1, 0xb5, 0xf4, 0x47, 0x57, 0xd7, 0xd8, 0xe4, 0xb6, 0xcd, 0x3b,
	0x52, 0xfb, 0x01, 0xd2, 0x53, 0x52, 0x9c, 0xc2, 0xda, 0xac, 0xea, 0x24, 0xd2, 0x92, 0xd6, 0x49,
	0x29, 0x8d, 0x67, 0x4b, 0xc8, 0x57, 0x93, 0x25, 0x94, 0x12, 0xe7, 0x29, 0xc6, 0xd8, 0x6c, 0xc1,
	0x8d, 0x8c, 0xee, 0xf6, 0x37, 0x3b, 0xdf, 0xde, 0xee, 0x7c, 0xfb, 0x7d, 0xe7, 0xdb, 0xcf, 0x7b,
	0xdf, 0xda, 0xee, 0x7d, 0xeb, 0x75, 0xef, 0x5b, 0x8f, 0xd7, 0x09, 0xc3, 0xc9, 0x72, 0x14, 0x8c,
	0xc5, 0x3c, 0x3c, 0x3a, 0x7a, 0xda, 0x09, 0x57, 0x47, 0x97, 0x5f, 0x2f, 0x40, 0x8d, 0xfe, 0x99,
	0x5b, 0x76, 0x3e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x0e, 0xa2, 0xfe, 0x21, 0x02, 0x00, 0x00,
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
	if len(m.BlockEntries) > 0 {
		for iNdEx := len(m.BlockEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BlockEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTimer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TimeEntries) > 0 {
		for iNdEx := len(m.TimeEntries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimeEntries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTimer(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.NextBlockTime != 0 {
		i = encodeVarintTimer(dAtA, i, uint64(m.NextBlockTime))
		i--
		dAtA[i] = 0x18
	}
	if m.NextBlockHeight != 0 {
		i = encodeVarintTimer(dAtA, i, uint64(m.NextBlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Version != 0 {
		i = encodeVarintTimer(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GenesisTimerEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisTimerEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisTimerEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTimer(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Value != 0 {
		i = encodeVarintTimer(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTimer(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTimer(dAtA []byte, offset int, v uint64) int {
	offset -= sovTimer(v)
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
	if m.Version != 0 {
		n += 1 + sovTimer(uint64(m.Version))
	}
	if m.NextBlockHeight != 0 {
		n += 1 + sovTimer(uint64(m.NextBlockHeight))
	}
	if m.NextBlockTime != 0 {
		n += 1 + sovTimer(uint64(m.NextBlockTime))
	}
	if len(m.TimeEntries) > 0 {
		for _, e := range m.TimeEntries {
			l = e.Size()
			n += 1 + l + sovTimer(uint64(l))
		}
	}
	if len(m.BlockEntries) > 0 {
		for _, e := range m.BlockEntries {
			l = e.Size()
			n += 1 + l + sovTimer(uint64(l))
		}
	}
	return n
}

func (m *GenesisTimerEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTimer(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovTimer(uint64(m.Value))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTimer(uint64(l))
	}
	return n
}

func sovTimer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTimer(x uint64) (n int) {
	return sovTimer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimer
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextBlockHeight", wireType)
			}
			m.NextBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextBlockTime", wireType)
			}
			m.NextBlockTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextBlockTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
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
				return ErrInvalidLengthTimer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTimer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimeEntries = append(m.TimeEntries, GenesisTimerEntry{})
			if err := m.TimeEntries[len(m.TimeEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockEntries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
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
				return ErrInvalidLengthTimer
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTimer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockEntries = append(m.BlockEntries, GenesisTimerEntry{})
			if err := m.BlockEntries[len(m.BlockEntries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimer
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
func (m *GenesisTimerEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimer
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
			return fmt.Errorf("proto: GenesisTimerEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisTimerEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
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
				return ErrInvalidLengthTimer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTimer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTimer
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTimer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTimer
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
func skipTimer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimer
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
					return 0, ErrIntOverflowTimer
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
					return 0, ErrIntOverflowTimer
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
				return 0, ErrInvalidLengthTimer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTimer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTimer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTimer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTimer = fmt.Errorf("proto: unexpected end of group")
)

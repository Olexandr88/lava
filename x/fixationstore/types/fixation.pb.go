// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/fixationstore/fixation.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/lavanet/lava/v4/x/timerstore/types"
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

type Entry struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Block    uint64 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	StaleAt  uint64 `protobuf:"varint,3,opt,name=stale_at,json=staleAt,proto3" json:"stale_at,omitempty"`
	Refcount uint64 `protobuf:"varint,4,opt,name=refcount,proto3" json:"refcount,omitempty"`
	Data     []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	DeleteAt uint64 `protobuf:"varint,6,opt,name=delete_at,json=deleteAt,proto3" json:"delete_at,omitempty"`
	IsLatest bool   `protobuf:"varint,7,opt,name=is_latest,json=isLatest,proto3" json:"is_latest,omitempty"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_872ca3cbd39027a3, []int{0}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return m.Size()
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Entry) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *Entry) GetStaleAt() uint64 {
	if m != nil {
		return m.StaleAt
	}
	return 0
}

func (m *Entry) GetRefcount() uint64 {
	if m != nil {
		return m.Refcount
	}
	return 0
}

func (m *Entry) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Entry) GetDeleteAt() uint64 {
	if m != nil {
		return m.DeleteAt
	}
	return 0
}

func (m *Entry) GetIsLatest() bool {
	if m != nil {
		return m.IsLatest
	}
	return false
}

type GenesisEntries struct {
	Index   string  `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	IsLive  bool    `protobuf:"varint,2,opt,name=is_live,json=isLive,proto3" json:"is_live,omitempty"`
	Entries []Entry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries"`
}

func (m *GenesisEntries) Reset()         { *m = GenesisEntries{} }
func (m *GenesisEntries) String() string { return proto.CompactTextString(m) }
func (*GenesisEntries) ProtoMessage()    {}
func (*GenesisEntries) Descriptor() ([]byte, []int) {
	return fileDescriptor_872ca3cbd39027a3, []int{1}
}
func (m *GenesisEntries) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisEntries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisEntries.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisEntries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisEntries.Merge(m, src)
}
func (m *GenesisEntries) XXX_Size() int {
	return m.Size()
}
func (m *GenesisEntries) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisEntries.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisEntries proto.InternalMessageInfo

func (m *GenesisEntries) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *GenesisEntries) GetIsLive() bool {
	if m != nil {
		return m.IsLive
	}
	return false
}

func (m *GenesisEntries) GetEntries() []Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type GenesisState struct {
	Version    uint64             `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Entries    []GenesisEntries   `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries"`
	Timerstore types.GenesisState `protobuf:"bytes,3,opt,name=timerstore,proto3" json:"timerstore"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_872ca3cbd39027a3, []int{2}
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

func (m *GenesisState) GetEntries() []GenesisEntries {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *GenesisState) GetTimerstore() types.GenesisState {
	if m != nil {
		return m.Timerstore
	}
	return types.GenesisState{}
}

func init() {
	proto.RegisterType((*Entry)(nil), "lavanet.lava.fixationstore.Entry")
	proto.RegisterType((*GenesisEntries)(nil), "lavanet.lava.fixationstore.GenesisEntries")
	proto.RegisterType((*GenesisState)(nil), "lavanet.lava.fixationstore.GenesisState")
}

func init() {
	proto.RegisterFile("lavanet/lava/fixationstore/fixation.proto", fileDescriptor_872ca3cbd39027a3)
}

var fileDescriptor_872ca3cbd39027a3 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x8e, 0xb7, 0x69, 0x93, 0xf5, 0xae, 0x38, 0x58, 0x2b, 0x61, 0x82, 0x14, 0x42, 0x11, 0x52,
	0xe0, 0x90, 0xa0, 0x85, 0x17, 0xe8, 0x4a, 0x08, 0x69, 0xd9, 0x53, 0xb8, 0x71, 0x59, 0xb9, 0xed,
	0x6c, 0xb1, 0x48, 0xe3, 0x2a, 0x9e, 0x46, 0xed, 0x95, 0x27, 0xe0, 0x71, 0xb8, 0x72, 0xeb, 0xb1,
	0x47, 0x4e, 0x08, 0xb5, 0x2f, 0x82, 0x62, 0xf7, 0x2f, 0x95, 0xba, 0x27, 0xcf, 0xe7, 0x99, 0xef,
	0x9b, 0xf9, 0xec, 0xa1, 0x6f, 0x72, 0x51, 0x89, 0x02, 0x30, 0xad, 0xcf, 0xf4, 0x41, 0xce, 0x04,
	0x4a, 0x55, 0x68, 0x54, 0x25, 0xec, 0x50, 0x32, 0x29, 0x15, 0x2a, 0x16, 0x6c, 0x4a, 0x93, 0xfa,
	0x4c, 0x1a, 0xa5, 0xc1, 0xd5, 0x48, 0x8d, 0x94, 0x29, 0x4b, 0xeb, 0xc8, 0x32, 0x82, 0x57, 0x0d,
	0x71, 0x94, 0x63, 0x28, 0xad, 0xb2, 0x09, 0x6d, 0x51, 0xf7, 0x17, 0xa1, 0xed, 0x8f, 0x05, 0x96,
	0x73, 0x76, 0x45, 0xdb, 0xb2, 0x18, 0xc2, 0x8c, 0x93, 0x88, 0xc4, 0xe7, 0x99, 0x05, 0xf5, 0x6d,
	0x3f, 0x57, 0x83, 0xef, 0xfc, 0x2c, 0x22, 0xb1, 0x9b, 0x59, 0xc0, 0x9e, 0x51, 0x5f, 0xa3, 0xc8,
	0xe1, 0x5e, 0x20, 0x6f, 0x99, 0x84, 0x67, 0x70, 0x0f, 0x59, 0x40, 0xfd, 0x12, 0x1e, 0x06, 0x6a,
	0x5a, 0x20, 0x77, 0x4d, 0x6a, 0x87, 0x19, 0xa3, 0xee, 0x50, 0xa0, 0xe0, 0xed, 0x88, 0xc4, 0x97,
	0x99, 0x89, 0xd9, 0x73, 0x7a, 0x3e, 0x84, 0x1c, 0xd0, 0x68, 0x75, 0x2c, 0xc1, 0x5e, 0xf4, 0xb0,
	0x4e, 0x4a, 0x7d, 0x9f, 0x0b, 0x04, 0x8d, 0xdc, 0x8b, 0x48, 0xec, 0x67, 0xbe, 0xd4, 0x77, 0x06,
	0x77, 0x7f, 0x10, 0xfa, 0xe4, 0x13, 0x14, 0xa0, 0xa5, 0xae, 0x1d, 0x48, 0xd0, 0x27, 0x3c, 0x3c,
	0xa5, 0x5e, 0xad, 0x22, 0x2b, 0x30, 0x2e, 0xfc, 0xac, 0x23, 0xf5, 0x9d, 0xac, 0x80, 0xf5, 0xa8,
	0x07, 0x96, 0xc9, 0x5b, 0x51, 0x2b, 0xbe, 0xb8, 0x7e, 0x99, 0x9c, 0x7e, 0xe5, 0xc4, 0x3c, 0xd3,
	0x8d, 0xbb, 0xf8, 0xfb, 0xc2, 0xc9, 0xb6, 0xbc, 0xee, 0x6f, 0x42, 0x2f, 0x37, 0x43, 0x7c, 0x41,
	0x81, 0xc0, 0x38, 0xf5, 0x2a, 0x28, 0xb5, 0x54, 0x85, 0x19, 0xc2, 0xcd, 0xb6, 0x90, 0xdd, 0xee,
	0xbb, 0x9d, 0x99, 0x6e, 0x6f, 0x1f, 0xeb, 0xd6, 0x74, 0x76, 0xd4, 0x96, 0x7d, 0xa6, 0x74, 0xff,
	0xa1, 0xe6, 0x0b, 0x2e, 0xae, 0x5f, 0x37, 0xe5, 0xf6, 0xf9, 0xe4, 0x70, 0xc0, 0x8d, 0xd2, 0x01,
	0xfd, 0xe6, 0x76, 0xb1, 0x0a, 0xc9, 0x72, 0x15, 0x92, 0x7f, 0xab, 0x90, 0xfc, 0x5c, 0x87, 0xce,
	0x72, 0x1d, 0x3a, 0x7f, 0xd6, 0xa1, 0xf3, 0xf5, 0xdd, 0x48, 0xe2, 0xb7, 0x69, 0x3f, 0x19, 0xa8,
	0x71, 0xda, 0xd8, 0xa6, 0xea, 0x43, 0x3a, 0x3b, 0xda, 0x57, 0x9c, 0x4f, 0x40, 0xf7, 0x3b, 0x66,
	0xad, 0xde, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xa4, 0x89, 0x38, 0xda, 0x02, 0x00, 0x00,
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Entry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsLatest {
		i--
		if m.IsLatest {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.DeleteAt != 0 {
		i = encodeVarintFixation(dAtA, i, uint64(m.DeleteAt))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintFixation(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Refcount != 0 {
		i = encodeVarintFixation(dAtA, i, uint64(m.Refcount))
		i--
		dAtA[i] = 0x20
	}
	if m.StaleAt != 0 {
		i = encodeVarintFixation(dAtA, i, uint64(m.StaleAt))
		i--
		dAtA[i] = 0x18
	}
	if m.Block != 0 {
		i = encodeVarintFixation(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFixation(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisEntries) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisEntries) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisEntries) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFixation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.IsLive {
		i--
		if m.IsLive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFixation(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
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
	{
		size, err := m.Timerstore.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintFixation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Entries) > 0 {
		for iNdEx := len(m.Entries) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Entries[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintFixation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Version != 0 {
		i = encodeVarintFixation(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintFixation(dAtA []byte, offset int, v uint64) int {
	offset -= sovFixation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Entry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFixation(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovFixation(uint64(m.Block))
	}
	if m.StaleAt != 0 {
		n += 1 + sovFixation(uint64(m.StaleAt))
	}
	if m.Refcount != 0 {
		n += 1 + sovFixation(uint64(m.Refcount))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovFixation(uint64(l))
	}
	if m.DeleteAt != 0 {
		n += 1 + sovFixation(uint64(m.DeleteAt))
	}
	if m.IsLatest {
		n += 2
	}
	return n
}

func (m *GenesisEntries) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFixation(uint64(l))
	}
	if m.IsLive {
		n += 2
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovFixation(uint64(l))
		}
	}
	return n
}

func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovFixation(uint64(m.Version))
	}
	if len(m.Entries) > 0 {
		for _, e := range m.Entries {
			l = e.Size()
			n += 1 + l + sovFixation(uint64(l))
		}
	}
	l = m.Timerstore.Size()
	n += 1 + l + sovFixation(uint64(l))
	return n
}

func sovFixation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFixation(x uint64) (n int) {
	return sovFixation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFixation
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
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
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
				return ErrInvalidLengthFixation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFixation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StaleAt", wireType)
			}
			m.StaleAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StaleAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Refcount", wireType)
			}
			m.Refcount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Refcount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
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
				return ErrInvalidLengthFixation
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFixation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeleteAt", wireType)
			}
			m.DeleteAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DeleteAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLatest", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLatest = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipFixation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFixation
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
func (m *GenesisEntries) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFixation
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
			return fmt.Errorf("proto: GenesisEntries: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisEntries: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
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
				return ErrInvalidLengthFixation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFixation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsLive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsLive = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
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
				return ErrInvalidLengthFixation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFixation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, Entry{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFixation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFixation
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
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFixation
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
					return ErrIntOverflowFixation
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entries", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
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
				return ErrInvalidLengthFixation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFixation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entries = append(m.Entries, GenesisEntries{})
			if err := m.Entries[len(m.Entries)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timerstore", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFixation
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
				return ErrInvalidLengthFixation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFixation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Timerstore.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFixation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFixation
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
func skipFixation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFixation
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
					return 0, ErrIntOverflowFixation
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
					return 0, ErrIntOverflowFixation
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
				return 0, ErrInvalidLengthFixation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFixation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFixation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFixation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFixation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFixation = fmt.Errorf("proto: unexpected end of group")
)

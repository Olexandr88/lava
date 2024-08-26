// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/epochstorage/endpoint.proto

package types

import (
	fmt "fmt"
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

type Endpoint struct {
	IPPORT        string   `protobuf:"bytes,1,opt,name=iPPORT,proto3" json:"iPPORT,omitempty"`
	Geolocation   int32    `protobuf:"varint,3,opt,name=geolocation,proto3" json:"geolocation,omitempty"`
	Addons        []string `protobuf:"bytes,4,rep,name=addons,proto3" json:"addons,omitempty"`
	ApiInterfaces []string `protobuf:"bytes,5,rep,name=api_interfaces,json=apiInterfaces,proto3" json:"api_interfaces,omitempty"`
	Extensions    []string `protobuf:"bytes,6,rep,name=extensions,proto3" json:"extensions,omitempty"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb18a6b0d300ae9, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return m.Size()
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetIPPORT() string {
	if m != nil {
		return m.IPPORT
	}
	return ""
}

func (m *Endpoint) GetGeolocation() int32 {
	if m != nil {
		return m.Geolocation
	}
	return 0
}

func (m *Endpoint) GetAddons() []string {
	if m != nil {
		return m.Addons
	}
	return nil
}

func (m *Endpoint) GetApiInterfaces() []string {
	if m != nil {
		return m.ApiInterfaces
	}
	return nil
}

func (m *Endpoint) GetExtensions() []string {
	if m != nil {
		return m.Extensions
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "lavanet.lava.epochstorage.Endpoint")
}

func init() {
	proto.RegisterFile("lavanet/lava/epochstorage/endpoint.proto", fileDescriptor_acb18a6b0d300ae9)
}

var fileDescriptor_acb18a6b0d300ae9 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xc8, 0x49, 0x2c, 0x4b,
	0xcc, 0x4b, 0x2d, 0xd1, 0x07, 0xd1, 0xfa, 0xa9, 0x05, 0xf9, 0xc9, 0x19, 0xc5, 0x25, 0xf9, 0x45,
	0x89, 0xe9, 0xa9, 0xfa, 0xa9, 0x79, 0x29, 0x05, 0xf9, 0x99, 0x79, 0x25, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x92, 0x50, 0x95, 0x7a, 0x20, 0x5a, 0x0f, 0x59, 0xa5, 0xd2, 0x4a, 0x46, 0x2e,
	0x0e, 0x57, 0xa8, 0x6a, 0x21, 0x31, 0x2e, 0xb6, 0xcc, 0x80, 0x00, 0xff, 0xa0, 0x10, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x48, 0x81, 0x8b, 0x3b, 0x3d, 0x35, 0x3f, 0x27, 0x3f,
	0x39, 0xb1, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x59, 0x81, 0x51, 0x83, 0x35, 0x08, 0x59, 0x08, 0xa4,
	0x33, 0x31, 0x25, 0x25, 0x3f, 0xaf, 0x58, 0x82, 0x45, 0x81, 0x19, 0xa4, 0x13, 0xc2, 0x13, 0x52,
	0xe5, 0xe2, 0x4b, 0x2c, 0xc8, 0x8c, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0x2d,
	0x96, 0x60, 0x05, 0xcb, 0xf3, 0x26, 0x16, 0x64, 0x7a, 0xc2, 0x05, 0x85, 0xe4, 0xb8, 0xb8, 0x52,
	0x2b, 0x4a, 0x52, 0xf3, 0x8a, 0x33, 0x41, 0x46, 0xb0, 0x81, 0x95, 0x20, 0x89, 0x78, 0xb1, 0x70,
	0x30, 0x09, 0x30, 0x3b, 0x79, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0x7e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x4a, 0xa8, 0x94, 0x19,
	0xe9, 0x57, 0xa0, 0x06, 0x4d, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x60, 0x8c, 0x01,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0xfe, 0x6a, 0xaa, 0x44, 0x01, 0x00, 0x00,
}

func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Endpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Extensions) > 0 {
		for iNdEx := len(m.Extensions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Extensions[iNdEx])
			copy(dAtA[i:], m.Extensions[iNdEx])
			i = encodeVarintEndpoint(dAtA, i, uint64(len(m.Extensions[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.ApiInterfaces) > 0 {
		for iNdEx := len(m.ApiInterfaces) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ApiInterfaces[iNdEx])
			copy(dAtA[i:], m.ApiInterfaces[iNdEx])
			i = encodeVarintEndpoint(dAtA, i, uint64(len(m.ApiInterfaces[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Addons) > 0 {
		for iNdEx := len(m.Addons) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addons[iNdEx])
			copy(dAtA[i:], m.Addons[iNdEx])
			i = encodeVarintEndpoint(dAtA, i, uint64(len(m.Addons[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Geolocation != 0 {
		i = encodeVarintEndpoint(dAtA, i, uint64(m.Geolocation))
		i--
		dAtA[i] = 0x18
	}
	if len(m.IPPORT) > 0 {
		i -= len(m.IPPORT)
		copy(dAtA[i:], m.IPPORT)
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.IPPORT)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEndpoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovEndpoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Endpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IPPORT)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if m.Geolocation != 0 {
		n += 1 + sovEndpoint(uint64(m.Geolocation))
	}
	if len(m.Addons) > 0 {
		for _, s := range m.Addons {
			l = len(s)
			n += 1 + l + sovEndpoint(uint64(l))
		}
	}
	if len(m.ApiInterfaces) > 0 {
		for _, s := range m.ApiInterfaces {
			l = len(s)
			n += 1 + l + sovEndpoint(uint64(l))
		}
	}
	if len(m.Extensions) > 0 {
		for _, s := range m.Extensions {
			l = len(s)
			n += 1 + l + sovEndpoint(uint64(l))
		}
	}
	return n
}

func sovEndpoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEndpoint(x uint64) (n int) {
	return sovEndpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPPORT", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEndpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPPORT = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Geolocation", wireType)
			}
			m.Geolocation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addons", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEndpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addons = append(m.Addons, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiInterfaces", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEndpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApiInterfaces = append(m.ApiInterfaces, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEndpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extensions = append(m.Extensions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEndpoint
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
func skipEndpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
				return 0, ErrInvalidLengthEndpoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEndpoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEndpoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEndpoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEndpoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEndpoint = fmt.Errorf("proto: unexpected end of group")
)

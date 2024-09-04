// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lavanet/lava/pairing/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	EpochBlocksOverlap                    uint64                                 `protobuf:"varint,8,opt,name=epochBlocksOverlap,proto3" json:"epochBlocksOverlap,omitempty" yaml:"epoch_blocks_overlap"`
	QoSWeight                             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,13,opt,name=QoSWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"QoSWeight" yaml:"data_reliability_reward"`
	RecommendedEpochNumToCollectPayment   uint64                                 `protobuf:"varint,14,opt,name=recommendedEpochNumToCollectPayment,proto3" json:"recommendedEpochNumToCollectPayment,omitempty" yaml:"recommended_epoch_num_to_collect_payment"`
	ReputationVarianceStabilizationPeriod int64                                  `protobuf:"varint,15,opt,name=reputation_variance_stabilization_period,json=reputationVarianceStabilizationPeriod,proto3" json:"reputation_variance_stabilization_period,omitempty" yaml:"reputation_variance_stabilization_period"`
	ReputationLatencyOverSyncFactor       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,16,opt,name=reputation_latency_over_sync_factor,json=reputationLatencyOverSyncFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reputation_latency_over_sync_factor" yaml:"reputation_latency_over_sync_factor"`
	ReputationHalfLifeFactor              int64                                  `protobuf:"varint,17,opt,name=reputation_half_life_factor,json=reputationHalfLifeFactor,proto3" json:"reputation_half_life_factor,omitempty" yaml:"reputation_half_life_factor"`
	ReputationRelayFailureCost            uint64                                 `protobuf:"varint,18,opt,name=reputation_relay_failure_cost,json=reputationRelayFailureCost,proto3" json:"reputation_relay_failure_cost,omitempty" yaml:"reputation_relay_failure_cost"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc338fce33b3b67a, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetEpochBlocksOverlap() uint64 {
	if m != nil {
		return m.EpochBlocksOverlap
	}
	return 0
}

func (m *Params) GetRecommendedEpochNumToCollectPayment() uint64 {
	if m != nil {
		return m.RecommendedEpochNumToCollectPayment
	}
	return 0
}

func (m *Params) GetReputationVarianceStabilizationPeriod() int64 {
	if m != nil {
		return m.ReputationVarianceStabilizationPeriod
	}
	return 0
}

func (m *Params) GetReputationHalfLifeFactor() int64 {
	if m != nil {
		return m.ReputationHalfLifeFactor
	}
	return 0
}

func (m *Params) GetReputationRelayFailureCost() uint64 {
	if m != nil {
		return m.ReputationRelayFailureCost
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "lavanet.lava.pairing.Params")
}

func init() { proto.RegisterFile("lavanet/lava/pairing/params.proto", fileDescriptor_fc338fce33b3b67a) }

var fileDescriptor_fc338fce33b3b67a = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6f, 0xd3, 0x3e,
	0x14, 0xc7, 0x9b, 0xcd, 0xeb, 0xbc, 0x6c, 0xfb, 0xfd, 0x42, 0xb4, 0x43, 0xb4, 0x89, 0x64, 0x04,
	0x98, 0xaa, 0x49, 0x34, 0x87, 0xdd, 0x76, 0xec, 0xc6, 0x84, 0xa2, 0x89, 0x95, 0x0c, 0x81, 0x84,
	0x90, 0x2c, 0xd7, 0x75, 0x5b, 0x6b, 0x4e, 0x1c, 0x39, 0x6e, 0x21, 0xdc, 0xb9, 0x71, 0xe0, 0xc8,
	0x91, 0x2b, 0xfc, 0x25, 0x3b, 0xee, 0x88, 0x38, 0x44, 0x68, 0xfb, 0x0f, 0xfa, 0x17, 0xa0, 0x38,
	0x65, 0xed, 0x28, 0x48, 0xe3, 0xf4, 0xa2, 0xbc, 0xcf, 0xf7, 0x7d, 0x9f, 0xfd, 0x6c, 0x9b, 0xf7,
	0x38, 0x1e, 0xe1, 0x84, 0xaa, 0xa0, 0x8c, 0x41, 0x8a, 0x99, 0x64, 0x49, 0x3f, 0x48, 0xb1, 0xc4,
	0x71, 0xd6, 0x4c, 0xa5, 0x50, 0xc2, 0xde, 0x98, 0x20, 0xcd, 0x32, 0x36, 0x27, 0xc8, 0xe6, 0x46,
	0x5f, 0xf4, 0x85, 0x06, 0x82, 0xf2, 0xab, 0x62, 0xfd, 0x2f, 0xcb, 0x66, 0xbd, 0xad, 0xc5, 0xf6,
	0x89, 0x69, 0xd3, 0x54, 0x90, 0x41, 0x8b, 0x0b, 0x72, 0x96, 0x9d, 0x8c, 0xa8, 0xe4, 0x38, 0x75,
	0xe0, 0xb6, 0xd1, 0x00, 0x2d, 0x6f, 0x5c, 0x78, 0x5b, 0x39, 0x8e, 0xf9, 0xbe, 0xaf, 0x19, 0xd4,
	0xd1, 0x10, 0x12, 0x15, 0xe5, 0x47, 0x7f, 0x90, 0xda, 0x89, 0xb9, 0xf2, 0x4c, 0x9c, 0xbe, 0xa4,
	0xac, 0x3f, 0x50, 0xce, 0xfa, 0xb6, 0xd1, 0x58, 0x69, 0xb5, 0xcf, 0x0b, 0xaf, 0xf6, 0xbd, 0xf0,
	0x76, 0xfa, 0x4c, 0x0d, 0x86, 0x9d, 0x26, 0x11, 0x71, 0x40, 0x44, 0x16, 0x8b, 0x6c, 0x12, 0x1e,
	0x65, 0xdd, 0xb3, 0x40, 0xe5, 0x29, 0xcd, 0x9a, 0x87, 0x94, 0x8c, 0x0b, 0xcf, 0xad, 0x5c, 0xbb,
	0x58, 0x61, 0x24, 0x29, 0x67, 0xb8, 0xc3, 0x38, 0x53, 0x39, 0x92, 0xf4, 0x0d, 0x96, 0x5d, 0x3f,
	0x9a, 0x5a, 0xd8, 0xef, 0x0d, 0xf3, 0xbe, 0xa4, 0x44, 0xc4, 0x31, 0x4d, 0xba, 0xb4, 0xfb, 0xb8,
	0xec, 0xe8, 0xe9, 0x30, 0x7e, 0x2e, 0x0e, 0x04, 0xe7, 0x94, 0xa8, 0x36, 0xce, 0x63, 0x9a, 0x28,
	0xe7, 0x3f, 0xbd, 0xa4, 0xbd, 0x71, 0xe1, 0x05, 0x55, 0xf1, 0x19, 0x11, 0xaa, 0x96, 0x97, 0x0c,
	0x63, 0xa4, 0x04, 0x22, 0x95, 0x10, 0xa5, 0x95, 0xd2, 0x8f, 0x6e, 0x53, 0xdf, 0xfe, 0x60, 0x98,
	0x0d, 0x49, 0xd3, 0xa1, 0xc2, 0x8a, 0x89, 0x04, 0x8d, 0xb0, 0x64, 0x38, 0x21, 0x14, 0x65, 0x4a,
	0x37, 0xff, 0xae, 0xfa, 0x9d, 0x52, 0xc9, 0x44, 0xd7, 0xf9, 0x7f, 0xdb, 0x68, 0x2c, 0xde, 0x6c,
	0xe6, 0x76, 0x4a, 0x3f, 0x7a, 0x38, 0x45, 0x5f, 0x4c, 0xc8, 0xd3, 0x59, 0xb0, 0xad, 0x39, 0xfb,
	0xab, 0xde, 0x96, 0xeb, 0xa2, 0x1c, 0x2b, 0x9a, 0x90, 0x5c, 0x8f, 0x0e, 0x65, 0x79, 0x42, 0x50,
	0x0f, 0x13, 0x25, 0xa4, 0x63, 0xe9, 0x09, 0xbd, 0xfe, 0xe7, 0x09, 0xed, 0xce, 0xf5, 0xfd, 0x37,
	0x0b, 0x3f, 0xf2, 0xa6, 0xd4, 0x71, 0x05, 0x95, 0x87, 0xe5, 0x34, 0x4f, 0xc8, 0x91, 0x26, 0x6c,
	0x6a, 0x6e, 0xcd, 0x14, 0x1a, 0x60, 0xde, 0x43, 0x9c, 0xf5, 0xe8, 0xaf, 0x1e, 0xef, 0xe8, 0xdd,
	0xda, 0x19, 0x17, 0x9e, 0x3f, 0xe7, 0xfa, 0x3b, 0xec, 0x47, 0xce, 0x34, 0xfb, 0x04, 0xf3, 0xde,
	0x31, 0xeb, 0xd1, 0x89, 0xcd, 0x99, 0x79, 0x77, 0x46, 0x29, 0x29, 0xc7, 0x39, 0xea, 0x61, 0xc6,
	0x87, 0x92, 0x22, 0x22, 0x32, 0xe5, 0xd8, 0xfa, 0x8c, 0x34, 0xc6, 0x85, 0xf7, 0x60, 0xce, 0x68,
	0x1e, 0xf7, 0xa3, 0xcd, 0x69, 0x3e, 0x2a, 0xd3, 0x47, 0x55, 0xf6, 0x40, 0x64, 0x6a, 0x1f, 0x7c,
	0xfa, 0xec, 0xd5, 0x42, 0x00, 0x0d, 0x6b, 0x21, 0x04, 0x70, 0xc1, 0x5a, 0x0c, 0x01, 0x5c, 0xb4,
	0x40, 0x08, 0x20, 0xb0, 0x96, 0x42, 0x00, 0x97, 0xac, 0x7a, 0x08, 0x60, 0xdd, 0x5a, 0x0e, 0x01,
	0x5c, 0xb6, 0x60, 0x08, 0xe0, 0x8a, 0x65, 0x86, 0x00, 0x9a, 0xd6, 0x6a, 0x08, 0xe0, 0xaa, 0xb5,
	0x16, 0x02, 0xb8, 0x66, 0xad, 0xb7, 0x0e, 0xcf, 0x2f, 0x5d, 0xe3, 0xe2, 0xd2, 0x35, 0x7e, 0x5c,
	0xba, 0xc6, 0xc7, 0x2b, 0xb7, 0x76, 0x71, 0xe5, 0xd6, 0xbe, 0x5d, 0xb9, 0xb5, 0x57, 0xbb, 0x33,
	0xc3, 0xba, 0xf1, 0x3e, 0x8c, 0xf6, 0x82, 0xb7, 0xd7, 0x8f, 0x84, 0x1e, 0x5a, 0xa7, 0xae, 0x2f,
	0xfe, 0xde, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xbb, 0xe1, 0xe7, 0x49, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReputationRelayFailureCost != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReputationRelayFailureCost))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	if m.ReputationHalfLifeFactor != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReputationHalfLifeFactor))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	{
		size := m.ReputationLatencyOverSyncFactor.Size()
		i -= size
		if _, err := m.ReputationLatencyOverSyncFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x82
	if m.ReputationVarianceStabilizationPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ReputationVarianceStabilizationPeriod))
		i--
		dAtA[i] = 0x78
	}
	if m.RecommendedEpochNumToCollectPayment != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RecommendedEpochNumToCollectPayment))
		i--
		dAtA[i] = 0x70
	}
	{
		size := m.QoSWeight.Size()
		i -= size
		if _, err := m.QoSWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.EpochBlocksOverlap != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.EpochBlocksOverlap))
		i--
		dAtA[i] = 0x40
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochBlocksOverlap != 0 {
		n += 1 + sovParams(uint64(m.EpochBlocksOverlap))
	}
	l = m.QoSWeight.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.RecommendedEpochNumToCollectPayment != 0 {
		n += 1 + sovParams(uint64(m.RecommendedEpochNumToCollectPayment))
	}
	if m.ReputationVarianceStabilizationPeriod != 0 {
		n += 1 + sovParams(uint64(m.ReputationVarianceStabilizationPeriod))
	}
	l = m.ReputationLatencyOverSyncFactor.Size()
	n += 2 + l + sovParams(uint64(l))
	if m.ReputationHalfLifeFactor != 0 {
		n += 2 + sovParams(uint64(m.ReputationHalfLifeFactor))
	}
	if m.ReputationRelayFailureCost != 0 {
		n += 2 + sovParams(uint64(m.ReputationRelayFailureCost))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochBlocksOverlap", wireType)
			}
			m.EpochBlocksOverlap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochBlocksOverlap |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QoSWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.QoSWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecommendedEpochNumToCollectPayment", wireType)
			}
			m.RecommendedEpochNumToCollectPayment = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RecommendedEpochNumToCollectPayment |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReputationVarianceStabilizationPeriod", wireType)
			}
			m.ReputationVarianceStabilizationPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReputationVarianceStabilizationPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReputationLatencyOverSyncFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReputationLatencyOverSyncFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReputationHalfLifeFactor", wireType)
			}
			m.ReputationHalfLifeFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReputationHalfLifeFactor |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReputationRelayFailureCost", wireType)
			}
			m.ReputationRelayFailureCost = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReputationRelayFailureCost |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feemarket/feemarket/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// Params contains the required set of parameters for the EIP1559 fee market
// plugin implementation.
type Params struct {
	// Alpha is the amount we additively increase the learning rate
	// when it is above or below the target +/- threshold.
	Alpha cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=alpha,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"alpha"`
	// Beta is the amount we multiplicatively decrease the learning rate
	// when it is within the target +/- threshold.
	Beta cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=beta,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"beta"`
	// Theta is the threshold for the learning rate. If the learning rate is
	// above or below the target +/- threshold, we additively increase the
	// learning rate by Alpha. Otherwise, we multiplicatively decrease the
	// learning rate by Beta.
	Theta cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=theta,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"theta"`
	// MinLearningRate is the lower bound for the learning rate.
	MinLearningRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=min_learning_rate,json=minLearningRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"min_learning_rate"`
	// MaxLearningRate is the upper bound for the learning rate.
	MaxLearningRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,5,opt,name=max_learning_rate,json=maxLearningRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"max_learning_rate"`
	// TargetBlockUtilization is the target block utilization.
	TargetBlockUtilization uint64 `protobuf:"varint,6,opt,name=target_block_utilization,json=targetBlockUtilization,proto3" json:"target_block_utilization,omitempty"`
	// MaxBlockUtilization is the maximum block utilization.
	MaxBlockUtilization uint64 `protobuf:"varint,7,opt,name=max_block_utilization,json=maxBlockUtilization,proto3" json:"max_block_utilization,omitempty"`
	// Window defines the window size for calculating an adaptive learning rate
	// over a moving window of blocks.
	Window uint64 `protobuf:"varint,8,opt,name=window,proto3" json:"window,omitempty"`
	// Enabled is a boolean that determines whether the EIP1559 fee market is
	// enabled.
	Enabled bool `protobuf:"varint,9,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// DefaultFeeDenom is the default fee denom for the EIP1559 fee market
	// used to simulate transactions if there are no fees specified
	DefaultFeeDenom string `protobuf:"bytes,10,opt,name=default_fee_denom,json=defaultFeeDenom,proto3" json:"default_fee_denom,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3907de4df2e1c66e, []int{0}
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

func (m *Params) GetTargetBlockUtilization() uint64 {
	if m != nil {
		return m.TargetBlockUtilization
	}
	return 0
}

func (m *Params) GetMaxBlockUtilization() uint64 {
	if m != nil {
		return m.MaxBlockUtilization
	}
	return 0
}

func (m *Params) GetWindow() uint64 {
	if m != nil {
		return m.Window
	}
	return 0
}

func (m *Params) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Params) GetDefaultFeeDenom() string {
	if m != nil {
		return m.DefaultFeeDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "feemarket.feemarket.v1.Params")
}

func init() {
	proto.RegisterFile("feemarket/feemarket/v1/params.proto", fileDescriptor_3907de4df2e1c66e)
}

var fileDescriptor_3907de4df2e1c66e = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x13, 0x4d, 0xd3, 0x76, 0x2e, 0xa5, 0x51, 0x97, 0xb1, 0x42, 0xba, 0xe8, 0x65, 0x11,
	0x9a, 0xb0, 0x7a, 0xf1, 0xbc, 0xac, 0x8a, 0xd0, 0x83, 0x04, 0xbc, 0x08, 0x12, 0x5e, 0x92, 0xb7,
	0xc9, 0x90, 0xcc, 0x4c, 0x48, 0x66, 0xb7, 0xa9, 0x9f, 0xc2, 0x0f, 0xe3, 0x87, 0xe8, 0xb1, 0x78,
	0x12, 0x0f, 0x45, 0x76, 0x8f, 0x7e, 0x09, 0xc9, 0x24, 0xba, 0xab, 0x7b, 0xcb, 0xed, 0xff, 0xde,
	0xff, 0xfd, 0x7f, 0xf3, 0x06, 0x1e, 0x79, 0xb6, 0x40, 0xe4, 0x50, 0xe5, 0xa8, 0xfc, 0xad, 0x5a,
	0x4d, 0xfd, 0x12, 0x2a, 0xe0, 0xb5, 0x57, 0x56, 0x52, 0x49, 0x67, 0xf4, 0xd7, 0xf2, 0xb6, 0x6a,
	0x35, 0x3d, 0x7b, 0x1c, 0xcb, 0x9a, 0xcb, 0x3a, 0xd4, 0x53, 0x7e, 0x57, 0x74, 0x91, 0xb3, 0x87,
	0xa9, 0x4c, 0x65, 0xd7, 0x6f, 0x55, 0xd7, 0x7d, 0xfa, 0xcb, 0x22, 0xf6, 0x7b, 0x4d, 0x76, 0xde,
	0x92, 0x03, 0x28, 0xca, 0x0c, 0xa8, 0x39, 0x36, 0x27, 0xc7, 0xb3, 0xe9, 0xcd, 0xdd, 0xb9, 0xf1,
	0xe3, 0xee, 0xfc, 0x49, 0x47, 0xa9, 0x93, 0xdc, 0x63, 0xd2, 0xe7, 0xa0, 0x32, 0xef, 0x12, 0x53,
	0x88, 0xaf, 0xe7, 0x18, 0x7f, 0xfb, 0x7a, 0x41, 0xfa, 0x47, 0xe6, 0x18, 0x07, 0x5d, 0xde, 0x79,
	0x4d, 0xac, 0x08, 0x15, 0xd0, 0x7b, 0x43, 0x39, 0x3a, 0xde, 0xee, 0xa3, 0xb2, 0x96, 0x73, 0x7f,
	0xf0, 0x3e, 0x3a, 0xef, 0x7c, 0x22, 0xa7, 0x9c, 0x89, 0xb0, 0x40, 0xa8, 0x04, 0x13, 0x69, 0x58,
	0x81, 0x42, 0x6a, 0x0d, 0x85, 0x9e, 0x70, 0x26, 0x2e, 0x7b, 0x54, 0x00, 0x0a, 0x35, 0x1e, 0x9a,
	0xff, 0xf0, 0x07, 0xc3, 0xf1, 0xd0, 0xfc, 0x83, 0x7f, 0x45, 0xa8, 0x82, 0x2a, 0x45, 0x15, 0x46,
	0x85, 0x8c, 0xf3, 0x70, 0xa9, 0x58, 0xc1, 0x3e, 0x83, 0x62, 0x52, 0x50, 0x7b, 0x6c, 0x4e, 0xac,
	0x60, 0xd4, 0xf9, 0xb3, 0xd6, 0xfe, 0xb0, 0x75, 0x9d, 0x17, 0xe4, 0x51, 0xbb, 0xd8, 0x7e, 0xec,
	0x50, 0xc7, 0x1e, 0x70, 0x68, 0xf6, 0x32, 0x23, 0x62, 0x5f, 0x31, 0x91, 0xc8, 0x2b, 0x7a, 0xa4,
	0x87, 0xfa, 0xca, 0xa1, 0xe4, 0x10, 0x05, 0x44, 0x05, 0x26, 0xf4, 0x78, 0x6c, 0x4e, 0x8e, 0x82,
	0x3f, 0xa5, 0xf3, 0x9c, 0x9c, 0x26, 0xb8, 0x80, 0x65, 0xa1, 0xc2, 0x05, 0x62, 0x98, 0xa0, 0x90,
	0x9c, 0x92, 0xf6, 0xfb, 0xc1, 0x49, 0x6f, 0xbc, 0x41, 0x9c, 0xb7, 0xed, 0xd9, 0xbb, 0x9b, 0xb5,
	0x6b, 0xde, 0xae, 0x5d, 0xf3, 0xe7, 0xda, 0x35, 0xbf, 0x6c, 0x5c, 0xe3, 0x76, 0xe3, 0x1a, 0xdf,
	0x37, 0xae, 0xf1, 0xd1, 0x4f, 0x99, 0xca, 0x96, 0x91, 0x17, 0x4b, 0xee, 0xd7, 0x39, 0x2b, 0x2f,
	0x38, 0xae, 0x76, 0xee, 0xbf, 0xd9, 0xd1, 0xea, 0xba, 0xc4, 0x3a, 0xb2, 0xf5, 0xfd, 0xbe, 0xfc,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0x72, 0x27, 0xdd, 0xb2, 0x2f, 0x03, 0x00, 0x00,
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
	if len(m.DefaultFeeDenom) > 0 {
		i -= len(m.DefaultFeeDenom)
		copy(dAtA[i:], m.DefaultFeeDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.DefaultFeeDenom)))
		i--
		dAtA[i] = 0x52
	}
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if m.Window != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.Window))
		i--
		dAtA[i] = 0x40
	}
	if m.MaxBlockUtilization != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxBlockUtilization))
		i--
		dAtA[i] = 0x38
	}
	if m.TargetBlockUtilization != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TargetBlockUtilization))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.MaxLearningRate.Size()
		i -= size
		if _, err := m.MaxLearningRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MinLearningRate.Size()
		i -= size
		if _, err := m.MinLearningRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Theta.Size()
		i -= size
		if _, err := m.Theta.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Beta.Size()
		i -= size
		if _, err := m.Beta.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Alpha.Size()
		i -= size
		if _, err := m.Alpha.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Alpha.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Beta.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.Theta.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinLearningRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxLearningRate.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.TargetBlockUtilization != 0 {
		n += 1 + sovParams(uint64(m.TargetBlockUtilization))
	}
	if m.MaxBlockUtilization != 0 {
		n += 1 + sovParams(uint64(m.MaxBlockUtilization))
	}
	if m.Window != 0 {
		n += 1 + sovParams(uint64(m.Window))
	}
	if m.Enabled {
		n += 2
	}
	l = len(m.DefaultFeeDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Alpha", wireType)
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
			if err := m.Alpha.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Beta", wireType)
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
			if err := m.Beta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Theta", wireType)
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
			if err := m.Theta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinLearningRate", wireType)
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
			if err := m.MinLearningRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLearningRate", wireType)
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
			if err := m.MaxLearningRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBlockUtilization", wireType)
			}
			m.TargetBlockUtilization = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetBlockUtilization |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBlockUtilization", wireType)
			}
			m.MaxBlockUtilization = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxBlockUtilization |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Window", wireType)
			}
			m.Window = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Window |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.Enabled = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultFeeDenom", wireType)
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
			m.DefaultFeeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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

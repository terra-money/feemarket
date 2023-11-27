// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feemarket/feemarket/v1/genesis.proto

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

// GenesisState defines the feemarket module's genesis state.
type GenesisState struct {
	// Params are the parameters for the feemarket module. These parameters
	// can be utilized to implement both the base EIP-1559 fee market and
	// and the AIMD EIP-1559 fee market.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// State contains the current state of the AIMD fee market.
	State State `protobuf:"bytes,2,opt,name=state,proto3" json:"state"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_2180652c84279298, []int{0}
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

func (m *GenesisState) GetState() State {
	if m != nil {
		return m.State
	}
	return State{}
}

// State is utilized to track the current state of the fee market. This includes
// the current base fee, learning rate, and block utilization within the
// specified AIMD window.
type State struct {
	// BaseFee is the current base fee. This is denominated in the fee per gas
	// unit.
	BaseFee cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=base_fee,json=baseFee,proto3,customtype=cosmossdk.io/math.Int" json:"base_fee"`
	// MinBaseFee is the minimum base fee that can be charged. This is denominated
	// in the fee per gas unit.
	MinBaseFee cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=min_base_fee,json=minBaseFee,proto3,customtype=cosmossdk.io/math.Int" json:"min_base_fee"`
	// LearningRate is the current learning rate.
	LearningRate cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=learning_rate,json=learningRate,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"learning_rate"`
	// Window contains a list of the last blocks' utilization values. This is used
	// to calculate the next base fee. This stores the number of units of gas
	// consumed per block.
	Window []uint64 `protobuf:"varint,4,rep,packed,name=window,proto3" json:"window,omitempty"`
	// Index is the index of the current block in the block utilization window.
	Index uint64 `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
	// MaxBlockUtilization is the maximum utilization of a given block. This is
	// denominated in the number of gas units consumed per block.
	MaxBlockUtilization uint64 `protobuf:"varint,6,opt,name=max_block_utilization,json=maxBlockUtilization,proto3" json:"max_block_utilization,omitempty"`
	// TargetBlockUtilization is the target utilization of a given block. This is
	// denominated in the number of gas units consumed per block.
	TargetBlockUtilization uint64 `protobuf:"varint,7,opt,name=target_block_utilization,json=targetBlockUtilization,proto3" json:"target_block_utilization,omitempty"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_2180652c84279298, []int{1}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_State.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return m.Size()
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetWindow() []uint64 {
	if m != nil {
		return m.Window
	}
	return nil
}

func (m *State) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *State) GetMaxBlockUtilization() uint64 {
	if m != nil {
		return m.MaxBlockUtilization
	}
	return 0
}

func (m *State) GetTargetBlockUtilization() uint64 {
	if m != nil {
		return m.TargetBlockUtilization
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "feemarket.feemarket.v1.GenesisState")
	proto.RegisterType((*State)(nil), "feemarket.feemarket.v1.State")
}

func init() {
	proto.RegisterFile("feemarket/feemarket/v1/genesis.proto", fileDescriptor_2180652c84279298)
}

var fileDescriptor_2180652c84279298 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0xdd, 0x24, 0xd5, 0x71, 0x7b, 0x89, 0xed, 0x12, 0x2b, 0xa6, 0x4b, 0xf5, 0xb0,
	0x20, 0x9b, 0xd0, 0xea, 0x41, 0xc1, 0x53, 0x90, 0xca, 0x82, 0x82, 0x44, 0x3c, 0xe8, 0x25, 0x4c,
	0xb2, 0xaf, 0xe9, 0x90, 0x9d, 0x99, 0x90, 0x99, 0x6e, 0x53, 0xbf, 0x80, 0x57, 0x2f, 0x7e, 0x93,
	0x7e, 0x88, 0x1e, 0x4b, 0x4f, 0xe2, 0xa1, 0xc8, 0xee, 0x17, 0x91, 0xcc, 0x8c, 0xdd, 0xa2, 0xdd,
	0x83, 0xb7, 0xf7, 0xe6, 0xff, 0xff, 0xff, 0xde, 0x7b, 0x30, 0xe8, 0xc9, 0x21, 0x00, 0xc5, 0x75,
	0x09, 0x32, 0x5a, 0x56, 0xb3, 0xbd, 0xa8, 0x00, 0x06, 0x82, 0x88, 0xb0, 0xaa, 0xb9, 0xe4, 0x5e,
	0xff, 0x5a, 0x0b, 0x97, 0xd5, 0x6c, 0x6f, 0x7b, 0xb3, 0xe0, 0x05, 0x57, 0x96, 0xa8, 0xad, 0xb4,
	0x7b, 0xfb, 0x41, 0xce, 0x05, 0xe5, 0x22, 0xd5, 0x82, 0x6e, 0x8c, 0xf4, 0x78, 0xc5, 0xb8, 0x0a,
	0xd7, 0x98, 0x1a, 0xd3, 0xee, 0x57, 0x0b, 0xf5, 0xde, 0xe8, 0xf9, 0x1f, 0x24, 0x96, 0xe0, 0xbd,
	0x42, 0xae, 0x36, 0xf8, 0xd6, 0xc0, 0x1a, 0xde, 0xdb, 0x0f, 0xc2, 0xdb, 0xf7, 0x09, 0xdf, 0x2b,
	0x57, 0x6c, 0x9f, 0x5f, 0xed, 0x74, 0x12, 0x93, 0xf1, 0x5e, 0x22, 0x47, 0xb4, 0x18, 0x7f, 0x4d,
	0x85, 0x1f, 0xad, 0x0a, 0xab, 0x59, 0x26, 0xab, 0x13, 0xbb, 0xdf, 0xbb, 0xc8, 0xd1, 0x2b, 0x1c,
	0xa0, 0x3b, 0x19, 0x16, 0x90, 0x1e, 0x02, 0xa8, 0x25, 0xee, 0xc6, 0x4f, 0x5b, 0xe3, 0xcf, 0xab,
	0x9d, 0x2d, 0x7d, 0xa0, 0x98, 0x94, 0x21, 0xe1, 0x11, 0xc5, 0xf2, 0x28, 0x1c, 0x33, 0x79, 0x79,
	0x36, 0x42, 0xe6, 0xf2, 0x31, 0x93, 0xc9, 0x7a, 0x1b, 0x3e, 0x00, 0xf0, 0xde, 0xa1, 0x1e, 0x25,
	0x2c, 0xbd, 0x66, 0xad, 0xfd, 0x3f, 0x0b, 0x51, 0xc2, 0x62, 0x83, 0xfb, 0x84, 0x36, 0xa6, 0x80,
	0x6b, 0x46, 0x58, 0x91, 0xd6, 0xed, 0x8d, 0x5d, 0xc5, 0x7b, 0x6e, 0x78, 0x0f, 0xff, 0xe5, 0xbd,
	0x85, 0x02, 0xe7, 0xa7, 0xaf, 0x21, 0xbf, 0x3c, 0x1b, 0x6d, 0x18, 0xaa, 0x7e, 0x4b, 0x7a, 0x7f,
	0x50, 0x49, 0x7b, 0x71, 0x1f, 0xb9, 0x27, 0x84, 0x4d, 0xf8, 0x89, 0x6f, 0x0f, 0xba, 0x43, 0x3b,
	0x31, 0x9d, 0xb7, 0x89, 0x1c, 0xc2, 0x26, 0xd0, 0xf8, 0xce, 0xc0, 0x1a, 0xda, 0x89, 0x6e, 0xbc,
	0x7d, 0xb4, 0x45, 0x71, 0x93, 0x66, 0x53, 0x9e, 0x97, 0xe9, 0xb1, 0x24, 0x53, 0xf2, 0x05, 0x4b,
	0xc2, 0x99, 0xef, 0x2a, 0xd7, 0x7d, 0x8a, 0x9b, 0xb8, 0xd5, 0x3e, 0x2e, 0x25, 0xef, 0x05, 0xf2,
	0x25, 0xae, 0x0b, 0x90, 0xb7, 0xc4, 0xd6, 0x55, 0xac, 0xaf, 0xf5, 0xbf, 0x93, 0xf1, 0xf8, 0x7c,
	0x1e, 0x58, 0x17, 0xf3, 0xc0, 0xfa, 0x35, 0x0f, 0xac, 0x6f, 0x8b, 0xa0, 0x73, 0xb1, 0x08, 0x3a,
	0x3f, 0x16, 0x41, 0xe7, 0x73, 0x54, 0x10, 0x79, 0x74, 0x9c, 0x85, 0x39, 0xa7, 0x91, 0x28, 0x49,
	0x35, 0xa2, 0x30, 0xbb, 0xf1, 0xd5, 0x9a, 0x1b, 0xb5, 0x3c, 0xad, 0x40, 0x64, 0xae, 0xfa, 0x73,
	0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x42, 0xf0, 0xf1, 0x69, 0x09, 0x03, 0x00, 0x00,
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
		size, err := m.State.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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

func (m *State) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *State) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *State) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TargetBlockUtilization != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TargetBlockUtilization))
		i--
		dAtA[i] = 0x38
	}
	if m.MaxBlockUtilization != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxBlockUtilization))
		i--
		dAtA[i] = 0x30
	}
	if m.Index != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Window) > 0 {
		dAtA4 := make([]byte, len(m.Window)*10)
		var j3 int
		for _, num := range m.Window {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintGenesis(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.LearningRate.Size()
		i -= size
		if _, err := m.LearningRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinBaseFee.Size()
		i -= size
		if _, err := m.MinBaseFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.BaseFee.Size()
		i -= size
		if _, err := m.BaseFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
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
	l = m.State.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *State) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.MinBaseFee.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.LearningRate.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Window) > 0 {
		l = 0
		for _, e := range m.Window {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
	}
	if m.Index != 0 {
		n += 1 + sovGenesis(uint64(m.Index))
	}
	if m.MaxBlockUtilization != 0 {
		n += 1 + sovGenesis(uint64(m.MaxBlockUtilization))
	}
	if m.TargetBlockUtilization != 0 {
		n += 1 + sovGenesis(uint64(m.TargetBlockUtilization))
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
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
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
			if err := m.State.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *State) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: State: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: State: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBaseFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBaseFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LearningRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LearningRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Window = append(m.Window, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Window) == 0 {
					m.Window = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Window = append(m.Window, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Window", wireType)
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBlockUtilization", wireType)
			}
			m.MaxBlockUtilization = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBlockUtilization", wireType)
			}
			m.TargetBlockUtilization = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feemarket/feemarket/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// ParamsRequest is the request type for the Query/Params RPC method.
type ParamsRequest struct {
}

func (m *ParamsRequest) Reset()         { *m = ParamsRequest{} }
func (m *ParamsRequest) String() string { return proto.CompactTextString(m) }
func (*ParamsRequest) ProtoMessage()    {}
func (*ParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d683b3b0d8494138, []int{0}
}
func (m *ParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsRequest.Merge(m, src)
}
func (m *ParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsRequest proto.InternalMessageInfo

// ParamsResponse is the response type for the Query/Params RPC method.
type ParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *ParamsResponse) Reset()         { *m = ParamsResponse{} }
func (m *ParamsResponse) String() string { return proto.CompactTextString(m) }
func (*ParamsResponse) ProtoMessage()    {}
func (*ParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d683b3b0d8494138, []int{1}
}
func (m *ParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsResponse.Merge(m, src)
}
func (m *ParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsResponse proto.InternalMessageInfo

func (m *ParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// StateRequest is the request type for the Query/State RPC method.
type StateRequest struct {
	FeeDenom string `protobuf:"bytes,1,opt,name=fee_denom,json=feeDenom,proto3" json:"fee_denom,omitempty"`
}

func (m *StateRequest) Reset()         { *m = StateRequest{} }
func (m *StateRequest) String() string { return proto.CompactTextString(m) }
func (*StateRequest) ProtoMessage()    {}
func (*StateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d683b3b0d8494138, []int{2}
}
func (m *StateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateRequest.Merge(m, src)
}
func (m *StateRequest) XXX_Size() int {
	return m.Size()
}
func (m *StateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StateRequest proto.InternalMessageInfo

func (m *StateRequest) GetFeeDenom() string {
	if m != nil {
		return m.FeeDenom
	}
	return ""
}

// StateResponse is the response type for the Query/State RPC method.
type StateResponse struct {
	States []State `protobuf:"bytes,1,rep,name=states,proto3" json:"states"`
}

func (m *StateResponse) Reset()         { *m = StateResponse{} }
func (m *StateResponse) String() string { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()    {}
func (*StateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d683b3b0d8494138, []int{3}
}
func (m *StateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StateResponse.Merge(m, src)
}
func (m *StateResponse) XXX_Size() int {
	return m.Size()
}
func (m *StateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StateResponse proto.InternalMessageInfo

func (m *StateResponse) GetStates() []State {
	if m != nil {
		return m.States
	}
	return nil
}

// BaseFeeRequest is the request type for the Query/BaseFee RPC method.
type BaseFeeRequest struct {
	FeeDenom string `protobuf:"bytes,1,opt,name=fee_denom,json=feeDenom,proto3" json:"fee_denom,omitempty"`
}

func (m *BaseFeeRequest) Reset()         { *m = BaseFeeRequest{} }
func (m *BaseFeeRequest) String() string { return proto.CompactTextString(m) }
func (*BaseFeeRequest) ProtoMessage()    {}
func (*BaseFeeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d683b3b0d8494138, []int{4}
}
func (m *BaseFeeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseFeeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseFeeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseFeeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseFeeRequest.Merge(m, src)
}
func (m *BaseFeeRequest) XXX_Size() int {
	return m.Size()
}
func (m *BaseFeeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseFeeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseFeeRequest proto.InternalMessageInfo

func (m *BaseFeeRequest) GetFeeDenom() string {
	if m != nil {
		return m.FeeDenom
	}
	return ""
}

// StateResponse is the response type for the Query/BaseFee RPC method.
type BaseFeeResponse struct {
	Fee types.DecCoin `protobuf:"bytes,1,opt,name=fee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoin" json:"fee"`
}

func (m *BaseFeeResponse) Reset()         { *m = BaseFeeResponse{} }
func (m *BaseFeeResponse) String() string { return proto.CompactTextString(m) }
func (*BaseFeeResponse) ProtoMessage()    {}
func (*BaseFeeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d683b3b0d8494138, []int{5}
}
func (m *BaseFeeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseFeeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseFeeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseFeeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseFeeResponse.Merge(m, src)
}
func (m *BaseFeeResponse) XXX_Size() int {
	return m.Size()
}
func (m *BaseFeeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseFeeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseFeeResponse proto.InternalMessageInfo

func (m *BaseFeeResponse) GetFee() types.DecCoin {
	if m != nil {
		return m.Fee
	}
	return types.DecCoin{}
}

func init() {
	proto.RegisterType((*ParamsRequest)(nil), "feemarket.feemarket.v1.ParamsRequest")
	proto.RegisterType((*ParamsResponse)(nil), "feemarket.feemarket.v1.ParamsResponse")
	proto.RegisterType((*StateRequest)(nil), "feemarket.feemarket.v1.StateRequest")
	proto.RegisterType((*StateResponse)(nil), "feemarket.feemarket.v1.StateResponse")
	proto.RegisterType((*BaseFeeRequest)(nil), "feemarket.feemarket.v1.BaseFeeRequest")
	proto.RegisterType((*BaseFeeResponse)(nil), "feemarket.feemarket.v1.BaseFeeResponse")
}

func init() {
	proto.RegisterFile("feemarket/feemarket/v1/query.proto", fileDescriptor_d683b3b0d8494138)
}

var fileDescriptor_d683b3b0d8494138 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xce, 0x1a, 0x1b, 0xed, 0xd4, 0x36, 0x38, 0x94, 0x22, 0x31, 0x6e, 0x74, 0x4d, 0x6d, 0x89,
	0x64, 0x87, 0xc4, 0xa3, 0x9e, 0x62, 0x11, 0x44, 0x11, 0x8d, 0x37, 0x2f, 0x61, 0xb2, 0x79, 0xbb,
	0x2e, 0xe9, 0xce, 0x6c, 0x33, 0x93, 0x60, 0x28, 0xbd, 0x78, 0xe8, 0x51, 0x04, 0x6f, 0xfe, 0x02,
	0xf1, 0xd4, 0x9f, 0xd1, 0x63, 0xc1, 0x8b, 0x27, 0x95, 0x44, 0xe8, 0xdf, 0x90, 0x99, 0x9d, 0x4d,
	0x13, 0x68, 0x12, 0x2f, 0xbb, 0x6f, 0xe7, 0x7d, 0xdf, 0xfb, 0xde, 0x7b, 0xdf, 0x2c, 0x72, 0x7c,
	0x80, 0x88, 0xf6, 0xba, 0x20, 0xc9, 0x45, 0x34, 0xa8, 0x91, 0x83, 0x3e, 0xf4, 0x86, 0x6e, 0xdc,
	0xe3, 0x92, 0xe3, 0xad, 0x49, 0xc6, 0xbd, 0x88, 0x06, 0xb5, 0xc2, 0x66, 0xc0, 0x03, 0xae, 0x21,
	0x44, 0x45, 0x09, 0xba, 0x50, 0x0c, 0x38, 0x0f, 0xf6, 0x81, 0xd0, 0x38, 0x24, 0x94, 0x31, 0x2e,
	0xa9, 0x0c, 0x39, 0x13, 0x26, 0x6b, 0x7b, 0x5c, 0x44, 0x5c, 0x90, 0x36, 0x15, 0x40, 0x06, 0xb5,
	0x36, 0x48, 0x5a, 0x23, 0x1e, 0x0f, 0x99, 0xc9, 0xdf, 0xa4, 0x51, 0xc8, 0x38, 0xd1, 0x4f, 0x73,
	0x74, 0x7f, 0x4e, 0x8b, 0x31, 0xed, 0xd1, 0x28, 0xad, 0x5b, 0x9e, 0x03, 0x0a, 0x80, 0x81, 0x08,
	0x0d, 0xca, 0xc9, 0xa3, 0xf5, 0xd7, 0x9a, 0xd5, 0x84, 0x83, 0x3e, 0x08, 0xe9, 0xbc, 0x42, 0x1b,
	0xe9, 0x81, 0x88, 0x39, 0x13, 0x80, 0x9f, 0xa0, 0x5c, 0x52, 0xf8, 0x96, 0x75, 0xd7, 0xda, 0x5d,
	0xab, 0xdb, 0xee, 0xe5, 0xd3, 0xbb, 0x09, 0xaf, 0x71, 0xf5, 0xf4, 0x57, 0x29, 0xd3, 0x34, 0x1c,
	0xe7, 0x21, 0xba, 0xf1, 0x56, 0x52, 0x09, 0xa6, 0x3e, 0xbe, 0x8d, 0x56, 0x7d, 0x80, 0x56, 0x07,
	0x18, 0x8f, 0x74, 0xc1, 0xd5, 0xe6, 0x75, 0x1f, 0x60, 0x4f, 0x7d, 0x3b, 0x2f, 0xd1, 0xba, 0x01,
	0x1b, 0xed, 0xc7, 0x28, 0x27, 0xd4, 0x81, 0xd2, 0xce, 0xee, 0xae, 0xd5, 0xef, 0xcc, 0xd3, 0xd6,
	0xb4, 0x54, 0x3a, 0xa1, 0x38, 0x55, 0xb4, 0xd1, 0xa0, 0x02, 0x9e, 0xc1, 0xff, 0x89, 0x7f, 0xb2,
	0x50, 0x7e, 0x82, 0x37, 0xfa, 0x87, 0x28, 0xeb, 0x03, 0x98, 0xc1, 0x8b, 0x6e, 0x62, 0x95, 0xab,
	0xac, 0x72, 0x8d, 0x55, 0xee, 0x1e, 0x78, 0x4f, 0x79, 0xc8, 0x1a, 0x2f, 0x94, 0xf6, 0xf7, 0xdf,
	0xa5, 0x4a, 0x10, 0xca, 0xf7, 0xfd, 0xb6, 0xeb, 0xf1, 0x88, 0x18, 0x6b, 0x93, 0x57, 0x55, 0x74,
	0xba, 0x44, 0x0e, 0x63, 0x10, 0x29, 0xe7, 0xeb, 0xf9, 0x49, 0x25, 0xbf, 0x0f, 0x01, 0xf5, 0x86,
	0xad, 0x0e, 0x78, 0x2d, 0xe5, 0xfa, 0xb7, 0xf3, 0x93, 0x8a, 0xd5, 0x54, 0xaa, 0xf5, 0xe3, 0x2c,
	0x5a, 0x79, 0xa3, 0x6e, 0x1d, 0xee, 0xa3, 0x5c, 0xb2, 0x5c, 0xbc, 0xbd, 0x78, 0xf9, 0x66, 0xd0,
	0xc2, 0x83, 0x65, 0xb0, 0x64, 0x3e, 0xa7, 0xf8, 0xf1, 0xc7, 0xdf, 0x2f, 0x57, 0xb6, 0xf0, 0xe6,
	0x65, 0x17, 0x09, 0x1f, 0xa1, 0x15, 0xbd, 0x57, 0x5c, 0x5e, 0xb8, 0xf6, 0x54, 0x74, 0x7b, 0x09,
	0xca, 0x68, 0xee, 0x68, 0xcd, 0x7b, 0xb8, 0x34, 0xab, 0xa9, 0x4d, 0x23, 0x87, 0x13, 0x7f, 0x8e,
	0xf0, 0xb1, 0x85, 0xae, 0x19, 0x43, 0xf0, 0xdc, 0x81, 0x66, 0x1d, 0x2e, 0xec, 0x2c, 0xc5, 0x99,
	0x2e, 0x2a, 0xba, 0x8b, 0x32, 0x76, 0x66, 0xbb, 0x50, 0xd6, 0xb6, 0x7c, 0x98, 0x69, 0xa4, 0xf1,
	0xfc, 0x74, 0x64, 0x5b, 0x67, 0x23, 0xdb, 0xfa, 0x33, 0xb2, 0xad, 0xcf, 0x63, 0x3b, 0x73, 0x36,
	0xb6, 0x33, 0x3f, 0xc7, 0x76, 0xe6, 0x1d, 0x99, 0x32, 0x5b, 0x74, 0xc3, 0xb8, 0x1a, 0xc1, 0x60,
	0xaa, 0xe0, 0x87, 0xa9, 0x58, 0x3b, 0xdf, 0xce, 0xe9, 0xdf, 0xee, 0xd1, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x72, 0x07, 0xd7, 0x4b, 0x66, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the current feemarket module parameters.
	Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error)
	// State returns the current feemarket module state.
	State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error)
	// BaseFee returns the current feemarket module base fee.
	BaseFee(ctx context.Context, in *BaseFeeRequest, opts ...grpc.CallOption) (*BaseFeeResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *ParamsRequest, opts ...grpc.CallOption) (*ParamsResponse, error) {
	out := new(ParamsResponse)
	err := c.cc.Invoke(ctx, "/feemarket.feemarket.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) State(ctx context.Context, in *StateRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := c.cc.Invoke(ctx, "/feemarket.feemarket.v1.Query/State", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BaseFee(ctx context.Context, in *BaseFeeRequest, opts ...grpc.CallOption) (*BaseFeeResponse, error) {
	out := new(BaseFeeResponse)
	err := c.cc.Invoke(ctx, "/feemarket.feemarket.v1.Query/BaseFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the current feemarket module parameters.
	Params(context.Context, *ParamsRequest) (*ParamsResponse, error)
	// State returns the current feemarket module state.
	State(context.Context, *StateRequest) (*StateResponse, error)
	// BaseFee returns the current feemarket module base fee.
	BaseFee(context.Context, *BaseFeeRequest) (*BaseFeeResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *ParamsRequest) (*ParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) State(ctx context.Context, req *StateRequest) (*StateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method State not implemented")
}
func (*UnimplementedQueryServer) BaseFee(ctx context.Context, req *BaseFeeRequest) (*BaseFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BaseFee not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feemarket.feemarket.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*ParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_State_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).State(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feemarket.feemarket.v1.Query/State",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).State(ctx, req.(*StateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BaseFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BaseFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feemarket.feemarket.v1.Query/BaseFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BaseFee(ctx, req.(*BaseFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feemarket.feemarket.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "State",
			Handler:    _Query_State_Handler,
		},
		{
			MethodName: "BaseFee",
			Handler:    _Query_BaseFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "feemarket/feemarket/v1/query.proto",
}

func (m *ParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *StateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeDenom) > 0 {
		i -= len(m.FeeDenom)
		copy(dAtA[i:], m.FeeDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.FeeDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *StateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.States) > 0 {
		for iNdEx := len(m.States) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.States[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BaseFeeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseFeeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseFeeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeDenom) > 0 {
		i -= len(m.FeeDenom)
		copy(dAtA[i:], m.FeeDenom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.FeeDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BaseFeeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseFeeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseFeeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Fee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *StateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeeDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *StateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.States) > 0 {
		for _, e := range m.States {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *BaseFeeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeeDenom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *BaseFeeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Fee.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *ParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: ParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *StateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: StateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *StateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: StateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field States", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.States = append(m.States, State{})
			if err := m.States[len(m.States)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *BaseFeeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: BaseFeeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseFeeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *BaseFeeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: BaseFeeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseFeeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

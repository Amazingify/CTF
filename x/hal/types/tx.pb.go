// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gaia/hal/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgMintHAL defines a SDK message for the Msg/MintHAL request type.
type MsgMintHAL struct {
	// address is the Bech32-encoded address of the target account.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// collateral_amount are collateral token that should be exchanged to HAL.
	CollateralAmount []types.Coin `protobuf:"bytes,2,rep,name=collateral_amount,json=collateralAmount,proto3" json:"collateral_amount" yaml:"collateral_amount"`
}

func (m *MsgMintHAL) Reset()         { *m = MsgMintHAL{} }
func (m *MsgMintHAL) String() string { return proto.CompactTextString(m) }
func (*MsgMintHAL) ProtoMessage()    {}
func (*MsgMintHAL) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa557fc55cc32b2, []int{0}
}
func (m *MsgMintHAL) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintHAL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintHAL.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintHAL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintHAL.Merge(m, src)
}
func (m *MsgMintHAL) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintHAL) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintHAL.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintHAL proto.InternalMessageInfo

// MsgMintHALResponse defines the Msg/MintHAL response type.
type MsgMintHALResponse struct {
	// minted_amount is a minted HAL token.
	MintedAmount types.Coin `protobuf:"bytes,1,opt,name=minted_amount,json=mintedAmount,proto3" json:"minted_amount" yaml:"minted_amount"`
	// collaterals_amount are collateral tokens used for the mint.
	CollateralsAmount []types.Coin `protobuf:"bytes,2,rep,name=collaterals_amount,json=collateralsAmount,proto3" json:"collaterals_amount" yaml:"collaterals_amount"`
}

func (m *MsgMintHALResponse) Reset()         { *m = MsgMintHALResponse{} }
func (m *MsgMintHALResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintHALResponse) ProtoMessage()    {}
func (*MsgMintHALResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa557fc55cc32b2, []int{1}
}
func (m *MsgMintHALResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintHALResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintHALResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintHALResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintHALResponse.Merge(m, src)
}
func (m *MsgMintHALResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintHALResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintHALResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintHALResponse proto.InternalMessageInfo

// MsgRedeemCollateral defines a SDK message for the Msg/RedeemCollateral request type.
type MsgRedeemCollateral struct {
	// address is the Bech32-encoded address of the target account.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// hal_amount is the HAL token that should be exchanged to collateral tokens.
	HalAmount types.Coin `protobuf:"bytes,2,opt,name=hal_amount,json=halAmount,proto3" json:"hal_amount" yaml:"amount"`
}

func (m *MsgRedeemCollateral) Reset()         { *m = MsgRedeemCollateral{} }
func (m *MsgRedeemCollateral) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemCollateral) ProtoMessage()    {}
func (*MsgRedeemCollateral) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa557fc55cc32b2, []int{2}
}
func (m *MsgRedeemCollateral) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemCollateral) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemCollateral.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemCollateral) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemCollateral.Merge(m, src)
}
func (m *MsgRedeemCollateral) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemCollateral) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemCollateral.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemCollateral proto.InternalMessageInfo

// MsgMintHALResponse defines the Msg/RedeemCollateral response type.
type MsgRedeemCollateralResponse struct {
	// burned_amount is the HAL token converted amount (might be LT the requested amount).
	BurnedAmount types.Coin `protobuf:"bytes,1,opt,name=burned_amount,json=burnedAmount,proto3" json:"burned_amount" yaml:"burned_amount"`
	// redeemed_amount are collateral tokens that are transferred to an account after the redeeming timout.
	RedeemedAmount []types.Coin `protobuf:"bytes,2,rep,name=redeemed_amount,json=redeemedAmount,proto3" json:"redeemed_amount" yaml:"redeemed_amount"`
	// completion_time defines the redeeming period end time.
	CompletionTime time.Time `protobuf:"bytes,3,opt,name=completion_time,json=completionTime,proto3,stdtime" json:"completion_time" yaml:"completion_time"`
}

func (m *MsgRedeemCollateralResponse) Reset()         { *m = MsgRedeemCollateralResponse{} }
func (m *MsgRedeemCollateralResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemCollateralResponse) ProtoMessage()    {}
func (*MsgRedeemCollateralResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_caa557fc55cc32b2, []int{3}
}
func (m *MsgRedeemCollateralResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemCollateralResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemCollateralResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemCollateralResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemCollateralResponse.Merge(m, src)
}
func (m *MsgRedeemCollateralResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemCollateralResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemCollateralResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemCollateralResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgMintHAL)(nil), "gaia.hal.v1beta1.MsgMintHAL")
	proto.RegisterType((*MsgMintHALResponse)(nil), "gaia.hal.v1beta1.MsgMintHALResponse")
	proto.RegisterType((*MsgRedeemCollateral)(nil), "gaia.hal.v1beta1.MsgRedeemCollateral")
	proto.RegisterType((*MsgRedeemCollateralResponse)(nil), "gaia.hal.v1beta1.MsgRedeemCollateralResponse")
}

func init() { proto.RegisterFile("gaia/hal/v1beta1/tx.proto", fileDescriptor_caa557fc55cc32b2) }

var fileDescriptor_caa557fc55cc32b2 = []byte{
	// 556 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0xed, 0x46, 0xfa, 0x55, 0xdd, 0x5f, 0x9b, 0x06, 0x53, 0x50, 0x12, 0xaa, 0x4d, 0xb0,
	0x40, 0xaa, 0x10, 0x5d, 0xab, 0xe5, 0xc6, 0x8d, 0x84, 0x23, 0x11, 0x92, 0x4b, 0x2f, 0x08, 0xa9,
	0x5a, 0x3b, 0xcb, 0xc6, 0xc2, 0xf6, 0x5a, 0x9e, 0x75, 0xd5, 0xbe, 0x45, 0x0f, 0x3c, 0x02, 0x07,
	0x9e, 0x82, 0x73, 0x8e, 0x3d, 0x72, 0x4a, 0x21, 0x79, 0x83, 0x3e, 0x01, 0xb2, 0x77, 0x1d, 0x37,
	0x69, 0xd5, 0x16, 0x6e, 0xde, 0x9d, 0x3f, 0x9f, 0xef, 0xcc, 0x78, 0x16, 0xb5, 0x38, 0x0d, 0xa8,
	0x93, 0x81, 0xef, 0x1c, 0xef, 0x79, 0x4c, 0xd2, 0x3d, 0x47, 0x9e, 0x90, 0x24, 0x15, 0x52, 0x58,
	0x8d, 0xdc, 0x44, 0x32, 0xf0, 0x89, 0x36, 0xb5, 0xb7, 0xb8, 0xe0, 0xa2, 0x30, 0x3a, 0xf9, 0x97,
	0xf2, 0x6b, 0x77, 0xb8, 0x10, 0x3c, 0x64, 0x4e, 0x71, 0xf2, 0xb2, 0xcf, 0x8e, 0x0c, 0x22, 0x06,
	0x92, 0x46, 0x89, 0x76, 0xc0, 0xbe, 0x80, 0x48, 0x80, 0xe3, 0x51, 0x60, 0x73, 0x8c, 0x2f, 0x82,
	0x58, 0xd9, 0xed, 0x6f, 0x26, 0x42, 0x03, 0xe0, 0x83, 0x20, 0x96, 0x87, 0x07, 0x7d, 0xeb, 0x25,
	0x5a, 0xa5, 0xc3, 0x61, 0xca, 0x00, 0x9a, 0x66, 0xd7, 0xdc, 0x59, 0xeb, 0x59, 0x97, 0x93, 0x4e,
	0xfd, 0x94, 0x46, 0xe1, 0x6b, 0x5b, 0x1b, 0x6c, 0xb7, 0x74, 0xb1, 0x46, 0xe8, 0x81, 0x2f, 0xc2,
	0x90, 0x4a, 0x96, 0xd2, 0xf0, 0x88, 0x46, 0x22, 0x8b, 0x65, 0x73, 0xa5, 0x5b, 0xdb, 0xf9, 0x7f,
	0xbf, 0x45, 0x14, 0x98, 0xe4, 0xe0, 0xb2, 0x08, 0xd2, 0x17, 0x41, 0xdc, 0xeb, 0x8e, 0x27, 0x1d,
	0xe3, 0x72, 0xd2, 0x69, 0xaa, 0xb4, 0xd7, 0x32, 0xd8, 0x6e, 0xa3, 0xba, 0x7b, 0xa3, 0xae, 0x2e,
	0x4c, 0x64, 0x55, 0x32, 0x5d, 0x06, 0x89, 0x88, 0x81, 0x59, 0x9f, 0xd0, 0x46, 0x14, 0xc4, 0x92,
	0x0d, 0x4b, 0x78, 0x2e, 0xfa, 0x56, 0xf8, 0xb6, 0x86, 0x6f, 0x29, 0xf8, 0x42, 0xb4, 0xed, 0xae,
	0xab, 0xb3, 0x82, 0x5a, 0x5f, 0x90, 0x55, 0x09, 0x81, 0x7b, 0xd7, 0xf7, 0x54, 0x23, 0x5a, 0xcb,
	0xf5, 0xc1, 0x9c, 0x73, 0xa5, 0x6d, 0xa0, 0x2b, 0xfc, 0x6a, 0xa2, 0x87, 0x03, 0xe0, 0x2e, 0x1b,
	0x32, 0x16, 0xf5, 0xe7, 0xe6, 0xbf, 0x9c, 0xc8, 0x7b, 0x84, 0x32, 0xf0, 0x2b, 0xa9, 0x77, 0x74,
	0xe3, 0x91, 0x96, 0xba, 0xa1, 0xf3, 0x69, 0x79, 0x6b, 0x19, 0xf8, 0x5a, 0xd6, 0x78, 0x05, 0x3d,
	0xb9, 0x41, 0xd6, 0xd5, 0x09, 0x78, 0x59, 0x1a, 0xff, 0xfb, 0x04, 0x16, 0xa2, 0x6d, 0x77, 0x5d,
	0x9d, 0xf5, 0x04, 0x3c, 0xb4, 0x99, 0x16, 0xe4, 0x2a, 0xff, 0x9d, 0xed, 0xc7, 0x3a, 0xff, 0x63,
	0x95, 0x7f, 0x29, 0xde, 0x76, 0xeb, 0xe5, 0x8d, 0x66, 0x70, 0xb4, 0xe9, 0x8b, 0x28, 0x09, 0x99,
	0x0c, 0x44, 0x7c, 0x94, 0xef, 0x4f, 0xb3, 0x56, 0xd4, 0xd0, 0x26, 0x6a, 0xb9, 0x48, 0xb9, 0x5c,
	0xe4, 0x43, 0xb9, 0x5c, 0x3d, 0x7b, 0x11, 0xb2, 0x94, 0xc0, 0x3e, 0xbb, 0xe8, 0x98, 0x6e, 0xbd,
	0xba, 0xcd, 0x03, 0xf7, 0x7f, 0x98, 0xa8, 0x36, 0x00, 0x6e, 0x0d, 0xd0, 0x6a, 0xb9, 0x6e, 0xdb,
	0x64, 0x79, 0xcf, 0x49, 0xf5, 0x97, 0xb7, 0x9f, 0xdd, 0x66, 0x9d, 0x4f, 0x60, 0x84, 0x1a, 0xd7,
	0x7e, 0x9a, 0xe7, 0x37, 0x46, 0x2e, 0xbb, 0xb5, 0x77, 0xef, 0xe5, 0x56, 0x92, 0x7a, 0xef, 0xc6,
	0xbf, 0xb1, 0xf1, 0x7d, 0x8a, 0x8d, 0xf1, 0x14, 0x9b, 0xe7, 0x53, 0x6c, 0xfe, 0x9a, 0x62, 0xf3,
	0x6c, 0x86, 0x8d, 0xf3, 0x19, 0x36, 0x7e, 0xce, 0xb0, 0xf1, 0xf1, 0x05, 0x0f, 0xe4, 0x28, 0xf3,
	0x88, 0x2f, 0x22, 0xe7, 0xf0, 0xa0, 0xbf, 0xfb, 0x96, 0x1d, 0xb3, 0x50, 0x24, 0x2c, 0x05, 0xa7,
	0x78, 0xeb, 0x4e, 0x8a, 0xd7, 0x4e, 0x9e, 0x26, 0x0c, 0xbc, 0xff, 0x8a, 0xb6, 0xbe, 0xfa, 0x13,
	0x00, 0x00, 0xff, 0xff, 0x10, 0xf6, 0x0c, 0x8f, 0x06, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// MintHAL defines a method for converting collateral coins into HAL coin.
	MintHAL(ctx context.Context, in *MsgMintHAL, opts ...grpc.CallOption) (*MsgMintHALResponse, error)
	// RedeemCollateral defines a method for converting HAL coin into collateral coins from the module pool.
	RedeemCollateral(ctx context.Context, in *MsgRedeemCollateral, opts ...grpc.CallOption) (*MsgRedeemCollateralResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MintHAL(ctx context.Context, in *MsgMintHAL, opts ...grpc.CallOption) (*MsgMintHALResponse, error) {
	out := new(MsgMintHALResponse)
	err := c.cc.Invoke(ctx, "/gaia.hal.v1beta1.Msg/MintHAL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RedeemCollateral(ctx context.Context, in *MsgRedeemCollateral, opts ...grpc.CallOption) (*MsgRedeemCollateralResponse, error) {
	out := new(MsgRedeemCollateralResponse)
	err := c.cc.Invoke(ctx, "/gaia.hal.v1beta1.Msg/RedeemCollateral", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// MintHAL defines a method for converting collateral coins into HAL coin.
	MintHAL(context.Context, *MsgMintHAL) (*MsgMintHALResponse, error)
	// RedeemCollateral defines a method for converting HAL coin into collateral coins from the module pool.
	RedeemCollateral(context.Context, *MsgRedeemCollateral) (*MsgRedeemCollateralResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MintHAL(ctx context.Context, req *MsgMintHAL) (*MsgMintHALResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintHAL not implemented")
}
func (*UnimplementedMsgServer) RedeemCollateral(ctx context.Context, req *MsgRedeemCollateral) (*MsgRedeemCollateralResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RedeemCollateral not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MintHAL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintHAL)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintHAL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaia.hal.v1beta1.Msg/MintHAL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintHAL(ctx, req.(*MsgMintHAL))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RedeemCollateral_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRedeemCollateral)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RedeemCollateral(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gaia.hal.v1beta1.Msg/RedeemCollateral",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RedeemCollateral(ctx, req.(*MsgRedeemCollateral))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gaia.hal.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintHAL",
			Handler:    _Msg_MintHAL_Handler,
		},
		{
			MethodName: "RedeemCollateral",
			Handler:    _Msg_RedeemCollateral_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gaia/hal/v1beta1/tx.proto",
}

func (m *MsgMintHAL) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintHAL) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintHAL) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CollateralAmount) > 0 {
		for iNdEx := len(m.CollateralAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollateralAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintHALResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintHALResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintHALResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CollateralsAmount) > 0 {
		for iNdEx := len(m.CollateralsAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollateralsAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.MintedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgRedeemCollateral) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemCollateral) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemCollateral) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HalAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRedeemCollateralResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemCollateralResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemCollateralResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintTx(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	if len(m.RedeemedAmount) > 0 {
		for iNdEx := len(m.RedeemedAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RedeemedAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.BurnedAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMintHAL) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.CollateralAmount) > 0 {
		for _, e := range m.CollateralAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgMintHALResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MintedAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	if len(m.CollateralsAmount) > 0 {
		for _, e := range m.CollateralsAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgRedeemCollateral) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.HalAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgRedeemCollateralResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BurnedAmount.Size()
	n += 1 + l + sovTx(uint64(l))
	if len(m.RedeemedAmount) > 0 {
		for _, e := range m.RedeemedAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMintHAL) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintHAL: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintHAL: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralAmount = append(m.CollateralAmount, types.Coin{})
			if err := m.CollateralAmount[len(m.CollateralAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgMintHALResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintHALResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintHALResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MintedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralsAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralsAmount = append(m.CollateralsAmount, types.Coin{})
			if err := m.CollateralsAmount[len(m.CollateralsAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRedeemCollateral) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRedeemCollateral: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemCollateral: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HalAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HalAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRedeemCollateralResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgRedeemCollateralResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemCollateralResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedeemedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RedeemedAmount = append(m.RedeemedAmount, types.Coin{})
			if err := m.RedeemedAmount[len(m.RedeemedAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

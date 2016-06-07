// Code generated by protoc-gen-go.
// source: bfcta.proto
// DO NOT EDIT!

/*
Package bfcta is a generated protocol buffer package.

It is generated from these files:
	bfcta.proto

It has these top-level messages:
	BfModelData
	BfRobotData
	BfGatewayData
	BfOrderExData
*/
package bfcta

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import bfgateway "github.com/sunwangme/bfgo/api/bfgateway"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 模型实现采用的语言和发布方式
type BfModelLangType int32

const (
	BfModelLangType_MODELLANG_UNKNOWN      BfModelLangType = 0
	BfModelLangType_MODELLANG_PYTHONSCRIPT BfModelLangType = 1
	BfModelLangType_MODELLANG_GOLANGEXE    BfModelLangType = 2
	BfModelLangType_MODELLANG_CPPEXE       BfModelLangType = 3
	BfModelLangType_MODELLANG_CPPDLL       BfModelLangType = 4
)

var BfModelLangType_name = map[int32]string{
	0: "MODELLANG_UNKNOWN",
	1: "MODELLANG_PYTHONSCRIPT",
	2: "MODELLANG_GOLANGEXE",
	3: "MODELLANG_CPPEXE",
	4: "MODELLANG_CPPDLL",
}
var BfModelLangType_value = map[string]int32{
	"MODELLANG_UNKNOWN":      0,
	"MODELLANG_PYTHONSCRIPT": 1,
	"MODELLANG_GOLANGEXE":    2,
	"MODELLANG_CPPEXE":       3,
	"MODELLANG_CPPDLL":       4,
}

func (x BfModelLangType) String() string {
	return proto.EnumName(BfModelLangType_name, int32(x))
}
func (BfModelLangType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 模型信息类
type BfModelData struct {
	ModelId  string          `protobuf:"bytes,1,opt,name=modelId" json:"modelId,omitempty"`
	LangType BfModelLangType `protobuf:"varint,2,opt,name=langType,enum=bfcta.BfModelLangType" json:"langType,omitempty"`
	Path     string          `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
}

func (m *BfModelData) Reset()                    { *m = BfModelData{} }
func (m *BfModelData) String() string            { return proto.CompactTextString(m) }
func (*BfModelData) ProtoMessage()               {}
func (*BfModelData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 机器人信息类
type BfRobotData struct {
	RobotId   string `protobuf:"bytes,1,opt,name=robotId" json:"robotId,omitempty"`
	ModelId   string `protobuf:"bytes,2,opt,name=modelId" json:"modelId,omitempty"`
	GatewayId string `protobuf:"bytes,3,opt,name=gatewayId" json:"gatewayId,omitempty"`
	Symbol    string `protobuf:"bytes,4,opt,name=symbol" json:"symbol,omitempty"`
	Exchange  string `protobuf:"bytes,5,opt,name=exchange" json:"exchange,omitempty"`
	Status    string `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
}

func (m *BfRobotData) Reset()                    { *m = BfRobotData{} }
func (m *BfRobotData) String() string            { return proto.CompactTextString(m) }
func (*BfRobotData) ProtoMessage()               {}
func (*BfRobotData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 网关信息类
type BfGatewayData struct {
	GatewayId string `protobuf:"bytes,1,opt,name=gatewayId" json:"gatewayId,omitempty"`
	Ip        string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	Port      int32  `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
	Status    string `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
}

func (m *BfGatewayData) Reset()                    { *m = BfGatewayData{} }
func (m *BfGatewayData) String() string            { return proto.CompactTextString(m) }
func (*BfGatewayData) ProtoMessage()               {}
func (*BfGatewayData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 委托<-->机器人映射信息类
type BfOrderExData struct {
	BfOrderId string `protobuf:"bytes,1,opt,name=bfOrderId" json:"bfOrderId,omitempty"`
	RobotId   string `protobuf:"bytes,2,opt,name=robotId" json:"robotId,omitempty"`
}

func (m *BfOrderExData) Reset()                    { *m = BfOrderExData{} }
func (m *BfOrderExData) String() string            { return proto.CompactTextString(m) }
func (*BfOrderExData) ProtoMessage()               {}
func (*BfOrderExData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*BfModelData)(nil), "bfcta.BfModelData")
	proto.RegisterType((*BfRobotData)(nil), "bfcta.BfRobotData")
	proto.RegisterType((*BfGatewayData)(nil), "bfcta.BfGatewayData")
	proto.RegisterType((*BfOrderExData)(nil), "bfcta.BfOrderExData")
	proto.RegisterEnum("bfcta.BfModelLangType", BfModelLangType_name, BfModelLangType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for BfCtaService service

type BfCtaServiceClient interface {
	// 请求push
	Connect(ctx context.Context, in *bfgateway.BfConnectReq, opts ...grpc.CallOption) (BfCtaService_ConnectClient, error)
	// 断开push
	Disconnect(ctx context.Context, in *bfgateway.BfVoid, opts ...grpc.CallOption) (*bfgateway.BfVoid, error)
	// 活跃检测
	Ping(ctx context.Context, in *bfgateway.BfPingData, opts ...grpc.CallOption) (*bfgateway.BfPingData, error)
	// 获取策略信息，如position，用kv最方便
	GetRobotInfo(ctx context.Context, in *bfgateway.BfKvData, opts ...grpc.CallOption) (*bfgateway.BfKvData, error)
	// 发单
	SendOrder(ctx context.Context, in *bfgateway.BfSendOrderReq, opts ...grpc.CallOption) (*bfgateway.BfSendOrderResp, error)
	// 撤单
	CancelOrder(ctx context.Context, in *bfgateway.BfCancelOrderReq, opts ...grpc.CallOption) (*bfgateway.BfVoid, error)
}

type bfCtaServiceClient struct {
	cc *grpc.ClientConn
}

func NewBfCtaServiceClient(cc *grpc.ClientConn) BfCtaServiceClient {
	return &bfCtaServiceClient{cc}
}

func (c *bfCtaServiceClient) Connect(ctx context.Context, in *bfgateway.BfConnectReq, opts ...grpc.CallOption) (BfCtaService_ConnectClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BfCtaService_serviceDesc.Streams[0], c.cc, "/bfcta.BfCtaService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &bfCtaServiceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BfCtaService_ConnectClient interface {
	Recv() (*google_protobuf.Any, error)
	grpc.ClientStream
}

type bfCtaServiceConnectClient struct {
	grpc.ClientStream
}

func (x *bfCtaServiceConnectClient) Recv() (*google_protobuf.Any, error) {
	m := new(google_protobuf.Any)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bfCtaServiceClient) Disconnect(ctx context.Context, in *bfgateway.BfVoid, opts ...grpc.CallOption) (*bfgateway.BfVoid, error) {
	out := new(bfgateway.BfVoid)
	err := grpc.Invoke(ctx, "/bfcta.BfCtaService/Disconnect", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) Ping(ctx context.Context, in *bfgateway.BfPingData, opts ...grpc.CallOption) (*bfgateway.BfPingData, error) {
	out := new(bfgateway.BfPingData)
	err := grpc.Invoke(ctx, "/bfcta.BfCtaService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) GetRobotInfo(ctx context.Context, in *bfgateway.BfKvData, opts ...grpc.CallOption) (*bfgateway.BfKvData, error) {
	out := new(bfgateway.BfKvData)
	err := grpc.Invoke(ctx, "/bfcta.BfCtaService/GetRobotInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) SendOrder(ctx context.Context, in *bfgateway.BfSendOrderReq, opts ...grpc.CallOption) (*bfgateway.BfSendOrderResp, error) {
	out := new(bfgateway.BfSendOrderResp)
	err := grpc.Invoke(ctx, "/bfcta.BfCtaService/SendOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bfCtaServiceClient) CancelOrder(ctx context.Context, in *bfgateway.BfCancelOrderReq, opts ...grpc.CallOption) (*bfgateway.BfVoid, error) {
	out := new(bfgateway.BfVoid)
	err := grpc.Invoke(ctx, "/bfcta.BfCtaService/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BfCtaService service

type BfCtaServiceServer interface {
	// 请求push
	Connect(*bfgateway.BfConnectReq, BfCtaService_ConnectServer) error
	// 断开push
	Disconnect(context.Context, *bfgateway.BfVoid) (*bfgateway.BfVoid, error)
	// 活跃检测
	Ping(context.Context, *bfgateway.BfPingData) (*bfgateway.BfPingData, error)
	// 获取策略信息，如position，用kv最方便
	GetRobotInfo(context.Context, *bfgateway.BfKvData) (*bfgateway.BfKvData, error)
	// 发单
	SendOrder(context.Context, *bfgateway.BfSendOrderReq) (*bfgateway.BfSendOrderResp, error)
	// 撤单
	CancelOrder(context.Context, *bfgateway.BfCancelOrderReq) (*bfgateway.BfVoid, error)
}

func RegisterBfCtaServiceServer(s *grpc.Server, srv BfCtaServiceServer) {
	s.RegisterService(&_BfCtaService_serviceDesc, srv)
}

func _BfCtaService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(bfgateway.BfConnectReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BfCtaServiceServer).Connect(m, &bfCtaServiceConnectServer{stream})
}

type BfCtaService_ConnectServer interface {
	Send(*google_protobuf.Any) error
	grpc.ServerStream
}

type bfCtaServiceConnectServer struct {
	grpc.ServerStream
}

func (x *bfCtaServiceConnectServer) Send(m *google_protobuf.Any) error {
	return x.ServerStream.SendMsg(m)
}

func _BfCtaService_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bfgateway.BfVoid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BfCtaServiceServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bfcta.BfCtaService/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BfCtaServiceServer).Disconnect(ctx, req.(*bfgateway.BfVoid))
	}
	return interceptor(ctx, in, info, handler)
}

func _BfCtaService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bfgateway.BfPingData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BfCtaServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bfcta.BfCtaService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BfCtaServiceServer).Ping(ctx, req.(*bfgateway.BfPingData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BfCtaService_GetRobotInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bfgateway.BfKvData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BfCtaServiceServer).GetRobotInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bfcta.BfCtaService/GetRobotInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BfCtaServiceServer).GetRobotInfo(ctx, req.(*bfgateway.BfKvData))
	}
	return interceptor(ctx, in, info, handler)
}

func _BfCtaService_SendOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bfgateway.BfSendOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BfCtaServiceServer).SendOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bfcta.BfCtaService/SendOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BfCtaServiceServer).SendOrder(ctx, req.(*bfgateway.BfSendOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BfCtaService_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(bfgateway.BfCancelOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BfCtaServiceServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bfcta.BfCtaService/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BfCtaServiceServer).CancelOrder(ctx, req.(*bfgateway.BfCancelOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BfCtaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bfcta.BfCtaService",
	HandlerType: (*BfCtaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Disconnect",
			Handler:    _BfCtaService_Disconnect_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _BfCtaService_Ping_Handler,
		},
		{
			MethodName: "GetRobotInfo",
			Handler:    _BfCtaService_GetRobotInfo_Handler,
		},
		{
			MethodName: "SendOrder",
			Handler:    _BfCtaService_SendOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _BfCtaService_CancelOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _BfCtaService_Connect_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor0 = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x53, 0xdd, 0x6e, 0xda, 0x4c,
	0x10, 0xc5, 0x0e, 0x90, 0x30, 0xe4, 0x4b, 0xc8, 0x26, 0x21, 0x8e, 0xbf, 0x5e, 0x54, 0x54, 0x95,
	0xaa, 0x5e, 0xd8, 0x15, 0xad, 0x7a, 0x51, 0x55, 0xaa, 0xc2, 0x8f, 0x28, 0x0a, 0x01, 0x64, 0xe8,
	0xef, 0x4d, 0xb5, 0x36, 0x6b, 0x63, 0x09, 0xbc, 0xae, 0x59, 0x92, 0xf0, 0x08, 0x7d, 0x96, 0x3e,
	0x63, 0xa5, 0xee, 0xae, 0x0d, 0xc6, 0x09, 0xbd, 0xf2, 0x9e, 0x73, 0x66, 0xe6, 0xcc, 0xce, 0x8e,
	0xa1, 0x6c, 0xbb, 0x0e, 0xc3, 0x46, 0x18, 0x51, 0x46, 0x51, 0x41, 0x02, 0xfd, 0xd8, 0x76, 0x3d,
	0xcc, 0xc8, 0x1d, 0x5e, 0xc5, 0xbc, 0x7e, 0xe9, 0x51, 0xea, 0xcd, 0x88, 0x29, 0x91, 0xbd, 0x74,
	0x4d, 0x1c, 0x24, 0x52, 0x8d, 0x42, 0xb9, 0xe1, 0xde, 0xd0, 0x09, 0x99, 0xb5, 0x30, 0xc3, 0x48,
	0x83, 0xfd, 0xb9, 0x00, 0xdd, 0x89, 0xa6, 0x3c, 0x55, 0x5e, 0x94, 0xac, 0x35, 0x44, 0x75, 0x38,
	0x98, 0xe1, 0xc0, 0x1b, 0xaf, 0x42, 0xa2, 0xa9, 0x5c, 0x3a, 0xaa, 0x57, 0x8d, 0xd8, 0x3b, 0xc9,
	0xef, 0x25, 0xaa, 0xb5, 0x89, 0x43, 0x08, 0xf2, 0x21, 0x66, 0x53, 0x6d, 0x4f, 0x96, 0x92, 0xe7,
	0xda, 0x6f, 0x45, 0x38, 0x5a, 0xd4, 0xa6, 0x6c, 0xed, 0x18, 0x09, 0x90, 0x3a, 0x26, 0x70, 0xbb,
	0x17, 0x35, 0xdb, 0xcb, 0x13, 0x28, 0x25, 0x17, 0xe4, 0x5a, 0x5c, 0x3c, 0x25, 0x50, 0x15, 0x8a,
	0x8b, 0xd5, 0xdc, 0xa6, 0x33, 0x2d, 0x2f, 0xa5, 0x04, 0x21, 0x1d, 0x0e, 0xc8, 0xbd, 0x33, 0xe5,
	0xcd, 0x11, 0xad, 0x20, 0x95, 0x0d, 0x96, 0x39, 0x0c, 0xb3, 0xe5, 0x42, 0x2b, 0x26, 0x39, 0x12,
	0xd5, 0x7c, 0xf8, 0xaf, 0xe1, 0x76, 0xe2, 0xd2, 0xb2, 0xdd, 0x8c, 0xb5, 0xf2, 0xd0, 0xfa, 0x08,
	0x54, 0x3f, 0x4c, 0xba, 0xe5, 0x27, 0x39, 0x00, 0x1a, 0x31, 0xd9, 0x63, 0xc1, 0x92, 0xe7, 0x2d,
	0xab, 0x7c, 0xc6, 0xaa, 0x23, 0xac, 0x06, 0xd1, 0x84, 0x44, 0xed, 0xfb, 0xb5, 0x95, 0x1d, 0x13,
	0xa9, 0xd5, 0x86, 0xd8, 0x9e, 0x9b, 0x9a, 0x99, 0xdb, 0xcb, 0x5f, 0x0a, 0x1c, 0x3f, 0x78, 0x13,
	0x74, 0x0e, 0x27, 0x37, 0x83, 0x56, 0xbb, 0xd7, 0xbb, 0xea, 0x77, 0x7e, 0x7c, 0xea, 0x5f, 0xf7,
	0x07, 0x5f, 0xfa, 0x95, 0x1c, 0x1f, 0x49, 0x35, 0xa5, 0x87, 0xdf, 0xc6, 0x1f, 0x07, 0xfd, 0x51,
	0xd3, 0xea, 0x0e, 0xc7, 0x15, 0x05, 0x5d, 0xc0, 0x69, 0xaa, 0x75, 0x06, 0xe2, 0xd3, 0xfe, 0xda,
	0xae, 0xa8, 0xe8, 0x0c, 0x2a, 0xa9, 0xd0, 0x1c, 0x0e, 0x05, 0xbb, 0xf7, 0x88, 0x6d, 0xf5, 0x7a,
	0x95, 0x7c, 0xfd, 0x8f, 0x0a, 0x87, 0x0d, 0xb7, 0xc9, 0xf0, 0x88, 0x44, 0xb7, 0xbe, 0x43, 0xd0,
	0x7b, 0xd8, 0x6f, 0xd2, 0x20, 0x20, 0x0e, 0x43, 0x17, 0x46, 0xba, 0xa7, 0x3c, 0x26, 0x66, 0x2d,
	0xf2, 0x53, 0x3f, 0x33, 0xe2, 0x7d, 0x35, 0xd6, 0xfb, 0x6a, 0x5c, 0x05, 0xab, 0x5a, 0xee, 0x95,
	0x82, 0xde, 0x00, 0xb4, 0xfc, 0x85, 0x93, 0x14, 0x38, 0xc9, 0x14, 0xf8, 0x4c, 0xfd, 0x89, 0xfe,
	0x98, 0xaa, 0xe5, 0xd0, 0x5b, 0xc8, 0x0f, 0xfd, 0xc0, 0x43, 0xe7, 0x19, 0x51, 0x50, 0x62, 0xce,
	0xfa, 0x6e, 0x9a, 0xe7, 0xbd, 0x83, 0xc3, 0x0e, 0x61, 0x72, 0x55, 0xbb, 0x81, 0x4b, 0xd1, 0x69,
	0x26, 0xf0, 0xfa, 0x56, 0x66, 0xef, 0x22, 0x79, 0x6e, 0x0b, 0x4a, 0x23, 0x12, 0x4c, 0xe4, 0x6b,
	0xa1, 0xcb, 0x4c, 0xcc, 0x86, 0x17, 0x77, 0xd5, 0xff, 0x25, 0x2d, 0x42, 0x5e, 0xe5, 0x03, 0x94,
	0x9b, 0x38, 0x70, 0xc8, 0x2c, 0xae, 0xf3, 0x7f, 0x76, 0x62, 0xa9, 0x22, 0x2a, 0xed, 0xba, 0x7a,
	0xe3, 0xf9, 0xf7, 0x67, 0x9e, 0xcf, 0xa6, 0x4b, 0xdb, 0x70, 0xe8, 0xdc, 0x5c, 0x2c, 0x83, 0x3b,
	0xbe, 0x11, 0x73, 0x62, 0xf2, 0x50, 0x6a, 0xe2, 0xd0, 0x37, 0xe5, 0x2f, 0x6c, 0x17, 0xe5, 0xa4,
	0x5f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xbf, 0xd5, 0x18, 0x4e, 0x04, 0x00, 0x00,
}
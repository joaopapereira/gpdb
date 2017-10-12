// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command.proto

/*
Package idl is a generated protocol buffer package.

It is generated from these files:
	command.proto

It has these top-level messages:
	TransmitStateRequest
	TransmitStateReply
	CheckUpgradeStatusRequest
	CheckUpgradeStatusReply
*/
package idl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type TransmitStateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TransmitStateRequest) Reset()                    { *m = TransmitStateRequest{} }
func (m *TransmitStateRequest) String() string            { return proto.CompactTextString(m) }
func (*TransmitStateRequest) ProtoMessage()               {}
func (*TransmitStateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TransmitStateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TransmitStateReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *TransmitStateReply) Reset()                    { *m = TransmitStateReply{} }
func (m *TransmitStateReply) String() string            { return proto.CompactTextString(m) }
func (*TransmitStateReply) ProtoMessage()               {}
func (*TransmitStateReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TransmitStateReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CheckUpgradeStatusRequest struct {
}

func (m *CheckUpgradeStatusRequest) Reset()                    { *m = CheckUpgradeStatusRequest{} }
func (m *CheckUpgradeStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusRequest) ProtoMessage()               {}
func (*CheckUpgradeStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CheckUpgradeStatusReply struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *CheckUpgradeStatusReply) Reset()                    { *m = CheckUpgradeStatusReply{} }
func (m *CheckUpgradeStatusReply) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusReply) ProtoMessage()               {}
func (*CheckUpgradeStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckUpgradeStatusReply) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *CheckUpgradeStatusReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*TransmitStateRequest)(nil), "idl.TransmitStateRequest")
	proto.RegisterType((*TransmitStateReply)(nil), "idl.TransmitStateReply")
	proto.RegisterType((*CheckUpgradeStatusRequest)(nil), "idl.CheckUpgradeStatusRequest")
	proto.RegisterType((*CheckUpgradeStatusReply)(nil), "idl.CheckUpgradeStatusReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CommandListener service

type CommandListenerClient interface {
	TransmitState(ctx context.Context, in *TransmitStateRequest, opts ...grpc.CallOption) (*TransmitStateReply, error)
	CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error)
}

type commandListenerClient struct {
	cc *grpc.ClientConn
}

func NewCommandListenerClient(cc *grpc.ClientConn) CommandListenerClient {
	return &commandListenerClient{cc}
}

func (c *commandListenerClient) TransmitState(ctx context.Context, in *TransmitStateRequest, opts ...grpc.CallOption) (*TransmitStateReply, error) {
	out := new(TransmitStateReply)
	err := grpc.Invoke(ctx, "/idl.CommandListener/TransmitState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandListenerClient) CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error) {
	out := new(CheckUpgradeStatusReply)
	err := grpc.Invoke(ctx, "/idl.CommandListener/CheckUpgradeStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CommandListener service

type CommandListenerServer interface {
	TransmitState(context.Context, *TransmitStateRequest) (*TransmitStateReply, error)
	CheckUpgradeStatus(context.Context, *CheckUpgradeStatusRequest) (*CheckUpgradeStatusReply, error)
}

func RegisterCommandListenerServer(s *grpc.Server, srv CommandListenerServer) {
	s.RegisterService(&_CommandListener_serviceDesc, srv)
}

func _CommandListener_TransmitState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransmitStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandListenerServer).TransmitState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CommandListener/TransmitState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandListenerServer).TransmitState(ctx, req.(*TransmitStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandListener_CheckUpgradeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUpgradeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandListenerServer).CheckUpgradeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CommandListener/CheckUpgradeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandListenerServer).CheckUpgradeStatus(ctx, req.(*CheckUpgradeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CommandListener_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.CommandListener",
	HandlerType: (*CommandListenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransmitState",
			Handler:    _CommandListener_TransmitState_Handler,
		},
		{
			MethodName: "CheckUpgradeStatus",
			Handler:    _CommandListener_CheckUpgradeStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "command.proto",
}

func init() { proto.RegisterFile("command.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x4c, 0xc9, 0x51, 0xd2,
	0xe2, 0x12, 0x09, 0x29, 0x4a, 0xcc, 0x2b, 0xce, 0xcd, 0x2c, 0x09, 0x2e, 0x49, 0x2c, 0x49, 0x0d,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0xf4, 0xb8, 0x84, 0xd0, 0xd4, 0x16, 0xe4, 0x54,
	0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xc3, 0x14, 0xc3, 0xb8, 0x4a, 0xd2,
	0x5c, 0x92, 0xce, 0x19, 0xa9, 0xc9, 0xd9, 0xa1, 0x05, 0xe9, 0x45, 0x89, 0x29, 0xa9, 0x20, 0x3d,
	0xa5, 0xc5, 0x50, 0x0b, 0x94, 0xdc, 0xb9, 0xc4, 0xb1, 0x49, 0x82, 0x4c, 0x14, 0xe3, 0x62, 0x2b,
	0x06, 0x73, 0xa1, 0x06, 0x42, 0x79, 0x42, 0x22, 0x5c, 0xac, 0xa9, 0x45, 0x45, 0xf9, 0x45, 0x12,
	0x4c, 0x60, 0x61, 0x08, 0xc7, 0x68, 0x1d, 0x23, 0x17, 0xbf, 0x33, 0xc4, 0x63, 0x3e, 0x99, 0xc5,
	0x25, 0xa9, 0x79, 0xa9, 0x45, 0x42, 0xae, 0x5c, 0xbc, 0x28, 0x2e, 0x15, 0x92, 0xd4, 0xcb, 0x4c,
	0xc9, 0xd1, 0xc3, 0xe6, 0x53, 0x29, 0x71, 0x6c, 0x52, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x42, 0x21,
	0x5c, 0x42, 0x98, 0x6e, 0x14, 0x92, 0x03, 0x6b, 0xc0, 0xe9, 0x33, 0x29, 0x19, 0x9c, 0xf2, 0x60,
	0x53, 0x93, 0xd8, 0xc0, 0xc1, 0x6f, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x12, 0x64, 0xef, 0x0a,
	0x8f, 0x01, 0x00, 0x00,
}

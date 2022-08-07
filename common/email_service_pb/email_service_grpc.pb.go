// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: common/proto/email_service.proto

package email_service_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EmailClient is the client API for Email service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailClient interface {
	GetEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error)
	SubscribeToInbox(ctx context.Context, in *SubscribeCom, opts ...grpc.CallOption) (*ComResponse, error)
}

type emailClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailClient(cc grpc.ClientConnInterface) EmailClient {
	return &emailClient{cc}
}

func (c *emailClient) GetEmail(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*EmailResponse, error) {
	out := new(EmailResponse)
	err := c.cc.Invoke(ctx, "/email_proto.Email/GetEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailClient) SubscribeToInbox(ctx context.Context, in *SubscribeCom, opts ...grpc.CallOption) (*ComResponse, error) {
	out := new(ComResponse)
	err := c.cc.Invoke(ctx, "/email_proto.Email/SubscribeToInbox", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmailServer is the server API for Email service.
// All implementations must embed UnimplementedEmailServer
// for forward compatibility
type EmailServer interface {
	GetEmail(context.Context, *EmailRequest) (*EmailResponse, error)
	SubscribeToInbox(context.Context, *SubscribeCom) (*ComResponse, error)
	mustEmbedUnimplementedEmailServer()
}

// UnimplementedEmailServer must be embedded to have forward compatible implementations.
type UnimplementedEmailServer struct {
}

func (UnimplementedEmailServer) GetEmail(context.Context, *EmailRequest) (*EmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEmail not implemented")
}
func (UnimplementedEmailServer) SubscribeToInbox(context.Context, *SubscribeCom) (*ComResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeToInbox not implemented")
}
func (UnimplementedEmailServer) mustEmbedUnimplementedEmailServer() {}

// UnsafeEmailServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailServer will
// result in compilation errors.
type UnsafeEmailServer interface {
	mustEmbedUnimplementedEmailServer()
}

func RegisterEmailServer(s grpc.ServiceRegistrar, srv EmailServer) {
	s.RegisterService(&Email_ServiceDesc, srv)
}

func _Email_GetEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).GetEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_proto.Email/GetEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).GetEmail(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Email_SubscribeToInbox_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeCom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailServer).SubscribeToInbox(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_proto.Email/SubscribeToInbox",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailServer).SubscribeToInbox(ctx, req.(*SubscribeCom))
	}
	return interceptor(ctx, in, info, handler)
}

// Email_ServiceDesc is the grpc.ServiceDesc for Email service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Email_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "email_proto.Email",
	HandlerType: (*EmailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEmail",
			Handler:    _Email_GetEmail_Handler,
		},
		{
			MethodName: "SubscribeToInbox",
			Handler:    _Email_SubscribeToInbox_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/email_service.proto",
}

// NewEmailNotifClient is the client API for NewEmailNotif service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewEmailNotifClient interface {
	Notify(ctx context.Context, in *NewEmailCom, opts ...grpc.CallOption) (*Empty, error)
}

type newEmailNotifClient struct {
	cc grpc.ClientConnInterface
}

func NewNewEmailNotifClient(cc grpc.ClientConnInterface) NewEmailNotifClient {
	return &newEmailNotifClient{cc}
}

func (c *newEmailNotifClient) Notify(ctx context.Context, in *NewEmailCom, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/email_proto.NewEmailNotif/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewEmailNotifServer is the server API for NewEmailNotif service.
// All implementations must embed UnimplementedNewEmailNotifServer
// for forward compatibility
type NewEmailNotifServer interface {
	Notify(context.Context, *NewEmailCom) (*Empty, error)
	mustEmbedUnimplementedNewEmailNotifServer()
}

// UnimplementedNewEmailNotifServer must be embedded to have forward compatible implementations.
type UnimplementedNewEmailNotifServer struct {
}

func (UnimplementedNewEmailNotifServer) Notify(context.Context, *NewEmailCom) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedNewEmailNotifServer) mustEmbedUnimplementedNewEmailNotifServer() {}

// UnsafeNewEmailNotifServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewEmailNotifServer will
// result in compilation errors.
type UnsafeNewEmailNotifServer interface {
	mustEmbedUnimplementedNewEmailNotifServer()
}

func RegisterNewEmailNotifServer(s grpc.ServiceRegistrar, srv NewEmailNotifServer) {
	s.RegisterService(&NewEmailNotif_ServiceDesc, srv)
}

func _NewEmailNotif_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewEmailCom)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewEmailNotifServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/email_proto.NewEmailNotif/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewEmailNotifServer).Notify(ctx, req.(*NewEmailCom))
	}
	return interceptor(ctx, in, info, handler)
}

// NewEmailNotif_ServiceDesc is the grpc.ServiceDesc for NewEmailNotif service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewEmailNotif_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "email_proto.NewEmailNotif",
	HandlerType: (*NewEmailNotifServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _NewEmailNotif_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/email_service.proto",
}

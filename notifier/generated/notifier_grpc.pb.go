// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// NotifierClient is the client API for Notifier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifierClient interface {
	Push(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error)
}

type notifierClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifierClient(cc grpc.ClientConnInterface) NotifierClient {
	return &notifierClient{cc}
}

func (c *notifierClient) Push(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error) {
	out := new(NotificationReply)
	err := c.cc.Invoke(ctx, "/notifier.Notifier/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotifierServer is the server API for Notifier service.
// All implementations must embed UnimplementedNotifierServer
// for forward compatibility
type NotifierServer interface {
	Push(context.Context, *NotificationRequest) (*NotificationReply, error)
	mustEmbedUnimplementedNotifierServer()
}

// UnimplementedNotifierServer must be embedded to have forward compatible implementations.
type UnimplementedNotifierServer struct {
}

func (UnimplementedNotifierServer) Push(context.Context, *NotificationRequest) (*NotificationReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedNotifierServer) mustEmbedUnimplementedNotifierServer() {}

// UnsafeNotifierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifierServer will
// result in compilation errors.
type UnsafeNotifierServer interface {
	mustEmbedUnimplementedNotifierServer()
}

func RegisterNotifierServer(s grpc.ServiceRegistrar, srv NotifierServer) {
	s.RegisterService(&_Notifier_serviceDesc, srv)
}

func _Notifier_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotifierServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notifier.Notifier/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotifierServer).Push(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Notifier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "notifier.Notifier",
	HandlerType: (*NotifierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Notifier_Push_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifier/generated/notifier.proto",
}

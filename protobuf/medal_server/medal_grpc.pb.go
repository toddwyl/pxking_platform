// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package medal_server

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

// MedalServiceClient is the client API for MedalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MedalServiceClient interface {
	// Sends a greeting
	GetMedalList(ctx context.Context, in *GetMedalListReq, opts ...grpc.CallOption) (*GetMedalListResp, error)
}

type medalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMedalServiceClient(cc grpc.ClientConnInterface) MedalServiceClient {
	return &medalServiceClient{cc}
}

func (c *medalServiceClient) GetMedalList(ctx context.Context, in *GetMedalListReq, opts ...grpc.CallOption) (*GetMedalListResp, error) {
	out := new(GetMedalListResp)
	err := c.cc.Invoke(ctx, "/api.MedalService/GetMedalList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MedalServiceServer is the server API for MedalService service.
// All implementations must embed UnimplementedMedalServiceServer
// for forward compatibility
type MedalServiceServer interface {
	// Sends a greeting
	GetMedalList(context.Context, *GetMedalListReq) (*GetMedalListResp, error)
	mustEmbedUnimplementedMedalServiceServer()
}

// UnimplementedMedalServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMedalServiceServer struct {
}

func (UnimplementedMedalServiceServer) GetMedalList(context.Context, *GetMedalListReq) (*GetMedalListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMedalList not implemented")
}
func (UnimplementedMedalServiceServer) mustEmbedUnimplementedMedalServiceServer() {}

// UnsafeMedalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MedalServiceServer will
// result in compilation errors.
type UnsafeMedalServiceServer interface {
	mustEmbedUnimplementedMedalServiceServer()
}

func RegisterMedalServiceServer(s grpc.ServiceRegistrar, srv MedalServiceServer) {
	s.RegisterService(&MedalService_ServiceDesc, srv)
}

func _MedalService_GetMedalList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMedalListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MedalServiceServer).GetMedalList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.MedalService/GetMedalList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MedalServiceServer).GetMedalList(ctx, req.(*GetMedalListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MedalService_ServiceDesc is the grpc.ServiceDesc for MedalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MedalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.MedalService",
	HandlerType: (*MedalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMedalList",
			Handler:    _MedalService_GetMedalList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "medal.proto",
}

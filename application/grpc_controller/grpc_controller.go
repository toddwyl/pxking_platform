package grpc_controller

import (
	"github.com/go-eagle/eagle/application/grpc_controller/core"
	"github.com/go-eagle/eagle/global"
	"github.com/go-eagle/eagle/infrastructure/transport/grpc"
	"github.com/go-eagle/eagle/protobuf/helloworld"
	"github.com/go-eagle/eagle/protobuf/medal_service"
)

// NewGRPCServer creates a gRPC server
func NewGRPCServer(c *global.ServerConfig) *grpc.Server {

	grpcServer := grpc.NewServer(
		grpc.Network(c.Network),
		grpc.Address(c.Addr),
		grpc.Timeout(c.WriteTimeout),
	)

	// register biz service
	helloworld.RegisterGreeterServer(grpcServer, core.NewHelloWorldGrpcServer())
	medal_service.RegisterMedalServiceServer(grpcServer, core.NewMedalGrpcServer())

	return grpcServer
}

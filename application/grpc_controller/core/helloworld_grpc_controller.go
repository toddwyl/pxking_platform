package core

import (
	"context"
	"fmt"
	pb "github.com/go-eagle/eagle/protobuf/helloworld"
)

func NewHelloWorldGrpcServer() *HelloWorldGrpcServer {
	return &HelloWorldGrpcServer{}
}

type IHelloWorldGrpcServer interface {
	SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error)
}

// HelloWorldGrpcServer is used to implement helloworld.GreeterServer.
type HelloWorldGrpcServer struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *HelloWorldGrpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "" {
		return nil, fmt.Errorf("invalid argument %s", in.Name)
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

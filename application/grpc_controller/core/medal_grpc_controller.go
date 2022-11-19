package core

import (
	"context"
	"github.com/go-eagle/eagle/domain/medal/service"
	"github.com/go-eagle/eagle/global"
	"github.com/go-eagle/eagle/infrastructure/common"
	pb "github.com/go-eagle/eagle/protobuf/medal_service"
)

func NewMedalGrpcServer() *MedalGrpcServer {
	return &MedalGrpcServer{
		medalService: service.NewMedalService(),
	}
}

type IMedalGrpcServer interface {
	GetMedalList(ctx context.Context, in *pb.GetMedalListReq) (*pb.GetMedalListResp, error)
}

// MedalGrpcServer is used to implement helloworld.GreeterServer.
type MedalGrpcServer struct {
	pb.UnimplementedMedalServiceServer
	medalService service.IMedalService
}

// GetMedalList implements MedalGrpcServer
func (s *MedalGrpcServer) GetMedalList(ctx context.Context, in *pb.GetMedalListReq) (*pb.GetMedalListResp, error) {
	c := global.NewCommonContext(ctx)

	resp, err := s.medalService.GetMedalList(c, in)
	if err != nil {
		return common.ErrGrpcRespWrap(resp, err).(*pb.GetMedalListResp), nil
	}
	return common.SuccessGrpcRespWrap(resp).(*pb.GetMedalListResp), nil
}

package service

import (
	"github.com/go-eagle/eagle/domain/medal/persistance"
	"github.com/go-eagle/eagle/global"
	"github.com/go-eagle/eagle/infrastructure/common/errcode"
	pb "github.com/go-eagle/eagle/protobuf/medal_service"
)

type IMedalService interface {
	GetMedalList(ctx global.SysContext, req *pb.GetMedalListReq) (*pb.GetMedalListResp, *errcode.Error)
}

func NewMedalService() IMedalService {
	return &medalService{
		dao: persistance.NewMedalDAO(),
	}
}

type medalService struct {
	dao persistance.IMedalDAO
}

func (m *medalService) GetMedalList(ctx global.SysContext, req *pb.GetMedalListReq) (*pb.GetMedalListResp, *errcode.Error) {
	tabs, err := m.dao.GetMedalList(ctx, req.UserDid)
	if err != nil {
		return nil, err
	}
	resp := &pb.GetMedalListResp{
		//MedalInfoList: make([]*pb.MedalInfo, 0),
		Data: &pb.MedalList{
			MedalInfoList: make([]*pb.MedalInfo, 0),
		},
	}
	for _, tab := range tabs {
		resp.Data.MedalInfoList = append(resp.Data.MedalInfoList, &pb.MedalInfo{
			UserDid:     tab.UserDid,
			EventUid:    tab.EventUid,
			Teacher:     tab.Teacher,
			MedalStatus: int32(tab.MedalStatus),
			ChainHash:   tab.ChainHash,
			ChainStatus: int32(tab.ChainStatus),
		})
	}
	return resp, nil

}

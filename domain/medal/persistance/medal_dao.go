package persistance

import (
	"github.com/go-eagle/eagle/domain/medal/do"
	"github.com/go-eagle/eagle/global"
	errcode2 "github.com/go-eagle/eagle/infrastructure/common/errcode"
	"github.com/go-eagle/eagle/infrastructure/dbhelper"
)

type IMedalDAO interface {
	GetMedalList(ctx global.SysContext, userDid string) ([]*do.MedalInfoTab, *errcode2.Error)
}

func NewMedalDAO() IMedalDAO {
	return &medalDAO{}
}

type medalDAO struct {
}

func (m *medalDAO) GetMedalList(ctx global.SysContext, userDid string) ([]*do.MedalInfoTab, *errcode2.Error) {
	var res []*do.MedalInfoTab
	err := dbhelper.SearchAllData(ctx, &do.MedalInfoTab{}, &res, map[string]interface{}{
		"user_did": userDid,
	})
	if err != nil {
		return nil, errcode2.CloneErrorWithDetail(errcode2.ErrDatabase, err.Error())
	}
	return res, nil
}

package logic

import (
	"context"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShareBasicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicDetailLogic {
	return &ShareBasicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicDetailLogic) ShareBasicDetail(req *types.ShareBasicDetailRequest) (resp *types.ShareBasicDetailReply, err error) {
	err = l.svcCtx.DB.Exec("UPDATE share_basic SET click_num = click_num + 1 where identity = ?", req.Identity).Error
	if err != nil {
		return
	}
	db := l.svcCtx.DB
	shareBasic := new(models.ShareBasic)
	err = db.Table(shareBasic.TableName()).Select("repository_identity").Where("identity = ?", req.Identity).Take(shareBasic).Error
	if err != nil {
		return
	}

	ur := new(models.UserRepository)
	err = db.Table(ur.TableName()).Where("identity = ?", shareBasic.RepositoryIdentity).Take(ur).Error
	if err != nil {
		return
	}

	resp = &types.ShareBasicDetailReply{
		RepositoryIdentity: ur.Identity,
		Name:               ur.Name,
		Ext:                ur.Ext,
		Size:               ur.Size,
	}
	return
}

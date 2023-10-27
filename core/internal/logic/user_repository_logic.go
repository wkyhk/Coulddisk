package logic

import (
	"context"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/helper"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositoryLogic {
	return &UserRepositoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositoryLogic) UserRepository(req *types.UserRepositoryRequest, userIdentity string) (resp *types.UserRepositoryReply, err error) {
	ur := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	var size int64
	db := l.svcCtx.DB
	err = db.Table(models.RepositoryPool{}.TableName()).Select("size").Where("identity = ?", ur.RepositoryIdentity).Take(&size).Error
	if err != nil {
		return
	}
	ur.Size = size
	err = db.Table(ur.TableName()).Create(ur).Error
	if err != nil {
		return
	}

	return &types.UserRepositoryReply{}, err

}

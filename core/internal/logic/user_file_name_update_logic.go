package logic

import (
	"context"
	"errors"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) (resp *types.UserFileNameUpdateReply, err error) {
	db := l.svcCtx.DB
	var count int64
	err = db.Table(models.UserRepository{}.TableName()).Where("user_identity=? AND parent_id =(SELECT parent_id FROM user_repository  WHERE user_repository.identity=?) AND name=?", userIdentity, req.Identity, req.Name).Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		return nil, errors.New("文件名称已存在")
	}
	data := &models.UserRepository{
		Identity:     req.Identity,
		UserIdentity: userIdentity,
	}
	err = db.Model(&data).Where("identity=? AND user_identity=?", data.Identity, data.UserIdentity).Update("name", req.Name).Error
	if err != nil {
		return
	}
	return
}

package logic

import (
	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/helper"
	"context"
	"errors"

	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ShareBasicCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicCreateLogic {
	return &ShareBasicCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicCreateLogic) ShareBasicCreate(req *types.ShareBasicCreateRequest, userIdentity string) (resp *types.ShareBasicCreateReply, err error) {
	db := l.svcCtx.DB
	ur := new(models.UserRepository)
	err = db.Table(ur.TableName()).Where("identity = ? ", req.UserRepositoryIdentity).Take(ur).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("文件不存在")
	} else if err != nil {
		return
	}

	data := &models.ShareBasic{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		RepositoryIdentity: req.UserRepositoryIdentity,
		ExpiredTime:        req.ExpiredTime,
	}

	err = db.Table(data.TableName()).Create(data).Error
	if err != nil {
		return
	}
	resp = &types.ShareBasicCreateReply{
		Identity: data.Identity,
	}
	return
}

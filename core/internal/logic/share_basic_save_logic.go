package logic

import (
	"context"
	"errors"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/helper"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ShareBasicSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShareBasicSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShareBasicSaveLogic {
	return &ShareBasicSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShareBasicSaveLogic) ShareBasicSave(req *types.ShareBasicSaveRequest, userIdentity string) (resp *types.ShareBasicSaveReply, err error) {
	//获取资源详情
	ur := new(models.UserRepository)
	db := l.svcCtx.DB
	err = db.Table(ur.TableName()).Where("identity =(SELECT repository_identity FROM share_basic WHERE identity = ? )", req.ShareIdentity).Take(ur).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("资源不存在")
	} else if err != nil {
		return
	}
	err = db.Table(ur.TableName()).Where("id = ? AND user_identity =?", req.ParentId, userIdentity).Take(&models.UserRepository{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("文件夹不存在")
	} else if err != nil {
		return
	}
	ur1 := &models.UserRepository{
		Identity:           helper.UUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: ur.RepositoryIdentity,
		Ext:                ur.Ext,
		Name:               ur.Name,
		Size:               ur.Size,
	}
	//资源保存
	err = db.Table(ur.TableName()).Create(ur1).Error
	resp = new(types.ShareBasicSaveReply)
	resp.Identity = ur1.Identity
	return
}

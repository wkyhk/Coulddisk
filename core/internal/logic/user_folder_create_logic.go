package logic

import (
	"context"
	"errors"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/define"
	"CLOUDDISK/core/helper"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFolderCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFolderCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFolderCreateLogic {
	return &UserFolderCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFolderCreateLogic) UserFolderCreate(req *types.UserFolderCreateRequest, userIdentity string) (resp *types.UserFolderCreateReply, err error) {
	db := l.svcCtx.DB
	var count int64
	err = db.Table(models.UserRepository{}.TableName()).Where("user_identity = ? AND parent_id = ? AND ext = ? AND name=?", userIdentity, req.ParentId, "", req.Name).Count(&count).Error
	if err != nil {
		return
	}
	if count > 0 {
		return nil, errors.New("文件夹已经存在")
	}
	data := &models.UserRepository{
		Identity:     helper.UUID(),
		UserIdentity: userIdentity,
		ParentId:     req.ParentId,
		Name:         req.Name,
		Ext:          define.Folder,
	}
	err = db.Table(models.UserRepository{}.TableName()).Create(data).Error
	if err != nil {
		return
	}
	return
}

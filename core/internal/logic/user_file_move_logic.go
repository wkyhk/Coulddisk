package logic

import (
	"context"
	"errors"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/define"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserFileMoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileMoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileMoveLogic {
	return &UserFileMoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileMoveLogic) UserFileMove(req *types.UserFileMoveRequest, userIdentity string) (resp *types.UserFileMoveReply, err error) {
	parentData := new(models.UserRepository)
	db := l.svcCtx.DB
	err = db.Table(parentData.TableName()).Where("identity =? AND user_identity=? AND ext = ? ", req.ParentIdentity, userIdentity, define.Folder).First(&parentData).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("文件夹不存在")
	} else if err != nil {
		return
	}
	err = db.Table(parentData.TableName()).Where("identity =? AND user_identity=?", req.Identity, userIdentity).Updates(models.UserRepository{
		ParentId: int64(parentData.ID),
	}).Error
	if err != nil {
		return
	}

	return
}

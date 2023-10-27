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

type FileUploadPrepareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileUploadPrepareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileUploadPrepareLogic {
	return &FileUploadPrepareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileUploadPrepareLogic) FileUploadPrepare(req *types.FileUploadPrepareRequest) (resp *types.FileUploadPrepareReply, err error) {
	rp := new(models.RepositoryPool)
	db := l.svcCtx.DB
	resp = new(types.FileUploadPrepareReply)
	err = db.Table(rp.TableName()).Where("hash = ? ", req.Md5).Take(rp).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return
	} else if err == nil {
		resp.Identity = rp.Identity
		return
	}
	// 获取文件的upload_id
	key, uploadId, err := helper.CosInitPartUpload(req.Ext)
	if err != nil {
		return nil, err
	}
	resp.Key = key
	resp.UploadId = uploadId
	return
}

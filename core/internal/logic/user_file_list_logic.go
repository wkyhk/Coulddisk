package logic

import (
	"context"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/define"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListReply, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileListReply)
	var cnt int64
	//分页的参数
	size := req.Size
	if size == 0 {
		size = define.PageSize
	}

	page := req.Page
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * size
	db := l.svcCtx.DB
	err = db.Table(models.UserRepository{}.TableName()).Select("user_repository.id as id,user_repository.identity as identity,user_repository.repository_identity as repository_identity,user_repository.name as name,user_repository.ext as ext,repository_pool.path as path,user_repository.size as size").Where("parent_id =? AND user_identity =? AND user_repository.deleted_at IS NULL ", req.Id, userIdentity).Joins("LEFT JOIN repository_pool ON user_repository.repository_identity = repository_pool.identity").Limit(size).Offset(offset).Find(&uf).Error
	if err != nil {
		return
	}
	db.Table(models.UserRepository{}.TableName()).Where("parent_id =? AND user_identity =? AND user_repository.deleted_at IS NULL", req.Id, userIdentity).Count(&cnt)
	resp.List = uf
	resp.Count = cnt

	return
}

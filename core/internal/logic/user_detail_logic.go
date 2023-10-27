package logic

import (
	"context"
	"errors"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailReply, err error) {
	// todo: add your logic here and delete this line
	resp = &types.UserDetailReply{}
	ub := new(models.UserBasic)
	db := l.svcCtx.DB
	err = db.Where(models.UserBasic{Identity: req.Identity}, "identity").First(ub).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not found")
	} else if err != nil {
		return nil, err
	}
	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}

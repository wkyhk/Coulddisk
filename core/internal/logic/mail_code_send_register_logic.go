package logic

import (
	"context"
	"errors"
	"time"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/define"
	"CLOUDDISK/core/helper"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	var cnt int64
	err = l.svcCtx.DB.Table(models.UserBasic{}.TableName()).Where("email = ?", req.Email).Count(&cnt).Error
	if cnt > 0 {
		return nil, errors.New("该邮箱已被注册")
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	code, _ := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if len(code) != 6 {
		code = helper.RandCode()
		l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	}
	err = helper.MailSendCode(req.Email, code)
	if err != nil {
		return nil, err
	}
	return
}

package logic

import (
	"context"
	"errors"
	"log"

	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/helper"
	"CLOUDDISK/core/internal/svc"
	"CLOUDDISK/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterReply, err error) {
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, errors.New("未获取该邮箱验证码")
	}
	if code != req.Code {
		return nil, errors.New("验证码错误")
	}
	db := l.svcCtx.DB
	us := new(models.UserBasic)
	//判断用户名是否存在
	err = db.Where(&models.UserBasic{Name: req.Name}, "name").First(&us).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		us = &models.UserBasic{
			Identity: helper.UUID(),
			Name:     req.Name,
			Password: helper.Md5(req.Password),
			Email:    req.Email,
		}
		result := db.Table(us.TableName()).Create(us)
		log.Println("insert user row :%n", result.RowsAffected)
		return nil, nil
	} else if err == nil {
		return nil, errors.New("用户已经存在")
	} else if err != nil {
		return nil, err
	}

	return nil, err
}

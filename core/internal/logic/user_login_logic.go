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
	"gorm.io/gorm"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	//1.从数据库中查取用户信息
	user := new(models.UserBasic)
	db := l.svcCtx.DB
	err = db.Where(&models.UserBasic{Name: req.Name, Password: helper.Md5(req.Password)}, "name", "password").First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户名或密码错误")
	} else if err != nil {
		return nil, err
	}

	//db.where(models.UserBasic{Name: req.Name, Password: req.Password}, "name", "password").Find(&user)
	//2.生成token
	token, err := helper.GenerateToken(user.ID, user.Identity, user.Name, 20)
	if err != nil {
		return nil, err
	}
	//用于刷新token的token
	refreshToken, err := helper.GenerateToken(user.ID, user.Identity, user.Name, define.RefreshTokenExpire)
	if err != nil {
		return nil, err
	}
	resp = new(types.LoginReply)
	resp.Token = token
	resp.RefreshToken = refreshToken
	return
}

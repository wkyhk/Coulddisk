package define

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaim struct {
	Id       uint
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

// 验证码长度
var CodeLength = 6

// 验证码过期时间(s)
var CodeExpire = 300

var TencentSecretID = ""
var TencentSecretKey = ""
var CosBucket = ""

// PageSize 分页的默认参数
var PageSize = 20
var Folder = "文件夹"

// token 有效期
var TokenExpire int64 = 3600
var RefreshTokenExpire int64 = 7200

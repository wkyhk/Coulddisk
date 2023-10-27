package svc

import (
	models "CLOUDDISK/core/Models"
	"CLOUDDISK/core/internal/config"
	"CLOUDDISK/core/internal/middleware"

	"github.com/go-redis/redis/v8"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     models.GetDB(),
		RDB:    models.InitRedis(c.Redis.Addr),
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}

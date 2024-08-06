package global

import (
	"personal/ShopDev/Go-BE/package/logger"
	"personal/ShopDev/Go-BE/package/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)

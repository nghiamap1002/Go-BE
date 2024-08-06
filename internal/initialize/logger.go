package initialize

import (
	"personal/ShopDev/Go-BE/global"
	"personal/ShopDev/Go-BE/package/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}

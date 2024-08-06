package initialize

import (
	"fmt"
	"personal/ShopDev/Go-BE/global"
	"personal/ShopDev/Go-BE/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err)) // not understand zap error
		panic(err)
	}
}

func InitMysql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})

	checkErrorPanic(err, "init mysql error")
	global.Logger.Info("Mysql init success")
	global.Mdb = db

	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		fmt.Print("Mysql error: %s::", sqlDB)
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func migrateTables() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		fmt.Println("Migrating table error:", err)
	}
}

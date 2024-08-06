package benchmark

import (
	"fmt"
	"log"
	"personal/ShopDev/Go-BE/global"
	"testing"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   int
	Name string
}

func InserRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "concobebe"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	m := MySQLSetting{Host: "127.0.0.1", Port: 8811,
		Username: "root",
		Password: "concobebe",
		DBName:   "shopdevgo",
	}

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.DBName)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to coonect mysql", err)
	}

	if db.Migrator().HasTable(&User{}) {
		if err := db.Migrator().DropTable(&User{}); err != nil {
			log.Fatal("Failed to coonect mysql", err)
		}
	}

	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get sqlDB", err)
	}
	sqlDB.SetMaxOpenConns(10)
	defer sqlDB.Close()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			InserRecord(b, db)
		}
	})
}

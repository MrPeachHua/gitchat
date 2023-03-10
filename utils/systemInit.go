package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Client *redis.Client

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("ginchat\\config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println("config", viper.Get("app"))
}

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

var DB *gorm.DB

func InitSql() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	DB, _ = gorm.Open(mysql.Open(""), &gorm.Config{Logger: newLogger})
	print("db", DB)
}

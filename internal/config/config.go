package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppName   string
	AppEnv    string
	AppPort   string
	Database  string
	JWTSecret string
}

var Cfg Config

func Load() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("cannot load config:", err)
	}

	Cfg = Config{
		AppName:   viper.GetString("APP_NAME"),
		AppEnv:    viper.GetString("APP_ENV"),
		AppPort:   viper.GetString("APP_PORT"),
		Database:  viper.GetString("DATABASE_URL"),
		JWTSecret: viper.GetString("JWT_SECRET"),
	}
}

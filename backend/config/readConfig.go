package config

import (
	"time"
	"github.com/TimurZheksimbaev/Golang-webchat/utils"
	"github.com/spf13/viper"
)

type AppConfig struct {
	DatabaseURL string `mapstructure:"DB_URL"`
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
	JWTExpiration time.Duration `mapstructure:"TOKEN_EXPIRES_IN"`
	FrontendURL string `mapstructure:"FRONTEND_URL"`
}

func LoadEnv() (*AppConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, utils.ConfigError("Could not read config file", err)
	}
	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	return &appConfig, nil
}


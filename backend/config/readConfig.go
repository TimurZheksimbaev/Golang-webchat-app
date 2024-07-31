package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DatabaseURL string `mapstructure:"DB_URL"`
	ServerHost string `mapstructure:"SERVER_HOST"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
	JWTExpiration time.Duration `mapstructure:"TOKEN_EXPIRES_IN"`
}

func LoadEnv() (*AppConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.New("Could not read config file")
	}
	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	return &appConfig, nil
}


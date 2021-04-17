package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Val Config

type Config struct {
	Port string `mapstructure:"PORT"`
	LogFilePath string `mapstructure:"LOG_FILE_PATH"`
	LogFileName string `mapstructure:"LOG_FILE_NAME"`

	JWTTokenTtl int `mapstructure:"JWT_TOKEN_TTL"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

func Init() {
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %v", err))
	}
	if err := viper.Unmarshal(&Val); err != nil {
		panic(fmt.Errorf("Unable to decode into struct, %v", err))
	}
}

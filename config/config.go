package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Val Config

type Config struct {
	Port string `mapstructure:"PORT"`
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

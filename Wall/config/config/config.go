package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config = viper.New()

func init() {
	Config.SetConfigName("config")
	Config.SetConfigType("yaml")
	Config.AddConfigPath(".")
	Config.WatchConfig()
	err := Config.ReadInConfig()
	if err != nil {
		log.Fatal("error in config", err)
	}
}

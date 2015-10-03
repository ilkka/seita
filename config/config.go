package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.seita")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Could not read config file: %s", err)
	}
}

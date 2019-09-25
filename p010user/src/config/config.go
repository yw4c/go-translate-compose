package config

import (
	"github.com/spf13/viper"
	"log"
)

func LoadEnv(envPath string) {
	viper.SetConfigName("env")

	if envPath == "" {
		viper.AddConfigPath("$GOPATH/src/p010user")
	} else {
		viper.AddConfigPath(envPath)
	}

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("read env file failed. ", err.Error())
	}
}

package config

import (
	"github.com/spf13/viper"
	"os"
)

func Init() error {
	viper.AddConfigPath("./config")

	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	env := os.Getenv("HOST")
	if env == "" {
		env = "local"
	}

	viper.SetConfigName(env)
	return viper.MergeInConfig()
}
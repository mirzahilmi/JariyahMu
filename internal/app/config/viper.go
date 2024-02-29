package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViper() viper.Viper {
	viper := viper.New()

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	return *viper
}

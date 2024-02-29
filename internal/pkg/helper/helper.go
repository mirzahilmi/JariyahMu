package helper

import "github.com/spf13/viper"

func NewHelper(viper *viper.Viper) Helper {
	helper := Helper{}

	return helper
}

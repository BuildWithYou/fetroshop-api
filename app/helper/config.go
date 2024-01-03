package helper

import "github.com/spf13/viper"

func GetConfig() *viper.Viper {
	// Config
	config := viper.New()
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	PanicIfError(err)
	return config
}

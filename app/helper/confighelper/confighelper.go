package confighelper

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	// Config
	config := viper.New()
	config.SetConfigFile("config.yaml")
	err := config.ReadInConfig()
	errorhelper.PanicIfError(err)
	return config
}

package confighelper

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/projectpath"
	"github.com/spf13/viper"
)

func GetConfig() *viper.Viper {
	// Config
	config := viper.New()
	config.AddConfigPath(projectpath.Root)
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	// set up default value
	config.SetDefault("environment", "development")
	config.SetDefault("fiber.prefork", false)
	config.SetDefault("database.logLevel", "halo")

	err := config.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}

	return config
}

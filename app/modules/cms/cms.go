package cms

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/router"
)

//	@title			Fetroshop CMS API
//	@version		1.0
//	@description	Fetroshop API is a robust and efficient backend solution designed to power the online store app named Fetroshop. Developed using the Go programming language, this API serves as the backbone for managing the Content Management System (CMS) and handling various store-related functionalities.

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Use format 'Bearer YOUR_TOKEN'

func CmsServerConfigProvider(webRouter router.Router) *app.ServerConfig {
	config := helper.GetConfig()
	return &app.ServerConfig{
		Config: config,
		Host:   config.GetString("app.cms.host"),
		Port:   config.GetInt("app.cms.port"),
		Router: webRouter,
	}
}

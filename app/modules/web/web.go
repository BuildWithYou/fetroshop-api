package web

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/router"
)

//	@title			Fetroshop Web API
//	@version		1.0
//	@description	Fetroshop API is a robust and efficient backend solution designed to power the online store app named Fetroshop. Developed using the Go programming language, this API serves as the backbone for managing the Content Management System (CMS) and handling various store-related functionalities.

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Use format 'Bearer YOUR_TOKEN'

func WebServerConfigProvider(webRouter router.Router) *app.ServerConfig {
	config := app.GetConfig()
	return &app.ServerConfig{
		Config: config,
		Host:   config.GetString("app.host"),
		Port:   config.GetString("app.port"),
		Router: webRouter,
		Static: &map[string]string{
			"/swagger": "docs",
		},
	}
}

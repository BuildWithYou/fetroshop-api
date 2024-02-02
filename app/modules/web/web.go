package web

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/gofiber/template/html/v2"
)

//	@title         Fetroshop Web API
//	@version       1.0
//	@description	Fetroshop API is a robust and efficient backend solution designed to power the online store app named Fetroshop. Developed using the Go programming language, this API serves as the backbone for managing the Content Management System (CMS) and handling various store-related functionalities.

//	@securityDefinitions.apikey	Bearer
//	@in						header
//	@name						Authorization
//	@description			Use format 'Bearer YOUR_TOKEN'

func WebServerConfigProvider(webRouter app.Router, logger *logger.Logger) *app.ServerConfig {
	config := confighelper.GetConfig()
	logger.Info("Initializing web server")
	return &app.ServerConfig{
		Config: config,
		Host:   config.GetString("app.web.host"),
		Port:   config.GetInt("app.web.port"),
		Router: webRouter,
		Static: map[string]string{
			"/swagger": "docs",
		},
		Logger:     logger,
		ViewEngine: html.New("./docs/swagger-ui", ".gohtml"),
	}
}

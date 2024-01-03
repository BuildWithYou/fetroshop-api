package docs

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/router"
)

func DocsServerConfigProvider(webRouter router.Router) *app.ServerConfig {
	config := helper.GetConfig()
	return &app.ServerConfig{
		Config: config,
		Host:   config.GetString("app.docs.host"),
		Port:   config.GetInt("app.docs.port"),
		Router: webRouter,
		Static: map[string]string{
			"/swagger": "docs",
		},
	}
}

package docs

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

type Docs struct {
	Config *viper.Viper
}

func (d *Docs) SwaggerWeb() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.url")
	return swagger.New(swagger.Config{ // custom
		URL:         fmt.Sprintf("%s/swagger/web/swagger.json", url),
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	})
}

func (d *Docs) SwaggerCms() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.url")
	return swagger.New(swagger.Config{ // custom
		URL:         fmt.Sprintf("%s/swagger/cms/swagger.json", url),
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	})
}

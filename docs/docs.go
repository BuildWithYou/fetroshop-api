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
	})
}

func (d *Docs) SwaggerCms() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.url")
	return swagger.New(swagger.Config{ // custom
		URL:         fmt.Sprintf("%s/swagger/cms/swagger.json", url),
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
	})
}

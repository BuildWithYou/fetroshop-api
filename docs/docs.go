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

func (d *Docs) createSwagger(jsonUrl string) func(*fiber.Ctx) error {
	return swagger.New(swagger.Config{ // custom
		URL:          jsonUrl,
		DeepLinking:  false,
		DocExpansion: "none",
	})
}

func (d *Docs) SwaggerWeb() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.url")
	jsonUrl := fmt.Sprintf("%s/swagger/web/swagger.json", url)
	return d.createSwagger(jsonUrl)
}

func (d *Docs) SwaggerCms() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.url")
	jsonUrl := fmt.Sprintf("%s/swagger/cms/swagger.json", url)
	return d.createSwagger(jsonUrl)
}

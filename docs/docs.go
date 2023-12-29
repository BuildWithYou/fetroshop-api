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

type swaggerConfig struct {
	jsonUrl string
}

func (d *Docs) createSwagger(sc *swaggerConfig) func(*fiber.Ctx) error {
	return swagger.New(swagger.Config{ // custom
		URL:          sc.jsonUrl,
		DeepLinking:  d.Config.GetBool("swagger.deepLinking"),
		DocExpansion: d.Config.GetString("swagger.docExpansion"),
	})
}

func (d *Docs) SwaggerWeb() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.url")
	jsonUrl := fmt.Sprintf("%s/swagger/web/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		jsonUrl: jsonUrl,
	})
}

func (d *Docs) SwaggerCms() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.url")
	jsonUrl := fmt.Sprintf("%s/swagger/cms/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		jsonUrl: jsonUrl,
	})
}

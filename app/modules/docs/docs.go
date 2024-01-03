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
	title   string
}

func NewDocs(config *viper.Viper) *Docs {
	return &Docs{
		Config: config,
	}
}

func (d *Docs) createSwagger(sc *swaggerConfig) func(*fiber.Ctx) error {
	return swagger.New(swagger.Config{ // custom
		URL:          sc.jsonUrl,
		DeepLinking:  d.Config.GetBool("swagger.deepLinking"),
		DocExpansion: d.Config.GetString("swagger.docExpansion"),
		Title:        sc.title,
	})
}

func (d *Docs) SwaggerWeb() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.web.url")
	jsonUrl := fmt.Sprintf("%s/swagger/web/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		jsonUrl: jsonUrl,
		title:   "Fetroshop Web API",
	})
}

func (d *Docs) SwaggerCms() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.cms.url")
	jsonUrl := fmt.Sprintf("%s/swagger/cms/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		jsonUrl: jsonUrl,
		title:   "Fetroshop CMS API",
	})
}

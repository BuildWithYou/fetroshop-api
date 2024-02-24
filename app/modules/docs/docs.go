package docs

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Docs struct {
	Config *viper.Viper
}

type swaggerConfig struct {
	swaggerUiUrl   string
	title          string
	swaggerJsonUrl string
	deepLinking    bool
	docExpansion   string
}

func DocsProvider(config *viper.Viper) *Docs {
	return &Docs{
		Config: config,
	}
}

func (d *Docs) createSwagger(sc *swaggerConfig) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"title":          sc.title,
			"swaggerUiUrl":   sc.swaggerUiUrl,
			"swaggerJsonUrl": sc.swaggerJsonUrl,
			"deepLinking":    sc.deepLinking,
			"docExpansion":   sc.docExpansion,
		})
	}
}

func (d *Docs) SwaggerWeb() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.web.url")
	swaggerUiUrl := fmt.Sprintf("%s/swagger/swagger-ui", url)
	swaggerJsonUrl := fmt.Sprintf("%s/swagger/openapi2/web/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		swaggerUiUrl:   swaggerUiUrl,
		title:          "Fetroshop Web API",
		swaggerJsonUrl: swaggerJsonUrl,
		deepLinking:    d.Config.GetBool("swagger.deepLinking"),
		docExpansion:   d.Config.GetString("swagger.docExpansion"),
	})
}

func (d *Docs) SwaggerCms() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.cms.url")
	swaggerUiUrl := fmt.Sprintf("%s/swagger/swagger-ui", url)
	swaggerJsonUrl := fmt.Sprintf("%s/swagger/openapi2/cms/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		swaggerUiUrl:   swaggerUiUrl,
		title:          "Fetroshop CMS API",
		swaggerJsonUrl: swaggerJsonUrl,
		deepLinking:    d.Config.GetBool("swagger.deepLinking"),
		docExpansion:   d.Config.GetString("swagger.docExpansion"),
	})
}

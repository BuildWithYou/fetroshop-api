package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

type DocsRouter struct {
	Config *viper.Viper
}

func (d *DocsRouter) Init(app *fiber.App) {
	// root
	app.Get("/", d.welcome)

	// docs
	app.Get("/web/*", d.SwaggerWeb())
	app.Get("/cms/*", d.SwaggerCms())
}
func DocsRouterProvider(cfg *viper.Viper) Router {
	return &DocsRouter{
		Config: cfg,
	}
}

func (d *DocsRouter) welcome(ctx *fiber.Ctx) error {
	return ctx.SendString("Welcome to fetroshop-api docs!")
}

type swaggerConfig struct {
	jsonUrl string
	title   string
}

func (d *DocsRouter) createSwagger(sc *swaggerConfig) func(*fiber.Ctx) error {
	return swagger.New(swagger.Config{ // custom
		URL:          sc.jsonUrl,
		DeepLinking:  d.Config.GetBool("swagger.deepLinking"),
		DocExpansion: d.Config.GetString("swagger.docExpansion"),
		Title:        sc.title,
	})
}

func (d *DocsRouter) SwaggerWeb() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.docs.url")
	jsonUrl := fmt.Sprintf("%s/swagger/web/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		jsonUrl: jsonUrl,
		title:   "Fetroshop Web API",
	})
}

func (d *DocsRouter) SwaggerCms() func(*fiber.Ctx) error {
	url := d.Config.GetString("app.docs.url")
	jsonUrl := fmt.Sprintf("%s/swagger/cms/swagger.json", url)
	return d.createSwagger(&swaggerConfig{
		jsonUrl: jsonUrl,
		title:   "Fetroshop CMS API",
	})
}

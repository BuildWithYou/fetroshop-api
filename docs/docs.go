package docs

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Docs struct {
	Config *viper.Viper
}

func (d *Docs) Swagger(ctx *fiber.Ctx) error {
	return ctx.Render("swagger/index", fiber.Map{
		"swaggerUrl": fmt.Sprintf("%s/swagger", d.Config.GetString("app.url")),
	})
}

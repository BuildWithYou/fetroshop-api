//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	postgres.NewUserRepository,
	registration.NewRegistrationService,
	controller.NewRegistrationController,
)

func InitializeWebServer() error {
	wire.Build(
		validator.New,
		userSet,
		router.WebRouterProvider,
		web.WebServerConfigProvider,
		app.CreateFiber,
		app.StartFiber,
	)
	return nil
}

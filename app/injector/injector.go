//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	postgres.NewUserRepository,
	registration.NewRegistrationService,
	controller.NewRegistrationController,
)

func InitializeWebServer() error {
	wire.Build(
		helper.GetValidator,
		userSet,
		router.WebRouterProvider,
		web.WebServerConfigProvider,
		app.CreateFiber,
		app.StartFiber,
	)
	return nil
}

func InitializeCmsServer() error {
	wire.Build(
		router.CmsRouterProvider,
		cms.CmsServerConfigProvider,
		app.CreateFiber,
		app.StartFiber,
	)
	return nil
}

func InitializeDocsServer() error {
	wire.Build(
		helper.GetConfig,
		router.DocsRouterProvider,
		docs.DocsServerConfigProvider,
		app.CreateFiber,
		app.StartFiber,
	)
	return nil
}

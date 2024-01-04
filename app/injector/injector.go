//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	customerRepo "github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	userRepo "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	webController "github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	webRegistrationService "github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/BuildWithYou/fetroshop-api/db"
	"github.com/google/wire"
)

var userSet = wire.NewSet(
	userRepo.UserRepositoryProvider,
	webRegistrationService.RegistrationServiceProvider,
	webController.RegistrationControllerProvider,
)

var customerSet = wire.NewSet(
	customerRepo.CustomerRepositoryProvider,
	webRegistrationService.RegistrationServiceProvider,
	webController.RegistrationControllerProvider,
)

func InitializeWebServer() error {
	wire.Build(
		db.OpenConnection,
		helper.GetConfig,
		docs.DocsProvider,
		helper.GetValidator,
		customerSet,
		router.WebRouterProvider,
		web.WebServerConfigProvider,
		app.CreateFiber,
		app.StartFiber,
	)
	return nil
}

func InitializeCmsServer() error {
	wire.Build(
		docs.DocsProvider,
		helper.GetConfig,
		router.CmsRouterProvider,
		// userSet,
		cms.CmsServerConfigProvider,
		app.CreateFiber,
		app.StartFiber,
	)
	return nil
}

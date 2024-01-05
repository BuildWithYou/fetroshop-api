//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	customerRepo "github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	userRepo "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	webController "github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	webRegistrationService "github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/BuildWithYou/fetroshop-api/app/router"
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
		connection.OpenDBConnection,
		confighelper.GetConfig,
		docs.DocsProvider,
		validatorhelper.GetValidator,
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
		confighelper.GetConfig,
		router.CmsRouterProvider,
		// userSet,
		cms.CmsServerConfigProvider,
		app.CreateFiber,
		app.StartFiber,
	)
	return nil
}

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
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	webController "github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	webAuthService "github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	confighelper.GetConfig,
	docs.DocsProvider,
	middleware.JwtMiddlewareProvider,
	app.CreateFiber,
	app.StartFiber,
)

var userSet = wire.NewSet(
	userRepo.UserRepositoryProvider,
	webAuthService.AuthServiceProvider,
	webController.AuthControllerProvider,
)

var customerSet = wire.NewSet(
	customerRepo.CustomerRepositoryProvider,
	webAuthService.AuthServiceProvider,
	webController.AuthControllerProvider,
)

func InitializeWebServer() error {
	wire.Build(
		serverSet,
		connection.OpenDBConnection,
		validatorhelper.GetValidator,
		webController.WebControllerProvider,
		customerSet,
		router.WebRouterProvider,
		web.WebServerConfigProvider,
	)
	return nil
}

func InitializeCmsServer() error {
	wire.Build(
		serverSet,
		router.CmsRouterProvider,
		// userSet,
		cms.CmsServerConfigProvider,
	)
	return nil
}

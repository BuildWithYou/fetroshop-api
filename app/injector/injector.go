//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	categoryRepo "github.com/BuildWithYou/fetroshop-api/app/domain/categories/postgres"
	customerAccessRepo "github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses/postgres"
	customerRepo "github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	userAccessRepo "github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses/postgres"
	userRepo "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	cmsController "github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	cmsAuthService "github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/auth"
	cmsCategoryService "github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/category"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	webController "github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	webAuthService "github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth"
	webCategoryService "github.com/BuildWithYou/fetroshop-api/app/modules/web/service/category"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/google/wire"
)

var dbType connection.DBType = connection.DB_MAIN

var serverSet = wire.NewSet(
	wire.Value(dbType),
	confighelper.GetConfig,
	connection.OpenDBConnection,
	docs.DocsProvider,
	middleware.JwtMiddlewareProvider,
	middleware.DBMiddlewareProvider,
	validatorhelper.GetValidator,
	userAccessRepo.RepoProvider,
	customerAccessRepo.RepoProvider,
	app.CreateFiber,
)

// web dependencies
var webRepoSet = wire.NewSet(
	customerRepo.RepoProvider,
	categoryRepo.RepoProvider,
)

var webControllerSet = wire.NewSet(
	webController.WebControllerProvider,
	webController.AuthControllerProvider,
	webController.CategoryControllerProvider,
)

var webServiceSet = wire.NewSet(
	webAuthService.ServiceProvider,
	webCategoryService.ServiceProvider,
)

func InitializeWebServer() *app.Fetroshop {
	wire.Build(
		serverSet,
		webRepoSet,
		webControllerSet,
		webServiceSet,
		router.WebRouterProvider,
		web.WebServerConfigProvider,
	)
	return nil
}

// cms dependencies
var cmsRepoSet = wire.NewSet(
	userRepo.RepoProvider,
	categoryRepo.RepoProvider,
)

var cmsControllerSet = wire.NewSet(
	cmsController.CmsControllerProvider,
	cmsController.AuthControllerProvider,
	cmsController.CategoryControllerProvider,
)

var cmsServiceSet = wire.NewSet(
	cmsAuthService.ServiceProvider,
	cmsCategoryService.ServiceProvider,
)

func InitializeCmsServer() *app.Fetroshop {
	wire.Build(
		serverSet,
		cmsRepoSet,
		cmsControllerSet,
		cmsServiceSet,
		router.CmsRouterProvider,
		cms.CmsServerConfigProvider,
	)
	return nil
}

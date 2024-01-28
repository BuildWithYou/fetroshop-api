//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	brandRepo "github.com/BuildWithYou/fetroshop-api/app/domain/brands/postgres"
	categoryRepo "github.com/BuildWithYou/fetroshop-api/app/domain/categories/postgres"
	cityRepo "github.com/BuildWithYou/fetroshop-api/app/domain/cities/postgres"
	customerAccessRepo "github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses/postgres"
	customerRepo "github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	districtRepo "github.com/BuildWithYou/fetroshop-api/app/domain/districts/postgres"
	provinceRepo "github.com/BuildWithYou/fetroshop-api/app/domain/provinces/postgres"
	storeRepo "github.com/BuildWithYou/fetroshop-api/app/domain/stores/postgres"
	subdistrictRepo "github.com/BuildWithYou/fetroshop-api/app/domain/subdistricts/postgres"
	userAccessRepo "github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses/postgres"
	userRepo "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	cmsController "github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	webController "github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/service/brand"
	"github.com/BuildWithYou/fetroshop-api/app/service/category"
	"github.com/BuildWithYou/fetroshop-api/app/service/location"
	"github.com/BuildWithYou/fetroshop-api/app/service/store"
	"github.com/google/wire"
)

var dbType connection.DBType = connection.DB_MAIN

var repoSet = wire.NewSet(
	customerRepo.RepoProvider,
	customerAccessRepo.RepoProvider,
	userRepo.RepoProvider,
	userAccessRepo.RepoProvider,
	categoryRepo.RepoProvider,
	brandRepo.RepoProvider,
	storeRepo.RepoProvider,
	provinceRepo.RepoProvider,
	cityRepo.RepoProvider,
	districtRepo.RepoProvider,
	subdistrictRepo.RepoProvider,
)

var serviceSet = wire.NewSet(
	auth.ServiceProvider,
	category.ServiceProvider,
	brand.ServiceProvider,
	store.ServiceProvider,
	location.ServiceProvider,
)

var serverSet = wire.NewSet(
	wire.Value(dbType),
	confighelper.GetConfig,
	connection.OpenDBConnection,
	docs.DocsProvider,
	middleware.JwtMiddlewareProvider,
	middleware.DBMiddlewareProvider,
	validatorhelper.GetValidator,
	repoSet,
	serviceSet,
	app.CreateFiber,
)

// web dependencies
var webControllerSet = wire.NewSet(
	webController.WebControllerProvider,
	webController.AuthControllerProvider,
	webController.CategoryControllerProvider,
	webController.BrandControllerProvider,
	webController.StoreControllerProvider,
	webController.LocationControllerProvider,
)

func InitializeWebServer() *app.Fetroshop {
	wire.Build(
		logger.NewWebLogger,
		serverSet,
		webControllerSet,
		web.RouterProvider,
		web.WebServerConfigProvider,
	)
	return nil
}

// cms dependencies
var cmsControllerSet = wire.NewSet(
	cmsController.CmsControllerProvider,
	cmsController.AuthControllerProvider,
	cmsController.CategoryControllerProvider,
	cmsController.BrandControllerProvider,
	cmsController.StoreControllerProvider,
	cmsController.LocationControllerProvider,
)

func InitializeCmsServer() *app.Fetroshop {
	wire.Build(
		logger.NewCmsLogger,
		serverSet,
		cmsControllerSet,
		cms.RouterProvider,
		cms.CmsServerConfigProvider,
	)
	return nil
}

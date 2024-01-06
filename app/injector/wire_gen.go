// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/BuildWithYou/fetroshop-api/app"
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers/postgres"
	postgres2 "github.com/BuildWithYou/fetroshop-api/app/domain/users/postgres"
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/middleware"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms"
	controller2 "github.com/BuildWithYou/fetroshop-api/app/modules/cms/controller"
	auth2 "github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/modules/docs"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/controller"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth"
	"github.com/BuildWithYou/fetroshop-api/app/router"
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializeWebServer() error {
	viper := confighelper.GetConfig()
	docsDocs := docs.DocsProvider(viper)
	validate := validatorhelper.GetValidator()
	db := connection.OpenDBConnection(viper)
	customerRepository := postgres.CustomerRepositoryProvider(db)
	authService := auth.AuthServiceProvider(db, viper, customerRepository)
	authController := controller.AuthControllerProvider(validate, authService)
	controllerController := controller.WebControllerProvider(authController)
	jwtMiddleware := middleware.JwtMiddlewareProvider(viper)
	routerRouter := router.WebRouterProvider(docsDocs, controllerController, jwtMiddleware)
	serverConfig := web.WebServerConfigProvider(routerRouter)
	fiberApp := app.CreateFiber(serverConfig)
	error2 := app.StartFiber(fiberApp, serverConfig)
	return error2
}

func InitializeCmsServer() error {
	viper := confighelper.GetConfig()
	docsDocs := docs.DocsProvider(viper)
	jwtMiddleware := middleware.JwtMiddlewareProvider(viper)
	db := connection.OpenDBConnection(viper)
	validate := validatorhelper.GetValidator()
	userRepository := postgres2.UserRepositoryProvider(db)
	authService := auth2.AuthServiceProvider(db, viper, validate, userRepository)
	authController := controller2.AuthControllerProvider(authService)
	controllerController := controller2.CmsControllerProvider(authController)
	routerRouter := router.CmsRouterProvider(docsDocs, jwtMiddleware, controllerController)
	serverConfig := cms.CmsServerConfigProvider(routerRouter)
	fiberApp := app.CreateFiber(serverConfig)
	error2 := app.StartFiber(fiberApp, serverConfig)
	return error2
}

// injector.go:

var serverSet = wire.NewSet(confighelper.GetConfig, connection.OpenDBConnection, docs.DocsProvider, middleware.JwtMiddlewareProvider, validatorhelper.GetValidator, app.CreateFiber, app.StartFiber)

// web dependencies
var webRepoSet = wire.NewSet(postgres.CustomerRepositoryProvider)

var webControllerSet = wire.NewSet(controller.WebControllerProvider, controller.AuthControllerProvider)

var webServiceSet = wire.NewSet(auth.AuthServiceProvider)

// cms dependencies
var cmsRepoSet = wire.NewSet(postgres2.UserRepositoryProvider)

var cmsControllerSet = wire.NewSet(controller2.CmsControllerProvider, controller2.AuthControllerProvider)

var cmsServiceSet = wire.NewSet(auth2.AuthServiceProvider)
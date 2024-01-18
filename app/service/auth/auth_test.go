package auth_test

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/confighelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/injector"
)

var fetroshopApp = injector.InitializeCmsServer()
var cmsLogger = logger.NewCmsLogger(confighelper.GetConfig())
var webLogger = logger.NewWebLogger(confighelper.GetConfig())

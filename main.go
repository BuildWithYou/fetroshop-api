package main

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/injector"
)

func main() {

	// TODO: check wether minio bucket exist or not

	// Run cms server
	go func() {
		cmsApp := injector.InitializeCmsServer()
		cmsApp.Logger.Info("Starting cms server...")
		err := cmsApp.FiberApp.Listen(fmt.Sprintf("%s:%d", cmsApp.Host, cmsApp.Port))
		if err != nil {
			cmsApp.Logger.Panic(err.Error())
		}
	}()

	// Run web server
	webApp := injector.InitializeWebServer()
	webApp.Logger.Info("Starting web server...")
	err := webApp.FiberApp.Listen(fmt.Sprintf("%s:%d", webApp.Host, webApp.Port))
	if err != nil {
		webApp.Logger.Panic(err.Error())
	}
}

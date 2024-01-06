package main

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/injector"
)

func main() {

	// Run cms server
	go func() {
		err := injector.InitializeCmsServer()
		errorhelper.PanicIfError(err)
	}()

	// Run web server
	err := injector.InitializeWebServer()
	errorhelper.PanicIfError(err)
}

package main

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/injector"
)

func main() {

	// Run cms server
	go func() {
		err := injector.InitializeCmsServer()
		helper.PanicIfError(err)
	}()

	// Run web server
	err := injector.InitializeWebServer()
	helper.PanicIfError(err)
}

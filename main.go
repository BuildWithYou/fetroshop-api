package main

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/injector"
)

func main() {

	go func() {
		err := injector.InitializeDocsServer()
		helper.PanicIfError(err)
	}()

	err := injector.InitializeWebServer()
	helper.PanicIfError(err)
}

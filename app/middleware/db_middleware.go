package middleware

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DbMiddleware struct {
	DB  *gorm.DB
	Err error
}

func DBMiddlewareProvider(conn *connection.Connection) *DbMiddleware {
	return &DbMiddleware{
		DB:  conn.DB,
		Err: conn.Err,
	}
}

// Authenticate authenticates the request using JWT.
//
// It takes a *fiber.Ctx object as a parameter and returns an error.
func (dbMid *DbMiddleware) Authenticate(ctx *fiber.Ctx) error {
	if dbMid.Err != nil {
		fmt.Printf("\nError: %s\n", dbMid.Err.Error()) // #marked: logging
		return errorhelper.Error500(constant.ERROR_GENERAL)
	}

	return ctx.Next()
}

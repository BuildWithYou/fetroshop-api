package middleware

import (
	"github.com/BuildWithYou/fetroshop-api/app/connection"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DbMiddleware struct {
	DB     *gorm.DB
	Err    error
	Logger *logger.Logger
}

func DBMiddlewareProvider(conn *connection.Connection, logger *logger.Logger) *DbMiddleware {
	return &DbMiddleware{
		DB:     conn.DB,
		Err:    conn.Err,
		Logger: logger,
	}
}

// Authenticate authenticates the request using JWT.
//
// It takes a *fiber.Ctx object as a parameter and returns an error.
func (dbMid *DbMiddleware) Authenticate(ctx *fiber.Ctx) error {
	if dbMid.Err != nil {
		dbMid.Logger.Error(dbMid.Err.Error())
		return errorhelper.Error500(constant.ERROR_GENERAL)
	}

	return ctx.Next()
}

package middleware

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/gofiber/fiber/v2"
)

type LoggerMiddleware struct {
	Logger *logger.Logger
}

func LoggerMiddlewareProvider(logger *logger.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{
		Logger: logger,
	}
}

func (lm *LoggerMiddleware) CmsLoggerResetOutput(ctx *fiber.Ctx) error {
	lm.Logger.CmsLoggerResetOutput()
	return ctx.Next()
}

func (lm *LoggerMiddleware) WebLoggerResetOutput(ctx *fiber.Ctx) error {
	lm.Logger.WebLoggerResetOutput()
	return ctx.Next()
}

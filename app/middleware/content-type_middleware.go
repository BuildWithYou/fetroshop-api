package middleware

import (
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/gofiber/fiber/v2"
)

func ContentTypeMiddleware(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if contentType == "" {
		return errorhelper.Error400("Content-Type header is missing")
	}

	// Check if the Content-Type is application/json
	if !strings.Contains(contentType, "application/json") && !strings.Contains(contentType, "application/x-www-form-urlencoded") {
		return errorhelper.ErrorCustom(fiber.StatusUnsupportedMediaType, "Unsupported Content-Type. Only application/json and application/x-www-form-urlencoded are allowed")
	}

	// If everything is fine, move on to the next middleware or route handler
	return c.Next()
}

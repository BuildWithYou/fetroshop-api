package helper

import (
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func Error500(ctx *fiber.Ctx, err error) error {
	ctx.Status(fiber.StatusInternalServerError)
	return ctx.JSON(model.GeneralResponse{
		Code:    fiber.ErrInternalServerError.Code,
		Status:  fiber.ErrInternalServerError.Message,
		Message: err.Error(),
	})

}

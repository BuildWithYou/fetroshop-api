package responsehelper

import (
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func SendResponse(ctx *fiber.Ctx, httpStatus int, md appModel.Response) error {
	return ctx.Status(httpStatus).JSON(appModel.Response{
		Code:    md.Code,
		Status:  md.Status,
		Message: md.Message,
	})
}

func SendResponse200(ctx *fiber.Ctx, md appModel.Response) error {
	return ctx.Status(fiber.StatusOK).JSON(appModel.Response{
		Code:    md.Code,
		Status:  md.Status,
		Message: md.Message,
	})
}

func SendResponse201(ctx *fiber.Ctx, md appModel.Response) error {
	return ctx.Status(fiber.StatusCreated).JSON(appModel.Response{
		Code:    md.Code,
		Status:  md.Status,
		Message: md.Message,
	})
}

func SendResponse400(ctx *fiber.Ctx, md appModel.Response) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(appModel.Response{
		Code:    md.Code,
		Status:  md.Status,
		Message: md.Message,
	})
}

func SendResponse401(ctx *fiber.Ctx, md appModel.Response) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(appModel.Response{
		Code:    md.Code,
		Status:  md.Status,
		Message: md.Message,
	})
}

func SendResponse500(ctx *fiber.Ctx, md appModel.Response) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(appModel.Response{
		Code:    md.Code,
		Status:  md.Status,
		Message: md.Message,
	})
}

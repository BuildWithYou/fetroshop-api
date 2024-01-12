package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *CategoryServiceImpl) Find(ctx *fiber.Ctx) (*model.Response, error) {
	// TODO: implement this
	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Unimplemented", // #marked: message
	}, nil
}

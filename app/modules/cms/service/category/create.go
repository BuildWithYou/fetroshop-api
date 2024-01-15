package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *CategoryServiceImpl) Create(ctx *fiber.Ctx) (*appModel.Response, error) {
	payload := new(model.UpsertCategoryRequest)
	err := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if err != nil {
		return nil, err
	}

	// TODO: implement me
	svc.Logger.Info("halo bosku")

	return &appModel.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: "Unimplemented", // #marked: message
	}, nil
}

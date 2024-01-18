package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *CategoryServiceImpl) Update(ctx *fiber.Ctx) (*model.Response, error) {
	pathPayload := new(model.CategoryPathRequest)
	errValidation, errParsing := validatorhelper.ValidateParamPayload(ctx, svc.Validate, pathPayload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	bodyPayload := new(model.UpsertCategoryRequest)
	errValidation, errParsing = validatorhelper.ValidateBodyPayload(ctx, svc.Validate, bodyPayload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// TODO: implement me

	return &model.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Unimplemented", // #marked: message
		Data: fiber.Map{
			"path": pathPayload,
			"body": bodyPayload,
		},
	}, nil
}

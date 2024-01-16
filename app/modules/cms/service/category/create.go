package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"gopkg.in/guregu/null.v3"
)

func (svc *CategoryServiceImpl) Create(ctx *fiber.Ctx) (*appModel.Response, error) {
	payload := new(model.UpsertCategoryRequest)
	errorMap, err := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if err != nil {
		return responsehelper.Response500(constant.ERROR_GENERAL, fiber.Map{"message": err.Error()}), nil
	}
	if errorMap != nil {
		return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"messages": errorMap}), nil
	}

	var (
		parentID null.Int
	)

	if payload.ParentCode != "" {
		parentCategory := new(categories.Category)
		result := svc.CategoryRepo.Find(parentCategory, fiber.Map{"code": payload.ParentCode})
		if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
			return nil, result.Error
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"message": "Parent category not found"}), nil
		}
	}

	svc.CategoryRepo.Create(&categories.Category{
		ParentID: parentID,
	})

	return &appModel.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: "Unimplemented", // #marked: message
	}, nil
}

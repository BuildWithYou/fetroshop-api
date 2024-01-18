package brand

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v3"
)

func (svc *brandService) Update(ctx *fiber.Ctx) (*model.Response, error) {
	//  TODO: implement me
	// parse param
	pathPayload := new(model.CategoryPathRequest)
	errValidation, errParsing := validatorhelper.ValidateParamPayload(ctx, svc.Validate, pathPayload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// parse body
	bodyPayload := new(model.UpsertCategoryRequest)
	errValidation, errParsing = validatorhelper.ValidateBodyPayload(ctx, svc.Validate, bodyPayload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// check category exists
	category := new(categories.Category)
	result := svc.CategoryRepo.Find(category, fiber.Map{"code": pathPayload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category not found"}), nil // #marked: message
	}

	// check code is unique
	result = svc.CategoryRepo.Find(&categories.Category{}, fiber.Map{
		"code": bodyPayload.Code,
		"id":   []any{"!=", category.ID},
	})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category code already used"}), nil // #marked: message
	}

	// update category
	category.Code = bodyPayload.Code
	category.Name = bodyPayload.Name
	category.Icon = null.NewString(bodyPayload.Icon, bodyPayload.Icon != "")
	category.IsActive = *bodyPayload.IsActive
	category.DisplayOrder = bodyPayload.DisplayOrder
	result = svc.CategoryRepo.Update(category,
		fiber.Map{"id": category.ID},
	)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to update category", nil), nil // #marked: message
	}

	return responsehelper.Response201(
		"Category updated successfully", // #marked: message
		category,                        // TODO: data return must be filtered
		nil), nil
}

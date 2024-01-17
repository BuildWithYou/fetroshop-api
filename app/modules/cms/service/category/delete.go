package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *CategoryServiceImpl) Delete(ctx *fiber.Ctx) (*appModel.Response, error) {
	pathPayload := new(model.CategoryPathRequest)
	errValidation, errParsing := validatorhelper.ValidateParamPayload(ctx, svc.Validate, pathPayload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	bodyPayload := new(model.DeleteCategoryRequest)
	errValidation, errParsing = validatorhelper.ValidateBodyPayload(ctx, svc.Validate, bodyPayload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	forceDelete := *bodyPayload.ForceDelete
	category := new(categories.Category)
	result := svc.CategoryRepo.Find(category, fiber.Map{"code": pathPayload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category not found"}), nil // #marked: message
	}

	// TODO: check wether the category has dependent products
	children := new(categories.Category)
	result = svc.CategoryRepo.Find(children, fiber.Map{"parent_id": category.ID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		if !forceDelete {
			return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category has children"}), nil // #marked: message
		}
		result = svc.CategoryRepo.Delete(map[string]any{"parent_id": category.ID})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			return nil, result.Error
		}
		if !gormhelper.HasAffectedRows(result) {
			return responsehelper.Response500("Failed to delete category children", nil), nil // #marked: message
		}
	}

	result = svc.CategoryRepo.Delete(map[string]any{"id": category.ID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to delete category", nil), nil // #marked: message
	}

	return responsehelper.Response200("Category deleted successfully", nil, nil), nil // #marked: message
}

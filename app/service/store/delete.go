package store

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *storeService) Delete(ctx *fiber.Ctx) (*model.Response, error) {
	// TODO: implement me
	pathPayload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateParamPayload(ctx, svc.Validate, pathPayload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	bodyPayload := new(model.DeleteRequest)
	errValidation, errParsing = validatorhelper.ValidateBodyPayload(ctx, svc.Validate, bodyPayload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// find category
	forceDelete := *bodyPayload.ForceDelete
	category := new(stores.Store)
	result := svc.StoreRepo.Find(category, fiber.Map{"code": pathPayload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category not found"}), nil // #marked: message
	}

	// TODO: check wether the category has dependent products
	children := new(stores.Store)
	result = svc.StoreRepo.Find(children, fiber.Map{"parent_id": category.ID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		if !forceDelete {
			return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category has children"}), nil // #marked: message
		}
		result = svc.StoreRepo.Delete(map[string]any{"parent_id": category.ID})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			svc.Logger.UseError(result.Error)
			return nil, result.Error
		}
		if !gormhelper.HasAffectedRows(result) {
			return responsehelper.Response500("Failed to delete category children", nil), nil // #marked: message
		}
	}

	result = svc.StoreRepo.Delete(map[string]any{"id": category.ID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to delete category", nil), nil // #marked: message
	}

	return responsehelper.Response200("Category deleted successfully", nil, nil), nil // #marked: message
}

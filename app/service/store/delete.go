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
	pathPayload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateParamPayload(ctx, svc.Validate, pathPayload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	queryPayload := new(model.DeleteRequest)
	errValidation, errParsing = validatorhelper.ValidateQueryPayload(ctx, svc.Validate, queryPayload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// find store
	_ = *queryPayload.ForceDelete // TODO: implement force delete in case store has products or transactions
	store := new(stores.Store)
	result := svc.StoreRepo.Find(store, fiber.Map{"code": pathPayload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Store not found"}), nil // #marked: message
	}

	// TODO: check wether the store has dependent products

	result = svc.StoreRepo.Delete(map[string]any{"id": store.ID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to delete store", nil), nil // #marked: message
	}

	return responsehelper.Response200("Store deleted successfully", nil, nil), nil // #marked: message
}

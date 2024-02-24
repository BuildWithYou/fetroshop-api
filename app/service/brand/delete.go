package brand

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *brandService) Delete(ctx *fiber.Ctx) (*model.Response, error) {
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

	// find brand
	_ = *queryPayload.ForceDelete // TODO: implement force delete in case brands has products
	brand := new(brands.Brand)
	result := svc.BrandRepo.Find(brand, fiber.Map{"code": pathPayload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Brand not found"}), nil // #marked: message
	}

	// TODO: check wether the brand has dependent products

	result = svc.BrandRepo.Delete(map[string]any{"id": brand.ID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to delete brand", nil), nil // #marked: message
	}

	return responsehelper.Response200("Brand deleted successfully", nil, nil), nil // #marked: message
}

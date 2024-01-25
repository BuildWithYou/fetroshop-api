package location

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *locationService) FindProvince(ctx *fiber.Ctx) (*model.Response, error) {
	// TODO: implement me
	payload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	category := new(stores.Store)
	result := svc.StoreRepo.Find(category, map[string]any{"code": payload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Invalid category code"}), nil // #marked: message
	}

	return responsehelper.Response200(
		"Successfuly got category", // #marked: message
		&model.CategoryResponse{
			Code:      category.Code,
			Name:      category.Name,
			IsActive:  category.IsActive,
			Icon:      category.Icon,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}, nil), nil
}

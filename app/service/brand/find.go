package brand

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *brandService) Find(ctx *fiber.Ctx) (*model.Response, error) {
	payload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	brand := new(brands.Brand)
	result := svc.BrandRepo.Find(brand, map[string]any{"code": payload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Invalid brand code"}), nil // #marked: message
	}

	return &model.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Successfuly got brand", // #marked: message
		Data: &model.BrandResponse{
			Code:      brand.Code,
			Name:      brand.Name,
			IsActive:  brand.IsActive,
			Icon:      brand.Icon,
			CreatedAt: brand.CreatedAt,
			UpdatedAt: brand.UpdatedAt,
		},
	}, nil
}

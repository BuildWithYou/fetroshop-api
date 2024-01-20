package brand

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v3"
)

func (svc *brandService) Update(ctx *fiber.Ctx) (*model.Response, error) {
	// parse param
	pathPayload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateParamPayload(ctx, svc.Validate, pathPayload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// parse body
	bodyPayload := new(model.UpsertBrandRequest)
	errValidation, errParsing = validatorhelper.ValidateBodyPayload(ctx, svc.Validate, bodyPayload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// check brand exists
	brand := new(brands.Brand)
	result := svc.BrandRepo.Find(brand, fiber.Map{"code": pathPayload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Brand not found"}), nil // #marked: message
	}

	// check code is unique
	result = svc.BrandRepo.Find(&brands.Brand{}, fiber.Map{
		"code": bodyPayload.Code,
		"id":   []any{"!=", brand.ID},
	})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Brand code already used"}), nil // #marked: message
	}

	// update brand
	brand.Code = bodyPayload.Code
	brand.Name = bodyPayload.Name
	brand.Icon = null.NewString(bodyPayload.Icon, bodyPayload.Icon != "")
	brand.IsActive = *bodyPayload.IsActive
	result = svc.BrandRepo.Update(brand,
		fiber.Map{"id": brand.ID},
	)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		svc.Logger.Error("Failed to update brand")
		return responsehelper.Response500("Failed to update brand", nil), nil // #marked: message
	}

	return responsehelper.Response201(
		"Brand updated successfully", // #marked: message
		model.BrandResponse{
			Code:      brand.Code,
			Name:      brand.Name,
			IsActive:  brand.IsActive,
			Icon:      brand.Icon,
			CreatedAt: brand.CreatedAt,
			UpdatedAt: brand.UpdatedAt,
		},
		nil), nil
}

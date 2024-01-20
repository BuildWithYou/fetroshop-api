package brand

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/brands"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"gopkg.in/guregu/null.v3"
)

func (svc *brandService) Create(ctx *fiber.Ctx) (*model.Response, error) {
	payload := new(model.UpsertBrandRequest)
	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	var (
		code     string
		name     string
		isActive bool
		icon     null.String
	)

	code = slug.Make(payload.Code)
	name = payload.Name
	isActive = *payload.IsActive
	icon = null.NewString(payload.Icon, payload.Icon != "")

	// check code is unique
	brandByCode := new(brands.Brand)
	result := svc.BrandRepo.Find(brandByCode, map[string]any{"code": code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category code has been taken"}), nil // #marked: message
	}

	// create new brand
	newBrand := &brands.Brand{
		Code:     code,
		Name:     name,
		IsActive: isActive,
		Icon:     icon,
	}
	result = svc.BrandRepo.Create(newBrand)
	if result.Error != nil && !gormhelper.IsErrDuplicatedKey(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrDuplicatedKey(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Brand code has been taken"}), nil // #marked: message
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to create brand", nil), nil // #marked: message
	}
	return responsehelper.Response201(
		"Brand created successfully", // #marked: message
		model.BrandResponse{
			Code:      newBrand.Code,
			Name:      newBrand.Name,
			IsActive:  newBrand.IsActive,
			Icon:      newBrand.Icon,
			CreatedAt: newBrand.CreatedAt,
			UpdatedAt: newBrand.UpdatedAt,
		},
		nil), nil
}

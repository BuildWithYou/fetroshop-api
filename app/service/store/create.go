package store

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"gopkg.in/guregu/null.v3"
)

func (svc *storeService) Create(ctx *fiber.Ctx) (*model.Response, error) {
	// TODO: implement me
	// parse body
	payload := new(model.UpsertCategoryRequest)
	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	var (
		parentCode   null.String
		code         string
		name         string
		isActive     bool
		icon         null.String
		displayOrder int64
	)

	code = slug.Make(payload.Code)
	name = payload.Name
	isActive = *payload.IsActive
	displayOrder = payload.DisplayOrder
	icon = null.NewString(payload.Icon, payload.Icon != "")

	// check parent category exists
	if payload.ParentCode != "" {
		parentCategory := new(stores.Store)
		result := svc.StoreRepo.Find(parentCategory, map[string]any{"code": payload.ParentCode})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			svc.Logger.UseError(result.Error)
			return nil, result.Error
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return responsehelper.ResponseErrorValidation(fiber.Map{"parentCode": "Invalid parent category code"}), nil // #marked: message
		}
		parentCode = null.StringFrom(parentCategory.Code)
	}

	// check display order is unique
	categoryByDisplayOrder := new(stores.Store)
	result := svc.StoreRepo.Find(categoryByDisplayOrder, map[string]any{"display_order": displayOrder})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"displayOrder": "Display order has been taken"}), nil // #marked: message
	}

	// check code is unique
	categoryByCode := new(stores.Store)
	result = svc.StoreRepo.Find(categoryByCode, map[string]any{"code": code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category code has been taken"}), nil // #marked: message
	}

	// create new category
	newCategory := &stores.Store{
		Code:     code,
		Name:     name,
		IsActive: isActive,
		Icon:     icon,
	}
	result = svc.StoreRepo.Create(newCategory)
	if result.Error != nil && !gormhelper.IsErrDuplicatedKey(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrDuplicatedKey(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category code has been taken"}), nil // #marked: message
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to create category", nil), nil // #marked: message
	}

	return responsehelper.Response201(
		"Category created successfully", // #marked: message
		model.CategoryResponse{
			Code:       newCategory.Code,
			ParentCode: parentCode,
			Name:       newCategory.Name,
			IsActive:   newCategory.IsActive,
			Icon:       newCategory.Icon,
			CreatedAt:  newCategory.CreatedAt,
			UpdatedAt:  newCategory.UpdatedAt,
		},
		nil), nil
}

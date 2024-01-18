package brand

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"gopkg.in/guregu/null.v3"
)

func (svc *brandService) Create(ctx *fiber.Ctx) (*model.Response, error) {
	//  TODO: implement me
	payload := new(model.UpsertCategoryRequest)
	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	var (
		parentID     null.Int
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
		parentCategory := new(categories.Category)
		result := svc.CategoryRepo.Find(parentCategory, map[string]any{"code": payload.ParentCode})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			return nil, result.Error
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return responsehelper.ResponseErrorValidation(fiber.Map{"parentCode": "Invalid parent category code"}), nil // #marked: message
		}
		parentID = null.IntFrom(parentCategory.ID)
	}

	// check display order is unique
	categoryByDisplayOrder := new(categories.Category)
	result := svc.CategoryRepo.Find(categoryByDisplayOrder, map[string]any{"display_order": displayOrder})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"displayOrder": "Display order has been taken"}), nil // #marked: message
	}

	// check code is unique
	categoryByCode := new(categories.Category)
	result = svc.CategoryRepo.Find(categoryByCode, map[string]any{"code": code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category code has been taken"}), nil // #marked: message
	}

	newCategory := &categories.Category{
		ParentID:     parentID,
		Code:         code,
		Name:         name,
		IsActive:     isActive,
		Icon:         icon,
		DisplayOrder: displayOrder,
	}
	result = svc.CategoryRepo.Create(newCategory)
	if result.Error != nil && !gormhelper.IsErrDuplicatedKey(result.Error) {
		return nil, result.Error
	}
	if gormhelper.IsErrDuplicatedKey(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category code has been taken"}), nil // #marked: message
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to create category", nil), nil // #marked: message
	}
	return responsehelper.Response201("Category created successfully", newCategory, nil), nil // #marked: message

}

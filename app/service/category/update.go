package category

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

func (svc *categoryService) Update(ctx *fiber.Ctx) (*model.Response, error) {
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
	bodyPayload := new(model.UpsertCategoryRequest)
	errValidation, errParsing = validatorhelper.ValidateBodyPayload(ctx, svc.Validate, bodyPayload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	var (
		parentID     null.Int
		parentCode   null.String
		code         string
		name         string
		isActive     bool
		icon         null.String
		displayOrder int64
	)

	code = slug.Make(bodyPayload.Code)
	name = bodyPayload.Name
	isActive = *bodyPayload.IsActive
	displayOrder = bodyPayload.DisplayOrder
	icon = null.NewString(bodyPayload.Icon, bodyPayload.Icon != "")

	// check category exists
	category := new(categories.Category)
	result := svc.CategoryRepo.Find(category, fiber.Map{"code": pathPayload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category not found"}), nil // #marked: message
	}

	// check parent category exists
	if bodyPayload.ParentCode != "" {
		parentCategory := new(categories.Category)
		result := svc.CategoryRepo.Find(parentCategory, map[string]any{"code": bodyPayload.ParentCode})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			svc.Logger.UseError(result.Error)
			return nil, result.Error
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return responsehelper.ResponseErrorValidation(fiber.Map{"parentCode": "Invalid parent category code"}), nil // #marked: message
		}
		parentID = null.IntFrom(parentCategory.ID)
		parentCode = null.StringFrom(parentCategory.Code)
	}

	// check display order is unique
	categoryByDisplayOrder := new(categories.Category)
	result = svc.CategoryRepo.Find(categoryByDisplayOrder, map[string]any{
		"display_order": displayOrder,
		"id":            []any{"!=", category.ID},
	})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"displayOrder": "Display order has been taken"}), nil // #marked: message
	}

	// check code is unique
	result = svc.CategoryRepo.Find(&categories.Category{}, fiber.Map{
		"code": code,
		"id":   []any{"!=", category.ID},
	})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Category code already used"}), nil // #marked: message
	}

	// update category
	category.ParentID = parentID
	category.Code = code
	category.Name = name
	category.Icon = icon
	category.IsActive = isActive
	category.DisplayOrder = displayOrder
	result = svc.CategoryRepo.Update(category,
		fiber.Map{"id": category.ID},
	)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to update category", nil), nil // #marked: message
	}

	return responsehelper.Response201(
		"Category updated successfully", // #marked: message
		model.CategoryResponse{
			Code:         category.Code,
			ParentCode:   parentCode,
			Name:         category.Name,
			IsActive:     category.IsActive,
			Icon:         category.Icon,
			DisplayOrder: category.DisplayOrder,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
		},
		nil), nil
}

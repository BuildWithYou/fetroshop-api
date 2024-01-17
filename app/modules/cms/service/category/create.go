package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/gosimple/slug"
	"gopkg.in/guregu/null.v3"
)

func (svc *CategoryServiceImpl) Create(ctx *fiber.Ctx) (*appModel.Response, error) {
	payload := new(model.UpsertCategoryRequest)
	errorMap, err := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if err != nil {
		return responsehelper.Response500(constant.ERROR_GENERAL, fiber.Map{"message": err.Error()}), nil
	}
	if errorMap != nil {
		return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"messages": errorMap}), nil
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
	isActive = payload.IsActive
	displayOrder = payload.DisplayOrder

	if payload.Icon != "" {
		icon = null.StringFrom(payload.Icon)
	}

	// check parent category exists
	if payload.ParentCode != "" {
		parentCategory := new(categories.Category)
		result := svc.CategoryRepo.Find(parentCategory, map[string]any{"code": payload.ParentCode})
		if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
			return nil, result.Error
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"message": "Parent category not found"}), nil
		}
		parentID = null.IntFrom(parentCategory.ID)
	}

	// check display order is unique
	categoryByDisplayOrder := new(categories.Category)
	result := svc.CategoryRepo.Find(categoryByDisplayOrder, map[string]any{"display_order": displayOrder})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"message": "Display order has been taken"}), nil
	}

	// check code is unique
	categoryByCode := new(categories.Category)
	result = svc.CategoryRepo.Find(categoryByCode, map[string]any{"code": code})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"message": "Code has been taken"}), nil
	}

	createdCategory := &categories.Category{
		ParentID:     parentID,
		Code:         code,
		Name:         name,
		IsActive:     isActive,
		Icon:         icon,
		DisplayOrder: displayOrder,
	}
	result = svc.CategoryRepo.Create(createdCategory)
	if gormhelper.HasAffectedRows(result) {
		return &appModel.Response{
			Code:    fiber.StatusCreated,
			Status:  utils.StatusMessage(fiber.StatusCreated),
			Message: "Category created successfully",
			Data:    createdCategory,
		}, nil
	} else {
		return &appModel.Response{
			Code:    fiber.StatusInternalServerError,
			Status:  utils.StatusMessage(fiber.StatusInternalServerError),
			Message: "Failed to create category",
			Data:    createdCategory,
		}, nil
	}

}

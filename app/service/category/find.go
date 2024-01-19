package category

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v3"
)

func (svc *categoryService) Find(ctx *fiber.Ctx) (*model.Response, error) {
	payload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	category := new(categories.Category)
	result := svc.CategoryRepo.Find(category, map[string]any{"code": payload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Invalid category code"}), nil // #marked: message
	}

	parentCode := ""
	if category.Parent != nil {
		parentCode = category.Parent.Code
	}

	return responsehelper.Response200(
		"Successfuly got category", // #marked: message
		&model.CategoryResponse{
			Code:         category.Code,
			ParentCode:   null.NewString(parentCode, parentCode != ""),
			Name:         category.Name,
			IsActive:     category.IsActive,
			Icon:         category.Icon,
			DisplayOrder: category.DisplayOrder,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
		}, nil), nil
}

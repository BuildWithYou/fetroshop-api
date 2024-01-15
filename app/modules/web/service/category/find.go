package category

import (
	ctEty "github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"gopkg.in/guregu/null.v3"
)

func (svc *CategoryServiceImpl) Find(ctx *fiber.Ctx) (*appModel.Response, error) {
	payload := new(model.FindCategoryRequest)
	errorMap, err := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if err != nil {
		return responsehelper.Response500(constant.ERROR_GENERAL, nil, map[string]string{"message": err.Error()}), nil
	}
	if errorMap != nil {
		return responsehelper.Response400(constant.ERROR_VALIDATION, nil, errorMap), nil
	}

	category := new(ctEty.Category)
	result := svc.CategoryRepo.Find(category, map[string]any{
		"code": payload.Code,
	})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.Error(result.Error.Error())
		return nil, errorhelper.Error500("Something went wrong") // #marked: message
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Invalid category code") // #marked: message
	}

	parentCode := ""
	if category.Parent != nil {
		parentCode = category.Parent.Code
	}

	return &appModel.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Successfuly got category", // #marked: message
		Data: &model.CategoryResponse{
			Code:         category.Code,
			ParentCode:   null.NewString(parentCode, parentCode != ""),
			Name:         category.Name,
			IsActive:     category.IsActive,
			Icon:         category.Icon,
			DisplayOrder: category.DisplayOrder,
			CreatedAt:    category.CreatedAt,
			UpdatedAt:    category.UpdatedAt,
		},
	}, nil
}

package category

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v3"
)

func (svc *categoryService) List(ctx *fiber.Ctx) (*model.Response, error) {
	var (
		categorySlice []categories.Category
		parentID      null.Int
	)
	payload := new(model.ListCategoriesRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	if payload.ParentCode == "" {
		parentID = null.NewInt(0, false)
	} else {
		parent := new(categories.Category)
		result := svc.CategoryRepo.Find(parent, map[string]any{
			"code": payload.ParentCode,
		})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			svc.Logger.UseError(result.Error)
			return nil, result.Error
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return responsehelper.ResponseErrorValidation(fiber.Map{"parentCode": "Invalid parent category code"}), nil // #marked: message
		}
		parentID = null.NewInt(parent.ID, parent.ID != 0)
	}

	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)
	result := svc.CategoryRepo.List(&categorySlice, fiber.Map{
		"parent_id": parentID,
	}, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	var list []*model.CategoryResponse
	for _, ct := range categorySlice {
		parentCode := ""
		if ct.Parent != nil {
			parentCode = ct.Parent.Code
		}

		category := &model.CategoryResponse{
			Code:         ct.Code,
			ParentCode:   null.NewString(parentCode, parentCode != ""),
			Name:         ct.Name,
			IsActive:     ct.IsActive,
			Icon:         ct.Icon,
			DisplayOrder: ct.DisplayOrder,
			CreatedAt:    ct.CreatedAt,
			UpdatedAt:    ct.UpdatedAt,
		}
		list = append(list, category)
	}

	return responsehelper.Response200(
		"Successfuly got list of categories", // #marked: message
		list,
		fiber.Map{"total": result.RowsAffected},
	), nil
}

package category

import (
	"fmt"

	ctEty "github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"gopkg.in/guregu/null.v3"
)

func (svc *CategoryServiceImpl) List(ctx *fiber.Ctx) (*model.Response, error) {
	var (
		categories []ctEty.Category
		parentID   null.Int
	)
	payload := new(model.ListCategoriesRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	if payload.ParentCode == "" {
		parentID = null.NewInt(0, false)
	} else {
		parent := new(ctEty.Category)
		result := svc.CategoryRepo.Find(parent, map[string]any{
			"code": payload.ParentCode,
		})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			return nil, result.Error
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return responsehelper.ResponseErrorValidation(fiber.Map{"parentCode": "Invalid parent category code"}), nil // #marked: message
		}
		parentID = null.NewInt(parent.ID, parent.ID != 0)
	}

	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)
	result := svc.CategoryRepo.List(&categories, fiber.Map{
		"parent_id": parentID,
	}, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}

	var list []*model.CategoryResponse
	for _, ct := range categories {
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

	return &model.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Successfuly got list of categories", // #marked: message
		Data:    list,
	}, nil
}

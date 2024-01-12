package category

import (
	"fmt"

	ctEty "github.com/BuildWithYou/fetroshop-api/app/domain/categories"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"gopkg.in/guregu/null.v3"
)

func (svc *CategoryServiceImpl) List(ctx *fiber.Ctx) (*appModel.Response, error) {
	var (
		categories []ctEty.Category
		parentID   null.Int
	)
	payload := new(model.ListCategoriesRequest)
	validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)

	parent := new(ctEty.Category)
	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)
	if payload.ParentCode == "" {
		// when parent code is empty string
		result := svc.CategoryRepo.List(&categories, map[string]any{
			"parent_id": null.NewInt(0, false),
		}, int(payload.Limit), int(payload.Offset), orderBy)
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			return nil, errorhelper.Error500("Something went wrong") // #marked: message
		}
	} else {
		// when parent code is not empty
		result := svc.CategoryRepo.Find(parent, map[string]any{
			"code": payload.ParentCode,
		})
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			return nil, errorhelper.Error500("Something went wrong") // #marked: message
		}
		if gormhelper.IsErrRecordNotFound(result.Error) {
			return nil, errorhelper.Error400("Invalid parent code") // #marked: message
		}

		if validatorhelper.IsNotZero(parent.ID) {
			parentID = null.NewInt(parent.ID, true)
		}
		result = svc.CategoryRepo.List(&categories, map[string]any{
			"parent_id": parentID,
		}, int(payload.Limit), int(payload.Offset), orderBy)
		if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
			return nil, errorhelper.Error500("Something went wrong") // #marked: message
		}
	}

	var list []*model.CategoryResponse
	for _, ct := range categories {
		parentCode := ""
		if ct.Parent != nil {
			parentCode = ct.Parent.Code
		}

		category := &model.CategoryResponse{
			Code:         ct.Code,
			ParentCode:   null.NewString(parentCode, false),
			Name:         ct.Name,
			IsActive:     ct.IsActive,
			Icon:         ct.Icon,
			DisplayOrder: ct.DisplayOrder,
			CreatedAt:    ct.CreatedAt,
			UpdatedAt:    ct.UpdatedAt,
		}
		list = append(list, category)
	}

	return &appModel.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Successfuly got list of categories", // #marked: message
		Data: map[string]any{
			"list":     list,
			"original": categories,
		},
	}, nil
}

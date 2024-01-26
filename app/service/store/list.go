package store

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/guregu/null.v3"
)

func (svc *storeService) List(ctx *fiber.Ctx) (*model.Response, error) {
	// TODO: implement me
	var (
		categorySlice []stores.Store
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
		parent := new(stores.Store)
		result := svc.StoreRepo.Find(parent, map[string]any{
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
	result := svc.StoreRepo.List(&categorySlice, fiber.Map{
		"parent_id": parentID,
	}, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	var list []*model.CategoryResponse
	for _, ct := range categorySlice {
		category := &model.CategoryResponse{
			Code:      ct.Code,
			Name:      ct.Name,
			IsActive:  ct.IsActive,
			Icon:      ct.Icon,
			CreatedAt: ct.CreatedAt,
			UpdatedAt: ct.UpdatedAt,
		}
		list = append(list, category)
	}

	return responsehelper.Response200(
		"Successfuly got list of categories", // #marked: message
		list,
		fiber.Map{"total": result.RowsAffected},
	), nil
}

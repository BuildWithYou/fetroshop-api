package location

import (
	"fmt"
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *locationService) ListProvinces(ctx *fiber.Ctx) (*model.Response, error) {
	var (
		provinceSlice             []provinces.Province
		selected, filtered, total int64
	)

	payload := new(model.ProvinceListRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	condition := fiber.Map{
		"UPPER(name)": []any{"like", fmt.Sprintf("%%%s%%", strings.ToUpper(payload.Name))},
	}
	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)

	// retrive data
	result := svc.ProvinceRepo.List(&provinceSlice, condition, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	selected = result.RowsAffected

	var list []*model.ProvinceDetail
	for _, ct := range provinceSlice {
		category := &model.ProvinceDetail{
			ID:        ct.ID,
			Name:      ct.Name,
			CreatedAt: ct.CreatedAt,
			UpdatedAt: ct.UpdatedAt,
		}
		list = append(list, category)
	}

	// count filtered
	result = svc.ProvinceRepo.Count(&filtered, condition)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	// count total
	result = svc.ProvinceRepo.Count(&total, fiber.Map{})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	return responsehelper.Response200(
		"Successfuly got list of categories", // #marked: message
		list,
		fiber.Map{
			"selected": selected,
			"filtered": filtered,
			"total":    total,
		},
	), nil
}
package location

import (
	"fmt"
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/districts"
	"github.com/BuildWithYou/fetroshop-api/app/domain/subdistricts"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *locationService) ListSubdistricts(ctx *fiber.Ctx) (*model.Response, error) {
	var (
		subdistrictSlice          []subdistricts.Subdistrict
		selected, filtered, total int64
	)

	payload := new(model.SubdistrictListRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	district := new(districts.District)
	result := svc.DistrictRepo.Find(district, fiber.Map{"id": payload.DistrictID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"districtId": "District not found"}), nil
	}

	condition := fiber.Map{
		"district_id": payload.DistrictID,
		"UPPER(name)": []any{"like", fmt.Sprintf("%%%s%%", strings.ToUpper(payload.Name))},
	}
	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)

	// retrieve data
	result = svc.SubdistrictRepo.List(&subdistrictSlice, condition, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	selected = result.RowsAffected

	var list []*model.Location
	for _, ct := range subdistrictSlice {
		category := &model.Location{
			ID:   ct.ID,
			Name: ct.Name,
		}
		list = append(list, category)
	}

	// count filtered
	result = svc.SubdistrictRepo.Count(&filtered, condition)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	// count total
	result = svc.SubdistrictRepo.Count(&total, fiber.Map{})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	return responsehelper.Response200(
		"Successfuly got list of subdistricts", // #marked: message
		list,
		fiber.Map{
			"selected": selected,
			"filtered": filtered,
			"total":    total,
		},
	), nil
}

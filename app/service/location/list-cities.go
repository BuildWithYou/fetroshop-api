package location

import (
	"fmt"
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *locationService) ListCities(ctx *fiber.Ctx) (*model.Response, error) {
	var (
		citySlice                 []cities.City
		selected, filtered, total int64
	)

	payload := new(model.CityListRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	province := new(provinces.Province)
	result := svc.ProvinceRepo.Find(province, fiber.Map{"id": payload.ProvinceID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"provinceId": "Province not found"}), nil
	}

	condition := fiber.Map{
		"province_id": payload.ProvinceID,
		"UPPER(name)": []any{"like", fmt.Sprintf("%%%s%%", strings.ToUpper(payload.Name))},
	}
	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)

	// retrieve data
	result = svc.CityRepo.List(&citySlice, condition, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	selected = result.RowsAffected

	var list []*model.IDName
	for _, ct := range citySlice {
		category := &model.IDName{
			ID:   ct.ID,
			Name: ct.Name,
		}
		list = append(list, category)
	}

	// count filtered
	result = svc.CityRepo.Count(&filtered, condition)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	// count total
	result = svc.CityRepo.Count(&total, fiber.Map{})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	return responsehelper.Response200(
		"Successfuly got list of cities", // #marked: message
		list,
		fiber.Map{
			"selected": selected,
			"filtered": filtered,
			"total":    total,
		},
	), nil
}

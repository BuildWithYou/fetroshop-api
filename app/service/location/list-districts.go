package location

import (
	"fmt"
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/domain/districts"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *locationService) ListDistricts(ctx *fiber.Ctx) (*model.Response, error) {
	var (
		districtSlice             []districts.District
		selected, filtered, total int64
	)

	payload := new(model.DistrictListRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	city := new(cities.City)
	result := svc.CityRepo.Find(city, fiber.Map{"id": payload.CityID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"cityId": "City not found"}), nil
	}

	condition := fiber.Map{
		"city_id":     payload.CityID,
		"UPPER(name)": []any{"like", fmt.Sprintf("%%%s%%", strings.ToUpper(payload.Name))},
	}
	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)

	// retrieve data
	result = svc.DistrictRepo.List(&districtSlice, condition, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	selected = result.RowsAffected

	var list []*model.IDName
	for _, ct := range districtSlice {
		category := &model.IDName{
			ID:   ct.ID,
			Name: ct.Name,
		}
		list = append(list, category)
	}

	// count filtered
	result = svc.DistrictRepo.Count(&filtered, condition)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	// count total
	result = svc.DistrictRepo.Count(&total, fiber.Map{})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	return responsehelper.Response200(
		"Successfuly got list of districts", // #marked: message
		list,
		fiber.Map{
			"selected": selected,
			"filtered": filtered,
			"total":    total,
		},
	), nil
}

package store

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *storeService) List(ctx *fiber.Ctx) (*model.Response, error) {
	var (
		storeSlice                []stores.Store
		selected, filtered, total int64
	)
	payload := new(model.StoreListRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	orderBy := fmt.Sprintf("%s %s", payload.OrderBy, payload.OrderDirection)
	result := svc.StoreRepo.List(&storeSlice, payload.Search, int(payload.Limit), int(payload.Offset), orderBy)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	// retrieve data
	var list []*model.StoreListData
	for _, st := range storeSlice {
		store := &model.StoreListData{
			Code:          st.Code,
			Name:          st.Name,
			IsActive:      st.IsActive,
			Icon:          st.Icon.Ptr(),
			Latitude:      st.Latitude.Ptr(),
			Longitude:     st.Longitude.Ptr(),
			Address:       st.Address,
			ProvinceID:    st.ProvinceID,
			CityID:        st.CityID,
			DistrictID:    st.DistrictID,
			SubdistrictID: st.SubdistrictID,
			PostalCode:    st.PostalCode,
		}
		list = append(list, store)
	}
	selected = result.RowsAffected

	// count filtered
	result = svc.StoreRepo.Count(&filtered, payload.Search)
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	// count total
	result = svc.StoreRepo.Count(&total, "")
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}

	return responsehelper.Response200(
		"Successfuly got list of stores", // #marked: message
		list,
		fiber.Map{
			"selected": selected,
			"filtered": filtered,
			"total":    total,
		},
	), nil
}

package store

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *storeService) Find(ctx *fiber.Ctx) (*model.Response, error) {
	payload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateParamPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	store := new(stores.Store)
	result := svc.StoreRepo.FindWithLocation(store, map[string]any{"code": payload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Invalid store code"}), nil // #marked: message
	}

	return responsehelper.Response200(
		"Successfuly got store", // #marked: message
		&model.StoreDetail{
			Code:      store.Code,
			Name:      store.Name,
			IsActive:  store.IsActive,
			Icon:      store.Icon.Ptr(),
			Latitude:  store.Latitude.Ptr(),
			Longitude: store.Longitude.Ptr(),
			Address:   store.Address,
			Province: model.Location{
				ID:   store.Province.ID,
				Name: store.Province.Name,
			},
			City: model.Location{
				ID:   store.City.ID,
				Name: store.City.Name,
			},
			District: model.Location{
				ID:   store.District.ID,
				Name: store.District.Name,
			},
			Subdistrict: model.Location{
				ID:   store.Subdistrict.ID,
				Name: store.Subdistrict.Name,
			},
			PostalCode: store.PostalCode,
		}, nil), nil
}

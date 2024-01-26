package location

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *locationService) FindProvince(ctx *fiber.Ctx) (*model.Response, error) {
	// TODO: implement me
	payload := new(model.FindByCodeRequest)
	errValidation, errParsing := validatorhelper.ValidateQueryPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	province := new(provinces.Province)
	result := svc.ProvinceRepo.Find(province, map[string]any{"code": payload.Code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Invalid province code"}), nil // #marked: message
	}

	return responsehelper.Response200(
		"Successfuly got province", // #marked: message
		nil, nil), nil
}

package store

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/cities"
	"github.com/BuildWithYou/fetroshop-api/app/domain/districts"
	"github.com/BuildWithYou/fetroshop-api/app/domain/provinces"
	"github.com/BuildWithYou/fetroshop-api/app/domain/stores"
	"github.com/BuildWithYou/fetroshop-api/app/domain/subdistricts"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gosimple/slug"
	"github.com/minio/minio-go/v7"
	"gopkg.in/guregu/null.v3"
)

func (svc *storeService) Create(ctx *fiber.Ctx) (*model.Response, error) {
	// parse body
	payload := new(model.UpsertStoreRequest)
	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	var (
		userID        int64
		code          string
		name          string
		isActive      bool
		icon          null.String
		latitude      null.String
		longitude     null.String
		address       string
		provinceID    int64
		cityID        int64
		districtID    int64
		subdistrictID int64
		postalCode    string
	)

	userID = jwt.GetUserID(ctx)
	code = slug.Make(payload.Code)
	name = payload.Name
	isActive = *payload.IsActive
	latitude = null.NewString(payload.Latitude, payload.Latitude != "")
	longitude = null.NewString(payload.Longitude, payload.Longitude != "")
	address = payload.Address
	provinceID = payload.ProvinceID
	cityID = payload.CityID
	districtID = payload.DistrictID
	subdistrictID = payload.SubdistrictID
	postalCode = payload.PostalCode

	// parse file input
	file, err := ctx.FormFile("icon")
	if err == nil {
		// get buffer
		buffer, err := file.Open()
		if err != nil {
			return responsehelper.ResponseErrorValidation(fiber.Map{"icon": err.Error()}), nil // #marked: generated message
		}
		defer buffer.Close()

		basePath := "store/icon/"
		fileName := slug.Make(file.Filename)
		filePath := fmt.Sprint(basePath, fileName)
		fileBuffer := buffer
		contentType := file.Header["Content-Type"][0]
		fileSize := file.Size

		// Upload the zip file with PutObject
		info, err := svc.Minio.Client.PutObject(ctx.Context(), svc.Minio.BucketName, filePath, fileBuffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			return responsehelper.Response500("Error on upload icon", fiber.Map{"icon": err.Error()}), nil // #marked: generated message
		}
		icon = null.NewString(info.Key, true)
	}

	// check user has store
	existingStore := new(stores.Store)
	result := svc.StoreRepo.Find(existingStore, fiber.Map{"user_id": userID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "User already has store"}), nil
	}

	// check code is unique
	storeByCode := new(stores.Store)
	result = svc.StoreRepo.Find(storeByCode, fiber.Map{"code": code})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Store code has been taken"}), nil // #marked: message
	}

	// check province is valid
	province := new(provinces.Province)
	result = svc.ProvinceRepo.Find(province, fiber.Map{"id": provinceID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"provinceId": "provinceId is invalid"}), nil // #marked: message
	}

	// check city is valid
	city := new(cities.City)
	result = svc.CityRepo.Find(city, fiber.Map{"province_id": provinceID, "id": cityID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"cityId": "cityId is invalid or not match with provinceId"}), nil // #marked: message
	}

	// check district is valid
	district := new(districts.District)
	result = svc.DistrictRepo.Find(district, fiber.Map{"city_id": cityID, "id": districtID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"districtId": "districtId is invalid or not match with cityId"}), nil // #marked: message
	}

	// check subdistrict is valid
	subdistrict := new(subdistricts.Subdistrict)
	result = svc.SubdistrictRepo.Find(subdistrict, fiber.Map{"district_id": districtID, "id": subdistrictID})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"subdistrictId": "subdistrictId is invalid or not match with districtId"}), nil // #marked: message
	}

	// create new store
	newStore := &stores.Store{
		UserID:        userID,
		Code:          code,
		Name:          name,
		IsActive:      isActive,
		Icon:          icon,
		Latitude:      latitude,
		Longitude:     longitude,
		Address:       address,
		ProvinceID:    provinceID,
		CityID:        cityID,
		DistrictID:    districtID,
		SubdistrictID: subdistrictID,
		PostalCode:    postalCode,
	}
	result = svc.StoreRepo.Create(newStore)
	if result.Error != nil && !gormhelper.IsErrDuplicatedKey(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if gormhelper.IsErrDuplicatedKey(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"code": "Store code has been taken"}), nil // #marked: message
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to create store", nil), nil // #marked: message
	}

	return responsehelper.Response201(
		"Store created successfully", // #marked: message
		model.StoreDetail{
			Code:      newStore.Code,
			Name:      newStore.Name,
			IsActive:  newStore.IsActive,
			Icon:      newStore.Icon.Ptr(),
			Latitude:  newStore.Latitude.Ptr(),
			Longitude: newStore.Longitude.Ptr(),
			Address:   newStore.Address,
			Province: model.Location{
				ID:   newStore.ProvinceID,
				Name: province.Name,
			},
			City: model.Location{
				ID:   newStore.CityID,
				Name: city.Name,
			},
			District: model.Location{
				ID:   newStore.DistrictID,
				Name: district.Name,
			},
			Subdistrict: model.Location{
				ID:   newStore.SubdistrictID,
				Name: subdistrict.Name,
			},
			PostalCode: newStore.PostalCode,
		},
		nil), nil
}

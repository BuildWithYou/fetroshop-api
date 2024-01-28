package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/location"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LocationController interface {
	ListProvinces(ctx *fiber.Ctx) (err error)
	FindProvince(ctx *fiber.Ctx) (err error)
	ListCities(ctx *fiber.Ctx) (err error)
	FindCity(ctx *fiber.Ctx) (err error)
	ListDistricts(ctx *fiber.Ctx) (err error)
	FindDistrict(ctx *fiber.Ctx) (err error)
	ListSubdistricts(ctx *fiber.Ctx) (err error)
	FindSubdistrict(ctx *fiber.Ctx) (err error)
}

type locationController struct {
	Validate        *validator.Validate
	LocationService location.LocationService
}

func LocationControllerProvider(vld *validator.Validate, svc location.LocationService) LocationController {
	return &locationController{
		Validate:        vld,
		LocationService: svc,
	}
}

// @Summary      List provinces
// @Description  Retrieve provinces list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ProvinceListRequest  true  "Request"
// @Success      200  {object}  model.ProvinceListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/province/list [get]
func (ctr *locationController) ListProvinces(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.ListProvinces)
}

// @Summary      Get detail province
// @Description  Retrieve province detail
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.FindByCodeRequest  true  "Request"
// @Success      200  {object}  model.StoreDetailResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/province/detail [get]
func (ctr *locationController) FindProvince(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.FindProvince)
}

// @Summary      List cities
// @Description  Retrieve cities list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.StoresListRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/city/list [get]
func (ctr *locationController) ListCities(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.ListCities)
}

// @Summary      Get detail city
// @Description  Retrieve city detail
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.FindByCodeRequest  true  "Request"
// @Success      200  {object}  model.StoreDetailResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/city/detail [get]
func (ctr *locationController) FindCity(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.FindCity)
}

// @Summary      List districts
// @Description  Retrieve districts list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.StoresListRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/district/list [get]
func (ctr *locationController) ListDistricts(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.ListDistricts)
}

// @Summary      Get detail district
// @Description  Retrieve district detail
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.FindByCodeRequest  true  "Request"
// @Success      200  {object}  model.StoreDetailResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/district/detail [get]
func (ctr *locationController) FindDistrict(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.FindDistrict)
}

// @Summary      List subdistricts
// @Description  Retrieve subdistricts list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.StoresListRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/subdistrict/list [get]
func (ctr *locationController) ListSubdistricts(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.ListSubdistricts)
}

// @Summary      Get detail subdistrict
// @Description  Retrieve subdistrict detail
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.FindByCodeRequest  true  "Request"
// @Success      200  {object}  model.StoreDetailResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/subdistrict/detail [get]
func (ctr *locationController) FindSubdistrict(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.LocationService.FindSubdistrict)
}

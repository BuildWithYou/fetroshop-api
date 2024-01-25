package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/store"
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
	Validate     *validator.Validate
	StoreService store.StoreService
}

func LocationControllerProvider(vld *validator.Validate, svc store.StoreService) LocationController {
	return &locationController{
		Validate:     vld,
		StoreService: svc,
	}
}

// @Summary      List provinces
// @Description  Retrieve provinces list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ListStoresRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/province/list [get]
func (ctr *locationController) ListProvinces(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.List)
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
// @Router       /api/location/province/find [get]
func (ctr *locationController) FindProvince(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Find)
}

// @Summary      List cities
// @Description  Retrieve cities list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ListStoresRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/city/list [get]
func (ctr *locationController) ListCities(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.List)
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
// @Router       /api/location/city/find [get]
func (ctr *locationController) FindCity(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Find)
}

// @Summary      List districts
// @Description  Retrieve districts list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ListStoresRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/district/list [get]
func (ctr *locationController) ListDistricts(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.List)
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
// @Router       /api/location/district/find [get]
func (ctr *locationController) FindDistrict(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Find)
}

// @Summary      List subdistricts
// @Description  Retrieve subdistricts list
// @Tags         Locations
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ListStoresRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/location/subdistrict/list [get]
func (ctr *locationController) ListSubdistricts(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.List)
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
// @Router       /api/location/subdistrict/find [get]
func (ctr *locationController) FindSubdistrict(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Find)
}

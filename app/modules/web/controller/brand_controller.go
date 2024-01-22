package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/brand"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BrandController interface {
	List(ctx *fiber.Ctx) (err error)
	ListByPrefix(ctx *fiber.Ctx) (err error)
	Find(ctx *fiber.Ctx) (err error)
}

type brandController struct {
	Validate     *validator.Validate
	BrandService brand.BrandService
}

func BrandControllerProvider(vld *validator.Validate, brSvc brand.BrandService) BrandController {
	return &brandController{
		Validate:     vld,
		BrandService: brSvc,
	}
}

// @Summary      List brands
// @Description  Retrieve brands list
// @Tags         Brands
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ListBrandsRequest  true  "Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/brand/list [get]
func (ctr *brandController) List(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.BrandService.List)
}

// @Summary      List brands by prefix
// @Description  Retrieve brands list by prefix
// @Tags         Brands
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ListBrandsByPrefixRequest  true  "Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/brand/list-by-prefix [get]
func (ctr *brandController) ListByPrefix(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.BrandService.ListByPrefix)
}

// @Summary      Get detail brand
// @Description  Retrieve categories detail
// @Tags         Brands
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.FindByCodeRequest  true  "Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/brand/find [get]
func (ctr *brandController) Find(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.BrandService.Find)
}

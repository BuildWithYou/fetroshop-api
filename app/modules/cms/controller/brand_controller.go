package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/brand"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type BrandController interface {
	Create(ctx *fiber.Ctx) (err error)
	Update(ctx *fiber.Ctx) (err error)
	Delete(ctx *fiber.Ctx) (err error)
	List(ctx *fiber.Ctx) (err error)
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

// @Summary      Create brand
// @Description
// @Tags         Brands
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        data    formData    model.UpsertBrandRequest  true  "Request"
// @Success      200     {object}    model.Response
// @Failure      400     {object}    model.Response
// @Failure      404     {object}    model.Response
// @Failure      500     {object}    model.Response
// @Router       /api/brand/create  [post]
// @Security Bearer
func (ctr *brandController) Create(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.BrandService.Create)
}

// @Summary      Update brand
// @Description
// @Tags         Brands
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code   path      string                       true  "Category Code"
// @Param        data   formData  model.UpsertBrandRequest  true  "Request"
// @Success      200    {object}  model.Response
// @Failure      400    {object}  model.Response
// @Failure      404    {object}  model.Response
// @Failure      500    {object}  model.Response
// @Router       /api/brand/{code} [put]
// @Security Bearer
func (ctr *brandController) Update(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.BrandService.Update)
}

// @Summary      Delete brand
// @Description
// @Tags         Brands
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code   path       string                         true  "Category Code"
// @Param        data   formData   model.DeleteRequest    true  "Request"
// @Success      200    {object}   model.Response
// @Failure      400    {object}   model.Response
// @Failure      404    {object}   model.Response
// @Failure      500    {object}   model.Response
// @Router       /api/brand/{code} [delete]
// @Security Bearer
func (ctr *brandController) Delete(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.BrandService.Delete)
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

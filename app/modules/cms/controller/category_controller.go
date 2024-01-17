package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/category"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Create(ctx *fiber.Ctx) (err error)
	Update(ctx *fiber.Ctx) (err error)
	Delete(ctx *fiber.Ctx) (err error)
}

type CategoryControllerImpl struct {
	Validate        *validator.Validate
	CategoryService category.CategoryService
}

func CategoryControllerProvider(vld *validator.Validate, catSvc category.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Validate:        vld,
		CategoryService: catSvc,
	}
}

// @Summary      Create category
// @Description
// @Tags         Categories
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        data    formData    model.UpsertCategoryRequest  true  "Request"
// @Success      200     {object}    model.Response
// @Failure      400     {object}    model.Response
// @Failure      404     {object}    model.Response
// @Failure      500     {object}    model.Response
// @Router       /api/category/create  [post]
func (ctr *CategoryControllerImpl) Create(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.CategoryService.Create)
}

// @Summary      Update category
// @Description
// @Tags         Categories
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code   path      string                       true  "Category Code"
// @Param        data   formData  model.UpsertCategoryRequest  true  "Request"
// @Success      200    {object}  model.Response
// @Failure      400    {object}  model.Response
// @Failure      404    {object}  model.Response
// @Failure      500    {object}  model.Response
// @Router       /api/category/{code} [put]
func (ctr *CategoryControllerImpl) Update(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.CategoryService.Update)
}

// @Summary      Delete category
// @Description
// @Tags         Categories
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code   path       string                         true  "Category Code"
// @Param        data   formData   model.DeleteCategoryRequest    true  "Request"
// @Success      200    {object}   model.Response
// @Failure      400    {object}   model.Response
// @Failure      404    {object}   model.Response
// @Failure      500    {object}   model.Response
// @Router       /api/category/{code} [delete]
func (ctr *CategoryControllerImpl) Delete(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.CategoryService.Delete)
}

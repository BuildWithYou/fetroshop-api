package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/category"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Create(ctx *fiber.Ctx) (err error)
	Update(ctx *fiber.Ctx) (err error)
	Delete(ctx *fiber.Ctx) (err error)
	List(ctx *fiber.Ctx) (err error)
	Find(ctx *fiber.Ctx) (err error)
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
// @Security Bearer
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
// @Security Bearer
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
// @Security Bearer
func (ctr *CategoryControllerImpl) Delete(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.CategoryService.Delete)
}

// @Summary      List categories
// @Description  Retrieve categories list
// @Tags         Categories
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ListCategoriesRequest  true  "Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/category/list [get]
func (ctr *CategoryControllerImpl) List(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.CategoryService.List)
}

// @Summary      Get detail category
// @Description  Retrieve categories detail
// @Tags         Categories
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.FindCategoryRequest  true  "Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/category/find [get]
func (ctr *CategoryControllerImpl) Find(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.CategoryService.Find)
}

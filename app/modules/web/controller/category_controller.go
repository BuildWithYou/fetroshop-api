package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/category"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
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

// @Summary      List categories
// @Description  Retrieve categories list
// @Tags         Categories
// @Accept       json
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

// @Summary      Get category
// @Description  Retrieve categories detail
// @Tags         Categories
// @Accept       json
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

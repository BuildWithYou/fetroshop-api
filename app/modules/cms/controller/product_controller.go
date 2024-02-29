package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/product"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	Create(ctx *fiber.Ctx) (err error)
	Update(ctx *fiber.Ctx) (err error)
	Delete(ctx *fiber.Ctx) (err error)
	List(ctx *fiber.Ctx) (err error)
	Find(ctx *fiber.Ctx) (err error)
}

type productController struct {
	Validate       *validator.Validate
	ProductService product.ProductService
}

func ProductControllerProvider(vld *validator.Validate, svc product.ProductService) ProductController {
	return &productController{
		Validate:       vld,
		ProductService: svc,
	}
}

// @Summary      Create product
// @Description
// @Tags         Products
// @Accept       mpfd
// @Produce      json
// @Param        data        formData    model.UpsertProductRequest  true   "Request"
// @Param        mediaFile   formData    file                        false  "accept image/png, image/jpeg"
// @Success      200         {object}    model.productDetailResponse
// @Failure      400         {object}    model.Response
// @Failure      404         {object}    model.Response
// @Failure      500         {object}    model.Response
// @Router       /api/product/create  [post]
// @Security Bearer
func (ctr *productController) Create(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.ProductService.Create)
}

// @Summary      Update product
// @Description
// @Tags         Products
// @Accept       mpfd
// @Produce      json
// @Param        code        path      string                       true   "Product Code"
// @Param        data        formData  model.UpsertProductRequest   true   "Request"
// @Param        mediaFile   formData  file                         false  "accept image/png, image/jpeg"
// @Success      200         {object}  model.productDetailResponse
// @Failure      400         {object}  model.Response
// @Failure      404         {object}  model.Response
// @Failure      500         {object}  model.Response
// @Router       /api/product/{code} [put]
// @Security Bearer
func (ctr *productController) Update(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.ProductService.Update)
}

// @Summary      Delete product
// @Description
// @Tags         Products
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code   path       string                 true  "Product Code"
// @Param        data   query      model.DeleteRequest    true  "Request"
// @Success      200    {object}   model.Response
// @Failure      400    {object}   model.Response
// @Failure      404    {object}   model.Response
// @Failure      500    {object}   model.Response
// @Router       /api/product/{code} [delete]
// @Security Bearer
func (ctr *productController) Delete(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.ProductService.Delete)
}

// @Summary      List products
// @Description  Retrieve products list
// @Tags         Products
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.ProductListRequest  true  "Request"
// @Success      200  {object}  model.productsListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/product/list [get]
func (ctr *productController) List(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.ProductService.List)
}

// @Summary      Get detail product
// @Description  Retrieve product detail
// @Tags         Products
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code path      string                     true  "Product Code"
// @Success      200  {object}  model.productDetailResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/product/{code} [get]
func (ctr *productController) Find(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.ProductService.Find)
}

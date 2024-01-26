package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/store"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StoreController interface {
	Create(ctx *fiber.Ctx) (err error)
	Update(ctx *fiber.Ctx) (err error)
	Delete(ctx *fiber.Ctx) (err error)
	List(ctx *fiber.Ctx) (err error)
	Find(ctx *fiber.Ctx) (err error)
}

type storeController struct {
	Validate     *validator.Validate
	StoreService store.StoreService
}

func StoreControllerProvider(vld *validator.Validate, svc store.StoreService) StoreController {
	return &storeController{
		Validate:     vld,
		StoreService: svc,
	}
}

// @Summary      Create store
// @Description
// @Tags         Stores
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        data    formData    model.UpsertStoreRequest  true  "Request"
// @Success      200     {object}    model.StoreDetailResponse
// @Failure      400     {object}    model.Response
// @Failure      404     {object}    model.Response
// @Failure      500     {object}    model.Response
// @Router       /api/store/create  [post]
// @Security Bearer
func (ctr *storeController) Create(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Create)
}

// @Summary      Update store
// @Description
// @Tags         Stores
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code   path      string                       true  "Store Code"
// @Param        data   formData  model.UpsertStoreRequest  true  "Request"
// @Success      200    {object}  model.StoreDetailResponse
// @Failure      400    {object}  model.Response
// @Failure      404    {object}  model.Response
// @Failure      500    {object}  model.Response
// @Router       /api/store/{code} [put]
// @Security Bearer
func (ctr *storeController) Update(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Update)
}

// @Summary      Delete store
// @Description
// @Tags         Stores
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        code   path       string                 true  "Store Code"
// @Param        data   formData   model.DeleteRequest    true  "Request"
// @Success      200    {object}   model.Response
// @Failure      400    {object}   model.Response
// @Failure      404    {object}   model.Response
// @Failure      500    {object}   model.Response
// @Router       /api/store/{code} [delete]
// @Security Bearer
func (ctr *storeController) Delete(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Delete)
}

// @Summary      List stores
// @Description  Retrieve stores list
// @Tags         Stores
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.StoresListRequest  true  "Request"
// @Success      200  {object}  model.StoresListResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/store/list [get]
func (ctr *storeController) List(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.List)
}

// @Summary      Get detail store
// @Description  Retrieve store detail
// @Tags         Stores
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.FindByCodeRequest  true  "Request"
// @Success      200  {object}  model.StoreDetailResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/store/find [get]
func (ctr *storeController) Find(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Find)
}

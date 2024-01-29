package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/store"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StoreController interface {
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

// @Summary      List stores
// @Description  Retrieve stores list
// @Tags         Stores
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        q     query    model.StoreListRequest  true  "Request"
// @Success      200  {object}  model.storesListResponse
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
// @Param        code path      string                     true  "Store Code"
// @Success      200  {object}  model.storeDetailResponse
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/store/{code} [get]
func (ctr *storeController) Find(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.StoreService.Find)
}

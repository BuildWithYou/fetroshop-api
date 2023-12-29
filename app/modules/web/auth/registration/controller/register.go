package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /accounts/{id} [get]
func (r *RegistrationControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	// TODO - Implement Register
	return ctx.JSON(model.Response{
		Code:    fiber.ErrInternalServerError.Code,
		Status:  fiber.ErrInternalServerError.Message,
		Message: "Not Implemented",
	})
}

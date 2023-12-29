package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/auth/registration"
	"github.com/gofiber/fiber/v2"
)

// @Summary      Register new user
// @Description
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        body  body      registration.RegistrationRequest  true  "Registration Request"
// @Success      200  {object}  model.GeneralResponse
// @Failure      400  {object}  model.GeneralResponse
// @Failure      404  {object}  model.GeneralResponse
// @Failure      500  {object}  model.GeneralResponse
// @Router       /api/web/register [post]
func (r *RegistrationControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	registerResponse, err := r.RegistrationService.Register(&registration.RegistrationRequest{
		Username: ctx.FormValue("username"),
		Phone:    ctx.FormValue("phone"),
		Email:    ctx.FormValue("email"),
		FullName: ctx.FormValue("fullname"),
	})
	helper.PanicIfError(err)
	return ctx.JSON(registerResponse)
}

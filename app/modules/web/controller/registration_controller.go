package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/gofiber/fiber/v2"
)

type RegistrationController interface {
	Register(ctx *fiber.Ctx) (err error)
}

type RegistrationControllerImpl struct {
	RegistrationService registration.RegistrationService
}

func New(registrationService registration.RegistrationService) RegistrationController {
	return &RegistrationControllerImpl{
		RegistrationService: registrationService,
	}
}

// @Summary      Register new user
// @Description
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        body  body      registration.RegistrationRequest  true  "Registration Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/web/register [post]
func (r *RegistrationControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	payload := new(model.RegistrationRequest)
	err = ctx.BodyParser(payload)
	helper.PanicIfError(err)
	registerResponse, err := r.RegistrationService.Register(&model.RegistrationRequest{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
	})
	helper.PanicIfError(err)
	return ctx.JSON(registerResponse)
}

package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth/registration"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RegistrationController interface {
	Register(ctx *fiber.Ctx) (err error)
}

type RegistrationControllerTransport struct {
	Validate            *validator.Validate
	RegistrationService registration.RegistrationService
}

type RegistrationControllerV1 struct {
	Validate            *validator.Validate
	RegistrationService registration.RegistrationService
}

func New(o *RegistrationControllerTransport) RegistrationController {
	return &RegistrationControllerV1{
		Validate:            o.Validate,
		RegistrationService: o.RegistrationService,
	}
}

// @Summary      Register new user
// @Description
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        body  body     registration.RegistrationRequest  true  "Registration Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/web/register [post]
func (r *RegistrationControllerV1) Register(ctx *fiber.Ctx) (err error) {
	payload := new(model.RegistrationRequest)

	err = ctx.BodyParser(payload)
	if helper.ErrorIsNotNil(err) {
		return err
	}

	err = r.Validate.Struct(payload)
	if helper.ErrorIsNotNil(err) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	registerResponse, err := r.RegistrationService.Register(&model.RegistrationRequest{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
	})
	if helper.ErrorIsNotNil(err) {
		return err
	}
	return ctx.JSON(registerResponse)
}

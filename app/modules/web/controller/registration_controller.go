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

type RegistrationControllerImpl struct {
	Validate            *validator.Validate
	RegistrationService registration.RegistrationService
}

func RegistrationControllerProvider(vld *validator.Validate, regSvc registration.RegistrationService) RegistrationController {
	return &RegistrationControllerImpl{
		Validate:            vld,
		RegistrationService: regSvc,
	}
}

// @Summary      Register new user
// @Description
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        body  body     model.RegistrationRequest  true  "Registration Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/register [post]
func (r *RegistrationControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	// TODO: validation need to be move to helper
	payload := new(model.RegistrationRequest)

	err = ctx.BodyParser(payload)
	if helper.IsNotNil(err) {
		return err
	}

	err = r.Validate.Struct(payload)
	if helper.IsNotNil(err) {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	registerResponse, err := r.RegistrationService.Register(&model.RegistrationRequest{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
		Password: payload.Password,
	})
	if helper.IsNotNil(err) {
		return err
	}
	return ctx.JSON(registerResponse)
}

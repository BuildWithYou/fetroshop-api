package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) (err error)
	Login(ctx *fiber.Ctx) (err error)
}

type AuthControllerImpl struct {
	Validate    *validator.Validate
	AuthService auth.AuthService
}

func AuthControllerProvider(vld *validator.Validate, regSvc auth.AuthService) AuthController {
	return &AuthControllerImpl{
		Validate:    vld,
		AuthService: regSvc,
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
func (r *AuthControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	payload := new(model.RegistrationRequest)
	validatorhelper.ValidatePayload(ctx, r.Validate, payload)

	registerResponse, err := r.AuthService.Register(&model.RegistrationRequest{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
		Password: payload.Password,
	})
	if validatorhelper.IsNotNil(err) {
		return err
	}
	return ctx.JSON(registerResponse)
}

// @Summary      Login for users
// @Description
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        body  body     model.LoginRequest  true  "Login Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/login [post]
func (r *AuthControllerImpl) Login(ctx *fiber.Ctx) (err error) {
	payload := new(model.LoginRequest)
	validatorhelper.ValidatePayload(ctx, r.Validate, payload)

	registerResponse, err := r.AuthService.Login(&model.LoginRequest{
		Username: payload.Username,
		Password: payload.Password,
	})
	if validatorhelper.IsNotNil(err) {
		return err
	}
	return ctx.JSON(registerResponse)
}

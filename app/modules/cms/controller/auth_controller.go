package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	execute(ctx *fiber.Ctx, handler func(ctx *fiber.Ctx) (*model.Response, error)) (err error)
	Register(ctx *fiber.Ctx) (err error)
	Login(ctx *fiber.Ctx) (err error)
	Logout(ctx *fiber.Ctx) (err error)
	Refresh(ctx *fiber.Ctx) (err error)
}

type AuthControllerImpl struct {
	AuthService auth.AuthService
}

func AuthControllerProvider(regSvc auth.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: regSvc,
	}
}

func (ctr *AuthControllerImpl) execute(ctx *fiber.Ctx, handler func(ctx *fiber.Ctx) (*model.Response, error)) (err error) {
	response, err := handler(ctx)
	if validatorhelper.IsNotNil(err) {
		return err
	}
	return ctx.JSON(response)
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
func (ctr *AuthControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	return ctr.execute(ctx, ctr.AuthService.Register)
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
func (ctr *AuthControllerImpl) Login(ctx *fiber.Ctx) (err error) {
	return ctr.execute(ctx, ctr.AuthService.Login)
}

// @Summary      Logout for users
// @Description
// @Tags         Authentication
// @Produce      json
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/logout [post]
// @Security Bearer
func (ctr *AuthControllerImpl) Logout(ctx *fiber.Ctx) (err error) {
	return ctr.execute(ctx, ctr.AuthService.Logout)
}

// @Summary      Refresh for customers
// @Description
// @Tags         Authentication
// @Produce      json
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/refresh [post]
// @Security Bearer
func (ctr *AuthControllerImpl) Refresh(ctx *fiber.Ctx) (err error) {
	return ctr.execute(ctx, ctr.AuthService.Refresh)
}

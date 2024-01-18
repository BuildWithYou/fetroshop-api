package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/auth"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	// execute(ctx *fiber.Ctx, handler func(ctx *fiber.Ctx) (*model.Response, error)) (err error)
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

// @Summary      Register new user
// @Description
// @Tags         Authentication
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        req  formData  model.CmsRegistrationRequest  true  "Registration Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/register [post]
func (ctr *AuthControllerImpl) Register(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.AuthService.CmsRegister)
}

// @Summary      Login for users
// @Description
// @Tags         Authentication
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        req  formData  model.CmsLoginRequest  true  "Login Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/login [post]
func (ctr *AuthControllerImpl) Login(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.AuthService.CmsLogin)
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
	return execute(ctx, ctr.AuthService.CmsLogout)
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
	return execute(ctx, ctr.AuthService.CmsRefresh)
}

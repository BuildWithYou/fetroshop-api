package controller

import (
	"github.com/BuildWithYou/fetroshop-api/app/service/auth"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(ctx *fiber.Ctx) (err error)
	Login(ctx *fiber.Ctx) (err error)
	Logout(ctx *fiber.Ctx) (err error)
	Refresh(ctx *fiber.Ctx) (err error)
}

type authController struct {
	Validate    *validator.Validate
	AuthService auth.AuthService
}

func AuthControllerProvider(vld *validator.Validate, regSvc auth.AuthService) AuthController {
	return &authController{
		AuthService: regSvc,
	}
}

// @Summary      Register new user
// @Description
// @Tags         Authentication
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        req  formData  model.RegistrationRequest  true  "Registration Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/register [post]
func (ctr *authController) Register(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.AuthService.WebRegister)
}

// @Summary      Login for customers
// @Description
// @Tags         Authentication
// @Accept       x-www-form-urlencoded,json
// @Produce      json
// @Param        req  formData  model.LoginRequest  true  "Login Request"
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/login [post]
func (ctr *authController) Login(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.AuthService.WebLogin)
}

// @Summary      Logout for customers
// @Description
// @Tags         Authentication
// @Produce      json
// @Success      200  {object}  model.Response
// @Failure      400  {object}  model.Response
// @Failure      404  {object}  model.Response
// @Failure      500  {object}  model.Response
// @Router       /api/auth/logout [post]
// @Security Bearer
func (ctr *authController) Logout(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.AuthService.WebLogout)
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
func (ctr *authController) Refresh(ctx *fiber.Ctx) (err error) {
	return execute(ctx, ctr.AuthService.WebRefresh)
}

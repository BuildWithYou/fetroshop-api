package controller

type Controller struct {
	Auth AuthController
}

func CmsControllerProvider(
	authController AuthController,
) *Controller {
	return &Controller{
		Auth: authController,
	}
}

package controller

type Controller struct {
	Auth AuthController
}

func WebControllerProvider(
	authController AuthController,
) *Controller {
	return &Controller{
		Auth: authController,
	}
}

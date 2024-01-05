package controller

type WebController struct {
	Auth AuthController
}

func WebControllerProvider(
	authController AuthController,
) *WebController {
	return &WebController{
		Auth: authController,
	}
}

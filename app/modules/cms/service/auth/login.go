package auth

import (
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (rg *AuthServiceImpl) Login(request *webModel.LoginRequest) (*model.Response, error) {
	var user users.User

	result := rg.UserRepository.Find(&user, &users.User{
		Username: request.Username,
	})
	if gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error401("Invalid email or password") // #marked: message
	}
	if err := password.Verify(user.Password, request.Password); validatorhelper.IsNotNil(err) {
		return nil, errorhelper.Error401("Invalid email or password")
	}

	token, expiration := jwt.Generate(&jwt.TokenPayload{
		ID:         user.ID,
		Expiration: rg.Config.GetString("security.jwt.expiration"),
		TokenKey:   rg.Config.GetString("security.jwt.tokenKey"),
	})

	fmt.Println("expiration : ", expiration)

	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Login success", // #marked: message
		Data: map[string]string{
			"token":     token,
			"createdAt": time.Now().Format("2006-01-02 15:04:05"),
			"expiredAt": expiration.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

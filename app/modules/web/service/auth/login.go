package auth

import (
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Login(ctx *fiber.Ctx) (*model.Response, error) {
	var customer customers.Customer

	payload := new(webModel.LoginRequest)
	validatorhelper.ValidatePayload(ctx, svc.Validate, payload)

	result := svc.CustomerRepository.Find(&customer, &customers.Customer{
		Username: payload.Username,
	})
	if gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error401("Invalid email or password") // #marked: message
	}
	if err := password.Verify(customer.Password, payload.Password); validatorhelper.IsNotNil(err) {
		return nil, errorhelper.Error401("Invalid email or password")
	}

	token, expiration := jwt.Generate(&jwt.TokenPayload{
		ID:         customer.ID,
		Expiration: svc.Config.GetString("security.jwt.expiration"),
		TokenKey:   svc.Config.GetString("security.jwt.tokenKey"),
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

package auth

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Register(ctx *fiber.Ctx) (*model.Response, error) {
	if validatorhelper.IsNotNil(svc.Err) {
		fmt.Print("Error: ", svc.Err.Error()) // #marked: logging
		return nil, errorhelper.Error500(constant.ERROR_GENERAL)
	}

	var (
		message                                        string
		existingUsername, existingPhone, existingEmail customers.Customer
	)

	payload := new(webModel.RegistrationRequest)
	validatorhelper.ValidatePayload(ctx, svc.Validate, payload)

	result := svc.CustomerRepo.Find(&existingUsername, &customers.Customer{
		Username: payload.Username,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Username already used") // #marked: message
	}

	result = svc.CustomerRepo.Find(&existingPhone, &customers.Customer{
		Phone: payload.Phone,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Phone already used") // #marked: message
	}

	result = svc.CustomerRepo.Find(&existingEmail, &customers.Customer{
		Email: payload.Email,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Email already used") // #marked: message
	}

	hashedPassword := password.Generate(payload.Password)

	result = svc.CustomerRepo.Create(&customers.Customer{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
		Password: hashedPassword,
	})
	if validatorhelper.IsNotNil(result.Error) {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		message = "User created successfully" // #marked: message
	}

	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: message,
	}, nil
}

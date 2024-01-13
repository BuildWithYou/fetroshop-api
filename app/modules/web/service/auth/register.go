package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Register(ctx *fiber.Ctx) (*appModel.Response, error) {
	var (
		message                                        string
		existingUsername, existingPhone, existingEmail customers.Customer
	)

	payload := new(webModel.RegistrationRequest)
	err := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if validatorhelper.IsNotNil(err) {
		return nil, err
	}

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

	return &appModel.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: message,
	}, nil
}

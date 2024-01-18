package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *authService) WebRegister(ctx *fiber.Ctx) (*model.Response, error) {
	var existingUsername, existingPhone, existingEmail customers.Customer

	payload := new(model.WebRegistrationRequest)
	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	result := svc.CustomerRepo.Find(&existingUsername, fiber.Map{"username": payload.Username})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"username": "Username already used"}), nil // #marked: message
	}

	result = svc.CustomerRepo.Find(&existingPhone, fiber.Map{"phone": payload.Phone})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"phone": "Phone already used"}), nil // #marked: message
	}

	result = svc.CustomerRepo.Find(&existingEmail, fiber.Map{"email": payload.Email})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"email": "Email already used"}), nil // #marked: message
	}

	hashedPassword := password.Generate(payload.Password)

	newCustomer := &customers.Customer{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
		Password: hashedPassword,
	}
	result = svc.CustomerRepo.Create(newCustomer)
	if result.Error != nil {
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return nil, errorhelper.Error500("Failed to create user") // #marked: message
	}

	return responsehelper.Response201(
		"Customer created successfully", // #marked: message
		newCustomer,                     // TODO: data return must be filtered
		nil), nil
}

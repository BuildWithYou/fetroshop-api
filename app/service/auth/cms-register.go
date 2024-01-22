package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *authService) CmsRegister(ctx *fiber.Ctx) (*model.Response, error) {
	var existingUsername, existingPhone, existingEmail users.User

	payload := new(model.RegistrationRequest)
	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	result := svc.UserRepo.Find(&existingUsername, fiber.Map{"username": payload.Username})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"username": "Username already used"}), nil // #marked: message
	}

	result = svc.UserRepo.Find(&existingPhone, fiber.Map{"phone": payload.Phone})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"phone": "Phone already used"}), nil // #marked: message
	}

	result = svc.UserRepo.Find(&existingEmail, fiber.Map{"email": payload.Email})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(fiber.Map{"email": "Email already used"}), nil // #marked: message
	}

	hashedPassword := password.Generate(payload.Password)

	newUser := &users.User{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
		Password: hashedPassword,
	}
	result = svc.UserRepo.Create(newUser)
	if result.Error != nil {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		svc.Logger.Error("Failed to create user")
		return nil, errorhelper.Error500("Failed to create user") // #marked: message
	}

	return responsehelper.Response201(
		"User created successfully", // #marked: message
		model.RegistrationResponseData{
			Username: newUser.Username,
			Phone:    newUser.Phone,
			Email:    newUser.Email,
			FullName: newUser.FullName,
		},
		nil), nil
}

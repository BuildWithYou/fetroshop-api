package auth

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	cmsModel "github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Register(ctx *fiber.Ctx) (*model.Response, error) {
	if validatorhelper.IsNotNil(svc.Err) {
		fmt.Printf("\nError: %s\n", svc.Err.Error()) // #marked: logging
		return nil, errorhelper.Error500(constant.ERROR_GENERAL)
	}

	var (
		message                                        string
		existingUsername, existingPhone, existingEmail users.User
	)
	payload := new(cmsModel.RegistrationRequest)
	validatorhelper.ValidatePayload(ctx, svc.Validate, payload)

	result := svc.UserRepo.Find(&existingUsername, &users.User{
		Username: payload.Username,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Username already used") // #marked: message
	}

	result = svc.UserRepo.Find(&existingPhone, &users.User{
		Phone: payload.Phone,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Phone already used") // #marked: message
	}

	result = svc.UserRepo.Find(&existingEmail, &users.User{
		Email: payload.Email,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Email already used") // #marked: message
	}

	hashedPassword := password.Generate(payload.Password)

	result = svc.UserRepo.Create(&users.User{
		Username: payload.Username,
		Phone:    payload.Phone,
		Email:    payload.Email,
		FullName: payload.FullName,
		Password: hashedPassword,
	})
	if validatorhelper.IsNotNil(result.Error) {
		return nil, result.Error
	}

	if gormhelper.HasAffectedRows(result) {
		message = "User created successfully" // #marked: message
	}

	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: message,
	}, nil
}

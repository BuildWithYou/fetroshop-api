package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Register(ctx *fiber.Ctx) (*appModel.Response, error) {
	var (
		message                                        string
		existingUsername, existingPhone, existingEmail users.User
	)
	payload := new(model.RegistrationRequest)
	validatorhelper.ValidatePayload(ctx, svc.Validate, payload)

	result := svc.UserRepository.Find(&existingUsername, &users.User{
		Username: payload.Username,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Username already used") // #marked: message
	}

	result = svc.UserRepository.Find(&existingPhone, &users.User{
		Phone: payload.Phone,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Phone already used") // #marked: message
	}

	result = svc.UserRepository.Find(&existingEmail, &users.User{
		Email: payload.Email,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Email already used") // #marked: message
	}

	hashedPassword := password.Generate(payload.Password)

	result = svc.UserRepository.Create(&users.User{
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

package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (rg *AuthServiceImpl) Register(request *webModel.RegistrationRequest) (*model.Response, error) {

	var (
		message                                        string
		existingUsername, existingPhone, existingEmail users.User
	)

	result := rg.UserRepository.Find(&existingUsername, &users.User{
		Username: request.Username,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Username already used") // #marked: message
	}

	result = rg.UserRepository.Find(&existingPhone, &users.User{
		Phone: request.Phone,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Phone already used") // #marked: message
	}

	result = rg.UserRepository.Find(&existingEmail, &users.User{
		Email: request.Email,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Email already used") // #marked: message
	}

	hashedPassword := password.Generate(request.Password)

	result = rg.UserRepository.Create(&users.User{
		Username: request.Username,
		Phone:    request.Phone,
		Email:    request.Email,
		FullName: request.FullName,
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

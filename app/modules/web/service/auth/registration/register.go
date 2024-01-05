package registration

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (rg *RegistrationServiceImpl) Register(request *webModel.RegistrationRequest) (*model.Response, error) {

	var (
		message                                        string
		existingUsername, existingPhone, existingEmail customers.Customer
	)

	result := rg.CustomerRepository.Find(&existingUsername, &customers.Customer{
		Username: request.Username,
	})
	if helper.IsNotNil(result.Error) && !gormhelper.IsNotFound(result.Error) {
		return nil, result.Error
	}

	if helper.IsNotZero64(existingUsername.ID) {
		// TODO: validation error should be move to helper
		return nil, helper.Error400("Username already used") // #marked: message
	}

	result = rg.CustomerRepository.Find(&existingPhone, &customers.Customer{
		Phone: request.Phone,
	})
	if helper.IsNotNil(result.Error) && !gormhelper.IsNotFound(result.Error) {
		return nil, result.Error
	}
	if helper.IsNotZero64(existingPhone.ID) {
		return nil, helper.Error400("Phone already used") // #marked: message
	}

	result = rg.CustomerRepository.Find(&existingEmail, &customers.Customer{
		Email: request.Email,
	})
	if helper.IsNotNil(result.Error) && !gormhelper.IsNotFound(result.Error) {
		return nil, result.Error
	}
	if helper.IsNotZero64(existingEmail.ID) {
		return nil, helper.Error400("Email already used") // #marked: message
	}

	// TODO: hash password before save

	result = rg.CustomerRepository.Create(&customers.Customer{
		Username: request.Username,
		Phone:    request.Phone,
		Email:    request.Email,
		FullName: request.FullName,
		Password: request.Password,
	})
	if helper.IsNotNil(result.Error) {
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

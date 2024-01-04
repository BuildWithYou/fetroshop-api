package registration

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (rg *RegistrationServiceImpl) Register(request *webModel.RegistrationRequest) (*model.Response, error) {

	var message string

	existingUsername := rg.CustomerRepository.Find(&customers.Customer{
		Username: request.Username,
	})
	if existingUsername.ID != 0 {
		// TODO: validation error should be move to helper
		return nil, fiber.NewError(fiber.StatusBadRequest, "Username already used") // #marked: message
	}

	existingPhone := rg.CustomerRepository.Find(&customers.Customer{
		Phone: request.Phone,
	})
	if existingPhone.ID != 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Phone already used") // #marked: message
	}

	existingEmail := rg.CustomerRepository.Find(&customers.Customer{
		Email: request.Email,
	})
	if existingEmail.ID != 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Email already used") // #marked: message
	}

	// TODO: hash password before save

	result := rg.CustomerRepository.Create(&customers.Customer{
		Username: request.Username,
		Phone:    request.Phone,
		Email:    request.Email,
		FullName: request.FullName,
		Password: request.Password,
	})
	err := result.Error
	if helper.IsNotNil(err) {
		return nil, err
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

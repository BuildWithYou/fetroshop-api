package registration

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
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
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsNotFound(result.Error) {
		return nil, result.Error
	}
	if validatorhelper.IsNotZero64(existingUsername.ID) {
		return nil, errorhelper.Error400("Username already used") // #marked: message
	}

	result = rg.CustomerRepository.Find(&existingPhone, &customers.Customer{
		Phone: request.Phone,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsNotFound(result.Error) {
		return nil, result.Error
	}
	if validatorhelper.IsNotZero64(existingPhone.ID) {
		return nil, errorhelper.Error400("Phone already used") // #marked: message
	}

	result = rg.CustomerRepository.Find(&existingEmail, &customers.Customer{
		Email: request.Email,
	})
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.IsNotFound(result.Error) {
		return nil, result.Error
	}
	if validatorhelper.IsNotZero64(existingEmail.ID) {
		return nil, errorhelper.Error400("Email already used") // #marked: message
	}

	// TODO: hash password before save

	result = rg.CustomerRepository.Create(&customers.Customer{
		Username: request.Username,
		Phone:    request.Phone,
		Email:    request.Email,
		FullName: request.FullName,
		Password: request.Password,
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

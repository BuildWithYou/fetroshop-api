package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
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
	errorMap, err := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if err != nil {
		return responsehelper.Response400(constant.ERROR_VALIDATION, fiber.Map{"messages": errorMap}), nil
	}
	/*
			   TODO:
		      Add validation and give proper messages:
		         - Username : required, unique
		         - Phone : required, unique, numeric
		         - Email : required, unique, valid email
		         - FullName : required
		         - Password : required, min 8
	*/

	result := svc.CustomerRepo.Find(&existingUsername, fiber.Map{"username": payload.Username})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Username already used") // #marked: message
	}

	result = svc.CustomerRepo.Find(&existingPhone, fiber.Map{"phone": payload.Phone})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Phone already used") // #marked: message
	}

	result = svc.CustomerRepo.Find(&existingEmail, fiber.Map{"email": payload.Email})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
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
	if result.Error != nil {
		return nil, result.Error
	}

	if gormhelper.HasAffectedRows(result) {
		message = "User created successfully" // #marked: message
	}

	return &appModel.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusCreated),
		Message: message,
	}, nil
}

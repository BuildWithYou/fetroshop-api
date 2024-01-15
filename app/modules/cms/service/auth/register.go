package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	cmsModel "github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Register(ctx *fiber.Ctx) (*appModel.Response, error) {
	svc.Logger.CmsLoggerResetOutput()
	var (
		message                                        string
		existingUsername, existingPhone, existingEmail users.User
	)
	payload := new(cmsModel.RegistrationRequest)
	err := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if err != nil {
		return nil, err
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

	result := svc.UserRepo.Find(&existingUsername, map[string]any{"username": payload.Username})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Username already used") // #marked: message
	}

	result = svc.UserRepo.Find(&existingPhone, map[string]any{"phone": payload.Phone})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, result.Error
	}
	if !gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error400("Phone already used") // #marked: message
	}

	result = svc.UserRepo.Find(&existingEmail, map[string]any{"email": payload.Email})
	if result.Error != nil && !gormhelper.IsErrRecordNotFound(result.Error) {
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

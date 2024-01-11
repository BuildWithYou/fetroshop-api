package auth

import (
	"fmt"

	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Logout(ctx *fiber.Ctx) (*model.Response, error) {
	if validatorhelper.IsNotNil(svc.Err) {
		fmt.Printf("\nError: %s\n", svc.Err.Error()) // #marked: logging
		return nil, errorhelper.Error500(constant.ERROR_GENERAL)
	}

	userID := jwt.GetUserID(ctx)
	identifier := jwt.GetAccessIdentifier(ctx)
	result := svc.UserAccessRepo.Delete(&user_accesses.UserAccess{
		Key:    identifier,
		UserID: userID})
	if !gormhelper.HasAffectedRows(result) {
		return nil, errorhelper.Error500("Failed to logout") // #marked: message
	}

	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Logout success", // #marked: message
	}, nil
}
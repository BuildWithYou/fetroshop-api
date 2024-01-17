package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Logout(ctx *fiber.Ctx) (*appModel.Response, error) {
	userID := jwt.GetUserID(ctx)
	identifier := jwt.GetAccessIdentifier(ctx)
	result := svc.UserAccessRepo.Delete(&user_accesses.UserAccess{
		Key:    identifier,
		UserID: userID})
	if result.Error != nil {
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return responsehelper.Response500("Failed to logout", nil), nil // #marked: message
	}

	return &appModel.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Logout success", // #marked: message
	}, nil
}

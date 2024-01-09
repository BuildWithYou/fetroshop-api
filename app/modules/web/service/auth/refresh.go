package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Refresh(ctx *fiber.Ctx) (*model.Response, error) {
	customerID := jwt.GetCustomerID(ctx)
	customerAccessID := jwt.GetIdentifierID(ctx)
	result := svc.CustomerAccessRepo.Delete(&customer_accesses.CustomerAccess{
		Key:        customerAccessID,
		CustomerID: customerID})
	if !gormhelper.HasAffectedRows(result) {
		return nil, errorhelper.Error500("Failed to logout") // #marked: message
	}

	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Logout success", // #marked: message
	}, nil
}

package auth

import (
	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *authService) WebLogout(ctx *fiber.Ctx) (*appModel.Response, error) {
	customerID := jwt.GetCustomerID(ctx)
	identifier := jwt.GetAccessIdentifier(ctx)
	result := svc.CustomerAccessRepo.Delete(&customer_accesses.CustomerAccess{
		Key:        identifier,
		CustomerID: customerID})
	if result.Error != nil {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		svc.Logger.Error("Failed to logout")
		return responsehelper.Response500("Failed to logout", nil), nil // #marked: message
	}

	return responsehelper.Response200("Logout success", nil, nil), nil // #marked: message

}

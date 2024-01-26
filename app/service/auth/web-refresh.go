package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *authService) WebRefresh(ctx *fiber.Ctx) (*model.Response, error) {
	customerID := jwt.GetCustomerID(ctx)
	identifier := jwt.GetAccessIdentifier(ctx)
	jwtTokenKey := svc.Config.GetString("security.jwt.tokenKey")
	jwtExpiration := svc.Config.GetString("security.jwt.expiration")

	additionalDuration, err := time.ParseDuration(jwtExpiration)
	if err != nil {
		svc.Logger.Panic("Invalid time duration. Should be time.ParseDuration string")
	}
	expiredAt := time.Now().Add(additionalDuration)

	accessToken := password.Generate(fmt.Sprintf(
		"%s::%s::%s",
		strconv.Itoa(int(customerID)),
		jwtTokenKey,
		time.Now().Format("2006-01-02 15:04:05"),
	))

	generatedJwt := jwt.Generate(&jwt.TokenPayload{
		AccessKey:  accessToken,
		TokenKey:   jwtTokenKey,
		Expiration: expiredAt,
		Type:       CUSTOMER_TYPE,
	})

	result := svc.CustomerAccessRepo.Update(
		&customer_accesses.CustomerAccess{
			Key:        accessToken,
			CustomerID: customerID,
			ExpiredAt:  expiredAt,
		},
		fiber.Map{
			"key":         identifier,
			"customer_id": customerID,
		},
	)
	if result.Error != nil {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		svc.Logger.Error("Failed to refresh token")
		return nil, errorhelper.Error500("Failed to refresh token") // #marked: message
	}

	return responsehelper.Response200(
		"Refresh success", // #marked: message
		fiber.Map{
			"token":     generatedJwt.Token,
			"createdAt": time.Now().Format("2006-01-02 15:04:05"),
			"expiredAt": generatedJwt.ExpiredAt.Format("2006-01-02 15:04:05"),
		}, nil), nil
}

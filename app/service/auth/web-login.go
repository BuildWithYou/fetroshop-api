package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *authService) WebLogin(ctx *fiber.Ctx) (*model.Response, error) {
	var customer customers.Customer

	payload := new(model.WebLoginRequest)
	jwtTokenKey := svc.Config.GetString("security.jwt.tokenKey")
	jwtExpiration := svc.Config.GetString("security.jwt.expiration")

	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// check is customer exist
	result := svc.CustomerRepo.Find(&customer, fiber.Map{"username": payload.Username})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, result.Error
	}
	invalidEmailPasswordMsg := fiber.Map{
		"username": "Invalid username or password",
		"password": "Invalid username or password",
	} // #marked: message
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(invalidEmailPasswordMsg), nil
	}
	if err := password.Verify(customer.Password, payload.Password); err != nil {
		return responsehelper.ResponseErrorValidation(invalidEmailPasswordMsg), nil
	}

	accessToken := password.Generate(fmt.Sprintf(
		"%s::%s::%s",
		strconv.Itoa(int(customer.ID)),
		jwtTokenKey,
		time.Now().Format("2006-01-02 15:04:05"),
	))

	additionalDuration, err := time.ParseDuration(jwtExpiration)
	if err != nil {
		svc.Logger.Panic("Invalid time duration. Should be time.ParseDuration string")
	}
	expiredAt := time.Now().Add(additionalDuration)

	result = svc.CustomerAccessRepo.UpdateOrCreate(&customer_accesses.CustomerAccess{
		Key:        accessToken,
		CustomerID: customer.ID,
		Platform:   ctx.Get("Sec-Ch-Ua-Platform"),
		UserAgent:  ctx.Get("User-Agent"),
		ExpiredAt:  expiredAt,
	},
		fiber.Map{
			"customer_id": customer.ID,
			"platform":    ctx.Get("Sec-Ch-Ua-Platform"),
			"user_agent":  ctx.Get("User-Agent"),
		},
	)
	if result.Error != nil && !gormhelper.HasAffectedRows(result) {
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		return nil, errorhelper.Error500("Failed to record user access") // #marked: message
	}

	generatedJwt := jwt.Generate(&jwt.TokenPayload{
		AccessKey:  accessToken,
		TokenKey:   jwtTokenKey,
		Expiration: expiredAt,
		Type:       CUSTOMER_TYPE,
	})

	return responsehelper.Response200(
		"Login success", // #marked: message
		fiber.Map{
			"token":     generatedJwt.Token,
			"createdAt": time.Now().Format("2006-01-02 15:04:05"),
			"expiredAt": generatedJwt.ExpiredAt.Format("2006-01-02 15:04:05"),
		}, nil), nil
}

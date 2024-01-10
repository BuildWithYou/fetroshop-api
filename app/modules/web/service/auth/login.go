package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/customers"
	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	webModel "github.com/BuildWithYou/fetroshop-api/app/modules/web/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Login(ctx *fiber.Ctx) (*model.Response, error) {
	if validatorhelper.IsNotNil(svc.Err) {
		fmt.Print("Error: ", svc.Err.Error()) // #marked: logging
		return nil, errorhelper.Error500(constant.ERROR_GENERAL)
	}

	var customer customers.Customer

	payload := new(webModel.LoginRequest)
	jwtTokenKey := svc.Config.GetString("security.jwt.tokenKey")
	jwtExpiration := svc.Config.GetString("security.jwt.expiration")

	validatorhelper.ValidatePayload(ctx, svc.Validate, payload)

	// check is customer exist
	result := svc.CustomerRepo.Find(&customer, &customers.Customer{
		Username: payload.Username,
	})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		return nil, errorhelper.Error500("Something went wrong") // #marked: message
	}
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return nil, errorhelper.Error401("Invalid email or password") // #marked: message
	}
	if err := password.Verify(customer.Password, payload.Password); validatorhelper.IsNotNil(err) {
		return nil, errorhelper.Error401("Invalid email or password")
	}

	accessToken := password.Generate(fmt.Sprintf(
		"%s::%s::%s",
		strconv.Itoa(int(customer.ID)),
		jwtTokenKey,
		time.Now().Format("2006-01-02 15:04:05"),
	))

	additionalDuration, err := time.ParseDuration(jwtExpiration)
	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}
	expiredAt := time.Now().Add(additionalDuration)

	result = svc.CustomerAccessRepo.UpdateOrCreate(&customer_accesses.CustomerAccess{
		Key:        accessToken,
		CustomerID: customer.ID,
		Platform:   ctx.Get("Sec-Ch-Ua-Platform"),
		UserAgent:  ctx.Get("User-Agent"),
		ExpiredAt:  expiredAt,
	},
		&customer_accesses.CustomerAccess{
			CustomerID: customer.ID,
			Platform:   ctx.Get("Sec-Ch-Ua-Platform"),
			UserAgent:  ctx.Get("User-Agent"),
		},
	)
	if validatorhelper.IsNotNil(result.Error) && !gormhelper.HasAffectedRows(result) {
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

	return &model.Response{
		Code:    fiber.StatusCreated,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Login success", // #marked: message
		Data: map[string]string{
			"token":     generatedJwt.Token,
			"createdAt": time.Now().Format("2006-01-02 15:04:05"),
			"expiredAt": generatedJwt.ExpiredAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

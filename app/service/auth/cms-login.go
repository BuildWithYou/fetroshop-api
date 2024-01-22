package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/users"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	"github.com/BuildWithYou/fetroshop-api/app/helper/responsehelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
)

func (svc *authService) CmsLogin(ctx *fiber.Ctx) (*model.Response, error) {
	var user users.User

	payload := new(model.LoginRequest)
	jwtTokenKey := svc.Config.GetString("security.jwt.tokenKey")
	jwtExpiration := svc.Config.GetString("security.jwt.expiration")

	errValidation, errParsing := validatorhelper.ValidateBodyPayload(ctx, svc.Validate, payload)
	if errParsing != nil {
		svc.Logger.UseError(errParsing)
		return nil, errParsing
	}
	if errValidation != nil {
		return responsehelper.ResponseErrorValidation(errValidation), nil
	}

	// check is customer exist
	result := svc.UserRepo.Find(&user, fiber.Map{"username": payload.Username})
	if gormhelper.IsErrNotNilNotRecordNotFound(result.Error) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	invalidEmailPasswordMsg := fiber.Map{
		"username": "Invalid username or password",
		"password": "Invalid username or password",
	} // #marked: message
	if gormhelper.IsErrRecordNotFound(result.Error) {
		return responsehelper.ResponseErrorValidation(invalidEmailPasswordMsg), nil
	}
	if err := password.Verify(user.Password, payload.Password); err != nil {
		return responsehelper.ResponseErrorValidation(invalidEmailPasswordMsg), nil
	}

	accessToken := password.Generate(fmt.Sprintf(
		"%s::%s::%s",
		strconv.Itoa(int(user.ID)),
		svc.Config.GetString("security.jwt.tokenKey"),
		time.Now().Format("2006-01-02 15:04:05"),
	))

	additionalDuration, err := time.ParseDuration(jwtExpiration)
	if err != nil {
		svc.Logger.Panic("Invalid time duration. Should be time.ParseDuration string")
	}
	expiredAt := time.Now().Add(additionalDuration)

	result = svc.UserAccessRepo.UpdateOrCreate(&user_accesses.UserAccess{
		Key:       accessToken,
		UserID:    user.ID,
		Platform:  ctx.Get("Sec-Ch-Ua-Platform"),
		UserAgent: ctx.Get("User-Agent"),
		ExpiredAt: expiredAt,
	},
		fiber.Map{
			"user_id":    user.ID,
			"platform":   ctx.Get("Sec-Ch-Ua-Platform"),
			"user_agent": ctx.Get("User-Agent"),
		},
	)
	if result.Error != nil && !gormhelper.HasAffectedRows(result) {
		svc.Logger.UseError(result.Error)
		return nil, result.Error
	}
	if !gormhelper.HasAffectedRows(result) {
		svc.Logger.Error("Failed to record user access")
		return nil, errorhelper.Error500("Failed to record user access") // #marked: message
	}

	generatedJwt := jwt.Generate(&jwt.TokenPayload{
		AccessKey:  accessToken,
		TokenKey:   jwtTokenKey,
		Expiration: expiredAt,
		Type:       USER_TYPE,
	})

	return responsehelper.Response200(
		"Login success", // #marked: message
		fiber.Map{
			"token":     generatedJwt.Token,
			"createdAt": time.Now().Format("2006-01-02 15:04:05"),
			"expiredAt": generatedJwt.ExpiredAt.Format("2006-01-02 15:04:05"),
		}, nil), nil
}

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
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/model"
	cmsModel "github.com/BuildWithYou/fetroshop-api/app/modules/cms/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Login(ctx *fiber.Ctx) (*model.Response, error) {
	var user users.User

	payload := new(cmsModel.LoginRequest)
	validatorhelper.ValidatePayload(ctx, svc.Validate, payload)

	result := svc.UserRepo.Find(&user, &users.User{
		Username: payload.Username,
	})
	if gormhelper.IsRecordNotFound(result.Error) {
		return nil, errorhelper.Error401("Invalid email or password") // #marked: message
	}
	if err := password.Verify(user.Password, payload.Password); validatorhelper.IsNotNil(err) {
		return nil, errorhelper.Error401("Invalid email or password")
	}

	accessToken := password.Generate(fmt.Sprintf(
		"%s::%s::%s",
		strconv.Itoa(int(user.ID)),
		svc.Config.GetString("security.jwt.tokenKey"),
		time.Now().Format("2006-01-02 15:04:05"),
	))

	result = svc.UserAccessRepo.Create(&user_accesses.UserAccess{
		Token:     accessToken,
		UserID:    user.ID,
		Platform:  ctx.Get("Sec-Ch-Ua-Platform"),
		UserAgent: ctx.Get("User-Agent"),
	})

	if gormhelper.HasAffectedRows(result) {
		return nil, errorhelper.Error500("Failed to record user access") // #marked: message
	}

	generatedJwt := jwt.Generate(&jwt.TokenPayload{
		Token:      accessToken,
		Expiration: svc.Config.GetString("security.jwt.expiration"),
		TokenKey:   svc.Config.GetString("security.jwt.tokenKey"),
		Type:       USER_TYPE,
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

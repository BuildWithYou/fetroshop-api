package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/password"
	appModel "github.com/BuildWithYou/fetroshop-api/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
)

func (svc *AuthServiceImpl) Refresh(ctx *fiber.Ctx) (*appModel.Response, error) {
	userID := jwt.GetUserID(ctx)
	identifier := jwt.GetAccessIdentifier(ctx)
	jwtTokenKey := svc.Config.GetString("security.jwt.tokenKey")
	jwtExpiration := svc.Config.GetString("security.jwt.expiration")

	additionalDuration, err := time.ParseDuration(jwtExpiration)
	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}
	expiredAt := time.Now().Add(additionalDuration)

	accessToken := password.Generate(fmt.Sprintf(
		"%s::%s::%s",
		strconv.Itoa(int(userID)),
		jwtTokenKey,
		time.Now().Format("2006-01-02 15:04:05"),
	))

	generatedJwt := jwt.Generate(&jwt.TokenPayload{
		AccessKey:  accessToken,
		TokenKey:   jwtTokenKey,
		Expiration: expiredAt,
		Type:       USER_TYPE,
	})

	result := svc.UserAccessRepo.Update(
		&user_accesses.UserAccess{
			Key:       accessToken,
			UserID:    userID,
			ExpiredAt: expiredAt,
		},
		&user_accesses.UserAccess{
			Key:    identifier,
			UserID: userID,
		},
	)
	if !gormhelper.HasAffectedRows(result) {
		return nil, errorhelper.Error500("Failed to refresh token") // #marked: message
	}

	return &appModel.Response{
		Code:    fiber.StatusOK,
		Status:  utils.StatusMessage(fiber.StatusOK),
		Message: "Refresh success", // #marked: message
		Data: map[string]string{
			"token":     generatedJwt.Token,
			"createdAt": time.Now().Format("2006-01-02 15:04:05"),
			"expiredAt": generatedJwt.ExpiredAt.Format("2006-01-02 15:04:05"),
		},
	}, nil
}

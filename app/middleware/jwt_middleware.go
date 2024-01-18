package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/BuildWithYou/fetroshop-api/app/service/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type JwtMiddleware struct {
	Config             *viper.Viper
	UserAccessRepo     user_accesses.UserAccessRepo
	CustomerAccessRepo customer_accesses.CustomerAccessRepo
	Logger             *logger.Logger
}

// NewJwtMiddleware creates a new JwtMiddleware instance.
//
// It takes a pointer to a viper.Viper object as the parameter config and returns a pointer to a JwtMiddleware object.
func JwtMiddlewareProvider(
	config *viper.Viper,
	userAccessRepo user_accesses.UserAccessRepo,
	customerAccessRepo customer_accesses.CustomerAccessRepo,
) *JwtMiddleware {
	return &JwtMiddleware{
		Config:             config,
		UserAccessRepo:     userAccessRepo,
		CustomerAccessRepo: customerAccessRepo,
	}
}

// Authenticate authenticates the request using JWT.
//
// It takes a *fiber.Ctx object as a parameter and returns an error.
func (jwtMid *JwtMiddleware) Authenticate(ctx *fiber.Ctx) error {
	var tokenString string
	authorization := ctx.Get("Authorization")

	if authorization == "" {
		return fiber.ErrUnauthorized
	}

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	}

	if tokenString == "" {
		return errorhelper.Error401("You are not logged in") // #marked: message
	}

	// Spliting the header
	chunks := strings.Split(authorization, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	// Verify the token which is in the chunks
	reversedToken, err := jwt.Reverse(jwtMid.Config.GetString("security.jwt.tokenKey"), chunks[1], jwtMid.Logger)
	if err != nil {
		jwtMid.Logger.Error(fmt.Sprintln("Error on reverse jwt token : ", err.Error()))
		return fiber.ErrUnauthorized
	}

	ctx.Locals(jwt.ACCESS_IDENTIFIER, reversedToken.AccessKey)
	switch reversedToken.Type {
	case auth.USER_TYPE:
		{
			userAccess := new(user_accesses.UserAccess)
			result := jwtMid.UserAccessRepo.Find(userAccess, fiber.Map{"key": reversedToken.AccessKey}) // TODO: implement redis caching to improve performance
			if gormhelper.IsErrRecordNotFound(result.Error) {
				return fiber.ErrUnauthorized
			}
			if userAccess.ExpiredAt.Before(time.Now()) {
				jwtMid.UserAccessRepo.Delete(userAccess)
				return fiber.ErrUnauthorized
			}
			ctx.Locals(jwt.CMS_IDENTIFIER, userAccess.UserID)
		}
	case auth.CUSTOMER_TYPE:
		{
			customerAccess := new(customer_accesses.CustomerAccess)
			result := jwtMid.CustomerAccessRepo.Find(customerAccess, fiber.Map{"key": reversedToken.AccessKey}) // TODO: implement redis caching to improve performance
			if gormhelper.IsErrRecordNotFound(result.Error) {
				return fiber.ErrUnauthorized
			}
			if customerAccess.ExpiredAt.Before(time.Now()) {
				jwtMid.CustomerAccessRepo.Delete(customerAccess)
				return fiber.ErrUnauthorized
			}
			ctx.Locals(jwt.WEB_IDENTIFIER, customerAccess.CustomerID)
		}
	default:
		return errorhelper.Error500("Invalid token type") // #marked: message
	}

	return ctx.Next()
}

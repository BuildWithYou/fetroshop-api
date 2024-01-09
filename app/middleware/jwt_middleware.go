package middleware

import (
	"strings"

	"github.com/BuildWithYou/fetroshop-api/app/domain/customer_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/domain/user_accesses"
	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/gormhelper"
	"github.com/BuildWithYou/fetroshop-api/app/helper/jwt"
	"github.com/BuildWithYou/fetroshop-api/app/helper/validatorhelper"
	cmsAuthSvc "github.com/BuildWithYou/fetroshop-api/app/modules/cms/service/auth"
	webAuthSvc "github.com/BuildWithYou/fetroshop-api/app/modules/web/service/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type JwtMiddleware struct {
	Config             *viper.Viper
	UserAccessRepo     user_accesses.UserAccessRepo
	CustomerAccessRepo customer_accesses.CustomerAccessRepo
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
	var (
		tokenString string
		userID      int64
	)
	authorization := ctx.Get("Authorization")

	if authorization == "" {
		return fiber.ErrUnauthorized
	}

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	}

	if tokenString == "" {
		// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
		return errorhelper.Error401("You are not logged in") // #amrked: message
	}

	// Spliting the header
	chunks := strings.Split(authorization, " ")

	// If header signature is not like `Bearer <token>`, then throw
	// This is also required, otherwise chunks[1] will throw out of bound error
	if len(chunks) < 2 {
		return fiber.ErrUnauthorized
	}

	// Verify the token which is in the chunks
	reversedToken, err := jwt.Reverse(jwtMid.Config.GetString("security.jwt.tokenKey"), chunks[1])
	if validatorhelper.IsNotNil(err) {
		return fiber.ErrUnauthorized
	}
	switch reversedToken.Type {
	case cmsAuthSvc.USER_TYPE:
		{
			userAccess := new(user_accesses.UserAccess)
			result := jwtMid.UserAccessRepo.Find(userAccess, &user_accesses.UserAccess{
				Token: reversedToken.Token,
			})
			if gormhelper.IsErrRecordNotFound(result.Error) {
				return fiber.ErrUnauthorized
			}
			userID = userAccess.UserID
		}
	case webAuthSvc.CUSTOMER_TYPE:
		{
			customerAccess := new(customer_accesses.CustomerAccess)
			result := jwtMid.CustomerAccessRepo.Find(customerAccess, &customer_accesses.CustomerAccess{
				Token: reversedToken.Token,
			})
			if gormhelper.IsErrRecordNotFound(result.Error) {
				return fiber.ErrUnauthorized
			}
		}
	default:
		return errorhelper.Error500("Invalid token type") // #marked: message
	}

	ctx.Locals("UserID", userID)

	return ctx.Next()
}

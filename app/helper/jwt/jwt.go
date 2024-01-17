package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/helper/constant"
	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// TODO: include in struct

const CMS_IDENTIFIER = "USER_ID"
const WEB_IDENTIFIER = "CUSTOMER_ID"
const ACCESS_IDENTIFIER = "ACCESS_IDENTIFIER"

// TokenPayload defines the payload for the token
type TokenPayload struct {
	AccessKey  string
	TokenKey   string
	Expiration time.Time
	Type       string
}

type TokenGenerated struct {
	Token     string
	ExpiredAt time.Time
}

type TokenReversed struct {
	AccessKey string
	Type      string
	ExpiredAt time.Time
}

// Generate generates the jwt token based on payload
func Generate(payload *TokenPayload) *TokenGenerated {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":       payload.Expiration.Unix(),
		"accessKey": payload.AccessKey,
		"type":      payload.Type,
	})

	token, err := t.SignedString([]byte(payload.TokenKey))

	if err != nil {
		panic(err)
	}

	return &TokenGenerated{
		Token:     token,
		ExpiredAt: payload.Expiration,
	}
}

func parse(tokenKey string, jwtToken string) (*jwt.Token, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	return jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenKey), nil
	})
}

// Verify verifies the jwt token against the secret
func Reverse(tokenKey string, jwtToken string, lg *logger.Logger) (*TokenReversed, error) {
	var expiredAt time.Time
	parsed, err := parse(tokenKey, jwtToken)

	if err != nil {
		lg.Error(fmt.Sprint("Error from jwt.parse : ", err.Error()))
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		lg.Error("Error from jwt.Reverse on parsed.Claims")
		return nil, err
	}

	accessKey, ok := claims["accessKey"].(string)
	if !ok {
		lg.Error("Error from jwt.Reverse when parsing accessKey")
		return nil, errors.New(constant.ERROR_GENERAL) // #marked: message
	}

	expiredAtFloat, ok := claims["exp"].(float64)
	if ok {
		seconds := int64(expiredAtFloat)
		expiredAt = time.Unix(seconds, 0)
	} else {
		lg.Error("Error from jwt.Reverse when parsing exp")
		return nil, errors.New(constant.ERROR_GENERAL) // #marked: message
	}

	tokenType, ok := claims["type"].(string)
	if !ok {
		lg.Error("Error from jwt.Reverse when parsing type")
		return nil, errors.New(constant.ERROR_GENERAL) // #marked: message
	}

	return &TokenReversed{
		AccessKey: accessKey,
		ExpiredAt: expiredAt,
		Type:      tokenType,
	}, nil
}

func GetCustomerID(ctx *fiber.Ctx) int64 {
	return ctx.Locals(WEB_IDENTIFIER).(int64)
}

func GetUserID(ctx *fiber.Ctx) int64 {
	return ctx.Locals(CMS_IDENTIFIER).(int64)
}

func GetAccessIdentifier(ctx *fiber.Ctx) string {
	return ctx.Locals(ACCESS_IDENTIFIER).(string)
}

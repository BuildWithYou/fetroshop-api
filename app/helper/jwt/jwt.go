package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const CMS_IDENTIFIER = "USER_ID"
const WEB_IDENTIFIER = "CUSTOMER_ID"
const ACCESS_IDENTIFIER = "ACCESS_IDENTIFIER"
const errorMessage = "something went wrong"

// TokenPayload defines the payload for the token
type TokenPayload struct {
	ID         string
	TokenKey   string
	Expiration string
	Type       string
}

type TokenGenerated struct {
	Token     string
	ExpiredAt time.Time
}

type TokenReversed struct {
	ID        string
	Type      string
	ExpiredAt time.Time
}

// Generate generates the jwt token based on payload
func Generate(payload *TokenPayload) *TokenGenerated {
	additionalDuration, err := time.ParseDuration(payload.Expiration)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	expiration := time.Now().Add(additionalDuration)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  expiration.Unix(),
		"id":   payload.ID,
		"type": payload.Type,
	})

	token, err := t.SignedString([]byte(payload.TokenKey))

	if err != nil {
		panic(err)
	}

	return &TokenGenerated{
		Token:     token,
		ExpiredAt: expiration,
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
func Reverse(tokenKey string, jwtToken string) (*TokenReversed, error) {
	var expiredAt time.Time
	parsed, err := parse(tokenKey, jwtToken)

	if err != nil {
		fmt.Println("Error from jwt.parse : ", err.Error()) // #marked: logging
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error from jwt.Reverse on parsed.Claims") // #marked: logging
		return nil, err
	}

	identifier, ok := claims["id"].(string)
	if !ok {
		fmt.Println("Error from jwt.Reverse when parsing id") // #marked: logging
		return nil, errors.New(errorMessage)                  // #marked: message
	}

	expiredAtFloat, ok := claims["exp"].(float64)
	if ok {
		seconds := int64(expiredAtFloat)
		expiredAt = time.Unix(seconds, 0)
	} else {
		fmt.Println("Error from jwt.Reverse when parsing exp") // #marked: logging
		return nil, errors.New(errorMessage)                   // #marked: message
	}

	tokenType, ok := claims["type"].(string)
	if !ok {
		fmt.Println("Error from jwt.Reverse when parsing type") // #marked: logging
		return nil, errors.New(errorMessage)                    // #marked: message
	}

	return &TokenReversed{
		ID:        identifier,
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

func GetIdentifierID(ctx *fiber.Ctx) string {
	return ctx.Locals(ACCESS_IDENTIFIER).(string)
}

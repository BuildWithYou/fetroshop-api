package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const errorMessage = "something went wrong"

// TokenPayload defines the payload for the token
type TokenPayload struct {
	Token      string
	TokenKey   string
	Expiration string
	Type       string
}

type TokenGenerated struct {
	Token     string
	ExpiredAt time.Time
}

type TokenReversed struct {
	Token     string
	Type      string
	ExpiredAt string
}

// Generate generates the jwt token based on payload
func Generate(payload *TokenPayload) *TokenGenerated {
	additionalDuration, err := time.ParseDuration(payload.Expiration)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	expiration := time.Now().Add(additionalDuration)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":   expiration.Unix(),
		"Token": payload.Token,
		"type":  payload.Type,
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
	parsed, err := parse(tokenKey, jwtToken)

	if err != nil {
		fmt.Println("Error from jwt.parse : ", err.Error()) // #marked: debug
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Error from parsed.Claims : ", err.Error()) // #marked: debug
		return nil, err
	}

	token, ok := claims["Token"].(string)
	if !ok {
		return nil, errors.New(errorMessage) // #marked: message
	}

	expiredAt, ok := claims["exp"].(string)
	if !ok {
		return nil, errors.New(errorMessage) // #marked: message
	}

	tokenType, ok := claims["type"].(string)
	if !ok {
		return nil, errors.New(errorMessage) // #marked: message
	}

	return &TokenReversed{
		Token:     token,
		ExpiredAt: expiredAt,
		Type:      tokenType,
	}, nil
}

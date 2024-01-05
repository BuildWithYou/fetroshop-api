package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TokenPayload defines the payload for the token
type TokenPayload struct {
	Expiration string
	TokenKey   string
	ID         int64
}

// Generate generates the jwt token based on payload
func Generate(payload *TokenPayload) (token string, expiration time.Time) {
	v, err := time.ParseDuration(payload.Expiration)

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	expiration = time.Now().Add(v)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  expiration.Unix(),
		"ID":   payload.ID,
		"role": "customer",
	})

	token, err = t.SignedString([]byte(payload.TokenKey))

	if err != nil {
		panic(err)
	}

	return token, expiration
}

func parse(tokenKey string, token string) (*jwt.Token, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(tokenKey), nil
	})
}

// Verify verifies the jwt token against the secret
func Verify(tokenKey string, token string) (*TokenPayload, error) {
	parsed, err := parse(tokenKey, token)

	if err != nil {
		return nil, err
	}

	// Parsing token claims
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	// Getting ID, it's an interface{} so I need to cast it to uint
	id, ok := claims["ID"].(int64)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{
		ID: id,
	}, nil
}

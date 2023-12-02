package jwt

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtHeaderToken = "X-AUTH-TOKEN"

var defaultSecretKey = "my-local-secret-key"

// Custom Claims implementation
type AppClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func (a AppClaims) Validate() error {
	if a.UserId == "" {
		return errors.New("id can't be empty")
	}

	return nil
}

func getSecretKey() []byte {
	secretKeyFromEnv := os.Getenv("JWT_SECRET_KEY")

	if secretKeyFromEnv == "" {
		return []byte(defaultSecretKey)
	}

	return []byte(secretKeyFromEnv)
}

func GenerateToken(userId string) (string, error) {
	// Create a new token with a signing method and claims
	claims := &AppClaims{
		userId,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(getSecretKey())

	if err != nil {
		fmt.Println(err)

		return "", fmt.Errorf("failed to generate token")
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *AppClaims, error) {
	// Parse the token using the secret key
	token, err := jwt.ParseWithClaims(tokenString, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return getSecretKey(), nil
	})

	if err != nil {
		return nil, nil, err
	}

	appClaims, ok := token.Claims.(*AppClaims)

	if !ok {
		return nil, nil, fmt.Errorf("failed to cast claims type")
	}

	return token, appClaims, nil
}

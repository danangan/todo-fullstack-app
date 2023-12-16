package tokenService

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

var activeTokenCacheNamespace = "user_active_token"
var tokenBlackListCacheNamespace = "token_blacklist"
var tokenBlackListCacheValue = "true"

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
	var defaultSecretKey = "my-local-secret-key"

	secretKeyFromEnv := os.Getenv("JWT_SECRET_KEY")

	if secretKeyFromEnv == "" {
		fmt.Println("JWT_SECRET_KEY is not found in the environment variable, using the default key for token manager")

		return []byte(defaultSecretKey)
	}

	return []byte(secretKeyFromEnv)
}

type TokenManager struct {
	secretKey   []byte
	redisClient *redis.Client
}

func New(redisClient *redis.Client) *TokenManager {
	return &TokenManager{
		secretKey: getSecretKey(),
	}
}

// Generate token automatically revoke previously active token associated with the user
func (t *TokenManager) GenerateToken(ctx context.Context, userId string) (string, error) {
	token, err := t.GetActiveToken(ctx, userId)

	if err != nil {
		return "", nil
	}

	// If currently there is an active token, revoke it before regenerating a new one
	if token != "" {
		_, claims, err := t.ParseToken(token)

		if err != nil {
			return "", nil
		}

		expiryTime := claims.ExpiresAt.Time

		t.RevokeToken(ctx, token, expiryTime)
	}

	newTokenExpiry := time.Now().Add(24 * time.Hour)

	newToken, err := t.createToken(userId, newTokenExpiry)

	if err != nil {
		return "", nil
	}

	err = t.UpdateActiveToken(ctx, newToken, userId, newTokenExpiry)

	if err != nil {
		return "", nil
	}

	return newToken, nil
}

func (t *TokenManager) createToken(userId string, expiry time.Time) (string, error) {
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

func (t *TokenManager) ParseToken(tokenString string) (*jwt.Token, *AppClaims, error) {
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

func (t *TokenManager) buildActiveTokenCacheKey(userId string) string {
	return fmt.Sprintf("%s_%s", activeTokenCacheNamespace, userId)
}

func (t *TokenManager) UpdateActiveToken(ctx context.Context, token string, userId string, expiry time.Time) error {
	result := t.redisClient.Set(ctx, t.buildActiveTokenCacheKey(userId), token, time.Until(expiry))

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func (t *TokenManager) GetActiveToken(ctx context.Context, userId string) (string, error) {
	result, err := t.redisClient.Get(ctx, t.buildActiveTokenCacheKey(userId)).Result()

	if err == nil {
		return result, nil
	}

	if errors.Is(err, redis.Nil) {
		return "", nil
	}

	return "", err
}

func (t *TokenManager) buildBlackListCacheKey(token string) string {
	return fmt.Sprintf("%s_%s", tokenBlackListCacheNamespace, token)
}

func (t *TokenManager) RevokeToken(ctx context.Context, tokenString string, expiry time.Time) (bool, error) {
	_, err := t.redisClient.Set(ctx, t.buildBlackListCacheKey(tokenString), tokenBlackListCacheValue, time.Until(expiry)).Result()

	if err != nil {
		return false, nil
	}

	return true, nil
}

func (t *TokenManager) IsTokenRevoked(ctx context.Context, tokenString string) (bool, error) {
	result, err := t.redisClient.Get(ctx, t.buildBlackListCacheKey(tokenString)).Result()

	if err == nil && result == tokenBlackListCacheValue {
		return true, nil
	}

	if errors.Is(err, redis.Nil) {
		return false, nil
	}

	return false, err
}

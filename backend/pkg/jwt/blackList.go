package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var tokenBlackListNamespace = "token_blacklist"
var tokenBlackListCacheValue = "true"

func buildBlackListCacheKey(token string) string {
	return fmt.Sprintf("%s_%s", tokenBlackListNamespace, token)
}

func RevokeToken(ctx context.Context, client *redis.Client, tokenString string, expiry time.Time) (bool, error) {
	_, err := client.Set(ctx, buildBlackListCacheKey(tokenString), tokenBlackListCacheValue, time.Until(expiry)).Result()

	if err != nil {
		return false, nil
	}

	return true, nil
}

func IsTokenRevoked(ctx context.Context, client *redis.Client, tokenString string) (bool, error) {
	result, err := client.Get(ctx, buildBlackListCacheKey(tokenString)).Result()

	if err == nil && result == tokenBlackListCacheValue {
		return true, nil
	}

	if errors.Is(err, redis.Nil) {
		return false, nil
	}

	return false, err
}

package jwt

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var tokenCacheNameSpace = "user_active_token"

func buildCacheKey(userId string) string {
	return fmt.Sprintf("%s_%s", tokenCacheNameSpace, userId)
}

func UpdateActiveToken(ctx context.Context, redisClient *redis.Client, token string, userId string, expiry time.Time) error {
	result := redisClient.Set(ctx, buildCacheKey(userId), token, time.Until(expiry))

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}

func GetActiveToken(ctx context.Context, redisClient *redis.Client, userId string) (string, error) {
	result, err := redisClient.Get(ctx, buildCacheKey(userId)).Result()

	if err == nil {
		return result, nil
	}

	if errors.Is(err, redis.Nil) {
		return "", nil
	}

	return "", err
}

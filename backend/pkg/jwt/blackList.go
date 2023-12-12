package jwt

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

var blackListNamespace = "jwt_blacklist"

func BlackListToken(ctx context.Context, client *redis.Client, tokenString string) (bool, error) {
	token, claims, err := ParseToken(tokenString)

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	expiredAt := claims.ExpiresAt.Time

	result := client.ZAdd(ctx, blackListNamespace, redis.Z{
		Score:  float64(expiredAt.Unix()),
		Member: tokenString,
	})

	if err := result.Err(); err != nil {
		return false, err
	}

	return true, nil
}

func IsTokenBlackListed(ctx context.Context, client *redis.Client, tokenString string) (bool, error) {
	_, err := client.ZScore(ctx, blackListNamespace, tokenString).Result()

	if err == nil {
		return true, nil
	}

	if err == redis.Nil {
		return false, nil
	}

	return false, err
}

// Todo: To create a background job that clean up the expired token in the black list
func CleanupBlackListedToken(ctx context.Context, client *redis.Client, time time.Time) error {
	_, err := client.ZRemRangeByScore(ctx, blackListNamespace, "-inf", strconv.FormatInt(time.Unix(), 10)).Result()

	return err
}

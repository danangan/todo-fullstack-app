package middleware

import (
	"app/pkg/appContext"
	"app/pkg/db/models"
	"app/pkg/tokenService"
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CreateAuthMiddleware(db *gorm.DB, redisClient *redis.Client, tokenManager *tokenService.TokenManager) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var currentUser *models.User

			stringToken := r.Header.Get("X-AUTH-TOKEN")

			if stringToken != "" {
				token, claims, err := tokenManager.ParseToken(stringToken)

				if err != nil || !token.Valid {
					http.Error(w, "invalid auth token", http.StatusUnauthorized)

					return
				}

				isTokenRevoked, err := tokenManager.IsTokenRevoked(r.Context(), stringToken)

				if err != nil {
					http.Error(w, "can't verify token", http.StatusInternalServerError)

					return
				}

				if isTokenRevoked {
					http.Error(w, "token is invalid", http.StatusUnauthorized)

					return
				}

				user := &models.User{}

				result := db.Where("id = ?", claims.UserId).First(user)

				if result.Error != nil {
					http.Error(w, "invalid auth token", http.StatusUnauthorized)

					return
				}

				currentUser = user
			}

			ctx := context.WithValue(r.Context(), appContext.CurrentUserKey, currentUser)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

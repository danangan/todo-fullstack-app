package middleware

import (
	"app/pkg/appContext"
	"app/pkg/db/models"
	"app/pkg/jwt"
	"context"
	"net/http"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CreateAuthMiddleware(db *gorm.DB, redisClient *redis.Client) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var currentUser *models.User

			token := r.Header.Get(jwt.JwtHeaderToken)

			if token != "" {
				_, claims, err := jwt.ParseToken(token)

				if err != nil {
					http.Error(w, "invalid auth token", http.StatusUnauthorized)

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

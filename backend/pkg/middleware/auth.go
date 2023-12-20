package middleware

import (
	"app/pkg/appContext"
	"app/pkg/db/models"
	"app/pkg/tokenService"
	"app/pkg/userService"
	"context"
	"encoding/json"
	"net/http"
)

func writeError(w http.ResponseWriter, message string, errorCode int) {
	w.WriteHeader(http.StatusInternalServerError)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"errors": []map[string]interface{}{
			{
				"message": message,
				"code":    errorCode,
			},
		},
	})
}

func CreateAuthMiddleware(tokenManager *tokenService.TokenManager, userService *userService.UserService) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var currentUser *models.User

			stringToken := r.Header.Get("X-AUTH-TOKEN")

			if stringToken != "" {
				isTokenRevoked, err := tokenManager.IsTokenRevoked(r.Context(), stringToken)

				if err != nil {
					writeError(w, "internal server error", http.StatusInternalServerError)

					return
				}

				if isTokenRevoked {
					writeError(w, "invalid token", http.StatusUnauthorized)

					return
				}

				token, claims, err := tokenManager.ParseToken(stringToken)

				if err != nil || !token.Valid {
					writeError(w, "invalid token", http.StatusUnauthorized)

					return
				}

				user, err := userService.GetUserById(claims.UserId)

				if err != nil || user == nil {
					writeError(w, "invalid token", http.StatusUnauthorized)

					return
				}

				currentUser = user
			}

			ctx := context.WithValue(r.Context(), appContext.CurrentUserKey, currentUser)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

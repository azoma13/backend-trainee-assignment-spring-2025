package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/azoma13/backend-trainee-assignment-spring-2025/configs"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/internal/service"
	"github.com/azoma13/backend-trainee-assignment-spring-2025/models"
	"github.com/golang-jwt/jwt"
)

func NewAuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				service.ResponseJSON(w, http.StatusUnauthorized, models.ErrorResponse{
					Errors: "Missing Authorization header",
				})
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if tokenString == authHeader || tokenString == "" {
				service.ResponseJSON(w, http.StatusUnauthorized, models.ErrorResponse{
					Errors: "Invalid Authorization header format",
				})
				return
			}

			token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("invalid signing method")
				}
				return []byte(configs.SecretJWTKey), nil
			})

			if err != nil || !token.Valid {
				service.ResponseJSON(w, http.StatusUnauthorized, models.ErrorResponse{
					Errors: "Invalid token",
				})
				return
			}

			claims, ok := token.Claims.(*models.Claims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), models.RoleKey, claims.Role)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

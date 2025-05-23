package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Ambil dari .env

// AuthMiddleware checks for valid JWT token
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		
		sub, ok := claims["sub"].(string)
		if !ok {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}
		userID, err := strconv.Atoi(sub)
		if err != nil {
			http.Error(w, "Invalid user ID format", http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := config.DB.First(&user, userID).Error; err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)
	fmt.Fprintf(w, "Halo %s, kamu berhasil login!", user.Name)
}

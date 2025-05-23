package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"

	"github.com/golang-jwt/jwt/v4"
)

// JWT secret key (harus disimpan di .env atau konfigurasi yang lebih aman)
var jwtSecret = []byte("your-secret-key-1234567890") // Ganti dengan secret key yang kuat dan rahasia

// AuthMiddleware adalah middleware untuk memvalidasi token JWT
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return jwtSecret, nil
		})

		// Jika token tidak valid
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Ambil claims dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Ambil user ID dari token
		userID, ok := claims["sub"].(float64)
		if !ok {
			http.Error(w, "Invalid user ID in token", http.StatusUnauthorized)
			return
		}

		// Cek apakah user ada di database
		var user models.User
		if err := config.DB.First(&user, uint(userID)).Error; err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Set user ID ke context (opsional)
		// r = r.WithContext(context.WithValue(r.Context(), "user_id", userID))

		// Lanjutkan ke handler berikutnya
		next(w, r)
	}
}

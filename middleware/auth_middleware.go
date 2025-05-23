package middleware

import (
	"tryoutcpns-project-backend/config"
	"tryoutcpns-project-backend/models"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validasi token (misalnya dengan JWT atau session)
		// Untuk sementara, kita asumsikan token valid
		next(w, r)
	}
}

package middleware

import (
	"net/http"

	"github.com/ulule/limiter/v3"
)

// RateLimitMiddleware returns a middleware that limits the number of requests per user
func RateLimitMiddleware(limiter *limiter.Limiter) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr

			if _, err := limiter.Get(r.Context(), ip); err != nil {
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

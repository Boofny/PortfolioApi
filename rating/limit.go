// Package rating is just to store the rate limiter for the api
package rating

import (
	"net/http"

	goliveMiddleware "github.com/Boofny/goLive/middleware"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 15)

func RateLimit(allowedOrigins... string) goliveMiddleware.Middleware {
	return func (next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "rate limit reached", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

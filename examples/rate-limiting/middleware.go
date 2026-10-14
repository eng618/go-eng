package ratelimiting

import (
	"net/http"
	"strconv"
	"time"
)

// Limiter gates one request. Both FixedWindow and TokenBucket implement it.
type Limiter interface {
	Allow() bool
}

// WithRateLimit wraps next with abuse prevention. Over-limit requests get
// 429 with a Retry-After hint.
func WithRateLimit(limiter Limiter, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.Header().Set("Retry-After", strconv.Itoa(int(time.Second.Seconds())))
			http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)

			return
		}

		next.ServeHTTP(w, r)
	})
}

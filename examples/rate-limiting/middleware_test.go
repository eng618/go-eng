package ratelimiting_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	ratelimiting "github.com/eng618/go-eng/examples/rate-limiting"
)

func ExampleWithRateLimit() {
	limiter := ratelimiting.NewFixedWindow(1, time.Second)
	handler := ratelimiting.WithRateLimit(
		limiter,
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			_, _ = fmt.Fprintln(w, "ok")
		}),
	)

	first := httptest.NewRecorder()
	handler.ServeHTTP(first, httptest.NewRequest(http.MethodGet, "/gated", nil))
	fmt.Println(first.Code)

	second := httptest.NewRecorder()
	handler.ServeHTTP(second, httptest.NewRequest(http.MethodGet, "/gated", nil))
	fmt.Println(second.Code)

	// Output:
	// 200
	// 429
}

func TestWithRateLimit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		limit      int
		requests   int
		wantStatus []int
	}{
		{
			name:       "passes then gates",
			limit:      1,
			requests:   2,
			wantStatus: []int{http.StatusOK, http.StatusTooManyRequests},
		},
		{
			name:       "zero limit always gates",
			limit:      0,
			requests:   2,
			wantStatus: []int{http.StatusTooManyRequests, http.StatusTooManyRequests},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			handler := ratelimiting.WithRateLimit(
				ratelimiting.NewFixedWindow(tt.limit, time.Minute),
				http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
					w.WriteHeader(http.StatusOK)
				}),
			)

			for i := range tt.requests {
				rec := httptest.NewRecorder()
				handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/gated", nil))
				if got := rec.Code; got != tt.wantStatus[i] {
					t.Errorf("request %d status = %v, want %v", i, got, tt.wantStatus[i])
				}
			}
		})
	}
}

func BenchmarkWithRateLimit(b *testing.B) {
	handler := ratelimiting.WithRateLimit(
		ratelimiting.NewFixedWindow(b.N, time.Minute),
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		}),
	)

	for range b.N {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/gated", nil))
	}
}

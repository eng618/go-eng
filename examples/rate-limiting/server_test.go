package ratelimiting_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	ratelimiting "github.com/eng618/go-eng/examples/rate-limiting"
)

func ExampleNewMux() {
	mux := ratelimiting.NewMux(ratelimiting.NewFixedWindow(1, time.Second))

	ping := httptest.NewRecorder()
	mux.ServeHTTP(ping, httptest.NewRequest(http.MethodGet, "/ping", nil))
	fmt.Println(ping.Code)

	first := httptest.NewRecorder()
	mux.ServeHTTP(first, httptest.NewRequest(http.MethodGet, "/gated", nil))
	fmt.Println(first.Code)

	second := httptest.NewRecorder()
	mux.ServeHTTP(second, httptest.NewRequest(http.MethodGet, "/gated", nil))
	fmt.Println(second.Code)

	// Output:
	// 200
	// 200
	// 429
}

func TestNewMux(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		path       string
		requests   int
		wantStatus []int
	}{
		{
			name:       "ping stays ungated",
			path:       "/ping",
			requests:   3,
			wantStatus: []int{http.StatusOK, http.StatusOK, http.StatusOK},
		},
		{
			name:       "gated gates abuse",
			path:       "/gated",
			requests:   3,
			wantStatus: []int{http.StatusOK, http.StatusOK, http.StatusTooManyRequests},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mux := ratelimiting.NewMux(ratelimiting.NewFixedWindow(2, time.Minute))
			for i := range tt.requests {
				rec := httptest.NewRecorder()
				mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, tt.path, nil))
				if got := rec.Code; got != tt.wantStatus[i] {
					t.Errorf("request %d to %s status = %v, want %v", i, tt.path, got, tt.wantStatus[i])
				}
			}
		})
	}
}

func BenchmarkNewMux(b *testing.B) {
	mux := ratelimiting.NewMux(ratelimiting.NewFixedWindow(b.N, time.Minute))
	for range b.N {
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/gated", nil))
	}
}

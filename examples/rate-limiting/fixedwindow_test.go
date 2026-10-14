package ratelimiting_test

import (
	"fmt"
	"testing"
	"time"

	ratelimiting "github.com/eng618/go-eng/examples/rate-limiting"
)

func ExampleNewFixedWindow() {
	limiter := ratelimiting.NewFixedWindow(1, time.Second)
	fmt.Println(limiter.Allow())
	fmt.Println(limiter.Allow())
	// Output:
	// true
	// false
}

func TestFixedWindow_Allow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		limit  int
		calls  int
		wanted []bool
	}{
		{name: "allows up to limit", limit: 2, calls: 3, wanted: []bool{true, true, false}},
		{name: "zero limit denies", limit: 0, calls: 2, wanted: []bool{false, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			limiter := ratelimiting.NewFixedWindow(tt.limit, time.Minute)
			for i := range tt.calls {
				if got := limiter.Allow(); got != tt.wanted[i] {
					t.Errorf("Allow() call %d = %v, want %v", i, got, tt.wanted[i])
				}
			}
		})
	}
}

func BenchmarkFixedWindow_Allow(b *testing.B) {
	limiter := ratelimiting.NewFixedWindow(b.N, time.Minute)
	for range b.N {
		limiter.Allow()
	}
}

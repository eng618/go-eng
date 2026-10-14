package ratelimiting_test

import (
	"fmt"
	"testing"
	"time"

	ratelimiting "github.com/eng618/go-eng/examples/rate-limiting"
)

func ExampleNewTokenBucket() {
	limiter := ratelimiting.NewTokenBucket(1, time.Second)
	fmt.Println(limiter.Allow())
	fmt.Println(limiter.Allow())
	// Output:
	// true
	// false
}

func TestTokenBucket_Allow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		capacity int
		calls    int
		wanted   []bool
	}{
		{name: "allows burst then denies", capacity: 2, calls: 3, wanted: []bool{true, true, false}},
		{name: "zero capacity denies", capacity: 0, calls: 2, wanted: []bool{false, false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			limiter := ratelimiting.NewTokenBucket(tt.capacity, time.Minute)
			for i := range tt.calls {
				if got := limiter.Allow(); got != tt.wanted[i] {
					t.Errorf("Allow() call %d = %v, want %v", i, got, tt.wanted[i])
				}
			}
		})
	}
}

func BenchmarkTokenBucket_Allow(b *testing.B) {
	limiter := ratelimiting.NewTokenBucket(b.N, time.Minute)
	for range b.N {
		limiter.Allow()
	}
}

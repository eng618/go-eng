package ratelimiting

import (
	"sync"
	"time"
)

// TokenBucket allows short bursts up to capacity, refilling one token every
// refillEvery. It is smoother than FixedWindow for abuse prevention because
// legitimate bursts pass while sustained abuse is still gated.
type TokenBucket struct {
	mu          sync.Mutex
	capacity    int
	tokens      float64
	refillEvery time.Duration
	lastRefill  time.Time
}

// NewTokenBucket returns a TokenBucket starting full with capacity tokens.
func NewTokenBucket(capacity int, refillEvery time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:    capacity,
		tokens:      float64(capacity),
		refillEvery: refillEvery,
		lastRefill:  time.Now(),
	}
}

// Allow reports whether one request may proceed.
func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.capacity <= 0 {
		return false
	}

	now := time.Now()
	if b.refillEvery > 0 {
		elapsed := now.Sub(b.lastRefill)
		b.tokens += elapsed.Seconds() / b.refillEvery.Seconds()
		if b.tokens > float64(b.capacity) {
			b.tokens = float64(b.capacity)
		}
		b.lastRefill = now
	}

	if b.tokens < 1 {
		return false
	}

	b.tokens--

	return true
}

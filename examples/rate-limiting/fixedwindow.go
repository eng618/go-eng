package ratelimiting

import (
	"sync"
	"time"
)

// FixedWindow allows up to limit requests per window.
// It is the simplest abuse-prevention gate: count requests in the current
// window, reset the count when the window expires.
type FixedWindow struct {
	mu     sync.Mutex
	limit  int
	window time.Duration
	count  int
	start  time.Time
}

// NewFixedWindow returns a FixedWindow that allows limit requests per window.
func NewFixedWindow(limit int, window time.Duration) *FixedWindow {
	return &FixedWindow{
		limit:  limit,
		window: window,
		start:  time.Now(),
	}
}

// Allow reports whether one request may proceed.
func (f *FixedWindow) Allow() bool {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.limit <= 0 {
		return false
	}

	now := time.Now()
	if now.Sub(f.start) >= f.window {
		f.start = now
		f.count = 0
	}

	if f.count >= f.limit {
		return false
	}

	f.count++

	return true
}

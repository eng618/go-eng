package breaker

import (
	"errors"
	"testing"
	"time"
)

// Example tests.
func ExampleCircuitBreaker_successful() {
	cb := NewCircuitBreaker(3, 5*time.Second)
	err := cb.Execute(func() error {
		return nil // successful operation
	})
	if err != nil {
		panic(err)
	}
	// Output:
}

func ExampleCircuitBreaker_failing() {
	cb := NewCircuitBreaker(3, 5*time.Second)
	err := cb.Execute(func() error {
		return errors.New("operation failed")
	})
	if err == nil {
		panic("expected error")
	}
	// Output:
}

// Unit tests.
func TestNewCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(3, 5*time.Second)
	if cb.state != StateClosed {
		t.Errorf("expected initial state to be Closed, got %v", cb.state)
	}
	if cb.failureThreshold != 3 {
		t.Errorf("expected failure threshold to be 3, got %d", cb.failureThreshold)
	}
	if cb.resetTimeout != 5*time.Second {
		t.Errorf("expected reset timeout to be 5s, got %v", cb.resetTimeout)
	}
}

func TestCircuitBreakerTransitions(t *testing.T) {
	tests := []struct {
		name          string
		operations    []error
		expectedState State
		sleepBetween  time.Duration
	}{
		{
			name:          "stays closed on success",
			operations:    []error{nil, nil, nil},
			expectedState: StateClosed,
		},
		{
			name:          "opens after threshold failures",
			operations:    []error{errors.New("err"), errors.New("err"), errors.New("err")},
			expectedState: StateOpen,
		},
		{
			name:          "mixed success and failure stays closed",
			operations:    []error{errors.New("err"), nil, errors.New("err")},
			expectedState: StateClosed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := NewCircuitBreaker(3, 5*time.Second)

			for _, err := range tt.operations {
				if execErr := cb.Execute(func() error {
					return err
				}); execErr != nil && err == nil {
					t.Errorf("unexpected error: %v", execErr)
				}
				if tt.sleepBetween > 0 {
					time.Sleep(tt.sleepBetween)
				}
			}

			if cb.State() != tt.expectedState {
				t.Errorf("expected state %v, got %v", tt.expectedState, cb.State())
			}
		})
	}
}

func TestCircuitBreakerTimeout(t *testing.T) {
	cb := NewCircuitBreaker(2, 100*time.Millisecond)

	// Force circuit to open
	if err := cb.Execute(func() error { return errors.New("err") }); err == nil {
		t.Error("expected error but got nil")
	}
	if err := cb.Execute(func() error { return errors.New("err") }); err == nil {
		t.Error("expected error but got nil")
	}

	if cb.State() != StateOpen {
		t.Fatal("circuit should be open")
	}

	// Wait for timeout
	time.Sleep(150 * time.Millisecond)

	// Next execution should be in half-open state
	err := cb.Execute(func() error { return nil })
	if err != nil {
		t.Error("expected successful execution after timeout")
	}
	if cb.State() != StateClosed {
		t.Errorf("expected state Closed after successful execution, got %v", cb.State())
	}
}

func TestConcurrentAccess(_ *testing.T) {
	cb := NewCircuitBreaker(100, time.Second)
	done := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				_ = cb.Execute(func() error {
					return errors.New("err")
				})
			}
			done <- true
		}()
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
}

// Fuzz testing.
func FuzzCircuitBreaker(f *testing.F) {
	f.Add(uint(3), int64(5000))
	f.Fuzz(func(t *testing.T, threshold uint, resetMs int64) {
		if threshold == 0 || resetMs < 0 {
			return // Skip invalid inputs
		}

		cb := NewCircuitBreaker(threshold, time.Duration(resetMs)*time.Millisecond)

		// Perform random operations
		for i := 0; i < 10; i++ {
			_ = cb.Execute(func() error {
				if i%2 == 0 {
					return errors.New("random error")
				}
				return nil
			})
		}

		// Verify the circuit breaker is in a valid state
		state := cb.State()
		if state != StateClosed && state != StateOpen && state != StateHalfOpen {
			t.Errorf("invalid state: %v", state)
		}
	})
}

// Benchmark tests.
func BenchmarkCircuitBreakerSuccess(b *testing.B) {
	cb := NewCircuitBreaker(3, 5*time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = cb.Execute(func() error {
			return nil
		})
	}
}

func BenchmarkCircuitBreakerFailure(b *testing.B) {
	cb := NewCircuitBreaker(3, 5*time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = cb.Execute(func() error {
			return errors.New("err")
		})
	}
}

func BenchmarkCircuitBreakerParallel(b *testing.B) {
	cb := NewCircuitBreaker(1000, 5*time.Second)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = cb.Execute(func() error {
				return nil
			})
		}
	})
}

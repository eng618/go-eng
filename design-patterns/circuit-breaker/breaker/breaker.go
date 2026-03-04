package breaker

import (
	"errors"
	"sync"
	"time"
)

// State represents the current state of a circuit breaker.
type State int

// Circuit breaker states.
const (
	// StateClosed allows requests to execute normally.
	StateClosed State = iota
	// StateOpen blocks requests until reset timeout elapses.
	StateOpen
	// StateHalfOpen allows trial requests after timeout.
	StateHalfOpen
)

// CircuitBreaker controls request execution based on recent failures.
type CircuitBreaker struct {
	mutex sync.RWMutex

	state            State
	failureThreshold uint
	failureCount     uint
	resetTimeout     time.Duration
	lastFailureTime  time.Time
}

// NewCircuitBreaker creates a CircuitBreaker with failure threshold and reset timeout.
func NewCircuitBreaker(failureThreshold uint, resetTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:            StateClosed,
		failureThreshold: failureThreshold,
		resetTimeout:     resetTimeout,
	}
}

// Execute runs the provided operation unless the circuit is currently open.
func (cb *CircuitBreaker) Execute(operation func() error) error {
	if !cb.shouldExecute() {
		return errors.New("circuit breaker is open")
	}

	err := operation()
	cb.recordResult(err)
	return err
}

func (cb *CircuitBreaker) shouldExecute() bool {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()

	switch cb.state {
	case StateClosed:
		return true
	case StateOpen:
		if time.Since(cb.lastFailureTime) >= cb.resetTimeout {
			cb.mutex.RUnlock()
			cb.mutex.Lock()
			cb.state = StateHalfOpen
			cb.mutex.Unlock()
			cb.mutex.RLock()
			return true
		}
		return false
	case StateHalfOpen:
		return true
	default:
		return false
	}
}

func (cb *CircuitBreaker) recordResult(err error) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if err != nil {
		cb.failureCount++
		cb.lastFailureTime = time.Now()

		if cb.state == StateHalfOpen || cb.failureCount >= cb.failureThreshold {
			cb.state = StateOpen
		}
	} else {
		if cb.state == StateHalfOpen {
			cb.state = StateClosed
			cb.failureCount = 0
		}
	}
}

// State returns the current circuit breaker state.
func (cb *CircuitBreaker) State() State {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.state
}

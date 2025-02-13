package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/eng618/go-eng/examples/circuit-breaker/breaker"
)

// simulateUnstableService simulates an external service that sometimes fails
func simulateUnstableService() error {
	// Simulate random failures
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Float64() < 0.6 { // 60% chance of failure
		return errors.New("service error: connection timeout")
	}
	return nil
}

func main() {
	// Create a circuit breaker that opens after 3 failures and resets after 5 seconds
	cb := breaker.NewCircuitBreaker(3, 5*time.Second)

	// Run several attempts to call the service
	for i := 0; i < 10; i++ {
		err := cb.Execute(simulateUnstableService)
		if err != nil {
			fmt.Printf("Attempt %d: Failed - %v (Circuit State: %v)\n", i+1, err, cb.State())
		} else {
			fmt.Printf("Attempt %d: Success (Circuit State: %v)\n", i+1, cb.State())
		}
		time.Sleep(time.Second)
	}
}

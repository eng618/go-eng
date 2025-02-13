package main

import (
	"fmt"
	"time"

	"github.com/eng618/go-eng/design-patterns/circuit-breaker/breaker"
)

func main() {
	// Create a circuit breaker that opens after 3 failures and resets after 5 seconds
	cb := breaker.NewCircuitBreaker(3, 5*time.Second)

	// Demonstrate success case
	fmt.Println("Trying successful operation...")
	err := cb.Execute(func() error {
		fmt.Println("Operation executed successfully")
		return nil
	})
	fmt.Printf("Result: %v, State: %v\n\n", err, cb.State())

	// Demonstrate failure case
	fmt.Println("Trying failing operations...")
	for i := 0; i < 4; i++ {
		err := cb.Execute(func() error {
			return fmt.Errorf("simulated error")
		})
		fmt.Printf("Attempt %d - Result: %v, State: %v\n", i+1, err, cb.State())
	}

	fmt.Println("\nWaiting for circuit breaker to reset...")
	time.Sleep(6 * time.Second)

	// Try again after reset
	fmt.Println("\nTrying operation after reset...")
	err = cb.Execute(func() error {
		fmt.Println("Operation executed successfully")
		return nil
	})
	fmt.Printf("Result: %v, State: %v\n", err, cb.State())
}

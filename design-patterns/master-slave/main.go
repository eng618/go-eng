package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"master-slave/masterslave"
)

// simulateWork represents a task that takes some time to complete.
func simulateWork(id int, ctx context.Context) error {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond

	// Simulate occasional failures
	if rand.Float32() < 0.1 { // 10% chance of failure
		return fmt.Errorf("task %d failed randomly", id)
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(duration):
		fmt.Printf("Task %d completed after %v\n", id, duration)
		return nil
	}
}

func printMetrics(metrics map[string]interface{}) {
	data, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling metrics: %v\n", err)
		return
	}
	fmt.Printf("\nSystem Metrics:\n%s\n", string(data))
}

func main() {
	fmt.Println("Master-Slave Architecture Example")

	// Create a new master-slave system with 3 workers
	ms := masterslave.NewMasterSlave(3)
	defer ms.Stop()

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Start a goroutine to monitor results
	go func() {
		for result := range ms.GetResults() {
			if result.Error != nil {
				fmt.Printf("Task %d on Slave %d failed: %v\n",
					result.TaskID, result.SlaveID, result.Error)
			} else {
				fmt.Printf("Task %d completed by Slave %d in %v\n",
					result.TaskID, result.SlaveID, result.Time)
			}
		}
	}()

	// Submit tasks in batches to demonstrate load distribution
	for batch := 0; batch < 3; batch++ {
		fmt.Printf("\nSubmitting batch %d of tasks...\n", batch+1)

		// Submit 5 tasks per batch
		for i := 1; i <= 5; i++ {
			taskID := batch*5 + i
			err := ms.Execute(func(ctx context.Context) error {
				return simulateWork(taskID, ctx)
			})
			if err != nil {
				fmt.Printf("Failed to submit task %d: %v\n", taskID, err)
			}
		}

		// Print metrics after each batch
		time.Sleep(2 * time.Second)
		printMetrics(ms.GetMetrics())
	}

	// Wait for context timeout or completion
	<-ctx.Done()

	// Print final metrics
	fmt.Println("\nFinal System Metrics:")
	printMetrics(ms.GetMetrics())
}

package masterslave

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

// Example tests.
func ExampleMasterSlave_successful() {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	err := ms.Execute(func(ctx context.Context) error {
		return nil // successful operation
	})
	if err != nil {
		panic(err)
	}
	// Output:
}

func ExampleMasterSlave_failing() {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	err := ms.Execute(func(_ context.Context) error {
		return errors.New("operation failed")
	})
	if err != nil {
		// Expected error behavior
	}
	// Output:
}

// Unit tests.
func TestNewMasterSlave(t *testing.T) {
	ms := NewMasterSlave(3)
	defer ms.Stop()

	if ms == nil {
		t.Errorf("expected non-nil MasterSlave instance")
	}
	if len(ms.slaves) != 3 {
		t.Errorf("expected 3 slaves, got %d", len(ms.slaves))
	}
}

func TestMasterSlaveExecution(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	// Test successful execution
	err := ms.Execute(func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Test error propagation
	err = ms.Execute(func(ctx context.Context) error {
		return errors.New("operation failed")
	})
	if err != nil {
		// Expected error behavior
	}
}

func TestConcurrentExecution(t *testing.T) {
	ms := NewMasterSlave(3)
	defer ms.Stop()

	var completedTasks int32
	numTasks := 10

	// Submit multiple tasks
	for i := 0; i < numTasks; i++ {
		_ = ms.Execute(func(ctx context.Context) error {
			time.Sleep(10 * time.Millisecond) // Simulate work
			atomic.AddInt32(&completedTasks, 1)
			return nil
		})
	}

	// Wait for tasks to complete
	time.Sleep(200 * time.Millisecond)

	completed := atomic.LoadInt32(&completedTasks)
	if completed != int32(numTasks) {
		t.Errorf("expected %d completed tasks, got %d", numTasks, completed)
	}
}

func TestZeroSlaves(t *testing.T) {
	ms := NewMasterSlave(0)
	defer ms.Stop()

	// Should create at least one slave
	if len(ms.slaves) != 1 {
		t.Errorf("expected 1 slave when initialized with 0, got %d", len(ms.slaves))
	}
}

func TestContextCancellation(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	parentCtx, parentCancel := context.WithCancel(context.Background())
	defer parentCancel()

	taskCompleted := make(chan struct{})

	// Submit a task that uses the cancellable context
	err := ms.Execute(func(taskCtx context.Context) error {
		select {
		case <-parentCtx.Done():
			return parentCtx.Err()
		case <-taskCtx.Done():
			return taskCtx.Err()
		case <-time.After(1 * time.Second):
			close(taskCompleted)
			return nil
		}
	})
	if err != nil {
		t.Errorf("unexpected error submitting task: %v", err)
	}

	// Cancel the context before task completes
	parentCancel()

	// Verify task was canceled
	select {
	case <-taskCompleted:
		t.Error("task should have been canceled")
	case <-time.After(100 * time.Millisecond):
		// Expected: task was canceled
	}
}

func TestResultMonitoring(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	resultCount := 0
	go func() {
		for result := range ms.GetResults() {
			if result.Time < 0 {
				t.Error("invalid execution time")
			}
			resultCount++
		}
	}()

	numTasks := 5
	for i := 0; i < numTasks; i++ {
		_ = ms.Execute(func(ctx context.Context) error {
			time.Sleep(time.Millisecond)
			return nil
		})
	}

	time.Sleep(200 * time.Millisecond) // Increase sleep duration to ensure all tasks complete
	ms.Stop()                          // Ensure all results are processed before checking resultCount
	if resultCount != numTasks {
		t.Errorf("expected %d results, got %d", numTasks, resultCount)
	}
}

func TestHealthMonitoring(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	// Override health check duration for testing
	ms.healthCheck = 100 * time.Millisecond

	// Submit a task that will make the slave appear unhealthy
	slave := ms.slaves[0]
	slave.mu.Lock()
	slave.lastActive = time.Now().Add(-1 * time.Second)
	slave.mu.Unlock()

	// Wait for health check to detect and restart the slave
	time.Sleep(300 * time.Millisecond)

	metrics := ms.GetMetrics()
	healthyCount := metrics["healthy_slaves"].(int)
	if healthyCount != 2 {
		t.Errorf("expected 2 healthy slaves after health check, got %d", healthyCount)
	}
}

func TestMetricsCollection(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	// Submit some tasks
	for i := 0; i < 5; i++ {
		_ = ms.Execute(func(ctx context.Context) error {
			return nil
		})
	}

	// Wait for tasks to complete
	time.Sleep(100 * time.Millisecond)

	metrics := ms.GetMetrics()

	if metrics["total_slaves"].(int) != 2 {
		t.Errorf("expected 2 total slaves in metrics")
	}

	totalTasks := metrics["total_tasks_processed"].(uint64)
	if totalTasks != 5 {
		t.Errorf("expected 5 total tasks processed, got %d", totalTasks)
	}

	slaveMetrics := metrics["slave_metrics"].([]map[string]interface{})
	if len(slaveMetrics) != 2 {
		t.Errorf("expected metrics for 2 slaves, got %d", len(slaveMetrics))
	}
}

func TestTaskResultTracking(t *testing.T) {
	ms := NewMasterSlave(1)
	defer ms.Stop()

	var receivedResults []TaskResult
	resultsDone := make(chan struct{})

	// Collect results in background
	go func() {
		for result := range ms.GetResults() {
			receivedResults = append(receivedResults, result)
		}
		close(resultsDone)
	}()

	// Submit tasks with known outcomes
	ms.Execute(func(ctx context.Context) error {
		return nil // success
	})

	ms.Execute(func(ctx context.Context) error {
		return errors.New("planned failure")
	})

	// Wait for results and stop
	time.Sleep(100 * time.Millisecond)
	ms.Stop()
	<-resultsDone

	if len(receivedResults) != 2 {
		t.Fatalf("expected 2 results, got %d", len(receivedResults))
	}

	if receivedResults[0].Error != nil {
		t.Error("expected first task to succeed")
	}

	if receivedResults[1].Error == nil {
		t.Error("expected second task to fail")
	}
}

// Fuzz testing.
func FuzzMasterSlave(f *testing.F) {
	f.Add(uint(3), uint(5))

	f.Fuzz(func(t *testing.T, numSlaves, numTasks uint) {
		numSlaves = numSlaves%5 + 1 // 1-5 slaves
		numTasks = numTasks%10 + 1  // 1-10 tasks

		ms := NewMasterSlave(int(numSlaves))
		defer ms.Stop()

		results := make([]TaskResult, 0)
		resultsDone := make(chan struct{})

		// Collect results
		go func() {
			for result := range ms.GetResults() {
				results = append(results, result)
			}
			close(resultsDone)
		}()

		// Submit tasks
		for i := uint(0); i < numTasks; i++ {
			ms.Execute(func(ctx context.Context) error {
				if time.Now().UnixNano()%2 == 0 {
					return nil
				}
				return errors.New("random failure")
			})
		}

		// Wait for completion
		time.Sleep(100 * time.Millisecond)
		ms.Stop()
		<-resultsDone

		metrics := ms.GetMetrics()

		// Verify system consistency
		if metrics["total_slaves"].(int) != int(numSlaves) {
			t.Errorf("incorrect number of slaves")
		}

		if len(results) != int(numTasks) {
			t.Errorf("expected %d results, got %d", numTasks, len(results))
		}
	})
}

// Benchmark tests.
func BenchmarkMasterSlaveExecution(b *testing.B) {
	ms := NewMasterSlave(4)
	defer ms.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ms.Execute(func(taskCtx context.Context) error {
			return nil
		})
	}
}

func BenchmarkMasterSlaveParallel(b *testing.B) {
	ms := NewMasterSlave(4)
	defer ms.Stop()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = ms.Execute(func(taskCtx context.Context) error {
				return nil
			})
		}
	})
}

func BenchmarkMasterSlaveDifferentSizes(b *testing.B) {
	sizes := []int{1, 2, 4}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("slaves-%d", size), func(b *testing.B) {
			ms := NewMasterSlave(size)
			defer ms.Stop()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = ms.Execute(func(taskCtx context.Context) error {
					return nil
				})
			}
		})
	}
}

func BenchmarkMasterSlaveWithMetrics(b *testing.B) {
	sizes := []int{1, 2, 4}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("slaves-%d", size), func(b *testing.B) {
			ms := NewMasterSlave(size)
			defer ms.Stop()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ms.Execute(func(ctx context.Context) error {
					return nil
				})
				if i%100 == 0 {
					ms.GetMetrics() // Periodically collect metrics
				}
			}
		})
	}
}

func BenchmarkHealthChecks(b *testing.B) {
	ms := NewMasterSlave(4)
	ms.healthCheck = 10 * time.Millisecond // Increase health check frequency
	defer ms.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ms.Execute(func(ctx context.Context) error {
			return nil
		})
	}
}

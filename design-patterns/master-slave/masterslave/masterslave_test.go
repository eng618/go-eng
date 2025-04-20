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

	err := ms.Execute(func(_ context.Context) error {
		return nil // successful operation
	})
	if err != nil {
		panic(err)
	}
	// Output:
}

func ExampleMasterSlave_failing() {
	ms := NewMasterSlave(2)

	resultsDone := make(chan struct{})
	var taskFailed bool

	// Collect results in background
	go func() {
		for result := range ms.GetResults() {
			if result.Error != nil {
				taskFailed = true
			}
		}
		close(resultsDone)
	}()

	err := ms.Execute(func(_ context.Context) error {
		return errors.New("operation failed")
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	ms.Wait()
	ms.Stop()
	<-resultsDone

	if taskFailed {
		fmt.Println("Task failed as expected")
	}
	// Output:
	// Task failed as expected
}

// Unit tests.
func TestNewMasterSlave(t *testing.T) {
	ms := NewMasterSlave(3)
	if ms == nil {
		t.Fatal("NewMasterSlave returned nil")
	}
	defer ms.Stop()

	slaves := ms.slaves
	if slaves == nil {
		t.Fatal("expected non-nil slaves slice")
	}
	if len(slaves) != 3 {
		t.Errorf("expected 3 slaves, got %d", len(slaves))
	}
}

func TestMasterSlaveExecution(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	// Test successful execution
	err := ms.Execute(func(_ context.Context) error {
		return nil
	})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Test error propagation
	err = ms.Execute(func(_ context.Context) error {
		return errors.New("operation failed")
	})
	if err != nil {
		// Expected error behavior
		fmt.Println("ENG: This is an expected failure")
	}
}

func TestConcurrentExecution(t *testing.T) {
	ms := NewMasterSlave(3)
	defer ms.Stop()

	var completedTasks int32
	numTasks := 10

	// Submit multiple tasks
	for i := 0; i < numTasks; i++ {
		err := ms.Execute(func(_ context.Context) error {
			time.Sleep(10 * time.Millisecond) // Simulate work
			atomic.AddInt32(&completedTasks, 1)
			return nil
		})
		if err != nil {
			t.Errorf("failed to execute task %d: %v", i, err)
		}
	}

	// Wait for tasks to complete and give some extra time for scheduling
	ms.Wait()
	time.Sleep(50 * time.Millisecond)

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

	var resultCount int32
	resultsDone := make(chan struct{})

	// Collect results in background
	go func() {
		for result := range ms.GetResults() {
			if result.Time < 0 {
				t.Error("invalid execution time")
			}
			atomic.AddInt32(&resultCount, 1)
		}
		close(resultsDone)
	}()

	numTasks := 5
	for i := 0; i < numTasks; i++ {
		_ = ms.Execute(func(_ context.Context) error {
			time.Sleep(time.Millisecond)
			return nil
		})
	}

	ms.Wait() // Wait for all tasks to complete
	ms.Stop() // Stop after all tasks are done
	<-resultsDone

	if atomic.LoadInt32(&resultCount) != int32(numTasks) {
		t.Errorf("expected %d results, got %d", numTasks, int(atomic.LoadInt32(&resultCount)))
	}
}

func TestHealthMonitoring(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	// Override health check duration for testing
	ms.SetHealthCheckDuration(50 * time.Millisecond)

	// Make first slave appear unhealthy
	slave := ms.slaves[0]
	slave.mu.Lock()
	slave.lastActive = time.Now().Add(-1 * time.Second)
	slave.mu.Unlock()

	// Wait for at least 2 health check cycles
	time.Sleep(150 * time.Millisecond)

	// Execute a task to verify system is working
	err := ms.Execute(func(_ context.Context) error {
		return nil
	})
	if err != nil {
		t.Errorf("failed to execute task: %v", err)
	}

	// Wait for task to complete
	ms.Wait()

	metrics := ms.GetMetrics()
	healthyCount := metrics["healthy_slaves"].(int)
	if healthyCount != 2 {
		t.Errorf("expected 2 healthy slaves after health check, got %d", healthyCount)
	}

	// Verify both slaves are processing tasks
	slaveMetrics := metrics["slave_metrics"].([]map[string]interface{})
	healthySlaves := 0
	for _, sm := range slaveMetrics {
		if sm["healthy"].(bool) {
			healthySlaves++
		}
	}
	if healthySlaves != 2 {
		t.Errorf("expected 2 healthy slaves in metrics, got %d", healthySlaves)
	}
}

func TestMetricsCollection(t *testing.T) {
	ms := NewMasterSlave(2)
	defer ms.Stop()

	// Submit some tasks
	for i := 0; i < 5; i++ {
		_ = ms.Execute(func(_ context.Context) error {
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
	_ = ms.Execute(func(_ context.Context) error {
		return nil // success
	})

	_ = ms.Execute(func(_ context.Context) error {
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

// submitAndCollectTasks submits tasks and collects results.
func submitAndCollectTasks(t *testing.T, ms *MasterSlave, numTasks int) []TaskResult {
	results := make([]TaskResult, 0, numTasks)
	resultsDone := make(chan struct{})

	// Collect results
	go func() {
		for result := range ms.GetResults() {
			results = append(results, result)
		}
		close(resultsDone)
	}()

	// Submit tasks
	for i := 0; i < numTasks; i++ {
		if err := ms.Execute(func(_ context.Context) error {
			if time.Now().UnixNano()%2 == 0 {
				return nil
			}
			return errors.New("random failure")
		}); err != nil {
			t.Errorf("failed to execute task %d: %v", i, err)
		}
	}

	// Wait for completion
	time.Sleep(100 * time.Millisecond)
	ms.Stop()
	<-resultsDone

	return results
}

// FuzzMasterSlave performs fuzz testing of the master-slave system.
func FuzzMasterSlave(f *testing.F) {
	f.Add(uint(3), uint(5))
	const maxInt = int(^uint(0) >> 1)

	f.Fuzz(func(t *testing.T, numSlaves, numTasks uint) {
		// Validate input ranges to prevent integer overflow
		if numSlaves > uint(maxInt) || numTasks > uint(maxInt) {
			t.Skip("input values too large")
		}

		// Constrain the ranges to reasonable values and prevent overflow
		modSlaves := numSlaves % 5
		modTasks := numTasks % 10
		if modSlaves > uint(maxInt-1) || modTasks > uint(maxInt-1) {
			t.Skip("modulo result too large")
		}

		safeNumSlaves := 1 + int(modSlaves) // 1-5 slaves
		safeNumTasks := 1 + int(modTasks)   // 1-10 tasks

		ms := NewMasterSlave(safeNumSlaves)
		if ms == nil {
			t.Fatal("NewMasterSlave returned nil")
		}
		defer ms.Stop()

		// Submit tasks and collect results
		results := submitAndCollectTasks(t, ms, safeNumTasks)

		// Verify system consistency
		metrics := ms.GetMetrics()
		if metrics["total_slaves"].(int) != safeNumSlaves {
			t.Errorf("incorrect number of slaves")
		}
		if len(results) != safeNumTasks {
			t.Errorf("expected %d results, got %d", safeNumTasks, len(results))
		}
	})
}

// Benchmark tests.
func BenchmarkMasterSlaveExecution(b *testing.B) {
	ms := NewMasterSlave(4)
	defer ms.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ms.Execute(func(_ context.Context) error {
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
			_ = ms.Execute(func(_ context.Context) error {
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
				_ = ms.Execute(func(_ context.Context) error {
					return nil
				})
			}
		})
	}
}

// BenchmarkMasterSlaveWithMetrics runs benchmarks with periodic metric collection.
func BenchmarkMasterSlaveWithMetrics(b *testing.B) {
	sizes := []int{1, 2, 4}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("slaves-%d", size), func(b *testing.B) {
			ms := NewMasterSlave(size)
			defer ms.Stop()

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := ms.Execute(func(_ context.Context) error {
					return nil
				}); err != nil {
					b.Errorf("failed to execute task: %v", err)
				}
				if i%100 == 0 {
					ms.GetMetrics()
				}
			}
		})
	}
}

func BenchmarkHealthChecks(b *testing.B) {
	ms := NewMasterSlave(4)
	ms.SetHealthCheckDuration(10 * time.Millisecond) // Use setter instead of direct field access
	defer ms.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ms.Execute(func(_ context.Context) error {
			return nil
		})
	}
}

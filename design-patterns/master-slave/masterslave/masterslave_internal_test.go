package masterslave

import (
	"context"
	"testing"
	"time"
)

// White-box tests below access unexported internals (slave list, mutexes,
// timestamps) and therefore live in the internal test package. All other
// tests use the external masterslave_test package.

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

func TestZeroSlaves(t *testing.T) {
	ms := NewMasterSlave(0)
	defer ms.Stop()

	// Should create at least one slave
	if len(ms.slaves) != 1 {
		t.Errorf("expected 1 slave when initialized with 0, got %d", len(ms.slaves))
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
	slaveMetrics := metrics["slave_metrics"].([]map[string]any)
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

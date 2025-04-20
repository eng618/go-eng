package masterslave

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TaskStatus int

const (
	StatusPending TaskStatus = iota
	StatusRunning
	StatusCompleted
	StatusFailed
)

type TaskResult struct {
	Error   error
	Time    time.Duration
	Status  TaskStatus
	SlaveID int
	TaskID  uint64
}

// Task represents a unit of work to be processed.
type Task func(ctx context.Context) error

type taskWrapper struct {
	id       uint64
	task     Task
	status   TaskStatus
	assigned time.Time
}

// Slave represents a worker that can process tasks.
type Slave struct {
	id         int
	taskChan   chan *taskWrapper
	resultChan chan TaskResult
	done       chan struct{}
	ctx        context.Context
	cancel     context.CancelFunc
	taskCount  uint64
	isHealthy  bool
	lastActive time.Time
	mu         sync.RWMutex
}

// MasterSlave represents the master-slave architecture.
type MasterSlave struct {
	slaves      []*Slave
	numSlaves   int
	resultsChan chan TaskResult
	wg          sync.WaitGroup
	mu          sync.RWMutex
	slavesMu    sync.RWMutex // New mutex specifically for slaves management
	ctx         context.Context
	cancel      context.CancelFunc
	taskCounter uint64
	healthCheck time.Duration
	isStopped   bool
	stopOnce    sync.Once
}

// NewMasterSlave creates a new MasterSlave instance with the specified number of slaves.
func NewMasterSlave(numSlaves int) *MasterSlave {
	if numSlaves <= 0 {
		numSlaves = 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	ms := &MasterSlave{
		slaves:      make([]*Slave, numSlaves),
		numSlaves:   numSlaves,
		resultsChan: make(chan TaskResult, numSlaves*20), // Further increased buffer size
		ctx:         ctx,
		cancel:      cancel,
		healthCheck: 5 * time.Second,
	}

	// Initialize and start slave workers
	for i := range ms.slaves {
		slaveCtx, slaveCancel := context.WithCancel(ctx)
		slave := &Slave{
			id:         i + 1,
			taskChan:   make(chan *taskWrapper, 10), // Further increased buffer size
			resultChan: ms.resultsChan,
			done:       make(chan struct{}),
			ctx:        slaveCtx,
			cancel:     slaveCancel,
			isHealthy:  true,
			lastActive: time.Now(),
		}
		ms.slaves[i] = slave
		go ms.runSlave(slave)
	}

	// Start health monitoring
	go ms.monitorSlaveHealth()

	return ms
}

// GetHealthCheckDuration returns the current health check duration
func (ms *MasterSlave) GetHealthCheckDuration() time.Duration {
	ms.mu.RLock()
	defer ms.mu.RUnlock()
	return ms.healthCheck
}

// SetHealthCheckDuration sets the health check duration
func (ms *MasterSlave) SetHealthCheckDuration(d time.Duration) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.healthCheck = d
}

// monitorSlaveHealth periodically checks slave health.
func (ms *MasterSlave) monitorSlaveHealth() {
	duration := ms.GetHealthCheckDuration()
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {
		case <-ms.ctx.Done():
			return
		case <-ticker.C:
			ms.checkSlaveHealth()
			// Update ticker duration in case it changed
			newDuration := ms.GetHealthCheckDuration()
			if newDuration != duration {
				ticker.Reset(newDuration)
				duration = newDuration
			}
		}
	}
}

// checkSlaveHealth verifies the health of all slaves.
func (ms *MasterSlave) checkSlaveHealth() {
	// Get a snapshot of slaves under lock
	ms.slavesMu.RLock()
	slaves := make([]*Slave, len(ms.slaves))
	copy(slaves, ms.slaves)
	ms.slavesMu.RUnlock()

	healthCheck := ms.GetHealthCheckDuration()
	now := time.Now()

	for _, slave := range slaves {
		if slave == nil {
			continue
		}

		slave.mu.RLock()
		lastActive := slave.lastActive
		slave.mu.RUnlock()

		if now.Sub(lastActive) > healthCheck*2 {
			slave.mu.Lock()
			slave.isHealthy = false
			slave.mu.Unlock()

			// Restart unhealthy slave
			ms.restartSlave(slave)
		}
	}
}

// restartSlave recreates and restarts an unhealthy slave.
func (ms *MasterSlave) restartSlave(oldSlave *Slave) {
	oldSlave.cancel() // Cancel the old context
	<-oldSlave.done   // Wait for old slave to finish

	// Create new slave with same ID
	slaveCtx, slaveCancel := context.WithCancel(ms.ctx)
	newSlave := &Slave{
		id:         oldSlave.id,
		taskChan:   make(chan *taskWrapper, 10),
		resultChan: ms.resultsChan,
		done:       make(chan struct{}),
		ctx:        slaveCtx,
		cancel:     slaveCancel,
		isHealthy:  true,
		lastActive: time.Now(),
	}

	// Replace old slave with new one using the slave-specific mutex
	ms.slavesMu.Lock()
	ms.slaves[oldSlave.id-1] = newSlave
	ms.slavesMu.Unlock()

	// Start new slave
	go ms.runSlave(newSlave)
}

// runSlave starts a slave worker that processes tasks.
func (ms *MasterSlave) runSlave(slave *Slave) {
	defer close(slave.done)

	for {
		select {
		case <-slave.ctx.Done():
			return
		case wrapper := <-slave.taskChan:
			if wrapper != nil {
				start := time.Now()

				slave.mu.Lock()
				slave.lastActive = time.Now()
				slave.taskCount++
				slave.mu.Unlock()

				err := wrapper.task(slave.ctx)

				result := TaskResult{
					Error:   err,
					Time:    time.Since(start),
					Status:  StatusCompleted,
					SlaveID: slave.id,
					TaskID:  wrapper.id,
				}

				if err != nil {
					result.Status = StatusFailed
				}

				select {
				case slave.resultChan <- result:
				case <-slave.ctx.Done():
					return
				}

				ms.wg.Done() // Decrement wait group when task is done
			}
		}
	}
}

// Execute runs the given task using the master-slave architecture.
func (ms *MasterSlave) Execute(task Task) error {
	ms.mu.Lock()
	taskID := ms.taskCounter
	ms.taskCounter++
	ms.mu.Unlock()

	wrapper := &taskWrapper{
		id:       taskID,
		task:     task,
		status:   StatusPending,
		assigned: time.Now(),
	}

	ms.wg.Add(1)

	// Keep trying until we successfully schedule the task or timeout
	timeout := time.After(10 * time.Second)
	for {
		// Get a snapshot of slaves under lock
		ms.slavesMu.RLock()
		slaves := make([]*Slave, len(ms.slaves))
		copy(slaves, ms.slaves)
		ms.slavesMu.RUnlock()

		// Try each slave in round-robin fashion
		for _, slave := range slaves {
			if slave == nil {
				continue
			}

			slave.mu.RLock()
			isHealthy := slave.isHealthy
			slave.mu.RUnlock()

			if isHealthy {
				select {
				case slave.taskChan <- wrapper:
					return nil
				case <-timeout:
					ms.wg.Done()
					return fmt.Errorf("timeout scheduling task %d", taskID)
				case <-ms.ctx.Done():
					ms.wg.Done()
					return fmt.Errorf("master-slave system shutting down")
				default:
					continue
				}
			}
		}

		select {
		case <-timeout:
			ms.wg.Done()
			return fmt.Errorf("timeout scheduling task %d", taskID)
		case <-ms.ctx.Done():
			ms.wg.Done()
			return fmt.Errorf("master-slave system shutting down")
		case <-time.After(time.Millisecond):
			continue
		}
	}
}

// Wait blocks until all submitted tasks have been processed.
func (ms *MasterSlave) Wait() {
	ms.wg.Wait()
}

// GetResults returns a channel that receives task results.
func (ms *MasterSlave) GetResults() <-chan TaskResult {
	return ms.resultsChan
}

// Stop gracefully shuts down all slave workers.
func (ms *MasterSlave) Stop() {
	ms.stopOnce.Do(func() {
		ms.mu.Lock()
		if ms.isStopped {
			ms.mu.Unlock()
			return
		}
		ms.isStopped = true
		ms.mu.Unlock()

		// Cancel context to stop accepting new tasks
		ms.cancel()

		// Create a channel to signal completion of shutdown
		done := make(chan struct{})

		go func() {
			// Wait for tasks with a timeout
			waitCh := make(chan struct{})
			go func() {
				ms.wg.Wait()
				close(waitCh)
			}()

			select {
			case <-waitCh:
				// Tasks completed normally
			case <-time.After(5 * time.Second):
				// Timeout occurred, proceed with shutdown anyway
			}

			// Get a snapshot of slaves under lock
			ms.slavesMu.RLock()
			slaves := make([]*Slave, len(ms.slaves))
			copy(slaves, ms.slaves)
			ms.slavesMu.RUnlock()

			// Close all slave task channels
			for _, slave := range slaves {
				if slave != nil {
					close(slave.taskChan)
				}
			}

			// Wait for all slaves to finish with timeout
			for _, slave := range slaves {
				if slave != nil {
					select {
					case <-slave.done:
					case <-time.After(1 * time.Second):
						// Skip if slave doesn't finish in time
					}
				}
			}

			// Finally close the results channel
			close(ms.resultsChan)
			close(done)
		}()

		// Wait for shutdown to complete or timeout
		select {
		case <-done:
		case <-time.After(10 * time.Second):
			// If shutdown takes too long, we return anyway
		}
	})
}

// GetMetrics returns current metrics about the master-slave system.
func (ms *MasterSlave) GetMetrics() map[string]interface{} {
	ms.mu.RLock()
	ms.slavesMu.RLock()
	defer ms.mu.RUnlock()
	defer ms.slavesMu.RUnlock()

	metrics := make(map[string]interface{})
	metrics["total_slaves"] = ms.numSlaves

	healthySlaves := 0
	totalTasks := uint64(0)
	slaveMetrics := make([]map[string]interface{}, ms.numSlaves)

	for i, slave := range ms.slaves {
		if slave == nil {
			continue
		}

		slave.mu.RLock()
		if slave.isHealthy {
			healthySlaves++
		}
		totalTasks += slave.taskCount
		slaveMetrics[i] = map[string]interface{}{
			"id":          slave.id,
			"healthy":     slave.isHealthy,
			"task_count":  slave.taskCount,
			"last_active": slave.lastActive,
		}
		slave.mu.RUnlock()
	}

	metrics["healthy_slaves"] = healthySlaves
	metrics["total_tasks_processed"] = totalTasks
	metrics["slave_metrics"] = slaveMetrics

	return metrics
}

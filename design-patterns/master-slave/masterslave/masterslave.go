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
	mu          sync.RWMutex
	ctx         context.Context
	cancel      context.CancelFunc
	taskCounter uint64
	healthCheck time.Duration
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
		resultsChan: make(chan TaskResult, numSlaves*2),
		ctx:         ctx,
		cancel:      cancel,
		healthCheck: 5 * time.Second,
	}

	// Initialize and start slave workers
	for i := range ms.slaves {
		slaveCtx, slaveCancel := context.WithCancel(ctx)
		slave := &Slave{
			id:         i + 1,
			taskChan:   make(chan *taskWrapper, 1),
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

// monitorSlaveHealth periodically checks slave health.
func (ms *MasterSlave) monitorSlaveHealth() {
	ticker := time.NewTicker(ms.healthCheck)
	defer ticker.Stop()

	for {
		select {
		case <-ms.ctx.Done():
			return
		case <-ticker.C:
			ms.checkSlaveHealth()
		}
	}
}

// checkSlaveHealth verifies the health of all slaves.
func (ms *MasterSlave) checkSlaveHealth() {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	now := time.Now()
	for _, slave := range ms.slaves {
		slave.mu.RLock()
		inactive := now.Sub(slave.lastActive) > ms.healthCheck*2
		slave.mu.RUnlock()

		if inactive {
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
		taskChan:   make(chan *taskWrapper, 1),
		resultChan: ms.resultsChan,
		done:       make(chan struct{}),
		ctx:        slaveCtx,
		cancel:     slaveCancel,
		isHealthy:  true,
		lastActive: time.Now(),
	}

	// Replace old slave with new one
	ms.mu.Lock()
	ms.slaves[oldSlave.id-1] = newSlave
	ms.mu.Unlock()

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

	// Try to find a healthy slave
	for _, slave := range ms.slaves {
		slave.mu.RLock()
		isHealthy := slave.isHealthy
		slave.mu.RUnlock()

		if isHealthy {
			select {
			case slave.taskChan <- wrapper:
				return nil
			default:
				continue
			}
		}
	}

	// If no healthy slaves available, try any slave
	return fmt.Errorf("no healthy slaves available to process task %d", taskID)
}

// GetResults returns a channel that receives task results.
func (ms *MasterSlave) GetResults() <-chan TaskResult {
	return ms.resultsChan
}

// Stop gracefully shuts down all slave workers.
func (ms *MasterSlave) Stop() {
	ms.cancel() // Cancel the main context

	// Wait for all slaves to finish
	for _, slave := range ms.slaves {
		<-slave.done
	}

	// Close the results channel
	close(ms.resultsChan)
}

// GetMetrics returns current metrics about the master-slave system.
func (ms *MasterSlave) GetMetrics() map[string]interface{} {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	metrics := make(map[string]interface{})
	metrics["total_slaves"] = ms.numSlaves

	healthySlaves := 0
	totalTasks := uint64(0)
	slaveMetrics := make([]map[string]interface{}, ms.numSlaves)

	for i, slave := range ms.slaves {
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

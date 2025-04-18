/*
Package masterslave implements the Master-Slave architectural pattern, which distributes tasks between a master coordinator and multiple slave workers.

Core Components:
  - Master: The coordinator that distributes tasks and manages slave lifecycle
  - Slaves: Workers that execute tasks concurrently
  - Tasks: Units of work that can be processed independently

Key Features:

1. Task Management:
  - Concurrent task processing with multiple slave workers
  - Round-robin task distribution for load balancing
  - Task status tracking (Pending, Running, Completed, Failed)
  - Detailed task results including execution time and error information

2. Health Monitoring:
  - Automatic health checks for slave workers
  - Detection of inactive or stuck workers
  - Automatic worker restart when unhealthy
  - Configurable health check intervals

3. Metrics Collection:
  - Real-time system metrics
  - Per-slave statistics including:
  - Task count
  - Health status
  - Last active timestamp
  - System-wide metrics including:
  - Total slaves
  - Healthy slave count
  - Total tasks processed

4. Context Support:
  - Context-aware task execution
  - Graceful cancellation handling
  - Timeout support
  - Clean shutdown capabilities

5. Error Handling:
  - Comprehensive error reporting
  - Task failure isolation
  - Worker failure recovery
  - System-wide error propagation

Example Usage:

	// Create a master-slave system with 3 workers
	ms := masterslave.NewMasterSlave(3)
	defer ms.Stop()

	// Execute a task with context support
	err := ms.Execute(func(ctx context.Context) error {
	    select {
	    case <-ctx.Done():
	        return ctx.Err()
	    default:
	        // Do work
	        return nil
	    }
	})

	// Monitor task results
	for result := range ms.GetResults() {
	    if result.Error != nil {
	        log.Printf("Task %d failed: %v", result.TaskID, result.Error)
	    }
	}

	// Get system metrics
	metrics := ms.GetMetrics()
	fmt.Printf("Health: %d/%d slaves\n",
	    metrics["healthy_slaves"], metrics["total_slaves"])

Implementation Details:
  - Thread-safe operations using mutex locks
  - Non-blocking task submission
  - Buffered channels for task and result communication
  - Automatic worker lifecycle management
  - Graceful system shutdown

Best Used For:
  - Parallel processing of independent tasks
  - CPU-intensive operations
  - Batch processing workloads
  - Systems requiring high availability
  - Distributed computing patterns
  - Load-balanced task processing
*/
package masterslave

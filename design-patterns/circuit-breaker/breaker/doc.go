/*
Package breaker implements the Circuit Breaker design pattern, which is used to detect failures and encapsulates the logic of preventing a failure from constantly recurring during maintenance, temporary external system failure, or unexpected system difficulties.

The Circuit Breaker pattern works similar to an electrical circuit breaker:

  - CLOSED State (Normal Operation):
    The circuit breaker forwards all requests to the service and monitors for failures.
    When failures reach a certain threshold, it transitions to OPEN state.

  - OPEN State (Failure Mode):
    All requests immediately return with an error without attempting to execute the real
    operation. After a timeout period, it transitions to HALF-OPEN state.

  - HALF-OPEN State (Testing Recovery):
    Allows a limited number of requests to pass through. If these requests succeed,
    the breaker switches to CLOSED state. If any fail, it returns to OPEN state.

Example Usage:

	cb := breaker.NewCircuitBreaker(5, 10*time.Second) // Opens after 5 failures, resets after 10s
	err := cb.Execute(func() error {
	    return someRiskyOperation()
	})

Implementation Details:
  - Uses a mutex to ensure thread-safety
  - Configurable failure threshold and reset timeout
  - Automatic state transitions based on success/failure patterns
  - Simple API with Execute() method that wraps potentially failing operations

This implementation is particularly useful for:
  - Protecting systems from cascading failures
  - Allowing failing services time to recover
  - Failing fast when a system is under stress
  - Providing fallback mechanisms during partial system failures
*/
package breaker

// Package contexts provides examples of how to use the context package in Go for managing deadlines,
// cancellations, and other request-scoped values. This package is named context(s) to avoid
// colliding with the standard library's context package.
package contexts

import (
	"context"
	"fmt"
	"time"
)

// runTask is a function that will simulate running a task for a given duration.
func runTask(ctx context.Context, duration time.Duration) error {
	select {
	case <-time.After(duration):
		fmt.Println("Your long running task has finished")
		return nil
	case <-ctx.Done():
		fmt.Println("Your long running task has been canceled")
		return ctx.Err()
	}
}

// TimeoutContextBeforeTaskIsFinished is an example of how to use context to cancel a task before it finishes.
func TimeoutContextBeforeTaskIsFinished(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	// Pass duration that is longer than the context timeout. This will cause the task to be canceled before it finishes.
	if err := runTask(ctx, 2*duration); err != nil {
		fmt.Println(err)
	}
}

// FinishTaskBeforeContextTimeout is an example of how to use context with timeout that should complete before the timeout.
func FinishTaskBeforeContextTimeout(duration time.Duration) {
	// Set the timeout to be longer than the task. This will cause the task to compleat before the timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 2*duration)
	defer cancel()

	if err := runTask(ctx, duration); err != nil {
		fmt.Println(err)
	}
}

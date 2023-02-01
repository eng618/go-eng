// Package contexts is a package that contains examples of how to use context.
// This is named context(s) so that it doesn't collide with the standard library.
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

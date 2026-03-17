package main

import (
	"context"
	"fmt"
	"time"
)

// simulateAPICall pretends to call an API, respecting context cancellation
func simulateAPICall(ctx context.Context, taskName string, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return fmt.Sprintf("%s completed successfully", taskName), nil
	case <-ctx.Done():
		return "", fmt.Errorf("%s cancelled: %w", taskName, ctx.Err())
	}
}

// withTimeout demonstrates context.WithTimeout
func withTimeout() {
	fmt.Println("-- context.WithTimeout --")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Fast task — finishes before timeout
	result, err := simulateAPICall(ctx, "fast-task", 500*time.Millisecond)
	if err != nil {
		fmt.Println("fast-task error:", err)
	} else {
		fmt.Println(result)
	}

	// Slow task — exceeds timeout
	result, err = simulateAPICall(ctx, "slow-task", 5*time.Second)
	if err != nil {
		fmt.Println("slow-task error:", err)
	} else {
		fmt.Println(result)
	}
}

// withCancel demonstrates context.WithCancel
func withCancel() {
	fmt.Println("\n-- context.WithCancel --")

	ctx, cancel := context.WithCancel(context.Background())

	// Cancel after 1 second from a goroutine
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Cancelling context...")
		cancel()
	}()

	result, err := simulateAPICall(ctx, "cancellable-task", 3*time.Second)
	if err != nil {
		fmt.Println("cancellable-task error:", err)
	} else {
		fmt.Println(result)
	}
}

// withValue demonstrates passing values through context (e.g. request ID)
func withValue() {
	fmt.Println("\n-- context.WithValue --")

	type contextKey string
	const requestIDKey contextKey = "requestID"

	ctx := context.WithValue(context.Background(), requestIDKey, "req-abc-123")

	// Retrieve value deep in a call chain
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Println("Request ID from context:", id)
	}
}

func main() {
	fmt.Println("=== Week 3: context.Context ===\n")

	withTimeout()
	withCancel()
	withValue()
}

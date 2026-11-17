package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"
)

// workerctx executes tasks from jobs channel with context cancellation support.
func workerctx(ctx context.Context, workerId int, jobs chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		case jobId, ok := <-jobs:
			if !ok {
				return
			}
			time.Sleep(100 * time.Millisecond) // to spead jobs evenly among workers
			fmt.Printf("run job #%d on worker #%d\n", jobId, workerId)
		}
	}
}

// process runs numWorkers workers to solve numJobs tasks with context cancellation support.
func process(ctx context.Context, numJobs, numWorkers int) error {
	jobs := make(chan int, numJobs)
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	wg := sync.WaitGroup{}
	for workerId := range numWorkers {
		wg.Go(func() {
			workerctx(ctx, workerId, jobs)
		})
	}
	wg.Wait()
	if err := ctx.Err(); err != nil {
		return err
	}
	return nil
}

func TestContextCancel(t *testing.T) {
	t.Run("Complete run", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		if err := process(ctx, 7, 3); err != nil {
			t.Logf("encountered error: %s", err)
			t.Fatal()
		}
	})
	t.Run("Cancelled run", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		go func() { cancel() }()
		if err := process(ctx, 4, 2); err != nil {
			t.Logf("encountered error: %s", err)
			if !errors.Is(err, context.Canceled) {
				t.Fatal()
			}
		}
	})
	t.Run("Deadline run", func(t *testing.T) {
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
		defer cancel()
		if err := process(ctx, 10, 7); err != nil {
			t.Logf("encountered error: %s", err)
			if !errors.Is(err, context.DeadlineExceeded) {
				t.Fatal()
			}
		}
	})
	t.Run("Timed out run", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		if err := process(ctx, 13, 4); err != nil {
			t.Logf("ecnountered error: %s", err)
			if !errors.Is(err, context.DeadlineExceeded) {
				t.Fatal()
			}
		}
	})
}

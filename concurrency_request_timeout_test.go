package main

import (
	"fmt"
	"testing"
	"time"

	"math/rand"
)

// requestTimeout executes request
func requestTimeout(request func(), timeout float64) float64 {
	done := make(chan float64)
	go func() {
		start := time.Now()
		request()
		done <- time.Since(start).Seconds()
	}()
	ticker := time.Tick(500 * time.Millisecond)
	for {
		select {
		case elapsed := <-done:
			return elapsed
		case <-time.After(time.Duration(timeout) * time.Second):
			return -1
		case <-ticker:
			fmt.Println("request handling")
		}
	}
}

func requestTimeoutTestHelper(t *testing.T, request func(), timeout float64, success bool) {
	t.Helper()
	if success {
		t.Logf("expect request to take less than %f seconds", timeout)
	} else {
		t.Logf("expect request to last more than %f seconds", timeout)
	}
	if elapsed := requestTimeout(request, timeout); success && elapsed >= timeout ||
		!success && elapsed < timeout {
		t.Fatal()
	}
}

func TestRequestTimeout(t *testing.T) {
	t.Run("Long lasting request", func(t *testing.T) {
		requestTimeoutTestHelper(t, func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			time.Sleep(time.Duration(1+r.Intn(3)) * time.Second)
		}, 1.0, false)
	})
	t.Run("Fast request", func(t *testing.T) {
		requestTimeoutTestHelper(t, func() {}, 1.0, true)
	})
}

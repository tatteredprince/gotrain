package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type pinger interface {
	ping()
}

// periodicPing periodically pings via interface unless stopped with returned cancel().
func periodicPing(pinger pinger, period time.Duration) (cancel func()) {
	ticker := time.Tick(period)
	done := make(chan struct{})
	cancel = sync.OnceFunc(func() {
		fmt.Println("cancelling ping")
		done <- struct{}{} // not necessary
		close(done)
	})
	go func() {
		for {
			select {
			case <-ticker:
				pinger.ping()
			case <-done:
				return
			}
		}
	}()
	return
}

type testPing struct{}

func (ping testPing) ping() {
	fmt.Println("pinging...")
}

func TestPeriodicPing(t *testing.T) {
	t.Run("Test ping", func(t *testing.T) {
		ping := testPing{}
		cancel := periodicPing(ping, 1*time.Second)
		time.Sleep(3 * time.Second)
		cancel()
		cancel()
		cancel()
	})
}

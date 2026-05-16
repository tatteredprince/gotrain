package main

import (
	"slices"
	"sync"
	"testing"
)

// fanIn returns channel that multiplexes values from input channels
func fanIn(chans ...chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func() {
			defer wg.Done()
			for value := range ch {
				out <- value
			}
		}()
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fanInTestHelper(t *testing.T, values []int) {
	t.Helper()
	chans := make([]chan int, 0, len(values))
	for range values {
		ch := make(chan int)
		chans = append(chans, ch)
	}
	for i, val := range values {
		go func() {
			chans[i] <- val
			close(chans[i])
		}()
	}
	got := make([]int, 0, len(values))
	for num := range fanIn(chans...) {
		got = append(got, num)
	}
	slices.Sort(got)
	slices.Sort(values)
	if !slices.Equal(got, values) {
		t.Fatalf("expected %v but got %v", values, got)
	}
}

func TestFanIn(t *testing.T) {
	t.Run("Count to ten", func(t *testing.T) { fanInTestHelper(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) })
}

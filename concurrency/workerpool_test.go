package main

import (
	"fmt"
	"slices"
	"sync"
	"testing"
)

// worker processes job and sends its results.
func worker(jobId int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("start worker #%d\n", jobId)
		results <- j * j
		fmt.Printf("end worker #%d\n", jobId)
	}
}

func TestWorkerPool(t *testing.T) {
	t.Run("Three workers for five jobs", func(t *testing.T) {
		jobsCount, workersCount := 5, 3
		jobs := make(chan int, jobsCount)
		results := make(chan int, jobsCount)
		wg := sync.WaitGroup{}
		wg.Add(workersCount)
		for w := range workersCount {
			go worker(w, jobs, results, &wg)
		}
		for j := range jobsCount {
			jobs <- j
		}
		close(jobs)
		wg.Wait()
		close(results)
		expect, got := make([]int, 0, jobsCount), make([]int, 0, jobsCount)
		for e := range jobsCount {
			expect = append(expect, e*e)
		}
		for r := range results {
			got = append(got, r)
		}
		slices.Sort(expect)
		slices.Sort(got)
		t.Logf("expect results %v but got %v", expect, got)
		if !slices.Equal(expect, got) {
			t.Fatal()
		}
	})
}

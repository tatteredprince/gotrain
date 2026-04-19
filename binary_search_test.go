package main

import "testing"

// binarySearch returns index of target value in sorted sequence
func binarySearch(seq []int, target int) int {
	return -1
}

func binarySearchTestHelper(t *testing.T, seq []int, target, expect int) {
	t.Helper()
	t.Logf("search %d in %v", target, seq)
	if got := binarySearch(seq, target); got != expect {
		t.Fatalf("got %d but expected %d", got, expect)
	}
}

func TestBinarySearch(t *testing.T) {
	t.Run("One to five", func(t *testing.T) { binarySearchTestHelper(t, []int{1, 2, 3, 4, 5}, 3, 2) })
	t.Run("Odd numbers", func(t *testing.T) { binarySearchTestHelper(t, []int{1, 3, 5, 7, 9}, 7, 3) })
	t.Run("Fibonacci numbers", func(t *testing.T) { binarySearchTestHelper(t, []int{0, 1, 1, 2, 3, 5, 8, 13, 34}, 8, 6) })
	t.Run("Search even number among odds", func(t *testing.T) { binarySearchTestHelper(t, []int{2, 4, 6, 8, 10}, 3, -1) })
}

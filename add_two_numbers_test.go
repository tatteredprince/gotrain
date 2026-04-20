package main

import (
	"slices"
	"testing"
)

// addTowNumbers returns sum of two non-negative numbers, all three written backwards
func addTwoNumbers(a, b []int) []int {
	sum := make([]int, 0)
	return sum
}

func addTwoNumbersTestHelper(t *testing.T, a, b, expect []int) {
	t.Helper()
	t.Logf("add numbers %v and %v to get %v", a, b, expect)
	if got := addTwoNumbers(a, b); !slices.Equal(got, expect) {
		t.Fatalf("got %v but expected %v", got, expect)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	t.Run("Two zero numbers", func(t *testing.T) { addTwoNumbersTestHelper(t, []int{0}, []int{0}, []int{0}) })
	t.Run("Two equal numbers", func(t *testing.T) {
		addTwoNumbersTestHelper(t, []int{5, 5}, []int{5, 5}, []int{0, 1, 1})
	})
	t.Run("Three-digit and two-digit", func(t *testing.T) {
		addTwoNumbersTestHelper(t, []int{3, 2, 1}, []int{7, 8}, []int{0, 1, 2})
	})
	t.Run("Two three-digit numbers", func(t *testing.T) {
		addTwoNumbersTestHelper(t, []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8})
	})
}

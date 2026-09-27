package main

import (
	"slices"
	"testing"
)

// sortArrays sorts array
func sortArray(seq []int) []int {
	sorted := []int{}
	return sorted
}

func sortArrayTestHelper(t *testing.T, seq []int, expect []int) {
	t.Helper()
	t.Logf("sort %v", seq)
	if got := sortArray(seq); !slices.Equal(got, expect) {
		t.Fatal("got %v but expected %v", got, expect)
	}
}

func TestSortArray(t *testing.T) {
	t.Run("Single zero", func(t *testing.T) { sortArrayTestHelper(t, []int{0}, []int{0}) })
	t.Run("Zeroes", func(t *testing.T) { sortArrayTestHelper(t, []int{0, 0, 0}, []int{0, 0, 0}) })
	t.Run("One and zero", func(t *testing.T) { sortArrayTestHelper(t, []int{1, 0}, []int{0, 1}) })
	t.Run("James Bond", func(t *testing.T) { sortArrayTestHelper(t, []int{0, 7, 0}, []int{0, 0, 7}) })
	t.Run("Negative and positive", func(t *testing.T) { sortArrayTestHelper(t, []int{0, -1, -1}, []int{-1, 0, 1}) })
	t.Run("Mix of numbers", func(t *testing.T) {
		sortArrayTestHelper(t, []int{5, 1, -4, 9, 3, -11, 7, 4, 0, -9}, []int{-11, -9, -4, 0, 1, 3, 4, 5, 7, 9})
	})
}

package main

import "testing"

func testHelper(t *testing.T, seq []int, expect int) {
	t.Helper()
	got := longestZeroes(seq)
	if got != expect {
		t.Fatalf("got %v but expected %v", got, expect)
	}
}

func TestLongestZeroes(t *testing.T) {
	t.Run("Empty sequence", func(t *testing.T) { testHelper(t, []int{}, 0) })
	t.Run("Sequence of single zero", func(t *testing.T) { testHelper(t, []int{0}, 1) })
	t.Run("Sequence of single one", func(t *testing.T) { testHelper(t, []int{1}, 0) })
	t.Run("Sequence of zero and one", func(t *testing.T) { testHelper(t, []int{0, 1}, 1) })
	t.Run("Sequence of zeroes", func(t *testing.T) { testHelper(t, []int{0, 0, 0, 0, 0}, 5) })
	t.Run("Sequence of ones", func(t *testing.T) { testHelper(t, []int{1, 1, 1, 1, 1}, 0) })
	t.Run("Harmonized ones and zeroes", func(t *testing.T) { testHelper(t, []int{1, 0, 1, 0, 1, 0}, 1) })
	t.Run("Non-harmonized ones and zeroes", func(t *testing.T) { testHelper(t, []int{1, 0, 0, 1, 0, 0, 0, 1}, 3) })
}

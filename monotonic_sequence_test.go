package main

import "testing"

// monotonicSequence indicates that sequence passed monotonically increases or descreases
func monotonicSequence(seq []int) bool {
	return false
}

func monotonicSequenceTestHelper(t *testing.T, seq []int, expect bool) {
	t.Helper()
	if expect {
		t.Logf("expecting monotonic %v", seq)
	} else {
		t.Logf("expecting non-monotonic %v", seq)
	}
	if got := monotonicSequence(seq); got != expect {
		t.Fail()
	}
}

func TestMonotonicSequence(t *testing.T) {
	t.Run("Empty sequence", func(t *testing.T) { monotonicSequenceTestHelper(t, []int{}, false) })
	t.Run("Sequence of single digit", func(t *testing.T) { monotonicSequenceTestHelper(t, []int{0}, false) })
	t.Run("Sequence of equal numbers", func(t *testing.T) {
		monotonicSequenceTestHelper(t, []int{1, 1, 1, 1, 1}, true)
	})
	t.Run("Monotonic sequence of ones and zeroes", func(t *testing.T) {
		monotonicSequenceTestHelper(t, []int{1, 1, 1, 1, 0}, true)
	})
	t.Run("Motonically increasing pair", func(t *testing.T) {
		monotonicSequenceTestHelper(t, []int{0, 1}, true)
	})
	t.Run("Motonically decreasing pair", func(t *testing.T) {
		monotonicSequenceTestHelper(t, []int{1, 0}, true)
	})
	t.Run("Monotonic sequence", func(t *testing.T) {
		monotonicSequenceTestHelper(t, []int{1, 2, 3, 4, 5}, true)
	})
	t.Run("Non-monotonic sequence", func(t *testing.T) {
		monotonicSequenceTestHelper(t, []int{6, 7, 8, 1, 2, 3}, false)
	})
}

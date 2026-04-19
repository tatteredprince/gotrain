package main

import "testing"

// monotonicSequence indicates that sequence passed monotonically increases or descreases
func monotonicSequence(seq []int) bool {
	return false
}

func expectFailure(t *testing.T, seq []int) {
	t.Helper()
	t.Logf("expecting non-monotonic %v", seq)
	if monotonicSequence(seq) {
		t.Fatalf("%v is monotonic", seq)
	}
}

func expectSuccess(t *testing.T, seq []int) {
	t.Helper()
	t.Logf("expecting monotonic %v", seq)
	if !monotonicSequence(seq) {
		t.Fatalf("%v is non-monotonic", seq)
	}
}

func TestMonotonicSequence(t *testing.T) {
	t.Run("Empty sequence", func(t *testing.T) { expectFailure(t, []int{}) })
	t.Run("Sequence of single digit", func(t *testing.T) { expectFailure(t, []int{0}) })
	t.Run("Sequence of equal numbers", func(t *testing.T) { expectSuccess(t, []int{1, 1, 1, 1, 1}) })
	t.Run("Monotonic sequence of ones and zeroes", func(t *testing.T) { expectSuccess(t, []int{1, 1, 1, 1, 0}) })
	t.Run("Motonically increasing pair", func(t *testing.T) { expectSuccess(t, []int{0, 1}) })
	t.Run("Motonically decreasing pair", func(t *testing.T) { expectSuccess(t, []int{1, 0}) })
	t.Run("Monotonic sequence", func(t *testing.T) { expectSuccess(t, []int{1, 2, 3, 4, 5}) })
	t.Run("Non-monotonic sequence", func(t *testing.T) { expectFailure(t, []int{6, 7, 8, 1, 2, 3}) })
}

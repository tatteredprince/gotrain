package main

import "testing"

// longestZeroesSequence calculates longest sequence of zeroes in slice
func longestZeroes(seq []int) int {
	if len(seq) == 1 && seq[0] == 0 {
		return 1
	}
	seqLen := 0
	for i := 1; i < len(seq); i++ {
		if seq[i-1] == 0 {
			_seqLen := 1
			j := i
			for ; j < len(seq); j++ {
				if seq[j] == 0 {
					_seqLen++
				} else {
					break
				}
			}
			i = j
			if _seqLen > seqLen {
				seqLen = _seqLen
			}
		}
	}
	return seqLen
}

func longestZeroesTestHelper(t *testing.T, seq []int, expect int) {
	t.Helper()
	t.Logf("calculate longest zeroes sequence in %v", seq)
	if got := longestZeroes(seq); got != expect {
		t.Fatalf("got %v but expected %v", got, expect)
	}
}

func TestLongestZeroes(t *testing.T) {
	t.Run("Empty sequence", func(t *testing.T) { longestZeroesTestHelper(t, []int{}, 0) })
	t.Run("Sequence of single zero", func(t *testing.T) { longestZeroesTestHelper(t, []int{0}, 1) })
	t.Run("Sequence of single one", func(t *testing.T) { longestZeroesTestHelper(t, []int{1}, 0) })
	t.Run("Sequence of zero and one", func(t *testing.T) { longestZeroesTestHelper(t, []int{0, 1}, 1) })
	t.Run("Sequence of zeroes", func(t *testing.T) { longestZeroesTestHelper(t, []int{0, 0, 0, 0, 0}, 5) })
	t.Run("Sequence of ones", func(t *testing.T) { longestZeroesTestHelper(t, []int{1, 1, 1, 1, 1}, 0) })
	t.Run("Harmonized ones and zeroes", func(t *testing.T) {
		longestZeroesTestHelper(t, []int{1, 0, 1, 0, 1, 0}, 1)
	})
	t.Run("Non-harmonized ones and zeroes", func(t *testing.T) {
		longestZeroesTestHelper(t, []int{1, 0, 0, 1, 0, 0, 0, 1}, 3)
	})
}

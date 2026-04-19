package main

import "testing"

// hasTwoComplements indicates that slice contains two values which in sum gives target value
func hasTwoComplements(seq []int, target int) bool {
	return false
}

func twoComplementsTestHelper(t *testing.T, seq []int, target int, expect bool) {
	t.Helper()
	if expect {
		t.Logf("expecting %v to have two complements for %d", seq, target)
	} else {
		t.Logf("expecting %v not to have two complements for %d", seq, target)
	}
	if got := hasTwoComplements(seq, target); got != expect {
		t.Fatal()
	}
}

func expectFailure(t *testing.T, target int, seq []int) {
	t.Helper()
	t.Logf("expecting %d to not have complements for %d", seq, target)
	if hasTwoComplements(target, seq) {
		t.Fatalf("%d has complement for %d", target, seq)
	}
}

func expectSuccess(t *testing.T, target int, seq []int) {
	t.Helper()
	t.Logf("expecting %d to have complements for %d", seq, target)
	if !hasTwoComplements(target, seq) {
		t.Fatalf("%d has no complement for %d", target, seq)
	}
}

func TestTwoComplements(t *testing.T) {
	t.Run("Binary set", func(t *testing.T) { twoComplementsTestHelper(t, []int{0, 1}, 1, true) })
	t.Run("Continious binary set", func(t *testing.T) { twoComplementsTestHelper(t, []int{1, 0, 1, 0, 1, 0, 1}, 9, false) })
	t.Run("Empty set", func(t *testing.T) { twoComplementsTestHelper(t, []int{}, 5, false) })
	t.Run("Zeroes set", func(t *testing.T) { twoComplementsTestHelper(t, []int{0, 0, 0, 0, 0}, 0, true) })
	t.Run("No complement", func(t *testing.T) { twoComplementsTestHelper(t, []int{3, 4, 5, 6, 7}, 3, false) })
}

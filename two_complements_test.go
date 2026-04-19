package main

import "testing"

// hasTwoComplements indicates that slice contains two values which in sum gives target value
func hasTwoComplements(target int, seq []int) bool {
	return false
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
	t.Run("Binary set", func(t *testing.T) { expectSuccess(t, 1, []int{0, 1}) })
	t.Run("Continious binary set", func(t *testing.T) { expectFailure(t, 9, []int{1, 0, 1, 0, 1, 0, 1}) })
	t.Run("Empty set", func(t *testing.T) { expectFailure(t, 5, []int{}) })
	t.Run("Zeroes set", func(t *testing.T) { expectSuccess(t, 0, []int{0, 0, 0, 0, 0}) })
	t.Run("No complement", func(t *testing.T) { expectFailure(t, 3, []int{3, 4, 5, 6, 7}) })
}

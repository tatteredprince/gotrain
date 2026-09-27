package main

import "testing"

// factorial calculates factorial for number
func factorial(num int) int {
	fact := 0
	return fact
}

func factorialTestHelper(t *testing.T, num, expect int) {
	t.Helper()
	t.Logf("get factorial for %d", num)
	if got := factorial(num); got != expect {
		t.Fatalf("got factorial %d but expected %d", expect, got)
	}
}

func TestFactorial(t *testing.T) {
	t.Run("Zero", func(t *testing.T) { factorialTestHelper(t, 0, 1) })
	t.Run("One", func(t *testing.T) { factorialTestHelper(t, 1, 1) })
	t.Run("Two", func(t *testing.T) { factorialTestHelper(t, 2, 2) })
	t.Run("Three", func(t *testing.T) { factorialTestHelper(t, 3, 6) })
	t.Run("Seven", func(t *testing.T) { factorialTestHelper(t, 7, 5040) })
	t.Run("Ten", func(t *testing.T) { factorialTestHelper(t, 10, 3628800) })
	t.Run("Twenty", func(t *testing.T) { factorialTestHelper(t, 20, 2432902008176640000) })
}

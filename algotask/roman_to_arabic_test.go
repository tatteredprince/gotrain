package main

import "testing"

// romanNumbers calculates roman number for arabic
func romanNumbers(arabic int) string {
	roman := ""
	return roman
}

func romanNumbersTestHelper(t *testing.T, arabic int, expect string) {
	t.Helper()
	t.Logf("calculate roman number for arabic %d", arabic)
	if got := romanNumbers(arabic); got != expect {
		t.Fatalf("got %s but expected %s")
	}
}

func TestRomanNumbers(t *testing.T) {
	t.Run("Arabic 1", func(t *testing.T) { romanNumbersTestHelper(t, 1, "I") })
	t.Run("Arabic 3", func(t *testing.T) { romanNumbersTestHelper(t, 3, "III") })
	t.Run("Arabic 4", func(t *testing.T) { romanNumbersTestHelper(t, 4, "IV") })
	t.Run("Arabic 5", func(t *testing.T) { romanNumbersTestHelper(t, 5, "V") })
	t.Run("Arabic 6", func(t *testing.T) { romanNumbersTestHelper(t, 6, "VI") })
	t.Run("Arabic 8", func(t *testing.T) { romanNumbersTestHelper(t, 8, "VIII") })
	t.Run("Arabic 9", func(t *testing.T) { romanNumbersTestHelper(t, 9, "IX") })
	t.Run("Arabic 10", func(t *testing.T) { romanNumbersTestHelper(t, 10, "X") })
	t.Run("Arabic 13", func(t *testing.T) { romanNumbersTestHelper(t, 13, "XIII") })
	t.Run("Arabic 15", func(t *testing.T) { romanNumbersTestHelper(t, 15, "XV") })
	t.Run("Arabic 50", func(t *testing.T) { romanNumbersTestHelper(t, 50, "L") })
	t.Run("Arabic 100", func(t *testing.T) { romanNumbersTestHelper(t, 100, "C") })
	t.Run("Arabic 500", func(t *testing.T) { romanNumbersTestHelper(t, 500, "D") })
	t.Run("Arabic 1000", func(t *testing.T) { romanNumbersTestHelper(t, 1000, "M") })
	t.Run("Arabic 39", func(t *testing.T) { romanNumbersTestHelper(t, 39, "XXXIX") })
	t.Run("Arabic 246", func(t *testing.T) { romanNumbersTestHelper(t, 246, "CCXLVI") })
	t.Run("Arabic 789", func(t *testing.T) { romanNumbersTestHelper(t, 789, "DCCLXXXIX") })
	t.Run("Arabic 2421", func(t *testing.T) { romanNumbersTestHelper(t, 2421, "MMCDXXI") })
	t.Run("Arabic 160", func(t *testing.T) { romanNumbersTestHelper(t, 160, "CLX") })
	t.Run("Arabic 207", func(t *testing.T) { romanNumbersTestHelper(t, 207, "CCVII") })
	t.Run("Year 1009", func(t *testing.T) { romanNumbersTestHelper(t, 1009, "MIX") })
	t.Run("Year 1066", func(t *testing.T) { romanNumbersTestHelper(t, 1066, "MLXVI") })
	t.Run("Year 1776", func(t *testing.T) { romanNumbersTestHelper(t, 1776, "MDCCLXXVI") })
	t.Run("Year 1918", func(t *testing.T) { romanNumbersTestHelper(t, 1918, "MCMXVIII") })
	t.Run("Year 1944", func(t *testing.T) { romanNumbersTestHelper(t, 1944, "MCMXLIV") })
	t.Run("Year 2025", func(t *testing.T) { romanNumbersTestHelper(t, 2025, "MMXXV") })
}

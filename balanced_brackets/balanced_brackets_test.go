package main

import "testing"

func testHelper(t *testing.T, str string, expect bool) {
	t.Helper()
	if expect {
		t.Logf("expecting balanced brackets for '%s'", str)
	} else {
		t.Logf("expecting non-balanced brackets for '%s'", str)
	}
	got := balancedBrackets(str)
	if got != expect {
		t.Fail()
	}
}

func TestBalancedBrackets(t *testing.T) {
	t.Run("Function declaration", func(t *testing.T) {
		testHelper(t, "func abs(num int) int {if num < 0 {return -num} else {return num}}", true)
	})
	t.Run("Array element conversion", func(t *testing.T) { testHelper(t, "if arr[i] != 0 {elem = string(arr[i])}", true) })
	t.Run("Clumsy array ", func(t *testing.T) { testHelper(t, "for {if arr[i {break}}", false) })
	t.Run("Badly decorated printing", func(t *testing.T) { testHelper(t, "{fmt.Print1, 2, 3)", false) })
	t.Run("Erroneous array definition", func(t *testing.T) { testHelper(t, "abc := []int{1, 2, 3", false) })
	t.Run("Unfinised iteration", func(t *testing.T) { testHelper(t, "for i := range arr {if arr[i] % 2 != 0 {return false}", false) })
	t.Run("Function cut", func(t *testing.T) { testHelper(t, "} return -1 }", false) })
}

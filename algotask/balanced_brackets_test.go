package main

import (
	"bytes"
	"testing"
)

// balancedBrackets indicates that brackets in string are balanced
func balancedBrackets(str string) bool {
	stack := make([]byte, 0, len(str))
	itop := -1
	for i := 0; i < len(str); i++ {
		if bytes.IndexByte([]byte{'(', ')', '{', '}', '[', ']'}, str[i]) != -1 {
			if str[i] == '(' || str[i] == '{' || str[i] == '[' {
				stack = append(stack, str[i])
				itop++
			} else if itop != -1 &&
				(str[i] == ')' && stack[itop] == '(' ||
					str[i] == '}' && stack[itop] == '{' ||
					str[i] == ']' && stack[itop] == '[') {
				stack = stack[:itop] // pop
				itop--
			} else {
				return false
			}
		}
	}
	if itop != -1 {
		return false
	}
	return true
}

func balancedBracketsTestHelper(t *testing.T, str string, expect bool) {
	t.Helper()
	if expect {
		t.Logf("expecting balanced brackets for '%s'", str)
	} else {
		t.Logf("expecting non-balanced brackets for '%s'", str)
	}
	if got := balancedBrackets(str); got != expect {
		t.Fail()
	}
}

func TestBalancedBrackets(t *testing.T) {
	t.Run("Function declaration", func(t *testing.T) {
		balancedBracketsTestHelper(t, "func abs(num int) int {if num < 0 {return -num} else {return num}}", true)
	})
	t.Run("Array element conversion", func(t *testing.T) {
		balancedBracketsTestHelper(t, "if arr[i] != 0 {elem = string(arr[i])}", true)
	})
	t.Run("Clumsy array ", func(t *testing.T) {
		balancedBracketsTestHelper(t, "for {if arr[i {break}}", false)
	})
	t.Run("Badly decorated printing", func(t *testing.T) {
		balancedBracketsTestHelper(t, "{fmt.Print1, 2, 3)", false)
	})
	t.Run("Erroneous array definition", func(t *testing.T) {
		balancedBracketsTestHelper(t, "abc := []int{1, 2, 3", false)
	})
	t.Run("Unfinised iteration", func(t *testing.T) {
		balancedBracketsTestHelper(t, "for i := range arr {if arr[i] % 2 != 0 {return false}", false)
	})
	t.Run("Function cut", func(t *testing.T) {
		balancedBracketsTestHelper(t, "} return -1 }", false)
	})
}

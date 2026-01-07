package main

import "bytes"

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

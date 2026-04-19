package main

import "testing"

// reverseString reverses string
func reverseString(str string) string {
	return ""
}

func reverseStringTestHelper(t *testing.T, str string, expect string) {
	t.Helper()
	t.Logf("reverse '%s'", str)
	if got := reverseString(str); got != expect {
		t.Fatalf("got '%s' but expected '%s'", got, expect)
	}
}

func TestReverseString(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) { reverseStringTestHelper(t, "", "") })
	t.Run("Simple word", func(t *testing.T) { reverseStringTestHelper(t, "ZYX", "XYZ") })
	t.Run("Palindrome", func(t *testing.T) { reverseStringTestHelper(t, "MOM", "MOM") })
	t.Run("Unicode string", func(t *testing.T) { reverseStringTestHelper(t, "Hello, 世界", "界世 ,olleH") })
}

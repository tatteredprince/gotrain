package main

import "testing"

func testHelper(t *testing.T, str string, expect string) {
	t.Helper()
	got := reverseString(str)
	if got != expect {
		t.Fatalf("got '%s' but expected '%s'", got, expect)
	}
}

func TestReverseString(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) { testHelper(t, "", "") })
	t.Run("Simple word", func(t *testing.T) { testHelper(t, "ZYX", "XYZ") })
	t.Run("Palindrome", func(t *testing.T) { testHelper(t, "MOM", "MOM") })
	t.Run("Unicode string", func(t *testing.T) { testHelper(t, "Hello, 世界", "界世 ,olleH") })
}

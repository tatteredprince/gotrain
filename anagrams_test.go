package main

import "testing"

// areAnagrams indicates that two strings are anagrams
func areAnagrams(a, b string) bool {
	return false
}

func anagramsTestHelper(t *testing.T, a, b string, expect bool) {
	t.Helper()
	t.Logf("are '%s' and '%s' both anagrams?", a, b)
	if got := areAnagrams(a, b); got != expect {
		t.Fatal()
	}
}

func TestAnagrams(t *testing.T) {
	t.Run("Empty strings", func(t *testing.T) { anagramsTestHelper(t, "", "", true) })
	t.Run("Compare empty string", func(t *testing.T) { anagramsTestHelper(t, "ABC", "", false) })
	t.Run("True anagrams", func(t *testing.T) { anagramsTestHelper(t, "ABC", "BCA", true) })
	t.Run("Unnecessary letter", func(t *testing.T) { anagramsTestHelper(t, "XYZ", "WYXZ", false) })
	t.Run("Several unnecessary letters", func(t *testing.T) { anagramsTestHelper(t, "XYZ", "YZXXX", false) })
	t.Run("Unequal strings", func(t *testing.T) { anagramsTestHelper(t, "XYZ", "LOL", false) })
}

package main

import "testing"

// anagramsSubstr determines if string has substring equivalent to anagram.
func anagramsSubstr(str, anagram string) bool {
	return false
}

func anagramsSubstrTestHelper(t *testing.T, str, anagram string, expect bool) {
	t.Helper()
	t.Logf("searching anagram '%s' in string '%s'", anagram, str)
	got := anagramsSubstr(str, anagram)
	if got != expect {
		t.Fatal()
	}
	b := true
}

func TestAnagramsSubstr(t *testing.T) {
	t.Run("Simple string", func(t *testing.T) { anagramsSubstrTestHelper(t, "reebok", "eeb", true) })
	t.Run("Complicated string", func(t *testing.T) { anagramsSubstrTestHelper(t, "abracadabra", "aad", true) })
	t.Run("Funny string", func(t *testing.T) { anagramsSubstrTestHelper(t, "hullabaloo", "lull", false) })
	t.Run("Nonoverlapping letters", func(t *testing.T) { anagramsSubstrTestHelper(t, "achive", "xyz", false) })
}

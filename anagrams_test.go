package main

import "testing"

// areAnagrams indicates that two strings are anagrams
func areAnagrams(a, b string) bool {
	amap := make(map[byte]int, 0)
	for i := 0; i < len(a); i++ {
		if _, ok := amap[a[i]]; ok {
			amap[a[i]]++
		} else {
			amap[a[i]] = 1
		}
	}
	bmap := make(map[byte]int, 0)
	for i := 0; i < len(b); i++ {
		if _, ok := bmap[b[i]]; ok {
			bmap[b[i]]++
		} else {
			bmap[b[i]] = 1
		}
	}
	if len(amap) != len(bmap) {
		return false
	}
	for ch, cnt := range amap {
		if bmap[ch] != cnt {
			return false
		}
	}
	return true
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
	t.Run("Music band", func(t *testing.T) { anagramsTestHelper(t, "ABBA", "BABA", true) })
	t.Run("Several unnecessary letters", func(t *testing.T) { anagramsTestHelper(t, "XYZ", "YZXXX", false) })
	t.Run("Unequal strings", func(t *testing.T) { anagramsTestHelper(t, "XYZ", "LOL", false) })
}

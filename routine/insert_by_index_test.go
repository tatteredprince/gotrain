package main

import (
	"slices"
	"testing"
)

// insertByIndex inserts value at index in slice.
func insertByIndex(s []int, i, v int) []int {
	if length := len(s); i < length {
		spare := make([]int, length-i)
		copy(spare, s[i:])
		s = append(s[:i], v)
		s = append(s, spare...)
	} else {
		s = append(s, v)
	}
	return s
}

func insertByIndexTestHelper(t *testing.T, s []int, i, v int, expect []int) {
	got := insertByIndex(s, i, v)
	t.Logf("got %v but expect %v", got, expect)
	if !slices.Equal(got, expect) {
		t.Fatal()
	}
}

func TestInsertByIndex(t *testing.T) {
	t.Run("Insert at the beginning", func(t *testing.T) {
		insertByIndexTestHelper(t, []int{2, 3, 4, 5}, 0, 1, []int{1, 2, 3, 4, 5})
	})
	t.Run("Insert in the middle", func(t *testing.T) {
		insertByIndexTestHelper(t, []int{1, 2, 4, 5}, 2, 3, []int{1, 2, 3, 4, 5})
	})
	t.Run("Insert at the end", func(t *testing.T) {
		insertByIndexTestHelper(t, []int{1, 2, 3, 4}, 4, 5, []int{1, 2, 3, 4, 5})
	})
}

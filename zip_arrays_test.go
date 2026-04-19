package main

import (
	"reflect"
	"testing"
)

// zipArrays returns array of arrays made of harnessed elements of passed slices
func zipArrays(arrays ...[]int) [][]int {
	zipped := make([][]int, 0)
	return zipped
}

func zipArraysTestHelper(t *testing.T, zipped [][]int, arrays ...[]int) {
	t.Helper()
	got := zipArrays(arrays...)
	if !reflect.DeepEqual(got, zipped) {
		t.Fatalf("got %v but expected %v", got, zipped)
	}
}

func TestZipArrays(t *testing.T) {
	t.Run("Empty slices", func(t *testing.T) { zipArraysTestHelper(t, [][]int{}) })
	t.Run("Empty single slice", func(t *testing.T) { zipArraysTestHelper(t, [][]int{}, []int{}) })
	t.Run("Single slice", func(t *testing.T) { zipArraysTestHelper(t, [][]int{{1}, {2}, {3}}, []int{1, 2, 3}) })
	t.Run("Two empty slices", func(t *testing.T) { zipArraysTestHelper(t, [][]int{}, []int{}, []int{}) })
	t.Run("One non-empty slice", func(t *testing.T) { zipArraysTestHelper(t, [][]int{}, []int{}, []int{0}) })
	t.Run("Two non-empty slices", func(t *testing.T) {
		zipArraysTestHelper(t, [][]int{{1, 2}, {11, 22}, {111, 222}}, []int{1, 11, 111}, []int{2, 22, 222, 2222, 22222})
	})
	t.Run("Odd and even numbers", func(t *testing.T) {
		zipArraysTestHelper(t, [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 3, 5}, []int{2, 4, 6})
	})
}

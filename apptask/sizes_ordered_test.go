package main

import (
	"maps"
	"reflect"
	"slices"
	"testing"
)

// sizesOrdered returns slice of sizes sorted by their frequencies in descending order.
func sizesOrdered(sizes []string) []map[string]int {
	sizesFrequencies := make(map[string]int, len(sizes))
	for _, sz := range sizes {
		sizesFrequencies[sz]++
	}
	sizesDescending := make([]map[string]int, 0, len(sizesFrequencies))
	for size, quantity := range sizesFrequencies {
		sizesDescending = append(sizesDescending, map[string]int{size: quantity})
	}
	slices.SortFunc(sizesDescending, func(a, b map[string]int) int {
		ak, bk := slices.Collect(maps.Keys(a)), slices.Collect(maps.Keys(b))
		if a[ak[0]] < b[bk[0]] {
			return 1
		} else if a[ak[0]] > b[bk[0]] {
			return -1
		}
		return 0
	})
	return sizesDescending
}

func TestSizesOrdered(t *testing.T) {
	t.Run("Shirts sizes", func(t *testing.T) {
		got := sizesOrdered([]string{"S", "M", "S", "XL", "S", "XL"})
		expect := []map[string]int{{"S": 3}, {"XL": 2}, {"M": 1}}
		t.Logf("got %v but expect %v", got, expect)
		if !reflect.DeepEqual(got, expect) {
			t.Fatal()
		}
	})
}

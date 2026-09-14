package main

import (
	"maps"
	"slices"
	"testing"
)

// uniqueGoods finds unique items sold by sellers.
func uniqueGoods(sellers map[int][]string) map[int][]string {
	goods := make(map[string][]int, 0) // items sold by sellers i.e. map[string][]int{"gloves" => {1, 2, 3}}
	for seller, inventories := range sellers {
		for _, item := range inventories {
			goods[item] = append(goods[item], seller)
		}
	}
	uniqueGoods := make(map[int][]string, len(sellers))
	for item, sellers := range goods {
		if len(sellers) == 1 {
			uniqueGoods[sellers[0]] = append(uniqueGoods[sellers[0]], item)
		}
	}
	return uniqueGoods
}

func uniqueGoodsTestHelper(t *testing.T, sellers, expect map[int][]string) {
	t.Helper()
	got := uniqueGoods(sellers)
	t.Logf("got %v but expected %v", got, expect)
	gotKeys := slices.Collect(maps.Keys(got))
	expectKeys := slices.Collect(maps.Keys(expect))
	slices.Sort(gotKeys)
	slices.Sort(expectKeys)
	if !slices.Equal(gotKeys, expectKeys) {
		t.Fatal("keys aren't equal")
	}
	for seller, gotGoods := range got {
		expectGoods := expect[seller]
		slices.Sort(gotGoods)
		slices.Sort(expectGoods)
		if !slices.Equal(gotGoods, expectGoods) {
			t.Fatal("goods aren't equal")
		}
	}
}

func TestUniqueGoods(t *testing.T) {
	t.Run("Many selleres", func(t *testing.T) {
		uniqueGoodsTestHelper(
			t,
			map[int][]string{
				1: {"gloves", "coat", "boots", "table", "hat", "scarf", "blouse", "shirt"},
				2: {"scarf", "sofa", "hat", "table", "spoon", "coat"},
				3: {"fork", "coat", "dish", "scarf", "coat", "table", "pot"},
				4: {"carpet", "sofa", "coat", "vacuum", "boots", "spoon", "table"},
			},
			map[int][]string{
				1: {"gloves", "shirt", "blouse"},
				3: {"fork", "dish", "pot"},
				4: {"carpet", "vacuum"},
			},
		)
	})
	t.Run("Single seller", func(t *testing.T) {
		uniqueGoodsTestHelper(
			t,
			map[int][]string{1: {"wardrobe", "sofa", "hat", "table", "spoon", "jacket"}},
			map[int][]string{1: {"wardrobe", "sofa", "hat", "table", "spoon", "jacket"}},
		)
	})
	t.Run("Duplicate goods", func(t *testing.T) {
		uniqueGoodsTestHelper(
			t,
			map[int][]string{
				1: {"wardrobe", "sofa", "hat", "table", "spoon", "jacket"},
				2: {"wardrobe", "sofa", "hat", "table", "spoon", "jacket"},
			},
			map[int][]string{},
		)
	})
}

package main

import (
	"testing"
)

type category struct {
	name     string
	children []category
}

// catalogSearch
func catalogSearch(cat category, item string) string {
	if cat.name == item {
		return " > " + item
	} else {
		for _, child := range cat.children {
			if child.name == item {
				return child.name + " > " + item
			} else {
				return child.name + " > " + catalogSearch(child, item)
			}
		}
	}
	return ""
}

func TestCatalogSearch(t *testing.T) {
	t.Run("Find OLED", func(t *testing.T) {
		root := category{
			"Root",
			[]category{
				{
					"Home appliances",
					[]category{
						{
							"TVs",
							[]category{
								{"CRT", []category{}}, {"IPS", []category{}}, {"IPS", []category{}},
							},
						},
						{
							"Refrigerators",
							[]category{{"Two-compartment", []category{}}, {"Mini", []category{}}},
						},
						{"Irons", []category{}},
					},
				},
				{
					"Plants",
					[]category{
						{"Indoor", []category{}},
						{"Gardern", []category{}},
					},
				},
			},
		}
		got, expect := catalogSearch(root, "OLED"), "Home appliances > TVs > CRT > OLED"
		t.Logf("got '%s' but expected '%s'", got, expect)
		if got != expect {
			t.Fatal()
		}
	})
}

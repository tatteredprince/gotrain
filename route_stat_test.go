package main

import (
	"reflect"
	"testing"
)

type route struct {
	From string
	To   string
}

type routeStat struct {
	MaxRoutes      []route
	MaxRoutesCount int
	Uniq           int
	Once           int
}

// analyzeRoutes calculates routes' statistics from argument: most frequent
// routes and their count, number of various routes and number of unique routes.
// Most frequent routes should be selected in order of appearance.
func analyzeRoutes(routes []route) routeStat {
	routesDigest := make(map[route]int)
	orderedRoutes := make([]route, 0, len(routes))
	for _, route := range routes {
		if _, ok := routesDigest[route]; ok {
			routesDigest[route]++
		} else {
			routesDigest[route] = 1
			orderedRoutes = append(orderedRoutes, route)
		}
	}
	var (
		maxRoutes      []route
		maxRoutesCount = 0
		once           = 0
	)
	for _, orderedRoute := range orderedRoutes {
		count := routesDigest[orderedRoute]
		if count == 1 {
			once += 1
		}
		if count > maxRoutesCount {
			maxRoutes = make([]route, 0, len(routesDigest))
			maxRoutes = append(maxRoutes, orderedRoute)
			maxRoutesCount = count
		} else if count == maxRoutesCount {
			maxRoutes = append(maxRoutes, orderedRoute)
		}
	}
	return routeStat{
		MaxRoutes:      maxRoutes,
		MaxRoutesCount: len(maxRoutes),
		Uniq:           len(routesDigest),
		Once:           once,
	}
}

func RouteStatTestHelper(t *testing.T, routes []route, expect routeStat) {
	t.Helper()
	got := analyzeRoutes(routes)
	t.Logf("got %v but expect %v", got, expect)
	if !reflect.DeepEqual(got, expect) {
		t.Fatal()
	}
}

func TestRouteStat(t *testing.T) {
	t.Run("Simple routes", func(t *testing.T) {
		RouteStatTestHelper(
			t,
			[]route{{"ABC", "XYZ"}, {"ABC", "XYZ"}, {"DEF", "ABC"}},
			routeStat{[]route{{"ABC", "XYZ"}}, 1, 2, 1},
		)
	})
	t.Run("Complex routes", func(t *testing.T) {
		RouteStatTestHelper(
			t,
			[]route{
				{"ABC", "XYZ"},
				{"XYZ", "ABC"},
				{"ABC", "XYZ"},
				{"DEF", "ABC"},
				{"ABC", "XYZ"},
				{"DEF", "ABC"},
				{"DEF", "ABC"},
			},
			routeStat{[]route{{"ABC", "XYZ"}, {"DEF", "ABC"}}, 2, 3, 1},
		)
	})
	t.Run("Duplicate route", func(t *testing.T) {
		RouteStatTestHelper(
			t,
			[]route{{"ABC", "DEF"}, {"ABC", "DEF"}, {"ABC", "XYZ"}, {"ABC", "XYZ"}},
			routeStat{[]route{{"ABC", "DEF"}, {"ABC", "XYZ"}}, 2, 2, 0},
		)
	})
	t.Run("Various routes", func(t *testing.T) {
		RouteStatTestHelper(
			t,
			[]route{{"ABC", "XYZ"}, {"XYZ", "ABC"}, {"DEF", "ABC"}, {"DEF", "XYZ"}},
			routeStat{[]route{{"ABC", "XYZ"}, {"XYZ", "ABC"}, {"DEF", "ABC"}, {"DEF", "XYZ"}}, 4, 4, 4},
		)
	})
}

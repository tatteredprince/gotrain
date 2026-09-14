package main

import (
	"reflect"
	"testing"
)

type Route struct {
	From string
	To   string
}

type RouteStat struct {
	MaxRoutes      []Route
	MaxRoutesCount int
	Uniq           int
	Once           int
}

// analyzeRoutes calculates routes' statistics from argument: most frequent
// routes and their count, number of various routes and number of unique routes.
// Most frequent routes should be selected in order of appearance.
func analyzeRoutes(routes []Route) RouteStat {
	routesDigest := make(map[Route]int)
	orderedRoutes := make([]Route, 0, len(routes))
	for _, route := range routes {
		if _, ok := routesDigest[route]; ok {
			routesDigest[route]++
		} else {
			routesDigest[route] = 1
			orderedRoutes = append(orderedRoutes, route)
		}
	}
	var (
		maxRoutes      []Route
		maxRoutesCount = 0
		once           = 0
	)
	for _, route := range orderedRoutes {
		count := routesDigest[route]
		if count == 1 {
			once += 1
		}
		if count > maxRoutesCount {
			maxRoutes = make([]Route, 0, len(routesDigest))
			maxRoutes = append(maxRoutes, route)
			maxRoutesCount = count
		} else if count == maxRoutesCount {
			maxRoutes = append(maxRoutes, route)
		}
	}
	return RouteStat{
		MaxRoutes:      maxRoutes,
		MaxRoutesCount: len(maxRoutes),
		Uniq:           len(routesDigest),
		Once:           once,
	}
}

func RouteStatTestHelper(t *testing.T, routes []Route, expect RouteStat) {
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
			[]Route{{"ABC", "XYZ"}, {"ABC", "XYZ"}, {"DEF", "ABC"}},
			RouteStat{[]Route{{"ABC", "XYZ"}}, 1, 2, 1},
		)
	})
	t.Run("Complex routes", func(t *testing.T) {
		RouteStatTestHelper(
			t,
			[]Route{
				{"ABC", "XYZ"},
				{"XYZ", "ABC"},
				{"ABC", "XYZ"},
				{"DEF", "ABC"},
				{"ABC", "XYZ"},
				{"DEF", "ABC"},
				{"DEF", "ABC"},
			},
			RouteStat{[]Route{{"ABC", "XYZ"}, {"DEF", "ABC"}}, 2, 3, 1},
		)
	})
	t.Run("Duplicate route", func(t *testing.T) {
		RouteStatTestHelper(
			t,
			[]Route{{"ABC", "DEF"}, {"ABC", "DEF"}, {"ABC", "XYZ"}, {"ABC", "XYZ"}},
			RouteStat{[]Route{{"ABC", "DEF"}, {"ABC", "XYZ"}}, 2, 2, 0},
		)
	})
	t.Run("Various routes", func(t *testing.T) {
		RouteStatTestHelper(
			t,
			[]Route{{"ABC", "XYZ"}, {"XYZ", "ABC"}, {"DEF", "ABC"}, {"DEF", "XYZ"}},
			RouteStat{[]Route{{"ABC", "XYZ"}, {"XYZ", "ABC"}, {"DEF", "ABC"}, {"DEF", "XYZ"}}, 4, 4, 4},
		)
	})
}

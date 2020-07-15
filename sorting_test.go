package main

import (
	"sort"
	"testing"
)

func TestSorting(t *testing.T) {
	routeOne := Route{"test", 12, 30}
	routeTwo := Route{"test", 12, 35}
	routeThree := Route{"test", 22, 20}
	routeFour := Route{"test", 25, 40}

	routesSlice := RoutesSlice{routeThree, routeFour, routeTwo, routeOne}

	sort.Sort(routesSlice)

	if routesSlice[0] != routeOne || routesSlice[1] != routeTwo || routesSlice[2] != routeThree || routesSlice[3] != routeFour {
		t.Error("Sorting was incorrect")
	}
}

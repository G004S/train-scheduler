package routing

import (
	"fmt"
	"pathfinder/internal/models"
	"sort"
	"time"
)

func FindAllRoutes(stations map[string]*models.Station, startSt, endSt string) ([]*models.Route, error) {

	shortestPath, err := BFS(stations, startSt, endSt)
	if err != nil {
		return nil, fmt.Errorf(
			"no path exists between %q and %q",
			startSt,
			endSt,
		)
	}

	maxPathLength := len(shortestPath) + 5
	//for saving the routes
	var Routes []*models.Route
	deadline := time.Now().Add(15 * time.Second)

	err = DFS(stations[startSt], stations[endSt], map[string]bool{}, &models.Route{}, &Routes, maxPathLength, deadline)
	if err != nil {
		return nil, err
	}

	if len(Routes) == 0 {
		return nil, fmt.Errorf("no path exists between %q and %q", startSt, endSt)
	}

	sort.Slice(Routes, func(i, j int) bool {
		return len(Routes[i].Stations) < len(Routes[j].Stations)
	})
	return Routes, nil
}

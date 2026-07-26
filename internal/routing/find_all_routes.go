package routing

import (
	"fmt"
	"pathfinder/internal/models"
	"sort"
)

func FindAllRoutes(stations map[string]*models.Station, startSt, endSt string) ([]*models.Route, error) {

	//for saving the routes
	var Routes []*models.Route

	DFS(stations[startSt], stations[endSt], map[string]bool{}, &models.Route{}, &Routes)

	if len(Routes) == 0 {
		return nil, fmt.Errorf("no path exists between %q and %q", startSt, endSt)
	}

	sort.Slice(Routes, func(i, j int) bool {
		return len(Routes[i].Stations) < len(Routes[j].Stations)
	})
	return Routes, nil
}

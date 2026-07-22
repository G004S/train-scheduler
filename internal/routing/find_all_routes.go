package routing

import (
	"pathfinder/internal/models"
	"sort"
)

func FindAllRoutes(stations map[string]*models.Station, startSt, endSt string) []*models.Route {

	//for saving the routes
	var Routes []*models.Route

	DFS(stations[startSt], stations[endSt], map[string]bool{}, &models.Route{}, &Routes)

	sort.Slice(Routes, func(i, j int) bool {
		return len(Routes[i].Stations) < len(Routes[j].Stations)
	})
	return Routes
}

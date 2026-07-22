package scheduling

import (
	"fmt"
	"math"
	"pathfinder/internal/models"
	"pathfinder/internal/routing"
	"pathfinder/internal/simulation"
	"strings"
)

func FindOptimalRouteGroup(stations map[string]*models.Station, start, end string, trainCount int) {
	//get ALL the possible routes
	allroutes := routing.FindAllRoutes(stations, start, end)

	//group routs into NON-connected groups
	GroupsOfUniqueRoutes := routing.FindUniqueRoutes(allroutes)

	//find best through simulation
	best := math.MaxInt
	var bestMoves [][]string

	for _, v := range GroupsOfUniqueRoutes {
		moves := simulation.MoveTrains(v, trainCount)
		if len(moves) < best {
			best = len(moves)
			bestMoves = moves
		}
	}

	for _, m := range bestMoves {
		fmt.Println(strings.Join(m, " "))
	}
}

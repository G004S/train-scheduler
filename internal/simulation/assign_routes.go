package simulation

import (
	"pathfinder/internal/models"
)

func GiveRoutes(routes []*models.Route, trainCount int) []models.Train {
	//stores amount of trains per path
	given := make([]int, len(routes))

	trains := make([]models.Train, trainCount)

	for i := 0; i < trainCount; i++ {

		//stores ID of current lightest path (considering its length and amount of trains on it)
		bestID := 0

		//calculates the lightest path
		for j := 1; j < len(routes); j++ {

			scoreBest := len(routes[bestID].Stations) + given[bestID]

			scoreCurrent := len(routes[j].Stations) + given[j]

			if scoreCurrent < scoreBest {
				bestID = j
			}
		}

		//asign path to a train
		trains[i] = models.Train{
			ID:       i + 1,
			Position: 0,
			Route:    routes[bestID].Stations,
		}

		//add weight to the path
		given[bestID]++
	}

	return trains
}

package simulation

import (
	"fmt"
	"pathfinder/internal/models"
)

func MoveTrains(routes []*models.Route, trainCount int) [][]string {

	var totalmoves [][]string

	trains := GiveRoutes(routes, trainCount)

	occupied := make(map[string]bool)

	for {

		moved := false
		moves := []string{}

		usedTracks := make(map[string]bool)

		for i := range trains {

			train := &trains[i]

			if train.Position >= len(train.Route)-1 {
				continue
			}

			currentStation := train.Route[train.Position]

			next := train.Position + 1
			nextStation := train.Route[next]

			isEnd := next == len(train.Route)-1

			track := currentStation.Name + "-" + nextStation.Name

			// The same track cannot be used twice during one turn.
			if usedTracks[track] {
				continue
			}

			if occupied[nextStation.Name] && !isEnd {
				continue
			}

			if train.Position != 0 {
				current := train.Route[train.Position]
				occupied[current.Name] = false
			}

			train.Position = next

			if !isEnd {
				occupied[nextStation.Name] = true
			}

			// Mark this track as used during the current turn.
			usedTracks[track] = true

			moves = append(
				moves,
				fmt.Sprintf("T%d-%s", train.ID, nextStation.Name),
			)

			moved = true
		}

		totalmoves = append(totalmoves, moves)

		if !moved {
			break
		}
	}

	return totalmoves
}

// // moving trains by one step
//
//	func MoveTrainOneStep(train *Train, path []*Station) bool {
//		if train.Position >= len(path)-1 {
//			return false //cannot move if train is on the end station
//		}
//		train.Position++
//		return true
//	}

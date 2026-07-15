package main

import (
	"fmt"
	"strings"
)

func MoveTrains(routes [][]*Station, trainCount int) {
	trains := make([]Train, trainCount)

	//Enumerate trains
	for i := range trains {
		trains[i] = Train{
			ID:       i + 1,
			Position: 0,
			Route:    routes[i%len(routes)],
		}
	}

	//Map for tracking stations' vacancy
	occupied := make(map[string]bool)

	//Big loop - turn
	for {
		moved := false
		moves := []string{}

		for i := range trains {
			train := &trains[i]
			if train.Position >= len(train.Route)-1 {
				continue
			}

			nextPosition := train.Position + 1
			nextStation := train.Route[nextPosition]

			isEnd := nextPosition == len(train.Route)-1

			if occupied[nextStation.Name] && !isEnd {
				continue
			}

			curStation := train.Route[train.Position]

			if train.Position != 0 {
				occupied[curStation.Name] = false
			}

			train.Position = nextPosition

			if !isEnd {
				occupied[nextStation.Name] = true
			}

			moves = append(moves, fmt.Sprintf("PosT%d-%s", train.ID, nextStation.Name))

			moved = true
		}

		if !moved {
			break
		}

		fmt.Println(strings.Join(moves, " "))
	}
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

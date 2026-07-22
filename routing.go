package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func DFS(current, end *Station, visited map[string]bool, path *Route, allPaths *[]*Route) {
	path.Stations = append(path.Stations, current)

	if current == end {
		copyPath := make([]*Station, len(path.Stations))
		copy(copyPath, path.Stations)

		*allPaths = append(*allPaths, &Route{
			Stations: copyPath,
		})

		path.Stations = path.Stations[:len(path.Stations)-1]
		return
	}

	visited[current.Name] = true

	for _, neighbour := range current.Neighbours {
		if visited[neighbour.Name] {
			continue
		}

		DFS(neighbour, end, visited, path, allPaths)
	}

	visited[current.Name] = false

	path.Stations = path.Stations[:len(path.Stations)-1]
}

func (S *Setup) FindAllRoutes(stations map[string]*Station) []*Route {

	//for saving the routes
	var Routes []*Route

	DFS(stations[S.Start], stations[S.End], map[string]bool{}, &Route{}, &Routes)

	sort.Slice(Routes, func(i, j int) bool {
		return len(Routes[i].Stations) < len(Routes[j].Stations)
	})
	return Routes
}

func FindUniqueRoutes(routes []*Route) [][]*Route {

	//result
	var RouteGroups [][]*Route

	//sort out connected routes
	for _, base := range routes {

		group := []*Route{base}

		occupied := make(map[string]bool)

		// mark used nodes for BASE ROUTE
		for i := 1; i < len(base.Stations)-1; i++ {
			occupied[base.Stations[i].Name] = true
		}

		//Mark all the routes as NOT interconnected until proven wrong
		for _, candidate := range routes {

			if candidate == base {
				continue
			}

			interconnected := false

			//check if any stations are used in marked routes
			for i := 1; i < len(candidate.Stations)-1; i++ {

				if occupied[candidate.Stations[i].Name] {
					interconnected = true
					break
				}
			}

			if interconnected {
				continue
			}

			//if candidate wasn't wiped out by the IF above then it's not connected
			group = append(group, candidate)

			//mark candidate's stations as used
			for i := 1; i < len(candidate.Stations)-1; i++ {
				occupied[candidate.Stations[i].Name] = true
			}
		}

		//Add group
		RouteGroups = append(RouteGroups, group)
	}

	return RouteGroups
}

// give route to each train
func GiveRoutes(routes []*Route, trainCount int) []Train {
	//stores amount of trains per path
	given := make([]int, len(routes))

	trains := make([]Train, trainCount)

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
		trains[i] = Train{
			ID:       i + 1,
			Position: 0,
			Route:    routes[bestID].Stations,
		}

		//add weight to the path
		given[bestID]++
	}

	return trains
}

func (S *Setup) FindOptimalRouteGroup(stations map[string]*Station) {
	//get ALL the possible routes
	allroutes := S.FindAllRoutes(stations)

	//group routs into NON-connected groups
	GroupsOfUniqueRoutes := FindUniqueRoutes(allroutes)

	//find best through simulation
	best := math.MaxInt
	var bestMoves [][]string

	for _, v := range GroupsOfUniqueRoutes {
		moves := MoveTrains(v, S.Trains)
		if len(moves) < best {
			best = len(moves)
			bestMoves = moves
		}
	}

	for _, m := range bestMoves {
		fmt.Println(strings.Join(m, " "))
	}
}

// func BFS(stationMap map[string]*Station, usedStations map[string]bool, start, end string) ([]*Station, error) {

// 	startStation, exists := stationMap[start]
// 	if !exists {
// 		return nil, errors.New("Error: start station not found")
// 	}
// 	endStation, exists := stationMap[end]
// 	if !exists {
// 		return nil, errors.New("Error: end station not found")
// 	}

// 	// to know what station to check next
// 	var ochered []*Station
// 	ochered = append(ochered, startStation)

// 	// to track already visited stations
// 	visited := make(map[string]bool)
// 	visited[start] = true

// 	// to know which station we got here from (for remembering the path)\
// 	parent := make(map[string]*Station)

// 	for len(ochered) > 0 {
// 		curStation := ochered[0]
// 		ochered = ochered[1:]

// 		if curStation.Name == endStation.Name {
// 			break
// 		}

// 		for _, neighbour := range curStation.Neighbours {
// 			if visited[neighbour.Name] || usedStations[neighbour.Name] {
// 				continue
// 			}

// 			parent[neighbour.Name] = curStation
// 			visited[neighbour.Name] = true
// 			ochered = append(ochered, neighbour)
// 		}
// 	}

// 	if !visited[endStation.Name] {
// 		return nil, errors.New("Error: Path not found")
// 	}

// 	path := restorePath(parent, startStation, endStation)

// 	return path, nil
// }

// func restorePath(parent map[string]*Station, startStation, endStation *Station) []*Station {

// 	path := []*Station{}

// 	for cur := endStation; cur != nil; cur = parent[cur.Name] {
// 		path = append(path, cur)

// 		if cur == startStation {
// 			break
// 		}
// 	}

// 	//reverse the order
// 	for first, last := 0, len(path)-1; first < last; first, last = first+1, last-1 {
// 		path[first], path[last] = path[last], path[first]
// 	}

// 	return path
// }

// func FindAllRoutes(stations map[string]*Station, start, end string) ([][]*Station, error) {

// 	//fpr excluding routes that have already been found
// 	usedStations := make(map[string]bool)

// 	//for specific cases where distance between end and start is 1 edge
// 	foundRoutes := make(map[string]bool)

// 	//result
// 	var routes [][]*Station

// 	for {
// 		route, err := BFS(stations, usedStations, start, end)

// 		if err != nil {
// 			break
// 		}

// 		key := RouteToString(route)

// 		if foundRoutes[key] {
// 			break
// 		}

// 		routes = append(routes, route)

// 		foundRoutes[key] = true

// 		for i := 1; i < len(route)-1; i++ {
// 			usedStations[route[i].Name] = true
// 		}
// 	}

// 	if len(routes) == 0 {
// 		return nil, errors.New("Error: no routes found")
// 	}

// 	return routes, nil
// }

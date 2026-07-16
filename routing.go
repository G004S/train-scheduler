package main

func DFS(current, end *Station, visited map[string]bool, path []*Station, allPaths *[][]*Station) {
	//add current to the path
	path = append(path, current)

	//end station found --> save the path
	if current == end {

		copyPath := make([]*Station, len(path))
		copy(copyPath, path)

		*allPaths = append(*allPaths, copyPath)
		return
	}

	//mark the sttion as visited to prevent infinite loops in current branch
	visited[current.Name] = true

	for _, neighbour := range current.Neighbours {

		if visited[neighbour.Name] {
			continue
		}
		//step forward to the next UNVISITED neighbour
		DFS(neighbour, end, visited, path, allPaths)
	}

	//set the station as unvisited for next possible paths
	visited[current.Name] = false
}

func FindRoutes(stations map[string]*Station, start, end string) [][]*Station {

	//for saving the routes
	var Routes [][]*Station

	DFS(stations[start], stations[end], map[string]bool{}, []*Station{}, &Routes)

	return Routes
}

// give route to each train
func GiveRoutes(routes [][]*Station, trainCount int) []Train {

	//stores amount of trains per path
	given := make([]int, len(routes))

	trains := make([]Train, trainCount)

	for i := 0; i < trainCount; i++ {

		//stores ID of current lightest path (considering its length and amount of trains on it)
		bestID := 0

		//calculates the lightest path
		for j := 1; j < len(routes); j++ {

			scoreBest := len(routes[bestID]) + given[bestID]

			scoreCurrent := len(routes[j]) + given[j]

			if scoreCurrent < scoreBest {
				bestID = j
			}
		}

		//asign path to a train
		trains[i] = Train{
			ID:       i + 1,
			Position: 0,
			Route:    routes[bestID],
		}

		//add weight to the path
		given[bestID]++
	}

	return trains
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

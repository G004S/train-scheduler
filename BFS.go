package main

import (
	"errors"
)

func BFS(stationMap map[string]*Station, usedStations map[string]bool, start, end string) ([]*Station, error) {

	startStation, exists := stationMap[start]
	if !exists {
		return nil, errors.New("Error: start station not found")
	}
	endStation, exists := stationMap[end]
	if !exists {
		return nil, errors.New("Error: end station not found")
	}

	// to know what station to check next
	var ochered []*Station
	ochered = append(ochered, startStation)

	// to track already visited stations
	visited := make(map[string]bool)
	visited[start] = true

	// to know which station we got here from (for remembering the path)\
	parent := make(map[string]*Station)

	for len(ochered) > 0 {
		curStation := ochered[0]
		ochered = ochered[1:]

		if curStation.Name == endStation.Name {
			break
		}

		for _, neighbour := range curStation.Neighbours {
			if visited[neighbour.Name] || usedStations[neighbour.Name] {
				continue
			}

			parent[neighbour.Name] = curStation
			visited[neighbour.Name] = true
			ochered = append(ochered, neighbour)
		}
	}

	if !visited[endStation.Name] {
		return nil, errors.New("Error: Path not found")
	}

	path := restorePath(parent, startStation, endStation)

	return path, nil
}

func restorePath(parent map[string]*Station, startStation, endStation *Station) []*Station {

	path := []*Station{}

	for cur := endStation; cur != nil; cur = parent[cur.Name] {
		path = append(path, cur)

		if cur == startStation {
			break
		}
	}

	//reverse the order
	for first, last := 0, len(path)-1; first < last; first, last = first+1, last-1 {
		path[first], path[last] = path[last], path[first]
	}

	return path
}

func FindAllRoutes(stations map[string]*Station, start, end string) ([][]*Station, error) {

	//fpr excluding routes that have already been found
	usedStations := make(map[string]bool)

	//result
	var routes [][]*Station

	for {
		route, err := BFS(stations, usedStations, start, end)

		if err != nil {
			break
		}

		routes = append(routes, route)

		for i := 1; i < len(route)-1; i++ {
			usedStations[route[i].Name] = true
		}
	}

	if len(routes) == 0 {
		return nil, errors.New("Error: no routes found")
	}

	return routes, nil
}

package routing

import (
	"errors"
	"pathfinder/internal/models"
)

func BFS(stationMap map[string]*models.Station, start, end string) ([]*models.Station, error) {

	startStation, exists := stationMap[start]
	if !exists {
		return nil, errors.New("Error: start station not found")
	}
	endStation, exists := stationMap[end]
	if !exists {
		return nil, errors.New("Error: end station not found")
	}

	// to know what station to check next
	var ochered []*models.Station
	ochered = append(ochered, startStation)

	// to track already visited stations
	visited := make(map[string]bool)
	visited[start] = true

	// to know which station we got here from (for remembering the path)\
	parent := make(map[string]*models.Station)

	for len(ochered) > 0 {
		curStation := ochered[0]
		ochered = ochered[1:]

		if curStation == endStation {
			break
		}

		for _, neighbour := range curStation.Neighbours {
			if visited[neighbour.Name] {
				continue
			}

			parent[neighbour.Name] = curStation
			visited[neighbour.Name] = true
			ochered = append(ochered, neighbour)
		}
	}

	if !visited[endStation.Name] {
		return nil, errors.New("path not found")
	}

	path := restorePath(parent, startStation, endStation)

	return path, nil
}

func restorePath(parent map[string]*models.Station, startStation, endStation *models.Station) []*models.Station {

	path := []*models.Station{}

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

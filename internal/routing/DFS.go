package routing

import (
	"errors"
	"pathfinder/internal/models"
	"time"
)

const maxRoutes = 100

var ErrSearchTimeout = errors.New("route search exceeded 15 seconds")

func DFS(current, end *models.Station, visited map[string]bool, path *models.Route, allPaths *[]*models.Route, maxPathLength int, deadline time.Time) error {

	if time.Now().After(deadline) {
		return ErrSearchTimeout
	}

	if len(*allPaths) >= maxRoutes {
		return nil
	}

	if len(path.Stations)+1 > maxPathLength {
		return nil
	}

	path.Stations = append(path.Stations, current)

	if current == end {
		copyPath := make([]*models.Station, len(path.Stations))
		copy(copyPath, path.Stations)

		*allPaths = append(*allPaths, &models.Route{
			Stations: copyPath,
		})

		path.Stations = path.Stations[:len(path.Stations)-1]
		return nil
	}

	visited[current.Name] = true

	for _, neighbour := range current.Neighbours {
		if visited[neighbour.Name] {
			continue
		}

		err := DFS(neighbour, end, visited, path, allPaths, maxPathLength, deadline)
		if err != nil {
			visited[current.Name] = false
			path.Stations = path.Stations[:len(path.Stations)-1]
			return err
		}
	}

	visited[current.Name] = false

	path.Stations = path.Stations[:len(path.Stations)-1]

	return nil
}

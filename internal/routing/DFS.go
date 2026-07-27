package routing

import (
	"pathfinder/internal/models"
)

const maxRoutes = 200

//const maxPathLength = 50

func DFS(current, end *models.Station, visited map[string]bool, path *models.Route, allPaths *[]*models.Route) {

	if len(*allPaths) >= maxRoutes {
		return
	}

	// if len(path.Stations) >= maxPathLength {
	// 	return
	// }

	path.Stations = append(path.Stations, current)

	if current == end {
		copyPath := make([]*models.Station, len(path.Stations))
		copy(copyPath, path.Stations)

		*allPaths = append(*allPaths, &models.Route{
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

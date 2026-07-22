package parsing

import (
	"fmt"
	"pathfinder/internal/models"
	"strings"
)

// ParseMap reads map content and creates stations with their connections.
func ParseMap(mapContent string) (map[string]*models.Station, error) {

	if strings.TrimSpace(mapContent) == "" {
		return nil, fmt.Errorf("map is empty")
	}

	// Initialize map for easy access to stations by name
	stationMap := make(map[string]*models.Station)
	connectionMap := make(map[string]bool)

	lines := strings.Split(mapContent, "\n")
	var stationsLines []string
	var connectionsLines []string

	section := ""

	hasStationsSection := false
	hasConnectionSection := false

	countStations := 0

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		if line == strings.TrimSpace("stations:") {
			section = "stations"
			hasStationsSection = true
			continue
		}

		if line == strings.TrimSpace("connections:") {
			// if !hasStationsSection {
			// 	return nil, fmt.Errorf("missing stations section")
			// }
			section = "connections"
			hasConnectionSection = true
			continue
		}

		switch section {
		case "stations":
			stationsLines = append(stationsLines, line)
		case "connections":
			connectionsLines = append(connectionsLines, line)
		default:
			return nil, fmt.Errorf("line outside of the section: %s", line)
		}

	}

	if !hasStationsSection {
		return nil, fmt.Errorf("missing stations section")
	}

	if !hasConnectionSection {
		return nil, fmt.Errorf("missing connections section")
	}

	if len(stationsLines) == 0 {
		return nil, fmt.Errorf("stations section is empty")
	}

	if len(connectionsLines) == 0 {
		return nil, fmt.Errorf("connections section is empty")
	}

	for _, line := range stationsLines {
		if err := ParseStations(line, stationMap); err != nil {
			return nil, err
		}
		countStations++
		if countStations > 10000 {
			return nil, fmt.Errorf("stations amount should be less than 10K. provided: %d", countStations)
		}
	}

	for _, line := range connectionsLines {
		if err := ParseConnections(line, stationMap, connectionMap); err != nil {
			return nil, err
		}
	}

	return stationMap, nil
}

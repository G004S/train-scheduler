package parsing

import (
	"fmt"
	"pathfinder/internal/models"
	"strings"
)

type NumberedLine struct {
	Text   string
	Number int
}

// ParseMap reads map content and creates stations with their connections.
func ParseMap(mapContent string) (map[string]*models.Station, error) {

	if strings.TrimSpace(mapContent) == "" {
		return nil, fmt.Errorf("map is empty")
	}

	// Initialize map for easy access to stations by name
	stationMap := make(map[string]*models.Station)
	connectionMap := make(map[string]bool)

	lines := strings.Split(mapContent, "\n")
	var stationsLines []NumberedLine
	var connectionsLines []NumberedLine

	section := ""

	hasStationsSection := false
	hasConnectionSection := false

	countStations := 0

	for i, line := range lines {

		lineNumber := i + 1

		line = strings.TrimSpace(line)
		line = strings.Split(line, "#")[0]
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

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
			stationsLines = append(stationsLines, NumberedLine{
				Text:   line,
				Number: lineNumber,
			})
		case "connections":
			connectionsLines = append(connectionsLines, NumberedLine{
				Text:   line,
				Number: lineNumber,
			})
		default:
			return nil, fmt.Errorf("line outside of the section: %d", lineNumber)
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

		if err := ParseStations(line.Text, stationMap); err != nil {
			return nil, fmt.Errorf("line %d: %w", line.Number, err)
		}
		countStations++
		if countStations > 10000 {
			return nil, fmt.Errorf("too many stations: maximum is 10000, provided %d", countStations)
		}
	}

	for _, line := range connectionsLines {
		if err := ParseConnections(line.Text, stationMap, connectionMap); err != nil {
			return nil, fmt.Errorf("line %d: %w", line.Number, err)
		}
	}

	return stationMap, nil
}

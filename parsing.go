package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseMap reads map content and creates stations with their connections.
func ParseMap(mapContent string) (map[string]*Station, error) {

	if strings.TrimSpace(mapContent) == "" {
		return nil, fmt.Errorf("map is empty")
	}

	// Initialize map for easy access to stations by name
	stationMap := make(map[string]*Station)
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

func ParseStations(line string, stationMap map[string]*Station) error {
	partsStations := strings.Split(line, ",")

	if len(partsStations) != 3 {
		return fmt.Errorf("invalid station format")
	}

	name := strings.TrimSpace(partsStations[0])
	if name == "" {
		return fmt.Errorf("empty station name")
	}

	if name != strings.ToLower(name) {
		return fmt.Errorf("station name %q must be in lower case", name)
	}

	for _, char := range name {
		if (char < 'a' || char > 'z') && char != '_' && (char < '0' || char > '9') {
			return fmt.Errorf("invalid station name %q", name)
		}
	}

	_, exists := stationMap[name]
	if exists {
		return fmt.Errorf("dublicate station %q", name)

	}

	x, err1 := strconv.Atoi(strings.TrimSpace(partsStations[1]))
	y, err2 := strconv.Atoi(strings.TrimSpace(partsStations[2]))

	if err1 != nil || err2 != nil {
		return fmt.Errorf("invalid coordinates")
	}

	if x <= 0 || y <= 0 {
		return fmt.Errorf("coordinates for station %q must be positive integers", name)
	}

	for _, station := range stationMap {
		if station.X == x && station.Y == y {
			return fmt.Errorf("stations %q and %q have the same coordinates", station.Name, name)
		}
	}

	stationMap[name] = &Station{
		Name: name,
		X:    x,
		Y:    y,
	}
	return nil
}

func ParseConnections(line string, stationMap map[string]*Station, connectionMap map[string]bool) error {
	partsConnect := strings.Split(line, "-")
	if len(partsConnect) != 2 {
		return fmt.Errorf("invalid connections format")

	}

	from := strings.TrimSpace(partsConnect[0])
	to := strings.TrimSpace(partsConnect[1])

	if from != strings.ToLower(from) || to != strings.ToLower(to) {
		return fmt.Errorf("station names in connection section must be in lower case")
	}

	if from == "" || to == "" {
		return fmt.Errorf("empty connection station")
	}

	if from == to {
		return fmt.Errorf("cannot connect to itself")
	}

	fromStation, ok1 := stationMap[from]
	toStation, ok2 := stationMap[to]

	if !ok1 || !ok2 {
		return fmt.Errorf("unknown station in connection between %q and %q", from, to)
	}

	connectionKey := from + "-" + to
	reverseKey := to + "-" + from

	if connectionMap[connectionKey] || connectionMap[reverseKey] {
		return fmt.Errorf("dublicate connection between %q and %q", from, to)
	}
	connectionMap[connectionKey] = true

	fromStation.Neighbours = append(fromStation.Neighbours, toStation)
	toStation.Neighbours = append(toStation.Neighbours, fromStation)

	return nil
}

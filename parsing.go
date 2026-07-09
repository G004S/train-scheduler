package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseMap reads map content and creates stations with their connections.
func ParseMap(mapContent string) (map[string]*Station, error) {

	// Initialize map for easy access to stations by name
	stationMap := make(map[string]*Station)

	lines := strings.Split(mapContent, "\n")

	section := ""

	for _, line := range lines {

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if line == strings.TrimSpace("stations:") {
			section = "stations"
			continue
		}

		if line == strings.TrimSpace("connections:") {
			section = "connections"
			continue
		}

		if section == "stations" {
			err := ParseStations(line, stationMap)
			if err != nil {
				return nil, err
			}
			continue
		}

		if section == "connections" {
			err := ParseConnections(line, stationMap)
			if err != nil {
				return nil, err
			}
			continue
		}
		return nil, fmt.Errorf("Line outside of the section: %s", line)
	}
	return stationMap, nil
}

func ParseStations(line string, stationMap map[string]*Station) error {
	partsStations := strings.Split(line, ",")

	if len(partsStations) != 3 {
		return fmt.Errorf("Invalid station line")

	}

	name := strings.TrimSpace(partsStations[0])
	if name == "" {
		return fmt.Errorf("Empty station name")

	}

	if _, exists := stationMap[name]; exists {
		return fmt.Errorf("Dublicate station")

	}

	x, err1 := strconv.Atoi(strings.TrimSpace(partsStations[1]))
	y, err2 := strconv.Atoi(strings.TrimSpace(partsStations[2]))

	if err1 != nil || err2 != nil {
		return fmt.Errorf("Invalid coordinates")

	}

	stationMap[name] = &Station{
		Name: name,
		X:    x,
		Y:    y,
	}
	return nil
}

func ParseConnections(line string, stationMap map[string]*Station) error {
	partsConnect := strings.Split(line, "-")
	if len(partsConnect) != 2 {
		return fmt.Errorf("Invalid station line")

	}

	from := strings.TrimSpace(partsConnect[0])
	to := strings.TrimSpace(partsConnect[1])

	if from == "" || to == "" {
		return fmt.Errorf("Empty connection station")

	}

	if from == to {
		return fmt.Errorf("Cannot connect to itself")

	}

	fromStation, ok1 := stationMap[from]
	toStation, ok2 := stationMap[to]

	if !ok1 || !ok2 {
		return fmt.Errorf("Unknown station in connection")

	}

	fromStation.Neighbours = append(fromStation.Neighbours, toStation)
	toStation.Neighbours = append(toStation.Neighbours, fromStation)

	return nil
}

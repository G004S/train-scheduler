package parsing

import (
	"fmt"
	"pathfinder/internal/models"
	"strconv"
	"strings"
)

func ParseStations(line string, stationMap map[string]*models.Station) error {
	partsStations := strings.Split(line, ",")

	if len(partsStations) != 3 {
		return fmt.Errorf("invalid station format %q: expected name,x,y", line)
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
			return fmt.Errorf("invalid station name %q: only lowercase letters, digits and '_' are allowed", name)
		}
	}

	_, exists := stationMap[name]
	if exists {
		return fmt.Errorf("duplicate station %q", name)

	}

	x, err1 := strconv.Atoi(strings.TrimSpace(partsStations[1]))
	y, err2 := strconv.Atoi(strings.TrimSpace(partsStations[2]))

	if err1 != nil {
		return fmt.Errorf("invalid coordinate X: %q is not an integer", strings.TrimSpace(partsStations[1]))
	}

	if err2 != nil {
		return fmt.Errorf("invalid coordinate Y %q: expected an integer", strings.TrimSpace(partsStations[2]))
	}

	if x < 0 || y < 0 {
		return fmt.Errorf("coordinates for station %q must be positive integers", name)
	}

	for _, station := range stationMap {
		if station.X == x && station.Y == y {
			return fmt.Errorf("stations %q and %q have the same coordinates (%d, %d)", station.Name, name, x, y)
		}
	}

	stationMap[name] = &models.Station{
		Name: name,
		X:    x,
		Y:    y,
	}
	return nil
}

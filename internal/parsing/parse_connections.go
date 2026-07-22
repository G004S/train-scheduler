package parsing

import (
	"fmt"
	"pathfinder/internal/models"
	"strings"
)

func ParseConnections(line string, stationMap map[string]*models.Station, connectionMap map[string]bool) error {
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
		return fmt.Errorf("invalid connection format %q: expected station1-station2", line)
	}

	if from == to {
		return fmt.Errorf("station %q cannot connect to itself", from)
	}

	fromStation, ok1 := stationMap[from]
	toStation, ok2 := stationMap[to]

	if !ok1 || !ok2 {
		return fmt.Errorf("unknown station in connection %q-%q", from, to)
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

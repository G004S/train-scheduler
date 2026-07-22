package validation

import (
	"fmt"
	"pathfinder/internal/models"
)

func ValidateInput(setup models.Setup, stationMap map[string]*models.Station) error {
	if _, exists := stationMap[setup.Start]; !exists {
		return fmt.Errorf("start station %q does not exists", setup.Start)
	}

	if _, exists := stationMap[setup.End]; !exists {
		return fmt.Errorf("end station %q does not exists", setup.End)
	}
	return nil
}

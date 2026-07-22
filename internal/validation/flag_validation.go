package validation

import (
	"fmt"
	"pathfinder/internal/models"
)

func ValidateFlags(setup models.Setup) error {

	if setup.MapPath == "" {
		return fmt.Errorf("missing map")
	}

	if setup.Start == setup.End {
		return fmt.Errorf("start and end station are the same")
	}

	if setup.Trains <= 0 {
		return fmt.Errorf(
			"invalid number of trains: %d is not a positive integer",
			setup.Trains,
		)
	}

	if setup.Trains >= 3000 {
		return fmt.Errorf("trains number should not be more than 3000")
	}

	return nil
}

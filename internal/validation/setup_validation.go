package validation

import (
	"fmt"
	"os"
	"pathfinder/internal/models"
	"strconv"
)

// PrepareSetup reads command-line flags, validates them,
// and finds the correct map file path.

func PrepareSetup() (models.Setup, error) {

	if len(os.Args) != 5 {
		return models.Setup{}, fmt.Errorf("invalid number of arguments")
	}

	trains, err := strconv.Atoi(os.Args[4])
	if err != nil {
		return models.Setup{}, fmt.Errorf("invalid train number %q", os.Args[4])
	}

	setup := models.Setup{
		MapPath: os.Args[1],
		Start:   os.Args[2],
		End:     os.Args[3],
		Trains:  trains,
	}

	if err := ValidateFlags(setup); err != nil {
		return models.Setup{}, err
	}

	validMapPath, err := FindMapPath(setup.MapPath)
	if err != nil {
		return models.Setup{}, err
	}

	setup.MapPath = validMapPath
	return setup, nil
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Setup struct {
	MapPath string
	Start   string
	End     string
	Trains  int
}

// PrepareSetup reads command-line flags, validates them,
// and finds the correct map file path.

func PrepareSetup() (Setup, error) {

	if len(os.Args) != 5 {
		return Setup{}, fmt.Errorf("invalid number of arguments")
	}

	trains, err := strconv.Atoi(os.Args[4])
	if err != nil {
		return Setup{}, fmt.Errorf("invalid train number %q", os.Args[4])
	}

	setup := Setup{
		MapPath: os.Args[1],
		Start:   os.Args[2],
		End:     os.Args[3],
		Trains:  trains,
	}

	if err := ValidateFlags(setup); err != nil {
		return Setup{}, err
	}

	validMapPath, err := FindMapPath(setup.MapPath)
	if err != nil {
		return Setup{}, err
	}

	setup.MapPath = validMapPath
	return setup, nil
}

func ValidateFlags(setup Setup) error {

	if setup.MapPath == "" {
		return fmt.Errorf("missing map")
	}

	// if setup.Start == "" {
	// 	return fmt.Errorf("missing start station")
	// }

	// if setup.End == "" {
	// 	return fmt.Errorf("missing end station")
	// }

	if setup.Start == setup.End {
		return fmt.Errorf("start and end station are the same")
	}

	if setup.Trains <= 0 {
		return fmt.Errorf("trains number should be more than 0")
	}

	return nil
}

// FindMapPath checks if the map file exists.
// If not, it also tries the same path with ".map" added.

func FindMapPath(mapPath string) (string, error) {
	inputPath, _ := os.Stat(mapPath)

	if inputPath == nil {
		inputPath, _ = os.Stat(mapPath + ".map")
		if inputPath == nil {
			return "", errors.New("file containing network map not found")
		}
		return mapPath + ".map", nil
	}
	return mapPath, nil
}

func ValidateInput(setup Setup, stationMap map[string]*Station) error {
	if _, exists := stationMap[setup.Start]; !exists {
		return fmt.Errorf("start station %q does not exists", setup.Start)
	}

	if _, exists := stationMap[setup.End]; !exists {
		return fmt.Errorf("end station %q does not exists", setup.End)
	}
	return nil
}

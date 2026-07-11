package main

import (
	"errors"
	"flag"
	"os"
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

	mapPath := flag.String("map", "", "path to map file")
	start := flag.String("start", "", "start station")
	end := flag.String("end", "", "end station")
	trains := flag.Int("trains", 0, "number of trains")

	flag.Parse()

	setup := Setup{
		MapPath: *mapPath,
		Start:   *start,
		End:     *end,
		Trains:  *trains,
	}

	err := ValidateFlags(setup)
	if err != nil {
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
		return errors.New("Error: missing map flag")
	}

	if setup.Start == "" {
		return errors.New("Error: missing start flag")
	}

	if setup.End == "" {
		return errors.New("Error: missing end flag")
	}

	if setup.Start == setup.End {
		return errors.New("Error: Start and end station are the same")
	}

	if setup.Trains <= 0 {
		return errors.New("Error: Invalid amount of trains")
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
			return "", errors.New("Error: file containing network map not found")
		}
		return mapPath + ".map", nil
	}
	return mapPath, nil
}

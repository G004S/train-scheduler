package validation

import (
	"errors"
	"os"
)

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

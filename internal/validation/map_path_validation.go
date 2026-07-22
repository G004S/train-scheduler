package validation

import (
	"fmt"
	"os"
)

// FindMapPath checks if the map file exists.
// If not, it also tries the same path with ".map" added.

func FindMapPath(mapPath string) (string, error) {
	fileInfo, err := os.Stat(mapPath)

	if err == nil {
		if fileInfo.IsDir() {
			return "", fmt.Errorf(
				"path %q is a directory, expected a map file",
				mapPath,
			)
		}

		return mapPath, nil
	}

	if !os.IsNotExist(err) {
		return "", fmt.Errorf(
			"cannot access map file %q: %w",
			mapPath,
			err,
		)
	}

	pathWithExtension := mapPath + ".map"

	fileInfo, err = os.Stat(pathWithExtension)

	if err == nil {
		if fileInfo.IsDir() {
			return "", fmt.Errorf(
				"path %q is a directory, expected a map file",
				pathWithExtension,
			)
		}

		return pathWithExtension, nil
	}

	if !os.IsNotExist(err) {
		return "", fmt.Errorf(
			"cannot access map file %q: %w",
			pathWithExtension,
			err,
		)
	}

	return "", fmt.Errorf(
		"map file not found: tried %q and %q",
		mapPath,
		pathWithExtension,
	)
}

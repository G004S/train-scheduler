package io

import (
	"fmt"
	"os"
)

func GetMapContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read map file %q: %w", path, err)
	}

	return string(data), nil
}

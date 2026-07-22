package io

import (
	"fmt"
	"os"
)

func GetMapContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("couldn't open file: %w", err)
	}

	return string(data), nil
}

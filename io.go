package main

import (
	"errors"
	"os"
)

func GetMapContent(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("Error: couldn't open file")
	}

	return string(data), nil
}

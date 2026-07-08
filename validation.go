package main

import (
	"errors"
	"os"
	"strconv"
)

// Run all the validators and return valid MAP or ERROR
func ValidateSetup() (string, error) {
	err := ValidateArgs()
	if err != nil {
		return "", err
	}

	mapPath, err := ValidateMapPath()
	if err != nil {
		return "", err
	}

	return mapPath, nil
}

// Validate the number of args
func ValidateArgs() error {
	if len(os.Args) != 5 {
		return errors.New("Error: Invalid amount of arguments")
	}

	if os.Args[2] == os.Args[3] {
		return errors.New("Error: Start and end station are the same")
	}

	trainNum, err := strconv.Atoi(os.Args[4])
	if err != nil {
		return errors.New("Error: Invalid train number format")
	}
	if trainNum <= 0 {
		return errors.New("Error: Invalid amount of trains")
	}
	return nil
}

// func ValidateMapContent(path string) (string, error) {
// 	data, err := os.ReadFile(path)
// 	if err != nil {
// 		return "", errors.New("Error: couldn't open file")
// 	}
// 	//check for duplicated stations

// 	return string(data), nil
// }

// Validate existance of map file
func ValidateMapPath() (string, error) {
	inputPath, _ := os.Stat(os.Args[1])
	if inputPath == nil {
		inputPath, _ = os.Stat(os.Args[1] + ".map")
		if inputPath == nil {
			return "", errors.New("Error: file containing network map not found")
		}
		return os.Args[1] + ".map", nil
	}
	return os.Args[1], nil
}

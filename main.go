package main

import "fmt"

func main() {

	//Validates setup: args, path to .map
	setup, err := PrepareSetup()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Get content from .map
	mapContent, err := GetMapContent(setup.MapPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	stations, err := ParseMap(mapContent)
	if err != nil {
		fmt.Println(err)
		return
	}

	//just to check
	for name, station := range stations {
		fmt.Printf("Stations: %s (%d, %d)\n", name, station.X, station.Y)
		fmt.Print("Neighbours: ")
		for _, neighbour := range station.Neighbours {
			fmt.Printf("%s ", neighbour.Name)
		}
		fmt.Println()
	}

}

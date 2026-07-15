package main

import (
	"fmt"
	"os"
)

func ExitError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func main() {

	//Validates setup: args, path to .map
	setup, err := PrepareSetup()
	ExitError(err)

	//Get content from .map
	mapContent, err := GetMapContent(setup.MapPath)
	ExitError(err)

	stations, err := ParseMap(mapContent)
	ExitError(err)

	err = ValidateInput(setup, stations)
	ExitError(err)

	//just to check
	sname := setup.Start
	ename := setup.End

	path, _ := BFS(stations, sname, ename)
	fmt.Println("Path:")

	for _, v := range path {
		fmt.Println(v.Name)
	}

	fmt.Println()
	MoveTrains(path, setup.Trains)

	// for name, station := range  stations{
	// 	fmt.Printf("Stations: %s (%d, %d)\n", name, station.X, station.Y)
	// 	fmt.Print("Neighbours: ")
	// 	for _, neighbour := range station.Neighbours {
	// 		fmt.Printf("%s ", neighbour.Name)
	// 	}
	// 	fmt.Println()
	// }

}

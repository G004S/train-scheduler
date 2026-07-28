package main

import (
	"fmt"
	"os"
	"pathfinder/internal/io"
	"pathfinder/internal/parsing"
	scheduling "pathfinder/internal/scheduler"
	"pathfinder/internal/validation"
)

func ExitError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func main() {

	//Validates setup: args, path to .map
	setup, err := validation.PrepareSetup()
	ExitError(err)

	//Get content from .map
	mapContent, err := io.GetMapContent(setup.MapPath)
	ExitError(err)

	stations, err := parsing.ParseMap(mapContent)
	ExitError(err)

	err = validation.ValidateInput(setup, stations)
	ExitError(err)

	// _, err = routing.FindAllRoutes(stations, setup.Start, setup.End)
	// ExitError(err)

	err = scheduling.FindOptimalRouteGroup(stations, setup.Start, setup.End, setup.Trains)
	ExitError(err)

	//////!!!! SIMULATIONS

	///////////////////////////////////////////////////////////////////////////////

	// for name, station := range  stations{
	// 	fmt.Printf("Stations: %s (%d, %d)\n", name, station.X, station.Y)
	// 	fmt.Print("Neighbours: ")
	// 	for _, neighbour := range station.Neighbours {
	// 		fmt.Printf("%s ", neighbour.Name)
	// 	}
	// 	fmt.Println()
	// }

}

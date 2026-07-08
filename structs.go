package main

// Contains everything related to a specific station
type Station struct {
	Name string
	X    int
	Y    int

	Neighbours []*Station
}

// Initialize map for easy access to stations by name
var stationMap map[string]*Station

// Train struct
type Train struct {
	ID int
}

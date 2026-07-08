package main

// Contains everything related to a specific station
type Station struct {
	Name string
	X    int
	Y    int

	Neighbours []*Station
}

// Train struct
type Train struct {
	ID int
}

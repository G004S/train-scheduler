package models

// Contains everything related to a specific route
type Route struct {
	ID                int
	Stations          []*Station
	ConflictingRoutes map[int]bool
}

// Contains everything related to a specific station
type Station struct {
	Name string
	X    int
	Y    int

	Neighbours []*Station
}

// Train struct
type Train struct {
	ID       int
	Position int
	Route    []*Station
}

type Setup struct {
	MapPath string
	Start   string
	End     string
	Trains  int
}

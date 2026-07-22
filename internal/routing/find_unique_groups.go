package routing

import (
	"pathfinder/internal/models"
)

func FindUniqueRoutes(routes []*models.Route) [][]*models.Route {

	//result
	var RouteGroups [][]*models.Route

	//sort out connected routes
	for _, base := range routes {

		group := []*models.Route{base}

		occupied := make(map[string]bool)

		// mark used nodes for BASE ROUTE
		for i := 1; i < len(base.Stations)-1; i++ {
			occupied[base.Stations[i].Name] = true
		}

		//Mark all the routes as NOT interconnected until proven wrong
		for _, candidate := range routes {

			if candidate == base {
				continue
			}

			interconnected := false

			//check if any stations are used in marked routes
			for i := 1; i < len(candidate.Stations)-1; i++ {

				if occupied[candidate.Stations[i].Name] {
					interconnected = true
					break
				}
			}

			if interconnected {
				continue
			}

			//if candidate wasn't wiped out by the IF above then it's not connected
			group = append(group, candidate)

			//mark candidate's stations as used
			for i := 1; i < len(candidate.Stations)-1; i++ {
				occupied[candidate.Stations[i].Name] = true
			}
		}

		//Add group
		RouteGroups = append(RouteGroups, group)
	}

	return RouteGroups
}

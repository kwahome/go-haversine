package main

import (
	"fmt"
	"github.com/kwahome/haversine/haversine"
)

func main() {
	origin := haversine.Coordinate{
		Latitude: 1,
		Longitude: 1,
	}

	remote := haversine.Coordinate{
		Latitude: 2,
		Longitude: 3,
	}

	fmt.Println(origin.DistanceTo(remote, haversine.KM))
}

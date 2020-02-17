package main

import (
	"fmt"
	"github.com/kwahome/go-haversine/pkg/haversine"
)

func main() {
	nairobi := haversine.Coordinate{
		Latitude: 1.2921,
		Longitude: 36.8219,
	}
	mombasa := haversine.Coordinate{
		Latitude: 4.0435,
		Longitude: 39.6682,
	}

	units := haversine.M
	distance := nairobi.DistanceTo(mombasa, units)
	fmt.Println("Distance from Nairobi =", nairobi, "to Mombasa =", mombasa, "in", units, "is", distance)

	units = haversine.KM
	distance = nairobi.DistanceTo(mombasa, units)
	fmt.Println("Distance from Nairobi =", nairobi, "to Mombasa =", mombasa, "in", units, "is", distance)

	units = haversine.MI
	distance = nairobi.DistanceTo(mombasa, units)
	fmt.Println("Distance from Nairobi =", nairobi, "to Mombasa =", mombasa, "in", units, "is", distance)
}

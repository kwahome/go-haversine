/*
The Haversine formula determines the great-circle distance between two points on a sphere given their latitudes and
longitudes. Important in navigation, it is a special case of a more general formula in spherical trigonometry,
the law of Haversines, that relates the sides and angles of spherical triangles.

Below is an example that shows how to calculate the shortest path between two coordinates on the surface of the Earth.


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
*/
package haversine

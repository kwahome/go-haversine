package haversine

import (
	"fmt"
	"math"
)

const (
	// degrees that constitute π radians
	DegreesInPiRadian = 180

	MetersPerKm = 1000

	// radius of the earth in miles
	EarthRadiusMiles = 3958

	// radius of the earth in kilometers
	EarthRadiusKm = 6371

	// radius of the earth in meters
	EarthRadiusMeters = EarthRadiusKm * MetersPerKm

	// kilometers unit
	KM Unit = "kilometers"

	// meters unit
	M Unit = "meters"

	// miles unit
	MI Unit = "miles"
)

type Unit string
type Distance float64
type Kilometers Distance
type Meters Distance
type Miles Distance

/*
	A coordinate represents a geographic coordinate.
*/
type Coordinate struct {
	Latitude  float64
	Longitude float64
}

/*
	A Coordinate representing the change in two coordinates
*/
type Delta Coordinate

/*
	Implementation of the Stringer interface for Coordinate type.
*/
func (coordinate Coordinate) String() string {
	return fmt.Sprintf("{latitude=%v, longitude=%v}", coordinate.Latitude, coordinate.Longitude)
}

/*
	Returns the radians equivalent of a Coordinate expressed in degrees.
*/
func (coordinate Coordinate) ToRadians() Coordinate {
	return Coordinate{
		Latitude:  degreesToRadians(coordinate.Latitude),
		Longitude: degreesToRadians(coordinate.Longitude),
	}
}

/*
	Method that returns the Haversine distance between it's receiver and a remote coordinate.
*/
func (coordinate Coordinate) DistanceTo(remote Coordinate, unit Unit) Distance {
	return toUnit(haversineDistance(coordinate, remote), unit)
}

/*
	Calculates and returns the Delta of a coordinate w.r.t another coordinate.
*/
func (coordinate Coordinate) Delta(origin Coordinate) Delta {
	return Delta{
		Latitude:  coordinate.Latitude - origin.Latitude,
		Longitude: coordinate.Longitude - origin.Longitude,
	}
}

/*
	degreesToRadians converts degrees to radians.
*/
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / DegreesInPiRadian
}

/*
	haversineDistance calculates the shortest path between two coordinates on
	the surface of the Earth.
*/
func haversineDistance(origin, remote Coordinate) Distance {
	origin, remote = origin.ToRadians(), remote.ToRadians()
	change := origin.Delta(remote)

	angle := math.Pow(math.Sin(change.Latitude/2), 2) + math.Cos(origin.Latitude)*math.Cos(remote.Latitude)*
		math.Pow(math.Sin(change.Longitude/2), 2)

	return Distance(2 * math.Atan2(math.Sqrt(angle), math.Sqrt(1-angle)))
}

/*

 */
func toUnit(distance Distance, unit Unit) Distance {
	var result Distance
	switch unit {
	case KM:
		result = Distance(Kilometers(distance * EarthRadiusKm))
	case M:
		result = Distance(Meters(distance * EarthRadiusMeters))
	case MI:
		result = Distance(Miles(distance * EarthRadiusMiles))
	}
	return result
}

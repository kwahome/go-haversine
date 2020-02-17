package haversine

import (
	"testing"
)

var testCases = []struct {
	from             Coordinate
	to               Coordinate
	expectedDistance []struct{
		unit Unit
		distance Distance
	}
	expectedDelta Delta
	expectedRadians struct{
		from Coordinate
		to Coordinate
	}
}{
	{
		from:             Coordinate{0, 0},
		to:               Coordinate{0, 0},
		expectedDistance: []struct {
			unit     Unit
			distance Distance
		}{
			{
				unit: M,
				distance: Distance(0),
			},
			{
				unit: KM,
				distance: Distance(0),
			},
			{
				unit: MI,
				distance: Distance(0),
			},
		},
		expectedDelta: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude: 0,
			Longitude: 0,
		},
		expectedRadians: struct {
			from Coordinate
			to   Coordinate
		}{
			from: Coordinate{
				Latitude:  0,
				Longitude: 0,
			},
			to: Coordinate{
				Latitude:  0,
				Longitude: 0,
			},
		},
	},
	{
		from:             Coordinate{0, 0},
		to:               Coordinate{-180, -180},
		expectedDistance: []struct {
			unit     Unit
			distance Distance
		}{
			{
				unit: M,
				distance: Distance(0),
			},
			{
				unit: KM,
				distance: Distance(0),
			},
			{
				unit: MI,
				distance: Distance(0),
			},
		},
		expectedDelta: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude: 180,
			Longitude: 180,
		},
		expectedRadians: struct {
			from Coordinate
			to   Coordinate
		}{
			from: Coordinate{
				Latitude:  0,
				Longitude: 0,
			},
			to: Coordinate{
				Latitude:  0,
				Longitude: 0,
			},
		},
	},
	{
		from:             Coordinate{22.55, 43.12},
		to:               Coordinate{13.45, 100.28},
		expectedDistance: []struct {
			unit     Unit
			distance Distance
		}{
			{
				unit: M,
				distance: Distance(6094544.408786774),
			},
			{
				unit: KM,
				distance: Distance(6094.544408786774),
			},
			{
				unit: MI,
				distance: Distance(3786.251258825624),
			},
		},
		expectedDelta: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude: 9.100000000000001,
			Longitude: -57.160000000000004,
		},
		expectedRadians: struct {
			from Coordinate
			to   Coordinate
		}{
			from: Coordinate{
				Latitude:  0.3935717463247213,
				Longitude: 0.7525859734599548,
			},
			to: Coordinate{
				Latitude:  0.3935717463247213,
				Longitude: 0.7525859734599548,
			},
		},
	},
}

func TestCoordinate_DistanceTo(t *testing.T) {
	for _, testCase := range testCases {
		for _, expectedDistance := range testCase.expectedDistance {
			result := testCase.from.DistanceTo(testCase.to, expectedDistance.unit)

			if expectedDistance.distance != result {
				t.Errorf("FAIL: Expected distance from %v to %v in %v to be: %v but got %v",
					testCase.from,
					testCase.to,
					expectedDistance.unit,
					expectedDistance.distance,
					result,
				)
			}
		}
	}
}

func TestCoordinate_ToRadians(t *testing.T) {
	for _, testCase := range testCases {

		fromInRadians := testCase.from.ToRadians()
		toInRadians := testCase.from.ToRadians()

		if testCase.expectedRadians.from != fromInRadians {
			t.Errorf("FAIL: Expected %v in radians to be: %v but got %v ",
				testCase.from,
				testCase.expectedRadians.from,
				fromInRadians,
			)
		}

		if testCase.expectedRadians.to != toInRadians {
			t.Errorf("FAIL: Expected %v in radians to be: %v but got %v ",
				testCase.to,
				testCase.expectedRadians.to,
				toInRadians,
			)
		}
	}
}

func TestCoordinate_Delta(t *testing.T) {
	for _, testCase := range testCases {

		result := testCase.from.Delta(testCase.to)

		if testCase.expectedDelta != result {
			t.Errorf("FAIL: Expected delta of %v with respect to %v to be: %v but got %v ",
				testCase.from,
				testCase.to,
				testCase.expectedDelta,
				result,
			)
		}
	}
}

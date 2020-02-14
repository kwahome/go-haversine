package haversine

import (
	"testing"
)

var testCases = []struct {
	from     Coordinate
	to       Coordinate
	unit     Unit
	expected Distance
}{
	{
		from:     Coordinate{0, 0},
		to:       Coordinate{0, 0},
		unit:     M,
		expected: Distance(0),
	},
	{
		from:     Coordinate{0, 0},
		to:       Coordinate{-180, -180},
		unit:     KM,
		expected: Distance(0),
	},
	{
		from:     Coordinate{22.55, 43.12},
		to:       Coordinate{13.45, 100.28},
		unit:     M,
		expected: Distance(6094544.408786774),
	},
	{
		from:     Coordinate{22.55, 43.12},
		to:       Coordinate{13.45, 100.28},
		unit:     KM,
		expected: Distance(6094.544408786774),
	},
	{
		from:     Coordinate{22.55, 43.12},
		to:       Coordinate{13.45, 100.28},
		unit:     MI,
		expected: Distance(3786.251258825624),
	},
}

func TestHaversineDistance(t *testing.T) {
	for _, testCase := range testCases {

		result := testCase.from.DistanceTo(testCase.to, testCase.unit)

		if testCase.expected != result {
			t.Errorf("FAIL: Expected distance from %v to %v to be: %v %v but got %v %v",
				testCase.from,
				testCase.to,
				testCase.expected,
				testCase.unit,
				result,
				testCase.unit,
			)
		}
	}
}

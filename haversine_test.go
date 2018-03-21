package main

import (
	"testing"
)

var tests = []struct {
	A        Point
	B        Point
	distance float64
}{
	{
		Point{Lat: 22.55, Lng: 43.12},  // Rio de Janeiro
		Point{Lat: 13.45, Lng: 100.28}, // Bangkok
		6094.544408786774,
	},
	{
		Point{Lat: 20.10, Lng: 57.30}, // Port Louis
		Point{Lat: 0.57, Lng: 100.21}, // Padang
		5145.525771394785,
	},
}

func TestHaversineDistance(t *testing.T) {

	for _, point := range tests {

		distance := point.A.KilometresTo(point.B)

		if point.distance != distance {
			t.Errorf("Fail: want %v got %v", point.distance, distance)
		}

	}

}

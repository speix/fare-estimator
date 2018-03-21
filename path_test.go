package main

import (
	"testing"
)

func TestPathDiscarded(t *testing.T) {

	taxi := FareService(&Taxi{})
	segments := Segments{}
	paths := Paths{
		Path{
			RideId:    1,
			Lat:       37.954302,
			Lng:       23.71337,
			Timestamp: 1405595284,
		}, {
			RideId:    1,
			Lat:       37.938042,
			Lng:       23.692308,
			Timestamp: 1405595362,
		},
	}

	p, _ := paths.Filter(taxi, segments)

	if p[0].Timestamp != 1405595284 {
		t.Errorf("Fail: incorrect path discarded expected 1405595284 got %v", p[0].Timestamp)
	}

}

func TestPathCutoff(t *testing.T) {

	taxi := FareService(&Taxi{})
	segments := Segments{}
	paths := Paths{
		Path{
			RideId:    1,
			Lat:       37.954302,
			Lng:       23.71337,
			Timestamp: 1405595284,
		}, {
			RideId:    1,
			Lat:       37.938042,
			Lng:       23.692308,
			Timestamp: 1405595362,
		},
	}

	p, _ := paths.Filter(taxi, segments)

	if len(p) != 1 {
		t.Errorf("Fail: path not discarded expected 1 got %v", len(p))
	}

}

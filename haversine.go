package main

import (
	"math"
)

/*
	Haversine Formula taken, tested and improved from:
		https://en.wikipedia.org/wiki/Haversine_formula
		https://www.movable-type.co.uk/scripts/latlong.html
		https://github.com/paultag/go-haversine
*/

type Point struct {
	Lat float64
	Lng float64
}

type Delta struct {
	Lat float64
	Lng float64
}

var earthRadiusMetres float64 = 6371000

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func (p Point) KilometresTo(remote Point) float64 {
	return Distance(p, remote)
}

func (p Point) Delta(point Point) Delta {
	return Delta{
		Lat: p.Lat - point.Lat,
		Lng: p.Lng - point.Lng,
	}
}

func (p Point) toRadians() Point {
	return Point{
		Lat: degreesToRadians(p.Lat),
		Lng: degreesToRadians(p.Lng),
	}
}

func Distance(origin, position Point) float64 {
	origin = origin.toRadians()
	position = position.toRadians()

	change := origin.Delta(position)

	a := math.Pow(math.Sin(change.Lat/2), 2) + math.Cos(origin.Lat)*math.Cos(position.Lat)*math.Pow(math.Sin(change.Lng/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return float64((earthRadiusMetres * c) / 1000)
}

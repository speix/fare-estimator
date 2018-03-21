package main

import (
	"strconv"
)

type Path struct {
	RideId    int64
	Lat       float64
	Lng       float64
	Timestamp int64
}

type Paths []Path

func (path Path) Add(line []string) Path {

	path.RideId, _ = strconv.ParseInt(line[0], 10, 64)
	path.Lat, _ = strconv.ParseFloat(line[1], 64)
	path.Lng, _ = strconv.ParseFloat(line[2], 64)
	path.Timestamp, _ = strconv.ParseInt(line[3], 10, 64)

	return path
}

/*
	Filter measures the speed of two consecutive paths.
	If the speed is greater than a given threshold, the 2nd path is discarded.
	If not, the Segment Fare is estimated and appended back to the Segment.
	The function returns the updated Paths and Segments back to the Handler.
*/

func (paths Paths) Filter(fs FareService, segments Segments) (Paths, Segments) {

	pathsLength := len(paths)

	if pathsLength < 2 {
		return paths, segments
	}

	params := fs.GetParameters()
	pointA := Point{Lat: paths[pathsLength-2].Lat, Lng: paths[pathsLength-2].Lng}
	pointB := Point{Lat: paths[pathsLength-1].Lat, Lng: paths[pathsLength-1].Lng}

	distanceKm := pointA.KilometresTo(pointB)
	durationHrs := float64(paths[pathsLength-1].Timestamp-paths[pathsLength-2].Timestamp) * params.SecondsToHourConversion
	speed := distanceKm / durationHrs

	if speed > params.SpeedFilter {
		paths = paths[:pathsLength-1]
	} else {

		segments = append(segments, Segment{
			Fare: fs.EstimateSegmentFare(paths[pathsLength-1].Timestamp, distanceKm, durationHrs),
		})

	}

	return paths, segments
}

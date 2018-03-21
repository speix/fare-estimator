package main

type Ride struct {
	RideId   int64
	Segments []Segment
}

/*
	Estimate calculates the total Price of a Ride.
	Given a Ride and a FareService the function returns the RideId and the total price.
*/

func (ride Ride) Estimate(fs FareService) (int64, float64) {

	params := fs.GetParameters()
	total := params.BaseFare

	for _, segment := range ride.Segments {
		total += segment.Fare
	}

	if total < params.MinFare {
		total = params.MinFare
	}

	return ride.RideId, total
}

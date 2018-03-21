package main

/*
	FareService is an abstraction to decouple models from service implementation.
	We can easily add different providers if needed and better test our code.
*/

type FareService interface {
	GetParameters() Parameters
	EstimateSegmentFare(time int64, distance, duration float64) float64
	Export(ride int64, price float64, output string)
}

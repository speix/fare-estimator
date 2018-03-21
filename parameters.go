package main

/*
	Representation of the available FareService parameters.
*/
type Parameters struct {
	BaseFare                 float64
	MinFare                  float64
	IdleFare                 float64
	NormalFarePrice          float64
	DoubleFarePrice          float64
	NormalFareThresholdSpeed float64
	NormalFarePriceStart     int
	NormalFarePriceStop      int
	SpeedFilter              float64
	SecondsToHourConversion  float64
}

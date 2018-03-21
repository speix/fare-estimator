package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

type Taxi struct{}

func (b *Taxi) GetParameters() Parameters {

	params := Parameters{
		BaseFare:                 1.30,
		MinFare:                  3.47,
		IdleFare:                 11.90,
		NormalFarePrice:          0.74,
		DoubleFarePrice:          1.30,
		NormalFareThresholdSpeed: 10,
		NormalFarePriceStart:     5,
		NormalFarePriceStop:      23,
		SpeedFilter:              100,
		SecondsToHourConversion:  0.000277778,
	}

	return params
}

/*
	EstimateSegmentFare calculates and returns the Segment estimation price based on our business logic.
*/
func (b *Taxi) EstimateSegmentFare(epochTime int64, distance, duration float64) float64 {

	params := b.GetParameters()
	fare := 0.0

	t := time.Unix(epochTime, 0)

	speed := distance / duration

	if speed > params.NormalFareThresholdSpeed {

		if t.Hour() >= params.NormalFarePriceStart && t.Hour() <= params.NormalFarePriceStop {

			fare = params.NormalFarePrice * distance

		} else {

			fare = params.DoubleFarePrice * distance

		}

	} else {

		fare = params.IdleFare * duration

	}

	return fare
}

/*
	Export saves an estimation to the output file.
*/
func (b *Taxi) Export(ride int64, price float64, output string) {

	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer file.Close()

	if err != nil {
		log.Fatal(err.Error())
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.Write([]string{strconv.Itoa(int(ride)), strconv.FormatFloat(price, 'f', 6, 32)})
	csvWriter.Flush()
}

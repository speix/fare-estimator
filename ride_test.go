package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestRideEstimationMinimumFare(t *testing.T) {

	taxi := FareService(&Taxi{})
	params := taxi.GetParameters()
	ride := Ride{
		RideId: 1,
		Segments: Segments{
			{
				Fare: 1.5,
			},
			{
				Fare: 0.5,
			},
		},
	}

	output := "files/testRideEstimationMinimumFare.csv"

	defer os.Remove(output)

	rideId, total := ride.Estimate(taxi)
	taxi.Export(rideId, total, output)

	file, _ := os.Open(output)
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			t.Errorf("Fail: %v", err.Error())
		}

		price, _ := strconv.ParseFloat(line[1], 64)

		if price != params.MinFare {
			t.Errorf("Fail: expected %v got %v", params.MinFare, price)
		}
	}
}

func TestRideEstimation(t *testing.T) {

	taxi := FareService(&Taxi{})
	ride := Ride{
		RideId: 1,
		Segments: Segments{
			{
				Fare: 1.5,
			},
			{
				Fare: 2.5,
			},
		},
	}

	output := "files/testRideEstimation.csv"

	defer os.Remove(output)

	rideId, total := ride.Estimate(taxi)
	taxi.Export(rideId, total, output)

	file, _ := os.Open(output)
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			t.Errorf("Fail: %v", err.Error())
		}

		price, _ := strconv.ParseFloat(line[1], 64)

		if price != 5.3 {
			t.Errorf("Fail: expected 5.3 got %v", price)
		}
	}
}

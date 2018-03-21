package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

// We declare a WaitGroup to wait for a collection of goroutines to finish executing.
var wg sync.WaitGroup

/*
	Start function takes a pointer to the open file descriptor (*os.File) and an output file name.
	Reading the file line by line, its purpose is to:
		1. Filter the Paths and calculate the Segment prices.
		2. Calculate and save the combination of a Ride and its respective Price.
	Each RidePriceEstimation/Export activities are processed leveraging Go concurrency with goroutines.
	The function stops when all goroutines finish.
*/
func StartEstimator(file *os.File, output string) {

	reader := csv.NewReader(bufio.NewReader(file))

	var currentRide, rideId = int64(1), int64(1)
	var paths Paths
	var segments Segments
	var finished = false

	taxiService := &Taxi{}

	for {
		line, err := reader.Read()
		if err == io.EOF {
			finished = true
		} else if err != nil {
			log.Fatal(err)
		}

		if !finished {
			rideId, _ = strconv.ParseInt(line[0], 10, 64)
		}

		/*
			If RideId changes or no more input is available (EOF):
				1. create a new Ride struct.
				2. start calculating the fare price.
				3. write (id_ride, fare_estimate) to the output file.
		*/
		if currentRide != rideId || finished {
			ride := Ride{
				RideId:   currentRide,
				Segments: segments,
			}

			wg.Add(1) // We add up 1 to the number of goroutines on WaitGroup.

			go estimateRide(&ride, taxiService, output) // Handle Ride estimation concurrently.

			paths, segments = nil, nil // Reset Paths and Segments for the next Ride.
			currentRide = rideId

			if finished {
				break
			}
		}

		path := Path{}
		paths = append(paths, path.Add(line))
		paths, segments = paths.Filter(taxiService, segments)
	}

	wg.Wait() // Wait for goroutines to finish until the WaitGroup counter is zero.
}

func estimateRide(ride *Ride, service FareService, output string) {

	defer wg.Done() // Subtracts one by one once each goroutine is finished.

	rideId, total := ride.Estimate(service)
	service.Export(rideId, total, output)
}

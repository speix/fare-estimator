package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"testing"
)

var test = struct {
	RideId int64
	Price  float64
}{
	1,
	5.50,
}

type fareTest struct {
	epoch    int64   //Timestamp
	distance float64 //Km
	duration float64 //Hour
	expected float64 //Expected Price
}

func TestEstimateSegmentFarePM(t *testing.T) {

	taxi := FareService(&Taxi{})
	params := taxi.GetParameters()

	ft := fareTest{
		epoch:    1518283887, //Before 00:00 and after 05:00
		distance: 20,         //Speed > 10km/h
		duration: 1,
		expected: params.NormalFarePrice,
	}

	farePerHour := taxi.EstimateSegmentFare(ft.epoch, ft.distance, ft.duration) / ft.distance

	if ft.expected != farePerHour {
		t.Errorf("Fail: want %v got %v", ft.expected, farePerHour)
	}

}

func TestEstimateSegmentFareAM(t *testing.T) {

	taxi := FareService(&Taxi{})
	params := taxi.GetParameters()

	ft := fareTest{
		epoch:    1518229887, //After 00:00 and before 05:00
		distance: 20,         //Speed > 10km/h
		duration: 1,
		expected: params.DoubleFarePrice,
	}

	farePerHour := taxi.EstimateSegmentFare(ft.epoch, ft.distance, ft.duration) / ft.distance

	if ft.expected != farePerHour {
		t.Errorf("Fail: want %v got %v", ft.expected, farePerHour)
	}

}

func TestEstimateSegmentFareIdle(t *testing.T) {

	taxi := FareService(&Taxi{})
	params := taxi.GetParameters()

	ft := fareTest{
		epoch:    1518229887,
		distance: 9,
		duration: 1,
		expected: params.IdleFare,
	}

	farePerHour := taxi.EstimateSegmentFare(ft.epoch, ft.distance, ft.duration)

	if ft.expected != farePerHour {
		t.Errorf("Fail: want %v got %v", ft.expected, farePerHour)
	}

}

func TestExportFileCreated(t *testing.T) {

	taxi := FareService(&Taxi{})
	output := "files/testExportFileCreated.csv"

	defer os.Remove(output)

	taxi.Export(test.RideId, test.Price, output)

	_, err := os.Stat(output)
	os.IsNotExist(err)

	if err != nil {
		t.Errorf("Fail: %v", err.Error())
	}

}

func TestExportResultsWritten(t *testing.T) {

	taxi := FareService(&Taxi{})
	output := "files/testExportResultsWritten.csv"

	defer os.Remove(output)

	taxi.Export(test.RideId, test.Price, output)

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

		rideId, _ := strconv.ParseInt(line[0], 10, 64)
		price, _ := strconv.ParseFloat(line[1], 64)

		if rideId != 1 || price != 5.50 {
			t.Errorf("Fail: line records missmatch")
		}
	}

}

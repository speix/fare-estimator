package main

import (
	"os"
)

/*
	Fare Estimation Script
	Consists of a simple entry point given an input and an output filename.
	Both arguments are passed to the Estimator to start the calculation process.
*/

func main() {

	input := "files/paths.csv"
	output := "files/output.csv"

	file, _ := os.Open(input)
	defer file.Close()

	StartEstimator(file, output)
}

package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"testing"
)

//Test that handler Process function produces an output with the expected number of results
func TestEstimatorOutputExists(t *testing.T) {

	inputName := "files/test_input.csv"
	outputName := "files/test_output.csv"

	defer os.Remove(outputName)

	fileInput, _ := os.Open(inputName)
	defer fileInput.Close()

	StartEstimator(fileInput, outputName)

	_, err := os.Stat(outputName)
	os.IsNotExist(err)
	if err != nil {
		t.Errorf("Fail: %v", err.Error())
	}

	fileOutput, _ := os.Open(outputName)
	defer fileOutput.Close()

	reader := csv.NewReader(bufio.NewReader(fileOutput))

	answer := 9
	counter := 0

	for {
		_, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Errorf("Fail: %v", err.Error())
		}

		counter += 1
	}

	if answer != counter {
		t.Errorf("Fail: expected %v got %v", answer, counter)
	}

}

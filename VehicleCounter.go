package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read command line input for vehicle data file
	if len(os.Args) != 2 {
		log.Fatal("Usage: $> ./VehicleCounter <inputfilename>")
	}

	counterData, err := getCounterData(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d data points read.\n", len(counterData))

	for _, dataPoint := range counterData {
		sensor, err := getSensor(dataPoint)
		if err != nil {
			fmt.Println(err)
			continue
		}

		miliseconds, err := getDataTime(dataPoint)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("DataPoint: %s, Sensor: %s, Miliseconds: %d\n", dataPoint, sensor, miliseconds)
	}
}

// given a raw data point (e.g. B1089810), return the miliseconds portion (e.g. "1089810")
func getDataTime(dataPoint string) (int, error) {
	dataPoint = strings.TrimSpace(dataPoint)
	time := -1

	if len(dataPoint) > 0 {
		time, err := strconv.Atoi(dataPoint[1:len(dataPoint)])
		if err != nil {
			return time, err
		}

		return time, err
	} else {
		return time, errors.New("getDataTime(): data point has zero length.")
	}
}

// given a raw data point (e.g. B1089810), return the sensor it was collected from (e.g. "B")
func getSensor(dataPoint string) (string, error) {
	sensor := "undefined"

	dataPoint = strings.TrimSpace(dataPoint)
	if len(dataPoint) > 0 {
		sensor = dataPoint[:1]
	} else {
		return sensor, errors.New("getSensor(): data point has zero length.")
	}

	if !(sensor == "A" || sensor == "B") {
		return sensor, errors.New(fmt.Sprintf("Invalid sensor (%s) for data point %s", sensor, dataPoint))
	}

	return sensor, nil
}

// return a string slice of raw counter data points read in from disk file
func getCounterData(inputFile string) ([]string, error) {
	results := []string{}

	file, err := os.Open(inputFile)
	if err != nil {
		return results, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		dataPoint := strings.TrimSpace(scanner.Text())
		results = append(results, dataPoint)
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}

	return results, err
}

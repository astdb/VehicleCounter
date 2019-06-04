package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: $> ./VehicleCounter <inputfilename>")
	}

	counterData, err := getCounterData(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d data points read.\n", len(counterData))
}

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
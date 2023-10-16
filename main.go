package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Number of times to repeat the operations
	const N = 100

	// Path to the input and output files
	inputFilePath := "housesInput.csv"
	outputFilePath := "housesOutputGo.txt"

	// Open output file for writing
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	for i := 0; i < N; i++ {
		// Read the CSV file
		data, err := readCSV(inputFilePath)
		if err != nil {
			panic(err)
		}

		// Compute the summary statistics
		stats := computeSummaryStats(data)

		// Write the summary statistics to the output file
		writeStats(outputFile, stats)
	}
}

func readCSV(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func computeSummaryStats(data [][]string) map[string]map[string]float64 {
	stats := make(map[string]map[string]float64)

	// Skipping header row and initializing stats
	for _, colName := range data[0] {
		stats[colName] = map[string]float64{
			"min":   1e9,
			"max":   -1e9,
			"total": 0,
		}
	}

	// Compute min, max, and total for each column
	for _, row := range data[1:] {
		for j, value := range row {
			colName := data[0][j]
			floatVal, _ := strconv.ParseFloat(value, 64)

			if floatVal < stats[colName]["min"] {
				stats[colName]["min"] = floatVal
			}
			if floatVal > stats[colName]["max"] {
				stats[colName]["max"] = floatVal
			}

			stats[colName]["total"] += floatVal
		}
	}

	// Compute mean for each column
	for _, colName := range data[0] {
		stats[colName]["mean"] = stats[colName]["total"] / float64(len(data)-1)
	}

	return stats
}

func writeStats(file *os.File, stats map[string]map[string]float64) {
	for colName, colStats := range stats {
		fmt.Fprintf(file, "Column: %s\n", colName)
		fmt.Fprintf(file, "Min: %.2f, Max: %.2f, Mean: %.2f\n\n", colStats["min"], colStats["max"], colStats["mean"])
	}
	fmt.Fprintln(file, "-----------------------------")
}

package main

import (
	"fmt"
	"os"

	"linear-stats/internal/data"
	"linear-stats/internal/stats"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <data file>")
		os.Exit(1)
	}

	dataFile := os.Args[1]
	data, err := data.ReadData(dataFile)
	if err != nil {
		fmt.Printf("Error reading data: %v\n", err)
		os.Exit(1)
	}

	if len(data) < 2 {
		fmt.Println("Not enough data to perform calculations")
		os.Exit(1)
	}

	slope, intercept, pearsonCorrelation := stats.CalculateStatistics(data)

	fmt.Printf("Linear Regression Line: y = %.6fx + %.6f\n", slope, intercept)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", pearsonCorrelation)
}

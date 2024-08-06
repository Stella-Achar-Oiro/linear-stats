package stats

import (
	"math"
)

func CalculateStatistics(data []float64) (float64, float64, float64) {
	n := float64(len(data))
	if n == 0 {
		return 0, 0, 0 // Handle empty data case
	}

	sumX := 0.0
	sumY := 0.0
	sumXY := 0.0
	sumX2 := 0.0
	sumY2 := 0.0

	for i, y := range data {
		x := float64(i)
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
		sumY2 += y * y
	}

	denominator := n*sumX2 - sumX*sumX
	if denominator == 0 {
		return 0, 0, 0 // Handle zero denominator case
	}
	slope := (n*sumXY - sumX*sumY) / denominator
	intercept := (sumY - slope*sumX) / n

	denomPearson := math.Sqrt((n*sumX2 - sumX*sumX) * (n*sumY2 - sumY*sumY))
	if denomPearson == 0 {
		return slope, intercept, 0 // Handle zero denominator for Pearson correlation
	}
	pearsonCorrelation := (n*sumXY - sumX*sumY) / denomPearson

	return slope, intercept, pearsonCorrelation
}

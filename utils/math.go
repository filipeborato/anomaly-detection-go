package utils

import "math"

func StandardDeviation(values []float64) float64 {
	n := len(values)
	variance := Variance(values)
	standardVariation := math.Sqrt(variance / float64(n))

	return standardVariation
}

func Mean(values []float64) float64 {
	var sum float64
	n := len(values)
	for _, x := range values {
		sum += x
	}
	mean := sum / float64(n)
	return mean
}

func Variance(values []float64) float64 {
	var variance float64
	mean := Mean(values)
	for _, x := range values {
		variance += math.Pow(x-mean, 2)
	}
	return variance
}

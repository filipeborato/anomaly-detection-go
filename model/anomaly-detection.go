package model

import (
	"anomaly-detection-go/utils"
	"fmt"
)

func AnomalyDetection(values []float64) {
	mean := utils.Mean(values)
	deviations := utils.Deviation(values)
	standardDeveiation := utils.StandardDeviation(values)

	for _, deviation := range deviations {
		if deviation/mean > 0.50 {
			fmt.Println(mean)
			fmt.Println(deviation)
			fmt.Println(deviation / mean)
			fmt.Println("Alert")
		} else if standardDeveiation < deviation {
			fmt.Println(mean)
			fmt.Println(deviation)
			fmt.Println(standardDeveiation)
			fmt.Println("Warn")
		} else {
			fmt.Println("Info")
		}
	}
}

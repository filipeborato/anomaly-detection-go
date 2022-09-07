package handler

import (
	"anomaly-detection-go/repository"
	"anomaly-detection-go/utils"
	"fmt"
	"strconv"
)

func AnomalyDetection() {
	var values []float64
	rootDir := utils.RootDir()
	data := repository.ExtractCsv(rootDir+"/data/machine_temperature_system_failure.csv", ',')

	for _, line := range data {
		value, _ := strconv.ParseFloat(line.Value, 64)
		values = append(values, value)
	}
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

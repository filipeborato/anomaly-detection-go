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
	standardDeveiation := utils.StandardDeviation(values)
	fmt.Println(standardDeveiation)
}

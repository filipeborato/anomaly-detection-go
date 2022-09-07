package model

import (
	"anomaly-detection-go/model"
	"anomaly-detection-go/repository"
	"anomaly-detection-go/utils"
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
	model.AnomalyDetection(values)
}

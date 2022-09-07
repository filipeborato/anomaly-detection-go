package handler

import (
	"anomaly-detection-go/repository"
	"anomaly-detection-go/utils"
	"log"
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
			log.Println(mean)
			log.Println(deviation)
			log.Println(deviation / mean)
			log.Println("Alert")
		} else if standardDeveiation < deviation {
			log.Println(mean)
			log.Println(deviation)
			log.Println(standardDeveiation)
			log.Println("Warn")
		} else {
			log.Println("Info")
		}
	}
}

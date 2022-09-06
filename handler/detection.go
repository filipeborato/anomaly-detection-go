package handler

import (
	"anomaly-detection-go/repository"
	"anomaly-detection-go/utils"
	"fmt"
	"math"
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
		if math.Abs(deviation-mean) > 49.999 {
			fmt.Println("Alert")
		} else if standardDeveiation > deviation*2 {
			fmt.Println("Warn")
		} else {
			fmt.Println("Info")
		}
	}

}

package handler

import (
	"anomaly-detection-go/model"
	"anomaly-detection-go/repository"
	"anomaly-detection-go/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func AnomalyDetection(w http.ResponseWriter, r *http.Request) {
	var values []float64
	rootDir := utils.RootDir()
	data := repository.ExtractCsv(rootDir+"/data/machine_temperature_system_failure.csv", ',')

	for _, line := range data {
		value, _ := strconv.ParseFloat(line.Value, 64)
		values = append(values, value)
	}
	anomalyDetection := model.AnomalyDetection(values, data)
	err := json.NewEncoder(w).Encode(anomalyDetection)
	utils.CheckErr(err)
}

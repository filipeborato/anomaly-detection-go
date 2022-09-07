package model

import (
	"anomaly-detection-go/entity"
	"anomaly-detection-go/utils"
)

func AnomalyDetection(values []float64, csv []entity.CSV) []entity.AnomalyDetection {
	var anomalyDetection []entity.AnomalyDetection
	var level string
	mean := utils.Mean(values)
	deviations := utils.Deviation(values)
	standardDeveiation := utils.StandardDeviation(values)

	for i, deviation := range deviations {
		if deviation/mean > 0.50 {
			level = "Alert"
		} else if standardDeveiation < deviation {
			level = "Warn"
		} else {
			level = "Info"
		}
		anomaly := entity.AnomalyDetection{
			Timestamp: csv[i].Timestamp,
			Value:     csv[i].Value,
			Level:     level,
		}
		anomalyDetection = append(anomalyDetection, anomaly)

	}
	return anomalyDetection
}

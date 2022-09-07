package entity

type AnomalyDetection struct {
	Timestamp string `json:"timestamp"`
	Value     string `json:"value"`
	Level     string `json:"level"`
}

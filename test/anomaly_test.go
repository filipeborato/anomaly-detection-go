package test

import (
	"anomaly-detection-go/repository"
	"anomaly-detection-go/utils"
	"testing"
)

func TestCsvRead(t *testing.T) {
	rootDir := utils.RootDir()
	filePath := rootDir + "/data/machine_temperature_system_failure.csv"
	repository.ExtractCsv(filePath, ',')
}

package test

import (
	"anomaly-detection-go/handler"
	"anomaly-detection-go/repository"
	"anomaly-detection-go/utils"
	"math/rand"
	"testing"
)

func TestCsvRead(t *testing.T) {
	rootDir := utils.RootDir()
	filePath := rootDir + "/data/machine_temperature_system_failure.csv"
	repository.ExtractCsv(filePath, ',')
}
func TestRootDir(t *testing.T) {
	utils.RootDir()
}
func TestCheckError(t *testing.T) {
	var err error
	utils.CheckErr(err)
}
func TestReaderCsv(t *testing.T) {
	rootDir := utils.RootDir()
	filePath := rootDir + "/data/machine_temperature_system_failure.csv"
	repository.ExtractCsv(filePath, ',')
}
func testGenerateValues() []float64 {
	values := make([]float64, 60)
	for n := 0; n < 60; n++ {
		values[n] = rand.Float64()
	}
	return values
}
func TestMean(t *testing.T) {
	values := testGenerateValues()
	utils.Mean(values)
}

func TestVariance(t *testing.T) {
	values := testGenerateValues()
	utils.Variance(values)
}
func TestStandardDesviation(t *testing.T) {
	values := testGenerateValues()
	utils.StandardDeviation(values)
}
func TestDeviation(t *testing.T) {
	values := testGenerateValues()
	utils.Deviation(values)
}
func TestAnomalyDetection(t *testing.T) {
	handler.AnomalyDetection()
}

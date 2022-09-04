package test

import (
	"anomaly-detection-go/repository"
	"testing"
)

func TestCsvRead(t *testing.T) {
	//fmt.Println(os.Getwd())
	repository.ReadCsv("/home/filipeborato/projetos/anomaly-detection-go/data/machine_temperature_system_failure.csv")

}

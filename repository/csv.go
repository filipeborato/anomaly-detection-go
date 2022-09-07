package repository

import (
	"anomaly-detection-go/entity"
	"anomaly-detection-go/utils"
	"io"
)

func ExtractCsv(filePath string, comma rune) []entity.CSV {
	reader := utils.CsvReader(filePath, comma)
	var data []entity.CSV
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			utils.CheckErr(err)
		}
		data = append(data, entity.CSV{
			Timestamp: line[0],
			Value:     line[1],
		})
	}
	data = append(data[:0], data[1:]...)
	return data
}

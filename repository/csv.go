package repository

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type CSV struct {
	Timestamp string `json:"timestamp"`
	Value     string `json:"value"`
}

func checkErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func ReadCsv(filePath string) []CSV {
	csvFile, err := os.Open(filePath)
	checkErr(err)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ','

	var data []CSV

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			checkErr(err)
		}
		data = append(data, CSV{
			Timestamp: line[0],
			Value:     line[1],
		})
	}
	return data
}

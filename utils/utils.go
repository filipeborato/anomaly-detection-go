package utils

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func CheckErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}
func CsvReader(filePath string, comma rune) *csv.Reader {
	csvFile, err := os.Open(filePath)
	CheckErr(err)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = comma

	return reader
}
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

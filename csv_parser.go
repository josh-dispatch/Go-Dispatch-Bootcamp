package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

const CSV_FILE = "data.csv"

type CsvRecord struct {
	Id   int
	Name string
}

func GetCsvData() ([]CsvRecord, error) {

	file, err := os.Open(CSV_FILE)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// remember to close the file at the end of the program
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var csvRecords []CsvRecord

	for _, line := range data {

		var rec CsvRecord
		for j, field := range line {
			if j == 0 {
				intField, err := strconv.Atoi(field)

				if err != nil {
					log.Println(err)
					return nil, err
				}
				rec.Id = intField
			} else if j == 1 {
				rec.Name = field
			}
		}
		csvRecords = append(csvRecords, rec)

	}
	return csvRecords, nil
}

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type BuildingIdea struct {
	Title       string
	SourceTitle string
	SourceUrl   string
	Frequency   int64
}

func LoadData(config Config) (int64, []BuildingIdea) {
	var err error

	dataFd, err := os.Open(config.DataPath)

	if err != nil {
		println("Could not open the data file at " + config.DataPath + "!")
		panic(err)
	}

	dataDecoder := csv.NewReader(dataFd)

	data, err := dataDecoder.ReadAll()

	if err != nil {
		println("Could not read the data file at " + config.DataPath + "!")
		panic(err)
	}

	return LoadDataFromStringArray(data, config)
}

func LoadDataFromStringArray(data [][]string, config Config) (int64, []BuildingIdea) {
	totalFrequency := int64(0)

	var err error
	ideas := []BuildingIdea{}

	for index, row := range data {
		if len(row) < 5 {
			panic(fmt.Sprintf("Row %d is invalid! It must be of length 5.", index+1))
		}

		use := true
		useStr := row[3]
		if len(useStr) > 0 {
			use, err = strconv.ParseBool(row[3])

			if err != nil {
				fmt.Printf("Row %d is invalid! Column D (use) must be a Boolean.", index+1)
				panic(err)
			}
		}

		if use {
			frequency, err := strconv.ParseInt(row[4], 10, 64)

			if err != nil {
				fmt.Printf("Row %d is invalid! Column E (frequency) must be a positive number.", index+1)
				panic(err)
			}

			if frequency <= 0 {
				panic(fmt.Sprintf("Row %d is invalid! Column E (frequency) must be strictly positive.", index+1))
			}

			idea := BuildingIdea{
				Title:       row[0],
				SourceTitle: row[1],
				SourceUrl:   row[2],
				Frequency:   frequency,
			}

			ideas = append(ideas, idea)
			totalFrequency += frequency
		}
	}

	return totalFrequency, ideas
}

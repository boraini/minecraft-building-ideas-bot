package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	DataPath   string
	WebhookUrl string
}

func LoadConfig() Config {
	var err error

	var config Config
	configFd, err := os.Open("config.json")

	if err != nil {
		println("Could not open the config.json file!")
		panic(err)
	}

	configDecoder := json.NewDecoder(configFd)

	err = configDecoder.Decode(&config)

	if err != nil {
		panic(err)
	}

	return config
}

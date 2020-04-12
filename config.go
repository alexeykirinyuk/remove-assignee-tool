package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration for the application
type Configuration struct {
	Login  string
	Token  string
	Domain string
}

// GetConfiguration is a function for get configurations from config.example.json
func GetConfiguration() (config Configuration, err error) {
	file, err := os.Open("config.json")

	if err != nil {
		err = fmt.Errorf("error when open configuration file: %s", err)
		return
	}

	defer func() {
		err = file.Close()
	}()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		err = fmt.Errorf("error when parsing configurations: %s", err)
		return
	}

	return
}

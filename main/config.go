package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Config struct {
	PackSizes []int `json:"packSizes"`
	Max       int   `json:"Max"`
}

// LoadConfig loads configuration from a given file
func LoadConfig(filename string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	if config.PackSizes == nil {
		return config, errors.New("PackSizes is nil after loading config")
	}

	return config, nil
}

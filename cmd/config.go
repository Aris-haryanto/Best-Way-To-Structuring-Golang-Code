package main

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/Aris-haryanto/Best-Way-To-Structuring-Golang-Code/services"
	"gopkg.in/yaml.v2"
)

// Yaml Config
func getConfig(filePath string) (*services.Configs, error) {
	config := services.Configs{}

	if _, err := os.Stat(filePath); err != nil {
		return &config, errors.New("config path not valid")
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &config, err
	}

	err = yaml.Unmarshal([]byte(data), &config)
	return &config, err
}

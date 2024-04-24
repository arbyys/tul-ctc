package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func printError(message string) {
	fmt.Printf("\n[!] Following error thrown by the app:\n%s\n", message)
}

func loadConfig(filePath string) (configFile, error) {
	config := configFile{}
	buffer, err := os.ReadFile(filePath)

	if err != nil {
		return configFile{}, err
	}

	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		return configFile{}, err
	}

	return config, nil
}

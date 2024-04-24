package main

import (
	"fmt"
	"os"
)

func printError(message string) {
	fmt.Printf("\n[!] Following error thrown by the app:\n%s\n", message)
}

func loadConfig(filePath string) (string, error) {
	buffer, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return "temp", nil
}

package main

import "fmt"

func main() {
	fmt.Printf("App %s starting", appName)

	conf, err := loadConfig(inputFile)
	if err != nil {
		printError(err)
		return
	}

	fmt.Println(conf.Cars.Count)

	// app logic
}

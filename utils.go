package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func printError(message error) {
	fmt.Printf("\n[!] Following error thrown by the app:\n%s\n", message)
}

func loadConfig(filePath string) (configFileStruct, error) {
	config := configFileStruct{}
	buffer, err := os.ReadFile(filePath)

	if err != nil {
		return configFileStruct{}, err
	}

	err = yaml.Unmarshal(buffer, &config)
	if err != nil {
		return configFileStruct{}, err
	}

	return config, nil
}

func getShortestQueue(dispatchers []*dispatcher) int {
	sqw := shortestQueueWorker{
		count: queueMoreThanMax,
		index: -1,
	}

	for i := range dispatchers {
		currentQueueLength := len(dispatchers[i].queue)
		if dispatchers[i].occupied {
			currentQueueLength += 1
		}

		if currentQueueLength < sqw.count {
			sqw.count = currentQueueLength
			sqw.index = i
		}
	}

	return sqw.index
}

func generateNewCars(count int, timeMin int, timeMax int, sq *sharedQ) {
	for i := 0; i < count; i++ {
		sleepRandomTime(timeMin, timeMax)
		c := car{
			fuel: getRandomFuel(),
		}
		//c.waitForSharedQueueStarted = time.Now()
		sq.queue <- c
	}
}

func updateStats(c car) {

}

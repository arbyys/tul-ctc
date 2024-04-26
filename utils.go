package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

func printErrorAndDie(message error) {
	fmt.Printf("\n[!] Following error thrown by the app:\n%s\n", message)
	os.Exit(1)
}

func printStats() error {
	prettyStats, err := json.MarshalIndent(sc.stats, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(prettyStats))
	return nil
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

func exportStats(filePath string) (returnVal error) {
	buffer, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer func(buffer *os.File) {
		err := buffer.Close()
		if err != nil {
			returnVal = err
		}
	}(buffer)

	yamlEncoder := yaml.NewEncoder(buffer)
	err = yamlEncoder.Encode(sc.stats)
	if err != nil {
		return err
	}
	return nil

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

func addNewTimeStat(statType int, value time.Duration) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	switch statType {
	case gas:
		sc.stats.Stations.Gas.TotalTime += value
		sc.stats.Stations.Gas.TotalCars += 1
		if value > sc.stats.Stations.Gas.MaxTime {
			sc.stats.Stations.Gas.MaxTime = value
		}
	case diesel:
		sc.stats.Stations.Diesel.TotalTime += value
		sc.stats.Stations.Diesel.TotalCars += 1
		if value > sc.stats.Stations.Diesel.MaxTime {
			sc.stats.Stations.Diesel.MaxTime = value
		}
	case LPG:
		sc.stats.Stations.LPG.TotalTime += value
		sc.stats.Stations.LPG.TotalCars += 1
		if value > sc.stats.Stations.LPG.MaxTime {
			sc.stats.Stations.LPG.MaxTime = value
		}
	case electric:
		sc.stats.Stations.Electric.TotalTime += value
		sc.stats.Stations.Electric.TotalCars += 1
		if value > sc.stats.Stations.Electric.MaxTime {
			sc.stats.Stations.Electric.MaxTime = value
		}
	case register:
		sc.stats.Registers.TotalTime += value
		sc.stats.Registers.TotalCars += 1
		if value > sc.stats.Registers.MaxTime {
			sc.stats.Registers.MaxTime = value
		}
	}
}

func recalculateAvgStats() {
	if sc.stats.Stations.Gas.TotalCars > 0 {
		sc.stats.Stations.Gas.AvgTime = time.Duration(float64(sc.stats.Stations.Gas.TotalTime.Milliseconds())/float64(sc.stats.Stations.Gas.TotalCars)) * time.Millisecond
	}

	if sc.stats.Stations.Diesel.TotalCars > 0 {
		sc.stats.Stations.Diesel.AvgTime = time.Duration(float64(sc.stats.Stations.Diesel.TotalTime.Milliseconds())/float64(sc.stats.Stations.Diesel.TotalCars)) * time.Millisecond
	}

	if sc.stats.Stations.LPG.TotalCars > 0 {
		sc.stats.Stations.LPG.AvgTime = time.Duration(float64(sc.stats.Stations.LPG.TotalTime.Milliseconds())/float64(sc.stats.Stations.LPG.TotalCars)) * time.Millisecond
	}

	if sc.stats.Stations.Electric.TotalCars > 0 {
		sc.stats.Stations.Electric.AvgTime = time.Duration(float64(sc.stats.Stations.Electric.TotalTime.Milliseconds())/float64(sc.stats.Stations.Electric.TotalCars)) * time.Millisecond
	}

	if sc.stats.Registers.TotalCars > 0 {
		sc.stats.Registers.AvgTime = time.Duration(float64(sc.stats.Registers.TotalTime.Milliseconds())/float64(sc.stats.Registers.TotalCars)) * time.Millisecond
	}
}

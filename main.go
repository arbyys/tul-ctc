package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var sc statsContainer = statsContainer{
	stats: outputFileStruct{},
}

func main() {
	fmt.Printf("\nWelcome to %s! Simulation is now running.", appName)

	conf, err := loadConfig(inputFilePath)
	if err != nil {
		printErrorAndDie(err)
	}

	// app logic

	cashRegisters := make([]*dispatcher, conf.Registers.Count)
	for i := range cashRegisters {
		cashRegisters[i] = &dispatcher{
			queue:    make(chan car, conf.Registers.QueueLengthMax),
			occupied: false,
			timeMin:  conf.Registers.HandleTimeMin,
			timeMax:  conf.Registers.HandleTimeMax,
		}
		go cashRegisters[i].cashRegisterProcess()
	}

	fuelTypeConfs := map[int]stationTypeRaw{
		gas:      stationTypeRaw(conf.Stations.Gas),
		diesel:   stationTypeRaw(conf.Stations.Diesel),
		LPG:      stationTypeRaw(conf.Stations.LPG),
		electric: stationTypeRaw(conf.Stations.Electric),
	}

	allFuelStands := make(map[int][]*dispatcher)

	for k, v := range fuelTypeConfs {
		allFuelStands[k] = make([]*dispatcher, v.Count)

		for i := range v.Count {
			allFuelStands[k][i] = &dispatcher{
				queue:    make(chan car, v.QueueLengthMax),
				occupied: false,
				timeMin:  v.ServeTimeMin,
				timeMax:  v.ServeTimeMax,
			}
			go allFuelStands[k][i].fuelStandProcess(cashRegisters)
		}
	}

	sq := sharedQ{make(chan car, conf.Cars.SharedQueueLengthMax)}
	go sq.sharedQueueProcess(allFuelStands)

	wg.Add(conf.Cars.Count)
	go generateNewCars(conf.Cars.Count, conf.Cars.ArrivalTimeMin, conf.Cars.ArrivalTimeMax, &sq)
	wg.Wait()

	recalculateAvgStats()

	err = exportStats(outputFilePath)
	if err != nil {
		printErrorAndDie(err)
	}

	fmt.Printf("\n\nAll cars gone. Following results have been exported to file %s.\n", outputFilePath)

	err = printStats()
	if err != nil {
		printErrorAndDie(err)
	}
}

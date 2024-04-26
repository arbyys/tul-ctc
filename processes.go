package main

import "time"

func (d *dispatcher) cashRegisterProcess() {
	for {
		c := <-d.queue
		_ = c
		d.occupied = true
		sleepRandomTime(d.timeMin, d.timeMax)

		currentTime := time.Now()
		go addNewTimeStat(register, currentTime.Sub(c.leftStandAt))
		go addNewTimeStat(c.fuel, c.leftStandAt.Sub(c.leftSharedQAt))
		d.occupied = false
		wg.Done()
	}
}

func (d *dispatcher) fuelStandProcess(registers []*dispatcher) {
	for {
		c := <-d.queue
		d.occupied = true
		defer func() {
			d.occupied = false
		}()

		sleepRandomTime(d.timeMin, d.timeMax)
		shortestQueueIndex := getShortestQueue(registers)
		c.leftStandAt = time.Now()

		registers[shortestQueueIndex].queue <- c
	}
}

func (sq *sharedQ) sharedQueueProcess(allFuelStands map[int][]*dispatcher) {
	for {
		c := <-sq.queue

		matchingFuelStands := allFuelStands[c.fuel]
		shortestQueueIndex := getShortestQueue(matchingFuelStands)
		c.leftSharedQAt = time.Now()

		matchingFuelStands[shortestQueueIndex].queue <- c
	}
}

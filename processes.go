package main

func (d *dispatcher) cashRegisterProcess() {
	for {
		c := <-d.queue
		_ = c
		d.occupied = true
		sleepRandomTime(d.timeMin, d.timeMax)
		//c.departureStarted = time.Now()
		//go updateStats(c)
		d.occupied = false
		wg.Done()
	}
}

func (d *dispatcher) fuelStandProcess(registers []*dispatcher) {
	for {
		c := <-d.queue
		d.occupied = true

		sleepRandomTime(d.timeMin, d.timeMax)
		shortestQueueIndex := getShortestQueue(registers)
		//c.waitForRegisterStarted = time.Now()

		registers[shortestQueueIndex].queue <- c
		d.occupied = false
	}
}

func (sq *sharedQ) sharedQueueProcess(allFuelStands map[int][]*dispatcher) {
	for {
		c := <-sq.queue
		matchingFuelStands := allFuelStands[c.fuel]
		shortestQueueIndex := getShortestQueue(matchingFuelStands)
		//c.waitForStandStarted = time.Now()
		matchingFuelStands[shortestQueueIndex].queue <- c
	}
}

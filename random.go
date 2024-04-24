package main

import (
	"math/rand"
	"time"
)

func getRandomFuel() int {
	return rand.Intn(4) + 1
}

func sleepRandomTime(timeMin, timeMax int) {
	generatedTime := timeMin + rand.Intn(timeMax-timeMin+1)

	time.Sleep(time.Duration(generatedTime) * time.Millisecond)
}

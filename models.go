package main

import (
	"time"
)

const (
	gas      = 1
	diesel   = 2
	LPG      = 3
	electric = 4
)

type car struct {
	fuel int
}

type sharedQ struct {
	queue chan car
}

// dispatcher is generic type for both stand and register
type dispatcher struct {
	queue    chan car
	occupied bool
	timeMin  int
	timeMax  int
}

package main

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

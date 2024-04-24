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

// config file structure:
type configFile struct {
	Cars      carConfig      `yaml:"cars"`
	Stations  stationConfig  `yaml:"stations"`
	Registers registerConfig `yaml:"registers"`
}

type carConfig struct {
	Count                int `yaml:"count"`
	ArrivalTimeMin       int `yaml:"arrival_time_min"`
	ArrivalTimeMax       int `yaml:"arrival_time_max"`
	SharedQueueLengthMax int `yaml:"shared_queue_length_max"`
}

type stationConfig struct {
	Gas      stationType `yaml:"gas"`
	Diesel   stationType `yaml:"diesel"`
	LPG      stationType `yaml:"lpg"`
	Electric stationType `yaml:"electric"`
}

type stationType struct {
	Count          int `yaml:"count"`
	ServeTimeMin   int `yaml:"serve_time_min"`
	ServeTimeMax   int `yaml:"serve_time_max"`
	QueueLengthMax int `yaml:"queue_length_max"`
}

type registerConfig struct {
	Count          int `yaml:"count"`
	HandleTimeMin  int `yaml:"handle_time_min"`
	HandleTimeMax  int `yaml:"handle_time_max"`
	QueueLengthMax int `yaml:"queue_length_max"`
}

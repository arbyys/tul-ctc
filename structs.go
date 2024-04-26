package main

import (
	"sync"
	"time"
)

type shortestQueueWorker struct {
	count int
	index int
}

type car struct {
	fuel          int
	leftSharedQAt time.Time
	leftStandAt   time.Time
}

type sharedQ struct {
	queue chan car
}

// dispatcher is generic type for both stand and cash register
type dispatcher struct {
	queue    chan car
	occupied bool
	timeMin  int
	timeMax  int
}

type stationTypeRaw struct {
	Count          int
	ServeTimeMin   int
	ServeTimeMax   int
	QueueLengthMax int
}

type statsContainer struct {
	mu    sync.Mutex
	stats outputFileStruct
	//statsTotal statsAllTypes
}

type statTimeRecord struct {
	detailedType statsAllTypes
	value        time.Duration
}

// config file structure:
type configFileStruct struct {
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

type statsAllTypes struct {
	registers time.Duration
	gas       time.Duration
	diesel    time.Duration
	LPG       time.Duration
	electric  time.Duration
}

// output file structure:
type outputFileStruct struct {
	Stations struct {
		Gas struct {
			TotalCars int           `yaml:"total_cars"`
			TotalTime time.Duration `yaml:"total_time"`
			AvgTime   time.Duration `yaml:"avg_time"`
			MaxTime   time.Duration `yaml:"max_time"`
		}
		Diesel struct {
			TotalCars int           `yaml:"total_cars"`
			TotalTime time.Duration `yaml:"total_time"`
			AvgTime   time.Duration `yaml:"avg_time"`
			MaxTime   time.Duration `yaml:"max_time"`
		}
		LPG struct {
			TotalCars int           `yaml:"total_cars"`
			TotalTime time.Duration `yaml:"total_time"`
			AvgTime   time.Duration `yaml:"avg_time"`
			MaxTime   time.Duration `yaml:"max_time"`
		}
		Electric struct {
			TotalCars int           `yaml:"total_cars"`
			TotalTime time.Duration `yaml:"total_time"`
			AvgTime   time.Duration `yaml:"avg_time"`
			MaxTime   time.Duration `yaml:"max_time"`
		}
	}
	Registers struct {
		TotalCars int           `yaml:"total_cars"`
		TotalTime time.Duration `yaml:"total_time"`
		AvgTime   time.Duration `yaml:"avg_time"`
		MaxTime   time.Duration `yaml:"max_time"`
	}
}

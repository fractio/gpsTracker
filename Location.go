package main

import "time"

// Location stores the GPS coords
type Location struct {
	StartTime          time.Time `json:"startTime" boltholdIndex:"StartTime"`
	ClientTimeStamp    time.Time `json:"clientTimeStamp" boltholdIndex:"ClientTimeStamp"`
	ServerTimeStamp    time.Time `json:"serverTimeStamp" boltholdIndex:"ServerTimeStamp"`
	Accuracy           float64   `json:"accuracy"`
	Lat                float64   `json:"lat"`
	Lng                float64   `json:"lng"`
	Altitude           float64   `json:"altitude"`
	Speed              float64   `json:"speed"`
	Serial             string    `json:"serial" boltholdIndex:"Serial"`
	NumberOfSatellites int       `json:"numberOfSatellites"`
	Direction          float64   `json:"direction"`
	Provider           string    `json:"provider"`
}

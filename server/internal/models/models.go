package models

import "time"

type PostReading struct {
	Timestamp string  `json:"timestamp,omitempty"`
	Pm10      float64 `json:"pm10,omitempty"`
}

type Reading struct {
	Time time.Time
	Pm10 float64
}

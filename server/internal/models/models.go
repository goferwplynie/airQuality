package models

import "time"

type PostReading struct {
	Timestamp   string  `json:"timestamp,omitempty" validate:"required"`
	Temperature float64 `json:"temperature,omitempty" validate:"required,gte=-50,lte=60"`
	Pressure    float64 `json:"pressure,omitempty" validate:"required,gte=870,lte=1100"`
	Humidity    int     `json:"humidity,omitempty" validate:"required,gte=0,lte=100"`
}

type Reading struct {
	Time        time.Time
	Temperature float64
	Pressure    float64
	Humidity    int
}

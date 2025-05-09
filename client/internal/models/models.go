package models

type ResponseModel struct {
	Hourly struct {
		Time        []string  `json:"time,omitempty"`
		Temperature []float64 `json:"temperature_2m,omitempty"`
		Pressure    []float64 `json:"surface_pressure,omitempty"`
		Humidity    []int     `json:"relative_humidity_2m,omitempty"`
	} `json:"hourly"`
}

type Reading struct {
	Timestamp   string  `json:"timestamp,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Pressure    float64 `json:"pressure,omitempty"`
	Humidity    int     `json:"humidity,omitempty"`
}

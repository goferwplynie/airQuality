package models

type ResponseModel struct {
	Hourly struct {
		Time []string  `json:"time,omitempty"`
		Pm10 []float64 `json:"pm10,omitempty"`
	} `json:"hourly"`
}

type Reading struct {
	Timestamp string
	Pm10      float64
}

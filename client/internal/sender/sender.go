package sender

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/goferwplynie/airQuality/client/internal/models"
)

func prepareData(data models.ResponseModel) []models.Reading {
	readings := make([]models.Reading, 0, len(data.Hourly.Time))
	for i := range len(data.Hourly.Time) {
		readings = append(readings, models.Reading{
			Timestamp: data.Hourly.Time[i],
			Pm10:      data.Hourly.Pm10[i],
		})
	}
	return readings
}

func Send(data models.ResponseModel) error {
	readings := prepareData(data)

	json, err := json.Marshal(readings)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(json)

	_, err = http.Post("http://localhost:8080/readings", "application/json", buff)

	if err != nil {
		return err
	}

	return nil
}

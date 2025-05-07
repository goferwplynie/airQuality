package sender

import "github.com/goferwplynie/airQuality/client/internal/models"

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

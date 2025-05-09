package services

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/goferwplynie/airQuality/server/internal/models"
)

type RepositoryInterface interface {
	SaveReading(models.Reading)
	GetAll() []models.Reading
}

type BuisnessLayer struct {
	repo RepositoryInterface
}

func New(repo RepositoryInterface) *BuisnessLayer {
	return &BuisnessLayer{
		repo: repo,
	}
}

func (b *BuisnessLayer) SaveReadings(readings []models.PostReading) []error {
	var errors []error
	validate := validator.New(validator.WithRequiredStructEnabled())

	for _, v := range readings {
		err := validate.Struct(v)
		if err == nil {
			convertedTime, err := convertTimestamp(v.Timestamp)
			if err != nil {
				errors = append(errors, err)
			} else {
				b.repo.SaveReading(models.Reading{
					Time:        convertedTime,
					Temperature: v.Temperature,
					Pressure:    v.Pressure,
					Humidity:    v.Humidity,
				})
			}

		} else {
			errors = append(errors, err)
		}
	}

	return errors
}

func (b *BuisnessLayer) GetReading(timestamp string) (models.Reading, error) {
	var reading models.Reading

	readings := b.repo.GetAll()

	convertedTime, err := convertTimestamp(timestamp)
	if err != nil {
		return reading, err
	}
	reading = getClosestReading(readings, convertedTime)

	return reading, nil
}

func getClosestReading(readings []models.Reading, target time.Time) models.Reading {
	var closest models.Reading
	closestDiff := time.Duration(10000000)

	for _, r := range readings {
		diff := r.Time.Sub(target)
		if diff < 0 {
			diff = -diff
		}
		if diff < closestDiff {
			closestDiff = diff
			closest = r
		}

	}
	return closest
}

func convertTimestamp(timestamp string) (time.Time, error) {
	var convertedTime time.Time
	layout := "2006-01-02T15:04"
	convertedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return convertedTime, err
	}
	return convertedTime, nil
}

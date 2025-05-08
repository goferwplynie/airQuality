package services

import (
	"time"

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

func (b *BuisnessLayer) SaveReadings(readings []models.PostReading) error {
	for _, v := range readings {
		convertedTime, err := convertTimestamp(v.Timestamp)
		if err != nil {
			return err
		}
		b.repo.SaveReading(models.Reading{
			Time: convertedTime,
		})
	}

	return nil
}

func convertTimestamp(timestamp string) (time.Time, error) {
	var convertedTime time.Time
	layout := "2006-01-02T15:04:05.000Z"
	convertedTime, err := time.Parse(layout, timestamp)
	if err != nil {
		return convertedTime, err
	}
	return convertedTime, nil
}

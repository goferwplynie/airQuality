package repository

import "github.com/goferwplynie/airQuality/server/internal/models"

type Repository struct {
	Readings []models.Reading
}

func New() *Repository {
	return &Repository{
		Readings: make([]models.Reading, 0),
	}
}

func (r *Repository) SaveReading(reading models.Reading) {
	r.Readings = append(r.Readings, reading)
}

func (r *Repository) GetAll() []models.Reading {
	return r.Readings
}

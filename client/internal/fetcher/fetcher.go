package fetcher

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goferwplynie/airQuality/client/internal/models"
)

func Fetch(latitude, longtitude float64) (models.ResponseModel, error) {
	url := fmt.Sprintf("https://air-quality-api.open-meteo.com/v1/air-quality?latitude=%v&longitude=%v&hourly=pm10", latitude, longtitude)
	resp, err := http.Get(url)

	var model models.ResponseModel

	if err != nil {
		return model, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&model)

	if err != nil {
		return model, err
	}

	return model, nil
}

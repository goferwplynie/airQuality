package main

import (
	"os"
	"strconv"

	"github.com/goferwplynie/airQuality/client/internal/fetcher"
	"github.com/goferwplynie/airQuality/client/internal/models"
	"github.com/goferwplynie/airQuality/client/internal/sender"
)

func main() {
	var resp models.ResponseModel
	var err error

	args := os.Args
	if len(args) > 2 {
		latitude, err := strconv.ParseFloat(args[1], 64)
		longtitude, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			panic(err)
		}
		resp, err = fetcher.Fetch(latitude, longtitude)
	} else {
		resp, err = fetcher.Fetch(52.23, 21.00)

	}

	if err != nil {
		panic(err)
	}

	err = sender.Send(resp)

	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"

	"github.com/goferwplynie/airQuality/client/internal/fetcher"
)

func main() {
	resp, err := fetcher.Fetch(52.23, 21.00)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

}

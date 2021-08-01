package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BTS struct {
	Ticker struct {
		Price float64 `json:"price,string"`
	} `json:"ticker"`
}

func BitCoinToUAHRate(btsResource string) (float64, error) {
	r, err := http.Get(btsResource)
	if err != nil {
		return -1, err
	}
	r.Header.Set("Accept", "application/json")
	if err != nil {
		return -1, err
	}
	var data BTS
	if err2 := json.NewDecoder(r.Body).Decode(&data); err2 != nil {
		return 0, err2
	}
	fmt.Println(data.Ticker.Price)
	return data.Ticker.Price, nil
}

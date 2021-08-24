package bitcoin

import (
	"encoding/json"
	"net/http"
)

type Convertor interface {
	ToUAH(btsResource string) (float64, error)
}

type BTS struct {
	Ticker struct {
		Price float64 `json:"price,string"`
	} `json:"ticker"`
}

type Conversion struct {
}

func (Conversion) ToUAH(btsResource string) (float64, error) {
	r, err := http.Get(btsResource)
	if err != nil {
		return -1, err
	}
	var data BTS
	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		return -1, err
	}
	return data.Ticker.Price, nil
}

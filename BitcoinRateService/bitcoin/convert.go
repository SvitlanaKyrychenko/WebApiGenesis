package bitcoin

import (
	grpcModel "WebApiGenesis/GRPCMessage/model"
	"context"
	"encoding/json"
	"net/http"
)

type ConversionCryptonator struct {
	grpcModel.UnimplementedConvertorServer
}

func (convertor *ConversionCryptonator) MustEmbedUnimplementedConvertorServer() {

}

func (convertor *ConversionCryptonator) ToUAH(context context.Context, input *grpcModel.BtsRateInput) (*grpcModel.BtsRateResponse, error) {
	r, err := http.Get("https://api.cryptonator.com/api/ticker/btc-uah")
	if err != nil {
		return &grpcModel.BtsRateResponse{Rate: -1.0}, err
	}
	var data BTS
	if err = json.NewDecoder(r.Body).Decode(&data); err != nil {
		return &grpcModel.BtsRateResponse{Rate: -1.0}, err
	}
	return &grpcModel.BtsRateResponse{Rate: float32(data.Ticker.Price)}, nil
}

type BTS struct {
	Ticker struct {
		Price float64 `json:"price,string"`
	} `json:"ticker"`
}

package app

import (
	"WebApiGenesis/bitcoin"
	"fmt"
	"net/http"
)

func BTSRateHandler(response http.ResponseWriter, request *http.Request) {
	var btsConvertor bitcoin.Convertor = bitcoin.Conversion{}
	if btsRate, err := btsConvertor.ToUAH("https://api.cryptonator.com/api/ticker/btc-uah"); err == nil {
		fmt.Fprintf(response, "BTS: 1\nUAH: "+fmt.Sprintf("%.2f", btsRate))
	}
}

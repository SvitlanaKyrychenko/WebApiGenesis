package main

import (
	. "WebApiGenesis/app"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", AuthenticationHandler)

	http.HandleFunc("/btsRate", func(response http.ResponseWriter, request *http.Request) {
		if btsRate, err := BitCoinToUAHRate("https://api.cryptonator.com/api/ticker/btc-uah"); err == nil {
			fmt.Fprintf(response, "BTS: 1\nUAH: " + fmt.Sprintf("%.2f", btsRate))
		}
	})

	http.HandleFunc("/user/create", RegistrationHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

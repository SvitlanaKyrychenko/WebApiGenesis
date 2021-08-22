package main

import (
	. "WebApiGenesis/app"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", AuthenticationHandler)
	http.HandleFunc("/btsRate", BTSRateHandler)

	http.HandleFunc("/user/create", RegistrationHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

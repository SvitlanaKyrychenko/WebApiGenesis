package main

import (
	"WebApiGenesis/app"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", app.AuthenticationHandler)
	http.HandleFunc("/btsRate", app.BTSRateHandler)
	http.HandleFunc("/user/create", app.RegistrationHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)


}

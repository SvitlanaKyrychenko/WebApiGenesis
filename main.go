package main

import (
	"WebApiGenesis/app"
	"WebApiGenesis/model"
	"WebApiGenesis/services"
	"WebApiGenesis/storage"
	"fmt"
	"net/http"
)

func main() {
	//prepare
	var convertor model.Convertor = model.JSONGConvertor{}
	var storage storage.Storage = storage.FileStorage{Convertor: convertor}
	var authService services.Authenticator = services.Authentication{Storage: storage}
	var regService services.Registrar = services.Registration{Storage: storage}
	var authentication app.Authentication = app.Authentication{AuthService: authService}
	var registration app.Registration = app.Registration{RegService: regService}

	http.HandleFunc("/user/login", authentication.AuthenticationHandler)
	http.HandleFunc("/btsRate", app.BTSRateHandler)
	http.HandleFunc("/user/create", registration.RegistrationHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}

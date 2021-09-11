package main

import (
	"WebApiGenesis/Api/app"
	"flag"
	"fmt"
	"net/http"
)

var (
	serverAddrAuth    = flag.String("server_addr_auth", "localhost:10003", "The server address in the format of host:port")
	serverAddrReg     = flag.String("server_addr_reg", "localhost:10004", "The server address in the format of host:port")
	serverAddrBtsRate = flag.String("server_addr_bts_rate", "localhost:10005", "The server address in the format of host:port")
)

func main() {
	//prepare
	var authentication app.Authentication = app.Authentication{AuthServer: serverAddrAuth}
	var registration app.Registration = app.Registration{RegServer: serverAddrReg}
	var btsRate app.BtsRate = app.BtsRate{BtsRateServer: serverAddrBtsRate}

	http.HandleFunc("/user/login", authentication.AuthenticationHandler)
	http.HandleFunc("/btsRate", btsRate.BTSRateHandler)
	http.HandleFunc("/user/create", registration.RegistrationHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)

}

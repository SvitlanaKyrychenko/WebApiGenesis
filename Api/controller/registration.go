package controller

import (
	grpcModel "Bitcoin/GRPCMessage"
	"context"
	"flag"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
)

type Registration struct {
	RegServer *string
}

func (reg Registration) RegistrationHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		postRegistrationLogic(response, request, reg)
	}
	if request.Method == "GET" {
		tmpl, _ := template.ParseFiles("../WebApiGenesis/Api/html/registration.html")
		tmpl.Execute(response, ViewData{""})
	}
}

func postRegistrationLogic(response http.ResponseWriter, request *http.Request, reg Registration) {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}
	if request.FormValue("signUp") == "Sign up" {
		singUpLogic(response, request, reg)
	} else if request.FormValue("goToSignIn") == "Go to sign in" {
		http.Redirect(response, request, "/user/login", 301)
	}
}

func singUpLogic(response http.ResponseWriter, request *http.Request, reg Registration) {
	var registrationUser grpcModel.RegistrationUser = createRegistrationUser(request)
	message, err := createRegConnection(reg.RegServer, registrationUser)
	if err != nil {
		log.Println(err)
		return
	}
	var viewData ViewData = ViewData{message}
	if viewData.Message == "" {
		http.Redirect(response, request, "/user/login", 301)
	}
	tmpl, _ := template.ParseFiles("../WebApiGenesis/Api/html/registration.html")
	tmpl.Execute(response, viewData)
}

func createRegistrationUser(request *http.Request) grpcModel.RegistrationUser {
	var registrationUser grpcModel.RegistrationUser
	registrationUser.Email = request.FormValue("email")
	registrationUser.Password = request.FormValue("password")
	registrationUser.ConfirmPassword = request.FormValue("confirmPassword")
	return registrationUser
}

func createRegConnection(serverAddr *string, registrationUser grpcModel.RegistrationUser) (string, error) {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := grpcModel.NewRegistrarClient(conn)
	authMessage, err := client.Register(context.Background(), &registrationUser)
	if err != nil {
		return "", err
	}
	return authMessage.Message, nil
}

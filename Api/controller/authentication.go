package controller

import (
	grpcModel "Bitcoin/GRPCMessage"
	"context"
	"flag"
	"fmt"
	grpc "google.golang.org/grpc"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Message string
}

type Authentication struct {
	AuthServer *string
}

func (auth Authentication) AuthenticationHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		postAuthenticationLogic(response, request, auth.AuthServer)
	}
	if request.Method == "GET" {
		tmpl, _ := template.ParseFiles("../WebApiGenesis/Api/html/authentication.html")
		tmpl.Execute(response, ViewData{""})
	}
}

func postAuthenticationLogic(response http.ResponseWriter, request *http.Request, authServer *string) {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}
	if request.FormValue("signIn") == "Sign in" {
		signInLogic(response, request, authServer)
	}
	if request.FormValue("signUp") == "Sign up" {
		http.Redirect(response, request, "/user/create", 301)
	}
}

func signInLogic(response http.ResponseWriter, request *http.Request, authServer *string) {
	var authenticationUser grpcModel.AuthenticationUser = createAuthenticationUser(request)
	message, err := createAuthConnection(authServer, authenticationUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	var viewData ViewData = ViewData{message}
	if viewData.Message == "" {
		http.Redirect(response, request, "/btsRate", 301)
	}
	tmpl, _ := template.ParseFiles("../WebApiGenesis/Api/html/authentication.html")
	tmpl.Execute(response, viewData)
}

func createAuthenticationUser(request *http.Request) grpcModel.AuthenticationUser {
	var authenticationUser grpcModel.AuthenticationUser
	authenticationUser.Password = request.FormValue("password")
	authenticationUser.Email = request.FormValue("email")
	return authenticationUser
}

func createAuthConnection(serverAddr *string, authenticationUser grpcModel.AuthenticationUser) (string, error) {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := grpcModel.NewAuthenticatorClient(conn)
	authMessage, err := client.Authenticate(context.Background(), &authenticationUser)
	if err != nil {
		return "", err
	}
	return authMessage.Message, nil
}

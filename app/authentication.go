package app

import (
	"WebApiGenesis/model"
	"WebApiGenesis/services"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Message string
}

func AuthenticationHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		postAuthenticationLogic(response, request)
	}
	if request.Method == "GET" {
		tmpl, _ := template.ParseFiles("../WebApiGenesis/html/authentication.html")
		tmpl.Execute(response, ViewData{""})
	}
}

func postAuthenticationLogic(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}
	if request.FormValue("signIn") == "Sign in" {
		signInLogic(response, request)
	}
	if request.FormValue("signUp") == "Sign up" {
		http.Redirect(response, request, "/user/create", 301)
	}
}

func signInLogic(response http.ResponseWriter, request *http.Request) {
	var convertor model.Convertor = model.JSONGConvertor{}
	var authService services.Authenticator = services.Authentication{Convertor: convertor}
	var authenticationUser model.AuthenticationUser = createAuthenticationUser(request)
	message, err := authService.Authenticate(authenticationUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	var viewData ViewData = ViewData{message}
	if viewData.Message == "" {
		http.Redirect(response, request, "/btsRate", 301)
	}
	tmpl, _ := template.ParseFiles("../WebApiGenesis/html/authentication.html")
	tmpl.Execute(response, viewData)
}

func createAuthenticationUser(request *http.Request) model.AuthenticationUser {
	var authenticationUser model.AuthenticationUser
	authenticationUser.Password = request.FormValue("password")
	authenticationUser.Email = request.FormValue("email")
	return authenticationUser
}



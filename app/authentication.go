package app

import (
	. "WebApiGenesis/model"
	. "WebApiGenesis/services"
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
	var viewData ViewData = ViewData{""}
	var authenticationUser AuthenticationUser = createAuthenticationUser(request)
	if valid, message := singIn(authenticationUser); !valid {
		viewData.Message += message.Message
	} else {
		http.Redirect(response, request, "/btsRate", 301)
	}
	tmpl, _ := template.ParseFiles("../WebApiGenesis/html/authentication.html")
	tmpl.Execute(response, viewData)
}

func createAuthenticationUser(request *http.Request) AuthenticationUser {
	var authenticationUser AuthenticationUser
	authenticationUser.Password = request.FormValue("password")
	authenticationUser.Email = request.FormValue("email")
	return authenticationUser
}

func singIn(authenticationUser AuthenticationUser) (bool, ViewData) {
	var viewData ViewData = ViewData{""}
	if authenticationUser.Email == "" || authenticationUser.Password == "" {
		viewData.Message += "Password and email must ve set. "
	}
	if valid, message := Authenticate(authenticationUser); !valid {
		viewData.Message += message
	}
	return viewData.Message == "", viewData
}

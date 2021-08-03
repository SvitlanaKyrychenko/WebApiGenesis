package app

import (
	. "WebApiGenesis/model"
	. "WebApiGenesis/services"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Message string
}

func AuthenticationHandler(response http.ResponseWriter, request *http.Request) {
	var viewData ViewData = ViewData{""}

	if request.Method == "GET" {
		err := request.ParseForm()
		if err != nil {
			log.Println(err)
		}
		viewData = getLogic(response, request)
		tmpl, _ := template.ParseFiles("../WebApiGenesis/html/authentication.html")
		tmpl.Execute(response, viewData)
	}
}

func getLogic(response http.ResponseWriter, request *http.Request) ViewData {
	var viewData ViewData = ViewData{""}
	var authenticationUser AuthenticationUser = createAuthenticationUser(request)

	if request.FormValue("signIn") == "Sign in" {
		if valid, message := singIn(authenticationUser); !valid {
			viewData.Message += message.Message
		} else {
			http.Redirect(response, request, "/btsRate", 301)
		}
	}
	if request.FormValue("signUp") == "Sign up" {
		http.Redirect(response, request, "/user/create", 301)
	}
	return viewData
}

func createAuthenticationUser(request *http.Request) AuthenticationUser {
	var authenticationUser AuthenticationUser
	if request.FormValue("password") != "" {
		hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), bcrypt.DefaultCost)
		if errHash != nil {
			fmt.Println(errHash)
		} else {
			authenticationUser.Password = string(hashedPassword)
		}
	}
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

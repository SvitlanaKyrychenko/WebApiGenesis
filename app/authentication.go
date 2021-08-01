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
	var authenticationUser AuthenticationUser
	authenticationUser.Email = request.FormValue("email")
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), bcrypt.DefaultCost)
	if errHash != nil {
		fmt.Println(errHash)
		return
	}
	authenticationUser.Password = string(hashedPassword)
	if request.Method == "GET" {
		err := request.ParseForm()
		if err != nil {
			log.Println(err)
		}
		signIn := request.FormValue("signIn")
		signUp := request.FormValue("signUp")
		if signIn == "Sign in" {
			if valid, message := singIn(authenticationUser); !valid {
				viewData = ViewData{message}
			} else {
				http.Redirect(response, request, "/btsRate", 301)
			}
		} else if signUp == "Sign up" {
			http.Redirect(response, request, "/user/create", 301)
			return
		}
		tmpl, _ := template.ParseFiles("../WebApiGenesis/html/authentication.html")
		tmpl.Execute(response, viewData)
	}
}

func singIn(authenticationUser AuthenticationUser) (bool, string) {
	if authenticationUser.Email == "" || authenticationUser.Password == "" {
		return false, "password and email must ve set"
	} else {
		if valid, message := Authenticate(authenticationUser); !valid {
			return false, message
		} else {
			return true, ""
		}
	}
}

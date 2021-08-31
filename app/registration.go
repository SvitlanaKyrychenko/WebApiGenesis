package app

import (
	"WebApiGenesis/model"
	"WebApiGenesis/services"
	"html/template"
	"log"
	"net/http"
)

type Registration struct {
	RegService services.Registrar
}

func (reg Registration) RegistrationHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		postRegistrationLogic(response, request, reg)
	}
	if request.Method == "GET" {
		tmpl, _ := template.ParseFiles("../WebApiGenesis/html/registration.html")
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
	var registrationUser model.RegistrationUser = createRegistrationUser(request)
	message, err := reg.RegService.Register(registrationUser)
	if err != nil {
		log.Println(err)
		return
	}
	var viewData ViewData = ViewData{message}
	if viewData.Message == "" {
		http.Redirect(response, request, "/user/login", 301)
	}
	tmpl, _ := template.ParseFiles("../WebApiGenesis/html/registration.html")
	tmpl.Execute(response, viewData)
}

func createRegistrationUser(request *http.Request) model.RegistrationUser {
	var registrationUser model.RegistrationUser
	registrationUser.Email = request.FormValue("email")
	registrationUser.Password = request.FormValue("password")
	registrationUser.ConfirmPassword = request.FormValue("confirmPassword")
	return registrationUser
}

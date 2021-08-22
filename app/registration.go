package app

import (
	. "WebApiGenesis/model"
	. "WebApiGenesis/services"
	"html/template"
	"log"
	"net/http"
	"net/mail"
)

func RegistrationHandler(response http.ResponseWriter, request *http.Request) {

	if request.Method == "POST" {
		postRegistrationLogic(response, request)
	}
	if request.Method == "GET" {
		tmpl, _ := template.ParseFiles("../WebApiGenesis/html/registration.html")
		tmpl.Execute(response,  ViewData{""})
	}
}

func postRegistrationLogic(response http.ResponseWriter, request *http.Request)  {
	err := request.ParseForm()
	if err != nil {
		log.Println(err)
	}
	if request.FormValue("signUp") == "Sign up" {
		singUpLogic(response, request)
	} else if request.FormValue("goToSignIn") == "Go to sign in" {
		http.ServeFile(response, request, "../WebApiGenesis/html/authentication.html")
		return
	}
}

func singUpLogic(response http.ResponseWriter, request *http.Request)  {
	var viewData ViewData = ViewData{""}
	var registrationUser RegistrationUser = createRegistrationUser(request)
	if valid, message := trySignUp(registrationUser); !valid {
		viewData = ViewData{message}
	} else {
		http.Redirect(response, request, "/user/login", 301)
	}
	tmpl, _ := template.ParseFiles("../WebApiGenesis/html/registration.html")
	tmpl.Execute(response, viewData)
}

func createRegistrationUser(request *http.Request) RegistrationUser {
	var registrationUser RegistrationUser
	registrationUser.Email = request.FormValue("email")
	registrationUser.Password = request.FormValue("password")
	registrationUser.ConfirmPassword = request.FormValue("confirmPassword")
	return registrationUser
}

func trySignUp(registrationUser RegistrationUser) (bool, string) {
	if registrationUser.Email == "" || registrationUser.Password == "" || registrationUser.ConfirmPassword == "" {
		return false, "Email, password and confirm password must be set"
	}
	if validEmail(registrationUser.Email) {
		return false, "Email is wrong"
	}
	if len(registrationUser.Password) < 6 {
		return false, "Password must be longer than 6"
	}
	if registrationUser.Password != registrationUser.ConfirmPassword {
		return false, "Confirm password and password can`t be different"
	}
	if !Registration(registrationUser) {
		return false, "This user has not registered yet"
	}
	return true, ""
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

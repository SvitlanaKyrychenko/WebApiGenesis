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
	var viewData ViewData = ViewData{""}
	var registrationUser RegistrationUser
	if request.Method == "GET" {
		err := request.ParseForm()
		if err != nil {
			log.Println(err)
		}
		registrationUser.Email = request.FormValue("email")
		registrationUser.Password = request.FormValue("password")
		registrationUser.ConfirmPassword = request.FormValue("confirmPassword")
		goToSignIn := request.FormValue("Go to sign in")
		signUp := request.FormValue("signUp")

		if signUp == "Sign up" {
			if valid, message := trySignUp(registrationUser); !valid{
				viewData = ViewData{message}
			}else {
				http.Redirect(response, request, "/user/login", 301)
			}
		} else if goToSignIn == "Go to sign in" {
			http.ServeFile(response, request, "../WebApiGenesis/html/authentication.html")
			return
		}
		tmpl, _ := template.ParseFiles("../WebApiGenesis/html/registration.html")
		tmpl.Execute(response, viewData)
	}
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

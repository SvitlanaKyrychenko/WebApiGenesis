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
			if registrationUser.Email == "" || registrationUser.Password == "" || registrationUser.ConfirmPassword == "" {
				viewData = ViewData{"Email, password and confirm password must be set"}
			} else if validEmail(registrationUser.Email) {
				viewData = ViewData{"Email os wrong"}
			} else if len(registrationUser.Password)<6 {
				viewData = ViewData{"Password must be longer than 6"}
			}else if registrationUser.Password != registrationUser.ConfirmPassword {
				viewData = ViewData{"Confirm password and password can`t be different"}
			}else{
				if Registration(registrationUser){
					http.Redirect(response, request, "/user/login", 301)
				}
			}
		} else if goToSignIn == "Go to sign in" {
			http.ServeFile(response, request, "../WebApiGenesis/html/authentication.html")
			return
		}
		tmpl, _ := template.ParseFiles("../WebApiGenesis/html/registration.html")
		tmpl.Execute(response, viewData)
	}
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}

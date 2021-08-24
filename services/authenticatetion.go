package services

import (
	"WebApiGenesis/model"
	"WebApiGenesis/storage"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
)

type Authenticator interface {
	Authenticate(authenticationUser model.AuthenticationUser) (string, error)
}

type Authentication struct {
	Convertor model.Convertor
}

func (authentication Authentication) Authenticate(authenticationUser model.AuthenticationUser) (string, error) {
	var message string
	if authenticationUser.Email == "" || authenticationUser.Password == "" {
		message += "Password and email must be set. "
	}
	inStorage, err := checkAuthenticationUserInStorage(authentication, authenticationUser)
	if err != nil {
		return message, err
	}
	message += inStorage
	return message, nil
}

func checkAuthenticationUserInStorage(authentication Authentication, authenticationUser model.AuthenticationUser) (string, error) {
	var storage storage.Storage = storage.FileStorage{Convertor: authentication.Convertor}
	users, err := storage.GetALLAsync(model.DBUser{})
	if err != nil {
		return "", err
	}
	for _, user := range users {
		if compareAuthenticationUsers(authenticationUser, user) {
			return "", nil
		}
	}
	return " Email or password are wrong. ", nil
}

func compareAuthenticationUsers(authenticationUser model.AuthenticationUser, user []byte) bool {
	var userGot model.DBUser
	err := json.Unmarshal(user, &userGot)
	if err != nil {
		return false
	}
	if userGot.Email == authenticationUser.Email &&
		bcrypt.CompareHashAndPassword([]byte(userGot.Password), []byte(authenticationUser.Password)) == nil {
		return true
	}
	return false
}

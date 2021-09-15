package server

import (
	"Bitcoin/CustomerService/model"
	"Bitcoin/CustomerService/storage"
	grpcModel "Bitcoin/GRPCMessage"
	"context"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationServer struct {
	grpcModel.UnimplementedAuthenticatorServer
	Storage storage.Storage
}

func (authentication *AuthenticationServer) MustEmbedUnimplementedRegistrarServer() {

}

func (authentication *AuthenticationServer) Authenticate(context context.Context, authenticationUser *grpcModel.AuthenticationUser) (*grpcModel.AuthenticationResponse, error) {
	var message string
	if authenticationUser.Email == "" || authenticationUser.Password == "" {
		message += "Password and email must be set. "
	}
	inStorage, err := checkAuthenticationUserInStorage(authentication, authenticationUser)
	if err != nil {
		return &grpcModel.AuthenticationResponse{Message: ""}, err
	}
	message += inStorage
	return &grpcModel.AuthenticationResponse{Message: message}, nil
}

func checkAuthenticationUserInStorage(authentication *AuthenticationServer, authenticationUser *grpcModel.AuthenticationUser) (string, error) {
	if authentication.Storage == nil {
		return "", errors.New("storage has not set")
	}
	users, err := authentication.Storage.GetALLAsync(model.DBUser{})
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

func compareAuthenticationUsers(authenticationUser *grpcModel.AuthenticationUser, user []byte) bool {
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

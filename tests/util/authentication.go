package util

import (
	model2 "WebApiGenesis/CustomerService/model"
	"WebApiGenesis/CustomerService/server"
	"WebApiGenesis/CustomerService/storage"
	"WebApiGenesis/GRPCMessage/model"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

func AddUserToStorageMock(authenticationUser model.AuthenticationUser, storage *storage.Storage) {
	guid, err := ksuid.NewRandom()
	if err != nil {
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authenticationUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	var userDB model2.DBUser = model2.DBUser{Email: authenticationUser.Email, Password: string(hashedPassword), Guid: guid}
	(*storage).AddOrUpdateAsync(userDB)
}

func PrepareAuthService() server.Authenticator {
	var storage storage.Storage = PrepareMockStorage()
	return server.AuthenticationServer{Storage: storage}
}

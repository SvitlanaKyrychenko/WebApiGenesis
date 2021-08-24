package util

import (
	"WebApiGenesis/model"
	"WebApiGenesis/services"
	"WebApiGenesis/storage"
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
	var userDB model.DBUser = model.DBUser{Email: authenticationUser.Email, Password: string(hashedPassword), Guid: guid}
	(*storage).AddOrUpdateAsync(userDB)
}

func PrepareAuthService() services.Authenticator {
	var storage storage.Storage = PrepareMockStorage()
	return services.Authentication{Storage: storage}
}

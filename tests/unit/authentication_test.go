package unit

import (
	"WebApiGenesis/model"
	"WebApiGenesis/services"
	"WebApiGenesis/storage"
	"WebApiGenesis/tests/mock"
	"WebApiGenesis/tests/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidAuthentication(t *testing.T) {
	t.Parallel()
	//Arrange
	var authenticationUser model.AuthenticationUser = model.AuthenticationUser{Email: "email@gmail.com", Password: "GreatPassword1"}
	var convertor model.Convertor = model.JSONGConvertor{}
	var mockStorage = mock.Storage{Convertor: convertor}
	mockStorage.Init()
	var storage storage.Storage = mockStorage
	util.AddUserToStorageMock(authenticationUser, &storage)
	var authService services.Authenticator = services.Authentication{Storage: storage}
	//Act
	message, err := authService.Authenticate(authenticationUser)
	//Assert
	require.Nil(t, err)
	require.Empty(t, message)
}

func TestWrongEmailFormat(t *testing.T) {
	t.Parallel()
	//Arrange
	var authenticationUser model.AuthenticationUser = model.AuthenticationUser{Email: "wrongEmail", Password: "password"}
	var authService services.Authenticator = util.PrepareAuthService()
	//Act
	message, err := authService.Authenticate(authenticationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestEmptyAuthenticationUser(t *testing.T) {
	t.Parallel()
	//Arrange
	var authenticationUser model.AuthenticationUser = model.AuthenticationUser{}
	var authService services.Authenticator = util.PrepareAuthService()
	//Act
	message, err := authService.Authenticate(authenticationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestAuthenticationUserNotExists(t *testing.T) {
	t.Parallel()
	//Arrange
	var authenticationUser model.AuthenticationUser = model.AuthenticationUser{Email: "email@gmail.com", Password: "GreatPassword1"}
	var authService services.Authenticator = util.PrepareAuthService()
	//Act
	message, err := authService.Authenticate(authenticationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestWrongPassword(t *testing.T) {
	t.Parallel()
	//Arrange
	var authenticationUserRegistered model.AuthenticationUser = model.AuthenticationUser{Email: "email@gmail.com", Password: "GreatPassword1"}
	var authenticationUserTryEnter model.AuthenticationUser = model.AuthenticationUser{Email: "email@gmail.com", Password: "WrongGreatPassword1"}
	var convertor model.Convertor = model.JSONGConvertor{}
	var mockStorage = mock.Storage{Convertor: convertor}
	mockStorage.Init()
	var storage storage.Storage = mockStorage
	util.AddUserToStorageMock(authenticationUserRegistered, &storage)
	var authService services.Authenticator = services.Authentication{Storage: storage}
	//Act
	message, err := authService.Authenticate(authenticationUserTryEnter)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}
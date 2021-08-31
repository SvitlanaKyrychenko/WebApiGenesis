package integrational

import (
	"WebApiGenesis/model"
	"WebApiGenesis/services"
	"WebApiGenesis/storage"
	"WebApiGenesis/tests/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidRegistrationAuthentication(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		Password:        "GreatPassword1",
		ConfirmPassword: "GreatPassword1"}
	var authenticationUser model.AuthenticationUser = model.AuthenticationUser{Email: "email@gmail.com",
		Password: "GreatPassword1"}
	var storage storage.Storage = util.PrepareMockStorage()
	var regService services.Registrar = services.Registration{Storage: storage}
	var authService services.Authenticator = services.Authentication{Storage: storage}
	//Act
	messageRegistration, errRegistration := regService.Register(registrationUser)
	messageAuthentication, errAuthentication := authService.Authenticate(authenticationUser)
	//Assert
	require.Nil(t, errRegistration)
	require.Nil(t, errAuthentication)
	require.Empty(t, messageRegistration)
	require.Empty(t, messageAuthentication)
}

func TestDifferentRegistrationUserAndAuthenticationUser(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		Password:        "GreatPassword1",
		ConfirmPassword: "GreatPassword1"}
	var authenticationUser model.AuthenticationUser = model.AuthenticationUser{Email: "email@gmail.com",
		Password: "WrongGreatPassword1"}
	var storage storage.Storage = util.PrepareMockStorage()
	var regService services.Registrar = services.Registration{Storage: storage}
	var authService services.Authenticator = services.Authentication{Storage: storage}
	//Act
	messageRegistration, errRegistration := regService.Register(registrationUser)
	messageAuthentication, errAuthentication := authService.Authenticate(authenticationUser)
	//Assert
	require.Nil(t, errRegistration)
	require.Nil(t, errAuthentication)
	require.Empty(t, messageRegistration)
	require.NotEmpty(t, messageAuthentication)
}

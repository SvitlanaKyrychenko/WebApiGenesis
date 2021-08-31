package unit

import (
	"WebApiGenesis/model"
	"WebApiGenesis/services"
	"WebApiGenesis/tests/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		Password:        "GreatPassword1",
		ConfirmPassword: "GreatPassword1"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.Empty(t, message)
}

func TestNilEmailRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{
		Password:        "GreatPassword1",
		ConfirmPassword: "GreatPassword1"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestNilPasswordRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		ConfirmPassword: "GreatPassword1"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestNilConfirmPasswordRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		Password: "GreatPassword1"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestInvalidEmailFormatRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		Password:        "pass",
		ConfirmPassword: "pass"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestShortPasswordRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "invalid",
		Password:        "GreatPassword1",
		ConfirmPassword: "GreatPassword1"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestConfirmPasswordNotEqualPasswordRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		Password:        "GreatPassword1",
		ConfirmPassword: "GreatPassword1"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	message, err = regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

func TestUserHasAlreadyRegisteredRegistration(t *testing.T) {
	t.Parallel()
	//Arrange
	var registrationUser model.RegistrationUser = model.RegistrationUser{Email: "email@gmail.com",
		Password:        "GreatPassword1",
		ConfirmPassword: "GreatPassword2"}
	var regService services.Registrar = util.PrepareRegService()
	//Act
	message, err := regService.Register(registrationUser)
	//Assert
	require.Nil(t, err)
	require.NotEmpty(t, message)
}

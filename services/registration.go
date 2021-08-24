package services

import (
	"WebApiGenesis/model"
	"WebApiGenesis/storage"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/mail"
)

type Registrar interface {
	Register(registrationUser model.RegistrationUser) (string, error)
}

type Registration struct {
	Convertor model.Convertor
}

func (registration Registration) Register(registrationUser model.RegistrationUser) (string, error) {
	if registrationUser.Email == "" || registrationUser.Password == "" || registrationUser.ConfirmPassword == "" {
		return "Email, password and confirm password must be set", nil
	}
	if validEmail(registrationUser.Email) {
		return "Email is wrong", nil
	}
	if len(registrationUser.Password) < 6 {
		return "Password must be longer than 6", nil
	}
	if registrationUser.Password != registrationUser.ConfirmPassword {
		return "Confirm password and password can`t be different", nil
	}
	if checkRegistrationUserInStorage(registration, registrationUser) {
		return "This user has already registered", nil
	}
	err := registration.saveInStorage(registrationUser)
	if err != nil {
		return "", err
	}
	return "", nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}
func (registration Registration) saveInStorage(registrationUser model.RegistrationUser) error {
	var storage storage.Storage = storage.FileStorage{Convertor: registration.Convertor}
	newDBUser, err := createDBUser(registrationUser)
	if err != nil {
		return err
	}
	return storage.AddOrUpdateAsync(newDBUser)
}

func createDBUser(registrationUser model.RegistrationUser) (model.DBUser, error) {
	var userDB model.DBUser
	guid, err := ksuid.NewRandom()
	if err != nil {
		return userDB, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registrationUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return userDB, err
	}
	userDB = model.DBUser{Guid: guid, Password: string(hashedPassword), Email: registrationUser.Email}
	return userDB, nil
}

func checkRegistrationUserInStorage(registration Registration, registrationUser model.RegistrationUser) bool {
	var storage storage.Storage = storage.FileStorage{Convertor: registration.Convertor}
	users, err := storage.GetALLAsync(model.DBUser{})
	if err != nil {
		log.Println(err)
	}
	for _, user := range users {
		if compareRegistrationUsers(registrationUser, user) {
			return true
		}
	}
	return false
}

func compareRegistrationUsers(registrationUser model.RegistrationUser, user []byte) bool {
	var userGot model.DBUser
	err := json.Unmarshal(user, &userGot)
	if err != nil {
		return false
	}
	if userGot.Email == registrationUser.Email &&
		bcrypt.CompareHashAndPassword([]byte(userGot.Password), []byte(registrationUser.Password)) == nil {
		return true
	}
	return false
}

package server

import (
	"WebApiGenesis/CustomerService/model"
	"WebApiGenesis/CustomerService/storage"
	grpcModel "WebApiGenesis/GRPCMessage/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/mail"
)

type Registration struct {
	grpcModel.UnimplementedRegistrarServer
	Storage storage.Storage
}

func (registration *Registration) MustEmbedUnimplementedRegistrarServer() {

}

func (registration *Registration) Register(context context.Context, registrationUser *grpcModel.RegistrationUser) (*grpcModel.RegistrationResponse, error) {
	if registrationUser.Email == "" || registrationUser.Password == "" || registrationUser.ConfirmPassword == "" {
		return &grpcModel.RegistrationResponse{Message: "Email, password and confirm password must be set"}, nil
	}
	if validEmail(registrationUser.Email) {
		return &grpcModel.RegistrationResponse{Message: "Email is wrong"}, nil
	}
	if len(registrationUser.Password) < 6 {
		return &grpcModel.RegistrationResponse{Message: "Password must be longer than 6"}, nil
	}
	if registrationUser.Password != registrationUser.ConfirmPassword {
		return &grpcModel.RegistrationResponse{Message: "Confirm password and password can`t be different"}, nil
	}
	if checkRegistrationUserInStorage(registration, registrationUser) {
		return &grpcModel.RegistrationResponse{Message: "This user has already registered"}, nil
	}
	err := saveInStorage(registrationUser, registration)
	if err != nil {
		return &grpcModel.RegistrationResponse{Message: ""}, err
	}
	return &grpcModel.RegistrationResponse{Message: ""}, nil
}

func validEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err != nil
}
func saveInStorage(registrationUser *grpcModel.RegistrationUser, registration *Registration) error {
	newDBUser, err := createDBUser(registrationUser)
	if err != nil {
		return err
	}
	if registration.Storage == nil {
		return errors.New("storage has not set")
	}
	return registration.Storage.AddOrUpdateAsync(newDBUser)
}

func createDBUser(registrationUser *grpcModel.RegistrationUser) (model.DBUser, error) {
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

func checkRegistrationUserInStorage(registration *Registration, registrationUser *grpcModel.RegistrationUser) bool {
	if registration.Storage == nil {
		return false
	}
	users, err := registration.Storage.GetALLAsync(model.DBUser{})
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

func compareRegistrationUsers(registrationUser *grpcModel.RegistrationUser, user []byte) bool {
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

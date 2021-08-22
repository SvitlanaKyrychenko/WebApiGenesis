package services

import (
	. "WebApiGenesis/dataStorage"
	. "WebApiGenesis/model"
	"fmt"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

func Registration(registrationUser RegistrationUser) bool {
	guid, e := ksuid.NewRandom()
	if e != nil {
		fmt.Println(e)
	}
	hashedPassword, errHash := bcrypt.GenerateFromPassword([]byte(registrationUser.Password), bcrypt.DefaultCost)
	if errHash == nil {
		var newUser DBUser = DBUser{Guid: guid, Password: string(hashedPassword), Email: registrationUser.Email}
		return AddOrUpdateAsync(newUser)
	}
	return false
}

package services

import (
	. "WebApiGenesis/dataStorage"
	. "WebApiGenesis/model"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Authenticate(authenticationUser AuthenticationUser) (bool, string) {
	var class ClassStorable = ClassStorable{NameClass: "DBUser"}
	valid, users := GetALLAsync(class)
	if valid {
		for _, user := range users {
			var userGot DBUser
			if err1 := json.Unmarshal(user, &userGot); err1 != nil {
				fmt.Println(err1)
			} else if userGot.Email == authenticationUser.Email &&
				bcrypt.CompareHashAndPassword([]byte(userGot.Password), []byte(authenticationUser.Password)) == nil {
				return true, ""
			}
		}
	}
	return false, "Email or password are wrong. "
}

package util

import (
	model2 "WebApiGenesis/BitcoinRateService/model"
	model3 "WebApiGenesis/CustomerService/model"
	"WebApiGenesis/CustomerService/storage"
	"WebApiGenesis/CustomerService/utils/file"
	"WebApiGenesis/GRPCMessage/model"
	"WebApiGenesis/tests/mock"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"path/filepath"
)

func PrepareStorage() storage.Storage {
	var convertor model2.Convertor = model2.JSONGConvertor{}
	return storage.FileStorage{Convertor: convertor}
}

func PrepareMockClass(message string) model3.Storable {
	guid, err := ksuid.NewRandom()
	if err != nil {
		return nil
	}
	return mock.StorableClass{Guid: guid, Message: message}
}

func DeleteClass(class model3.Storable) {
	src, err := GetClassPath(class)
	if err == nil {
		file.DeleteFile(src)
	}
}

func DeleteDBUser(registrationUser model.RegistrationUser, storage storage.Storage) {
	var needUser model3.DBUser = model3.DBUser{}
	usersByte, err := storage.GetALLAsync(needUser)
	if err != nil {
		return
	}
	for _, userByte := range usersByte {
		err = json.Unmarshal(userByte, &needUser)
		if err == nil {
			if needUser.Email == registrationUser.Email &&
				bcrypt.CompareHashAndPassword([]byte(needUser.Password), []byte(registrationUser.Password)) == nil {
				DeleteClass(needUser)
				return
			}
		}
	}
}

func GetClassPath(class model3.Storable) (string, error) {
	src, err := storage.GetDir(class)
	if err == nil {
		return filepath.Join(src, class.GetGuid().String()), nil
	}
	return "", err
}

func PrepareMockStorage() storage.Storage {
	var convertor model2.Convertor = model2.JSONGConvertor{}
	var mockStorage = mock.Storage{Convertor: convertor}
	mockStorage.Init()
	return mockStorage
}

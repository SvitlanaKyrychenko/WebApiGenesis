package util

import (
	"WebApiGenesis/model"
	"WebApiGenesis/storage"
	"WebApiGenesis/tests/mock"
	"WebApiGenesis/utils/file"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"path/filepath"
)

func PrepareStorage() storage.Storage {
	var convertor model.Convertor = model.JSONGConvertor{}
	return storage.FileStorage{Convertor: convertor}
}

func PrepareMockClass(message string) model.FileStorable {
	guid, err := ksuid.NewRandom()
	if err != nil {
		return nil
	}
	return mock.StorableClass{Guid: guid, Message: message}
}

func DeleteClass(class model.FileStorable) {
	src, err := GetClassPath(class)
	if err == nil {
		file.DeleteFile(src)
	}
}

func DeleteDBUser(registrationUser model.RegistrationUser, storage storage.Storage) {
	var needUser model.DBUser = model.DBUser{}
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

func GetClassPath(class model.FileStorable) (string, error) {
	src, err := storage.GetDir(class)
	if err == nil {
		return filepath.Join(src, class.GetGuid().String()), nil
	}
	return "", err
}

func PrepareMockStorage() storage.Storage {
	var convertor model.Convertor = model.JSONGConvertor{}
	var mockStorage = mock.Storage{Convertor: convertor}
	mockStorage.Init()
	return mockStorage
}

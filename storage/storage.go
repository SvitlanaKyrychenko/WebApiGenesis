package storage

import (
	"WebApiGenesis/model"
	"WebApiGenesis/utils/file"
	"errors"
	"fmt"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Storage interface {
	AddOrUpdateAsync(obj model.Storable) error
	GetAsync(obj model.Storable) ([]byte, error)
	GetALLAsync(obj model.Storable) (map[ksuid.KSUID][]byte, error)
}

type FileStorage struct {
	Convertor model.Convertor
}

func (storage FileStorage) AddOrUpdateAsync(obj model.Storable) error {
	if obj == nil {
		return errors.New("can not operate with nil object")
	}
	errorChan := make(chan error)
	go func() {
		if storage.Convertor == nil {
			errorChan <- errors.New("converter has not set")
			return
		}
		resJSON, err := storage.Convertor.ConvertByte(obj)
		if err != nil {
			errorChan <- err
			return
		}
		dir, err := GetDir(obj)
		if err != nil {
			errorChan <- err
			return
		}
		file.CreateIfNotExistDir(dir)

		file, err := os.Create(filepath.Join(dir, obj.GetGuid().String()))
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString(string(resJSON))
		errorChan <- nil
	}()
	return <-errorChan
}

func (FileStorage) GetAsync(obj model.Storable) ([]byte, error) {
	if obj == nil {
		return nil, errors.New("can not operate with nil object")
	}
	class := make(chan []byte)
	var err error
	go func() {
		var filePath string
		filePath, err = getFilePath(obj)
		if err != nil {
			return
		}
		data := file.ReadByte(filePath)
		class <- data
	}()
	return <-class, err
}

func getFilePath(obj model.Storable) (string, error) {
	dir, err := GetDir(obj)
	if err != nil {
		return "", err
	}
	var filePath = filepath.Join(dir, obj.GetGuid().String())
	_, err = os.Stat(filePath)
	if err != nil {
		return filePath, err
	}
	return filePath, nil
}

func (FileStorage) GetALLAsync(obj model.Storable) (map[ksuid.KSUID][]byte, error) {
	if obj == nil {
		return nil, errors.New("can not operate with nil object")
	}
	classes := make(chan map[ksuid.KSUID][]byte)
	var err error
	go func() {
		allClasses := make(map[ksuid.KSUID][]byte)
		allClasses, err = getClassesFromFiles(obj)
		classes <- allClasses
	}()
	return <-classes, err
}

func getClassesFromFiles(obj model.Storable) (map[ksuid.KSUID][]byte, error) {
	classes := make(map[ksuid.KSUID][]byte)
	files, err := getDirFiles(obj)
	if err != nil {
		return classes, err
	}
	for _, file := range files {
		guid, err := createGuidFromFile(file)
		if err != nil {
			return nil, err
		}
		currentClass, err := getClassFromFile(obj, guid)
		if err == nil {
			classes[guid] = currentClass
		}
	}
	return classes, nil
}

func getClassFromFile(obj model.Storable, guid ksuid.KSUID) ([]byte, error) {
	var class model.ClassStorable = model.ClassStorable{Guid: guid, NameClass: obj.Name()}
	var storage FileStorage
	currentClass, err := storage.GetAsync(class)
	if err != nil {
		return currentClass, err
	}
	return currentClass, nil
}

func createGuidFromFile(file os.FileInfo) (ksuid.KSUID, error) {
	var guid ksuid.KSUID
	err := guid.Set(file.Name())
	return guid, err
}

func getDirFiles(obj model.Storable) ([]os.FileInfo, error) {
	dir, err := GetDir(obj)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(dir)
	if err != nil {
		return nil, err
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, err
}

func GetDir(obj model.Storable) (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	var dir = filepath.Join(userConfigDir,
		"WebApiGenesisStorage", obj.Name())
	return dir, nil
}

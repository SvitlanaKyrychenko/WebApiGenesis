package storage

import (
	"WebApiGenesis/model"
	"WebApiGenesis/utils/file"
	"fmt"
	"github.com/segmentio/ksuid"
	"io/ioutil"
	"os"
	"path/filepath"
)

type FileStorage struct {
	Convertor model.Convertor
}

type Storage interface {
	AddOrUpdateAsync(obj model.FileStorable) error
	GetAsync(obj model.FileStorable) ([]byte, error)
	GetALLAsync(obj model.FileStorable) (map[ksuid.KSUID][]byte, error)
}

func (storage FileStorage) AddOrUpdateAsync(obj model.FileStorable) error {
	errorChan := make(chan error)
	go func() {
		resJSON, err := storage.Convertor.ConvertByte(obj)
		if err != nil {
			errorChan <- err
			return
		}
		dir, err := getDir(obj)
		if err != nil {
			errorChan <- err
			return
		}
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

func (FileStorage) GetAsync(obj model.FileStorable) ([]byte, error) {
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

func getFilePath(obj model.FileStorable) (string, error) {
	dir, err := getDir(obj)
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

func (FileStorage) GetALLAsync(obj model.FileStorable) (map[ksuid.KSUID][]byte, error) {
	classes := make(chan map[ksuid.KSUID][]byte)
	var err error
	go func() {
		allClasses := make(map[ksuid.KSUID][]byte)
		allClasses, err = getClassesFromFiles(obj)
		classes <- allClasses
	}()
	return <-classes, err
}

func getClassesFromFiles(obj model.FileStorable) (map[ksuid.KSUID][]byte, error) {
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

func getClassFromFile(obj model.FileStorable, guid ksuid.KSUID) ([]byte, error) {
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

func getDirFiles(obj model.FileStorable) ([]os.FileInfo, error) {
	dir, err := getDir(obj)
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

func getDir(obj model.FileStorable) (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	var dir = filepath.Join(userConfigDir,
		"WebApiGenesisStorage", obj.Name())
	return dir, nil
}

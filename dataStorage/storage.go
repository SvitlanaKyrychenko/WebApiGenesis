package dataStorage

import (
	"encoding/json"
	"fmt"
	"github.com/segmentio/ksuid"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func AddOrUpdateAsync(obj IsStorable) bool{
	done := make(chan bool)
	go func() {
		resJSON, err := json.Marshal(obj)
		if err != nil {
			done <- false
		}
		var dir string = getDir(obj)
		if !dirExists(dir) {
			os.MkdirAll(dir, os.ModePerm)
		}
		file, err := os.Create(filepath.Join(dir, obj.GetGuid().String()))
		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer file.Close()
		file.WriteString(string(resJSON))
		done <- true
	}()
	return <-done
}

func GetAsync(obj IsStorable)(bool, []byte) {
	done := make(chan bool)
	var data []byte
	go func() (){
		var dir string = filepath.Join(getDir(obj), obj.GetGuid().String())
		if !dirExists(dir) {
			done <- false
			return
		}
		data = readFile(dir)
		done <- true
	}()
	return <-done, data
}

func GetALLAsync(obj IsStorable)(bool, map[ksuid.KSUID][]byte) {
	done := make(chan bool)
	data:=make(map[ksuid.KSUID][]byte)
	go func() (){
		var dir string = filepath.Join(getDir(obj))
		if !dirExists(dir) {
			done <- false
			return
		}
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			var guid ksuid.KSUID
			valid := guid.Set(file.Name())
			if valid!=nil{
				continue
			}
			var class ClassStorable = ClassStorable{Guid: guid, NameClass: obj.Name()}
			validCurrent, dataCurrent := GetAsync(class)
			if validCurrent{
				data[guid] = dataCurrent
			}
		}
		done <- true
	}()
	return <-done, data
}

func getDir(obj IsStorable) string {
	userConfigDir, errConfig := os.UserConfigDir()
	var dir string = filepath.Join(userConfigDir,
		"WebApiGenesisStorage", obj.Name())
	if errConfig != nil {
		log.Fatal(errConfig)
	}
	return dir
}

func readFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	res := make([]byte, 0)
	var n int
	var err2 error
	for {
		n, err2 = file.Read(data)
		res = append(res, data[:n]...)
		if err2 == io.EOF {
			break
		}
	}
	return res
}

func dirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

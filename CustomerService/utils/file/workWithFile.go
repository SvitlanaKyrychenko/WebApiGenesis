package file

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func ReadByte(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	res := make([]byte, 0)
	for {
		n, err := file.Read(data)
		res = append(res, data[:n]...)
		if err == io.EOF {
			break
		}
	}
	return res
}

func ReadString(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		file, err = os.Create(filePath)
		return ""
	}
	defer file.Close()
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text() + "\n")
	}
	return wr.String()
}

func CreateIfNotExistDir(path string) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(path, 0777)
			if err != nil {
				panic(err)
			}
		}
	}
}

func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func DeleteFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
		return
	}
}

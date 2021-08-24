package file

import (
	"fmt"
	"io"
	"os"
)

func ReadByte (path string) []byte {
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

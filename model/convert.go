package model

import (
	"encoding/json"
)

type Convertor interface {
	ConvertByte(class FileStorable) ([]byte, error)
}

type JSONGConvertor struct {
}

func (JSONGConvertor) ConvertByte(class FileStorable) ([]byte, error) {
	res, err := json.Marshal(class)
	if err != nil {
		return nil, err
	}
	return res, nil
}

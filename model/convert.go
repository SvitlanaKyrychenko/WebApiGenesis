package model

import (
	"encoding/json"
)

type Convertor interface {
	ConvertByte(class Storable) ([]byte, error)
}

type JSONGConvertor struct {
}

func (JSONGConvertor) ConvertByte(class Storable) ([]byte, error) {
	res, err := json.Marshal(class)
	if err != nil {
		return nil, err
	}
	return res, nil
}

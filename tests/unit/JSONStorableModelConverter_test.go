package unit

import (
	"WebApiGenesis/model"
	"WebApiGenesis/tests/mock"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidStorableModelJSONConverter(t *testing.T) {
	//Arrange
	var convertor model.Convertor = model.JSONGConvertor{}
	var mockClass model.FileStorable = prepareMockClass("class")
	//Act
	classByte, errConvert := convertor.ConvertByte(mockClass)
	var classGot mock.StorableClass
	var errUnmarshal error
	if errConvert == nil {
		errUnmarshal = json.Unmarshal(classByte, &classGot)
	}
	//Assert
	require.Nil(t, errConvert)
	require.Nil(t, errUnmarshal)
	require.Equal(t, mockClass, classGot)
}

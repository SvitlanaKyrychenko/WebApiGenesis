package unit

import (
	model2 "WebApiGenesis/BitcoinRateService/model"
	"WebApiGenesis/CustomerService/model"
	"WebApiGenesis/tests/mock"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidStorableModelJSONConverter(t *testing.T) {
	//Arrange
	var convertor model2.Convertor = model2.JSONGConvertor{}
	var mockClass model.Storable = prepareMockClass("class")
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

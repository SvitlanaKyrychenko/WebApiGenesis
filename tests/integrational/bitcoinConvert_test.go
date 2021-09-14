package integrational

import (
	"WebApiGenesis/BitcoinRateService/bitcoin"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitcoinConvertCriptonator(t *testing.T) {
	t.Parallel()
	//Arrange
	var btsConvertor bitcoin.Convertor = bitcoin.ConversionCryptonator{}
	//Act
	btsRate, err := btsConvertor.ToUAH()
	//Assert
	require.Nil(t, err)
	require.Greater(t, btsRate, 0.0)
}

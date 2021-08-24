package integrational

import (
	"WebApiGenesis/bitcoin"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitcoinConvertURLUsedInProgram(t *testing.T) {
	t.Parallel()
	//Arrange
	var btsConvertor bitcoin.Convertor = bitcoin.Conversion{}
	//Act
	btsRate, err := btsConvertor.ToUAH("https://api.cryptonator.com/api/ticker/btc-uah")
	//Assert
	require.Nil(t, err)
	require.Greater(t, btsRate, 0.0)
}

func TestBitcoinConvertWrongURL(t *testing.T) {
	t.Parallel()
	//Arrange
	var btsConvertor bitcoin.Convertor = bitcoin.Conversion{}
	//Act
	btsRate, err := btsConvertor.ToUAH("wrongUrl")
	//Assert
	require.Error(t, err)
	require.Equal(t, -1.0, btsRate)
}

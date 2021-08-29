package unit

import (
	"WebApiGenesis/model"
	"WebApiGenesis/storage"
	"WebApiGenesis/tests/mock"
	"WebApiGenesis/tests/util"
	"WebApiGenesis/utils/file"
	"encoding/json"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidStorageAdd(t *testing.T) {
	//Arrange
	var mockClass model.Storable = util.PrepareMockClass("class")
	var fileStorage storage.Storage = util.PrepareStorage()
	defer util.DeleteClass(mockClass)
	//Act
	errAddOrUpdate := fileStorage.AddOrUpdateAsync(mockClass)
	//Assert
	path, err := util.GetClassPath(mockClass)
	require.Nil(t, errAddOrUpdate)
	require.Nil(t, err)
	require.True(t, file.IsFileExists(path))
}

func TestValidStorageAddAndThenGet(t *testing.T) {
	//Arrange
	var mockClass model.Storable = util.PrepareMockClass("class")
	var fileStorage storage.Storage = util.PrepareStorage()
	defer util.DeleteClass(mockClass)
	//Act
	errAddOrUpdate := fileStorage.AddOrUpdateAsync(mockClass)
	byteClass, errGet := fileStorage.GetAsync(mockClass)
	var classGot mock.StorableClass
	errUnmarshal := json.Unmarshal(byteClass, &classGot)
	//Assert
	require.Nil(t, errAddOrUpdate)
	require.Nil(t, errGet)
	require.Nil(t, errUnmarshal)
	require.Equal(t, mockClass, classGot)
}

func TestValidStorageAddAndThenGetAll(t *testing.T) {
	//Arrange
	var mockClass model.Storable = util.PrepareMockClass("class")
	var mockClass2 model.Storable = util.PrepareMockClass("class2")
	var fileStorage storage.Storage = util.PrepareStorage()
	defer util.DeleteClass(mockClass)
	defer util.DeleteClass(mockClass2)
	//Act
	errAddOrUpdate1 := fileStorage.AddOrUpdateAsync(mockClass)
	errAddOrUpdate2 := fileStorage.AddOrUpdateAsync(mockClass2)
	byteClasses, errGetAll := fileStorage.GetALLAsync(mockClass)
	var errUnmarshal error
	classes := make(map[ksuid.KSUID]model.Storable)
	for _, byteClass := range byteClasses {
		var classGot mock.StorableClass
		errUnmarshal = json.Unmarshal(byteClass, &classGot)
		classes[classGot.Guid] = classGot
	}

	//Assert
	require.Nil(t, errAddOrUpdate1)
	require.Nil(t, errAddOrUpdate2)
	require.Nil(t, errGetAll)
	require.Nil(t, errUnmarshal)
	require.Equal(t, 2, len(classes))
	require.Equal(t, mockClass, classes[mockClass.GetGuid()])
	require.Equal(t, mockClass2, classes[mockClass2.GetGuid()])
}
